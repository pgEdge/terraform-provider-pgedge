// Package fakeapi provides an in-process fake of the pgEdge HTTP API for use
// in acceptance tests. Tests point the provider at the fake's URL via
// PGEDGE_BASE_URL, which means the provider, the OpenAPI-generated client,
// and the JSON marshaling all run as in production — only the network
// endpoint is replaced.
package fakeapi

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/pgEdge/terraform-provider-pgedge/client/models"
)

// API is an in-process fake pgEdge API. Construct with New, seed with the
// Seed* helpers, and read the base URL via URL().
type API struct {
	Server *httptest.Server

	mu            sync.Mutex
	cloudAccounts map[string]*models.CloudAccount
	sshKeys       map[string]*models.SSHKey
	backupStores  map[string]*models.BackupStore
	clusters      map[string]*models.Cluster
	databases     map[string]*models.Database
	tasks         map[string]*models.Task
}

// New starts a fake API bound to a test. The server is closed automatically
// at the end of the test.
func New(t *testing.T) *API {
	t.Helper()
	a := &API{
		cloudAccounts: map[string]*models.CloudAccount{},
		sshKeys:       map[string]*models.SSHKey{},
		backupStores:  map[string]*models.BackupStore{},
		clusters:      map[string]*models.Cluster{},
		databases:     map[string]*models.Database{},
		tasks:         map[string]*models.Task{},
	}
	a.Server = httptest.NewServer(a.handler())
	t.Cleanup(a.Server.Close)
	return a
}

// URL returns the base URL the provider should use.
func (a *API) URL() string { return a.Server.URL }

// ConfigureProvider sets the env vars the provider reads at configure time.
// Use t.Setenv via this helper so the test framework restores them.
func (a *API) ConfigureProvider(t *testing.T) {
	t.Helper()
	t.Setenv("PGEDGE_BASE_URL", a.URL())
	t.Setenv("PGEDGE_CLIENT_ID", "test-client-id")
	t.Setenv("PGEDGE_CLIENT_SECRET", "test-client-secret")
}

// --- seed helpers ---------------------------------------------------------

// SeedCloudAccount inserts a cloud account. ID/timestamps are filled if zero.
func (a *API) SeedCloudAccount(c *models.CloudAccount) *models.CloudAccount {
	a.mu.Lock()
	defer a.mu.Unlock()
	fillID(&c.ID)
	fillTime(&c.CreatedAt)
	fillTime(&c.UpdatedAt)
	if c.Properties == nil {
		c.Properties = map[string]any{}
	}
	a.cloudAccounts[c.ID.String()] = c
	return c
}

// SeedSSHKey inserts an SSH key.
func (a *API) SeedSSHKey(s *models.SSHKey) *models.SSHKey {
	a.mu.Lock()
	defer a.mu.Unlock()
	fillID(&s.ID)
	fillTime(&s.CreatedAt)
	a.sshKeys[s.ID.String()] = s
	return s
}

// SeedBackupStore inserts a backup store.
func (a *API) SeedBackupStore(b *models.BackupStore) *models.BackupStore {
	a.mu.Lock()
	defer a.mu.Unlock()
	fillID(&b.ID)
	fillTime(&b.CreatedAt)
	fillTime(&b.UpdatedAt)
	if b.Status == nil {
		b.Status = ptr("available")
	}
	if b.CloudAccountType == nil {
		b.CloudAccountType = ptr("aws")
	}
	a.backupStores[b.ID.String()] = b
	return b
}

// SeedCluster inserts a cluster. Status defaults to "available".
func (a *API) SeedCluster(c *models.Cluster) *models.Cluster {
	a.mu.Lock()
	defer a.mu.Unlock()
	fillID(&c.ID)
	fillTime(&c.CreatedAt)
	if c.Status == nil {
		c.Status = ptr("available")
	}
	if c.NodeLocation == nil {
		c.NodeLocation = ptr("public")
	}
	a.clusters[c.ID.String()] = c
	return c
}

// SeedDatabase inserts a database. Status defaults to "available".
func (a *API) SeedDatabase(d *models.Database) *models.Database {
	a.mu.Lock()
	defer a.mu.Unlock()
	fillID(&d.ID)
	fillTime(&d.CreatedAt)
	fillTime(&d.UpdatedAt)
	if d.Status == nil {
		d.Status = ptr("available")
	}
	if d.Domain == "" {
		d.Domain = fmt.Sprintf("%s.fake.pgedge.io", *d.Name)
	}
	if d.PgVersion == "" {
		d.PgVersion = "16"
	}
	a.databases[d.ID.String()] = d
	return d
}

// --- routing --------------------------------------------------------------

func (a *API) handler() http.Handler {
	// The pgEdge client appends /v1 to any base URL that lacks it
	// (see client/client.go), so all routes live under that prefix.
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/oauth/token", a.handleOAuth)
	mux.HandleFunc("/v1/cloud-accounts", a.collection(a.listCloudAccounts, a.createCloudAccount))
	mux.HandleFunc("/v1/cloud-accounts/", a.item("/v1/cloud-accounts/", a.getCloudAccount, nil, a.deleteCloudAccount))
	mux.HandleFunc("/v1/ssh-keys", a.collection(a.listSSHKeys, a.createSSHKey))
	mux.HandleFunc("/v1/ssh-keys/", a.item("/v1/ssh-keys/", a.getSSHKey, nil, a.deleteSSHKey))
	mux.HandleFunc("/v1/backup-stores", a.collection(a.listBackupStores, a.createBackupStore))
	mux.HandleFunc("/v1/backup-stores/", a.item("/v1/backup-stores/", a.getBackupStore, nil, a.deleteBackupStore))
	mux.HandleFunc("/v1/clusters", a.collection(a.listClusters, a.createCluster))
	mux.HandleFunc("/v1/clusters/", a.handleClusterSubtree)
	mux.HandleFunc("/v1/databases", a.collection(a.listDatabases, a.createDatabase))
	mux.HandleFunc("/v1/databases/", a.item("/v1/databases/", a.getDatabase, a.patchDatabase, a.deleteDatabase))
	mux.HandleFunc("/v1/tasks", a.handleTasks)
	return mux
}

// recordTask simulates an async backend operation. The pgEdge client's
// PollTaskStatus loop first lists tasks by subject and locks onto the
// running one, then polls by ID until "succeeded". listTasks transitions
// running → succeeded after the unfiltered list returns, so a single
// poll cycle completes immediately.
func (a *API) recordTask(subjectID, subjectKind string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	id := newID().String()
	now := time.Now().UTC().Format(time.RFC3339)
	name := subjectKind + ".update"
	status := "running"
	a.tasks[id] = &models.Task{
		ID:          &id,
		Name:        &name,
		SubjectID:   &subjectID,
		SubjectKind: &subjectKind,
		Status:      &status,
		CreatedAt:   &now,
		UpdatedAt:   &now,
		Messages:    []*models.Message{},
	}
}

func (a *API) handleTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	q := r.URL.Query()
	subjectID := q.Get("subject_id")
	idFilter := q.Get("id")

	a.mu.Lock()
	defer a.mu.Unlock()

	out := make([]*models.Task, 0, len(a.tasks))
	for _, t := range a.tasks {
		if idFilter != "" && (t.ID == nil || *t.ID != idFilter) {
			continue
		}
		if subjectID != "" && (t.SubjectID == nil || *t.SubjectID != subjectID) {
			continue
		}
		out = append(out, t)
	}
	sortByCreatedAt(out, func(t *models.Task) string { return strDeref(t.CreatedAt) })

	writeJSON(w, http.StatusOK, out)

	// After a subject-scoped (no id) list, advance any running tasks to
	// "succeeded" so the next poll-by-id call completes the loop. When a
	// task succeeds, transition its subject from "creating" → "available"
	// to mirror the real backend's async provisioning lifecycle.
	if idFilter == "" {
		succeeded := "succeeded"
		for _, t := range out {
			if t.Status != nil && *t.Status == "running" {
				t.Status = &succeeded
				a.advanceSubjectStatus(t)
			}
		}
	}
}

// advanceSubjectStatus flips a "creating" cluster or database to "available"
// once its provisioning task succeeds. Caller holds a.mu.
func (a *API) advanceSubjectStatus(t *models.Task) {
	if t.SubjectID == nil || t.SubjectKind == nil {
		return
	}
	switch *t.SubjectKind {
	case "cluster":
		if c, ok := a.clusters[*t.SubjectID]; ok && c.Status != nil && *c.Status == "creating" {
			c.Status = ptr("available")
		}
	case "database":
		if d, ok := a.databases[*t.SubjectID]; ok && d.Status != nil && *d.Status == "creating" {
			d.Status = ptr("available")
		}
	}
}

type handlerFunc func(http.ResponseWriter, *http.Request)
type itemHandlerFunc func(http.ResponseWriter, *http.Request, string)

// collection dispatches GET → list, POST → create on a /resource path.
func (a *API) collection(list, create handlerFunc) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			list(w, r)
		case http.MethodPost:
			create(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// item dispatches GET → get, PATCH → patch (if non-nil), DELETE → del on a
// /resource/{id} path. Pass nil for patch on resources without updates.
func (a *API) item(prefix string, get, patch, del itemHandlerFunc) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, prefix)
		if id == "" || strings.Contains(id, "/") {
			http.NotFound(w, r)
			return
		}
		switch r.Method {
		case http.MethodGet:
			get(w, r, id)
		case http.MethodPatch:
			if patch == nil {
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
				return
			}
			patch(w, r, id)
		case http.MethodDelete:
			del(w, r, id)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

// /v1/clusters/ subtree: handles /v1/clusters/{id} and /v1/clusters/{id}/nodes.
func (a *API) handleClusterSubtree(w http.ResponseWriter, r *http.Request) {
	rest := strings.TrimPrefix(r.URL.Path, "/v1/clusters/")
	parts := strings.Split(rest, "/")
	if len(parts) == 1 {
		a.item("/v1/clusters/", a.getCluster, a.patchCluster, a.deleteCluster)(w, r)
		return
	}
	// /v1/clusters/{id}/nodes
	if len(parts) == 2 && parts[1] == "nodes" && r.Method == http.MethodGet {
		a.listClusterNodes(w, r, parts[0])
		return
	}
	http.NotFound(w, r)
}

// --- oauth ----------------------------------------------------------------

func (a *API) handleOAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"access_token": "fake-access-token",
		"token_type":   "Bearer",
		"expires_in":   3600,
	})
}

// --- cloud accounts -------------------------------------------------------

func (a *API) listCloudAccounts(w http.ResponseWriter, _ *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()
	out := make([]*models.CloudAccount, 0, len(a.cloudAccounts))
	for _, c := range a.cloudAccounts {
		out = append(out, c)
	}
	sortByCreatedAt(out, func(c *models.CloudAccount) string { return strDeref(c.CreatedAt) })
	writeJSON(w, http.StatusOK, out)
}

func (a *API) createCloudAccount(w http.ResponseWriter, r *http.Request) {
	var in models.CreateCloudAccountInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeAPIError(w, http.StatusBadRequest, "invalid body: "+err.Error())
		return
	}
	if in.Name == "" || in.Type == nil || *in.Type == "" {
		writeAPIError(w, http.StatusBadRequest, "name and type are required")
		return
	}
	props := map[string]any{}
	if in.Credentials != nil {
		props["credentials"] = in.Credentials
	}
	c := &models.CloudAccount{
		Name:        ptr(in.Name),
		Type:        ptr(*in.Type),
		Description: in.Description,
		Properties:  props,
	}
	a.SeedCloudAccount(c)
	writeJSON(w, http.StatusOK, c)
}

func (a *API) getCloudAccount(w http.ResponseWriter, _ *http.Request, id string) {
	a.mu.Lock()
	c, ok := a.cloudAccounts[id]
	a.mu.Unlock()
	if !ok {
		writeAPIError(w, http.StatusNotFound, "cloud account not found")
		return
	}
	writeJSON(w, http.StatusOK, c)
}

func (a *API) deleteCloudAccount(w http.ResponseWriter, _ *http.Request, id string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.cloudAccounts[id]; !ok {
		writeAPIError(w, http.StatusNotFound, "cloud account not found")
		return
	}
	delete(a.cloudAccounts, id)
	w.WriteHeader(http.StatusNoContent)
}

// --- ssh keys -------------------------------------------------------------

func (a *API) listSSHKeys(w http.ResponseWriter, _ *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()
	out := make([]*models.SSHKey, 0, len(a.sshKeys))
	for _, s := range a.sshKeys {
		out = append(out, s)
	}
	sortByCreatedAt(out, func(s *models.SSHKey) string { return strDeref(s.CreatedAt) })
	writeJSON(w, http.StatusOK, out)
}

func (a *API) createSSHKey(w http.ResponseWriter, r *http.Request) {
	var in models.CreateSSHKeyInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeAPIError(w, http.StatusBadRequest, "invalid body: "+err.Error())
		return
	}
	if in.Name == nil || *in.Name == "" || in.PublicKey == nil || *in.PublicKey == "" {
		writeAPIError(w, http.StatusBadRequest, "name and public_key are required")
		return
	}
	s := &models.SSHKey{Name: in.Name, PublicKey: in.PublicKey}
	a.SeedSSHKey(s)
	writeJSON(w, http.StatusOK, s)
}

func (a *API) getSSHKey(w http.ResponseWriter, _ *http.Request, id string) {
	a.mu.Lock()
	s, ok := a.sshKeys[id]
	a.mu.Unlock()
	if !ok {
		writeAPIError(w, http.StatusNotFound, "ssh key not found")
		return
	}
	writeJSON(w, http.StatusOK, s)
}

func (a *API) deleteSSHKey(w http.ResponseWriter, _ *http.Request, id string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.sshKeys[id]; !ok {
		writeAPIError(w, http.StatusNotFound, "ssh key not found")
		return
	}
	delete(a.sshKeys, id)
	w.WriteHeader(http.StatusNoContent)
}

// --- backup stores --------------------------------------------------------

func (a *API) listBackupStores(w http.ResponseWriter, _ *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()
	out := make([]*models.BackupStore, 0, len(a.backupStores))
	for _, b := range a.backupStores {
		out = append(out, b)
	}
	sortByCreatedAt(out, func(b *models.BackupStore) string { return strDeref(b.CreatedAt) })
	writeJSON(w, http.StatusOK, out)
}

func (a *API) createBackupStore(w http.ResponseWriter, r *http.Request) {
	var in models.CreateBackupStoreInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeAPIError(w, http.StatusBadRequest, "invalid body: "+err.Error())
		return
	}
	if in.Name == nil || *in.Name == "" || in.CloudAccountID == nil || *in.CloudAccountID == "" {
		writeAPIError(w, http.StatusBadRequest, "name and cloud_account_id are required")
		return
	}
	a.mu.Lock()
	cloudType := "aws"
	if ca, ok := a.cloudAccounts[*in.CloudAccountID]; ok && ca.Type != nil {
		cloudType = *ca.Type
	}
	a.mu.Unlock()
	b := &models.BackupStore{
		Name:             in.Name,
		CloudAccountID:   in.CloudAccountID,
		CloudAccountType: ptr(cloudType),
		Properties:       map[string]any{"region": in.Region},
	}
	a.SeedBackupStore(b)
	a.recordTask(b.ID.String(), "backup_store")
	writeJSON(w, http.StatusOK, b)
}

func (a *API) getBackupStore(w http.ResponseWriter, _ *http.Request, id string) {
	a.mu.Lock()
	b, ok := a.backupStores[id]
	a.mu.Unlock()
	if !ok {
		writeAPIError(w, http.StatusNotFound, "backup store not found")
		return
	}
	writeJSON(w, http.StatusOK, b)
}

func (a *API) deleteBackupStore(w http.ResponseWriter, _ *http.Request, id string) {
	a.mu.Lock()
	if _, ok := a.backupStores[id]; !ok {
		a.mu.Unlock()
		writeAPIError(w, http.StatusNotFound, "backup store not found")
		return
	}
	delete(a.backupStores, id)
	a.mu.Unlock()
	a.recordTask(id, "backup_store")
	w.WriteHeader(http.StatusNoContent)
}

// --- clusters -------------------------------------------------------------

func (a *API) listClusters(w http.ResponseWriter, _ *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()
	out := make([]*models.Cluster, 0, len(a.clusters))
	for _, c := range a.clusters {
		out = append(out, c)
	}
	sortByCreatedAt(out, func(c *models.Cluster) string { return strDeref(c.CreatedAt) })
	writeJSON(w, http.StatusOK, out)
}

func (a *API) createCluster(w http.ResponseWriter, r *http.Request) {
	var in models.CreateClusterInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeAPIError(w, http.StatusBadRequest, "invalid body: "+err.Error())
		return
	}
	if in.Name == nil || *in.Name == "" {
		writeAPIError(w, http.StatusBadRequest, "name is required")
		return
	}
	if in.NodeLocation == nil || *in.NodeLocation == "" {
		writeAPIError(w, http.StatusBadRequest, "node_location is required")
		return
	}
	cloudAcctID := in.CloudAccountID
	if cloudAcctID == "" && in.CloudAccount != nil && in.CloudAccount.ID != nil {
		cloudAcctID = *in.CloudAccount.ID
	}
	a.mu.Lock()
	if cloudAcctID == "" || a.cloudAccounts[cloudAcctID] == nil {
		a.mu.Unlock()
		writeAPIError(w, http.StatusBadRequest, "failed to read cloud account")
		return
	}
	ca := a.cloudAccounts[cloudAcctID]
	a.mu.Unlock()

	c := &models.Cluster{
		Name:           in.Name,
		Regions:        in.Regions,
		Nodes:          in.Nodes,
		Networks:       in.Networks,
		FirewallRules:  in.FirewallRules,
		BackupStoreIds: in.BackupStoreIds,
		Capacity:       in.Capacity,
		ResourceTags:   in.ResourceTags,
		SSHKeyID:       in.SSHKeyID,
		NodeLocation:   in.NodeLocation,
		CloudAccount: &models.CloudAccountProperties{
			ID:   ptr(ca.ID.String()),
			Name: strDeref(ca.Name),
			Type: strDeref(ca.Type),
		},
		// Real backend (saas/.../clusters/cluster.go:134-140) initializes new
		// clusters with status "creating"; the async provisioning task flips
		// it to "available" on completion. handleTasks mirrors that.
		Status: ptr("creating"),
	}
	a.SeedCluster(c)
	a.recordTask(c.ID.String(), "cluster")
	writeJSON(w, http.StatusOK, c)
}

func (a *API) getCluster(w http.ResponseWriter, _ *http.Request, id string) {
	a.mu.Lock()
	c, ok := a.clusters[id]
	a.mu.Unlock()
	if !ok {
		writeAPIError(w, http.StatusNotFound, "cluster not found")
		return
	}
	writeJSON(w, http.StatusOK, c)
}

func (a *API) patchCluster(w http.ResponseWriter, r *http.Request, id string) {
	a.mu.Lock()
	c, ok := a.clusters[id]
	a.mu.Unlock()
	if !ok {
		writeAPIError(w, http.StatusNotFound, "cluster not found")
		return
	}
	var in models.UpdateClusterInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeAPIError(w, http.StatusBadRequest, "invalid body: "+err.Error())
		return
	}
	a.mu.Lock()
	if len(in.Regions) > 0 {
		c.Regions = in.Regions
	}
	if len(in.Nodes) > 0 {
		c.Nodes = in.Nodes
	}
	if len(in.Networks) > 0 {
		c.Networks = in.Networks
	}
	if len(in.FirewallRules) > 0 {
		c.FirewallRules = in.FirewallRules
	}
	if len(in.BackupStoreIds) > 0 {
		c.BackupStoreIds = in.BackupStoreIds
	}
	a.mu.Unlock()
	a.recordTask(id, "cluster")
	writeJSON(w, http.StatusOK, c)
}

func (a *API) deleteCluster(w http.ResponseWriter, _ *http.Request, id string) {
	a.mu.Lock()
	if _, ok := a.clusters[id]; !ok {
		a.mu.Unlock()
		writeAPIError(w, http.StatusNotFound, "cluster not found")
		return
	}
	delete(a.clusters, id)
	a.mu.Unlock()
	a.recordTask(id, "cluster")
	w.WriteHeader(http.StatusNoContent)
}

func (a *API) listClusterNodes(w http.ResponseWriter, _ *http.Request, id string) {
	a.mu.Lock()
	c, ok := a.clusters[id]
	a.mu.Unlock()
	if !ok {
		writeAPIError(w, http.StatusNotFound, "cluster not found")
		return
	}
	out := make([]*models.ClusterNode, 0, len(c.Nodes))
	for _, n := range c.Nodes {
		nodeID := newID().String()
		out = append(out, &models.ClusterNode{
			ID:           &nodeID,
			Name:         &n.Name,
			Region:       n.Region,
			InstanceType: &n.InstanceType,
		})
	}
	writeJSON(w, http.StatusOK, out)
}

// --- databases ------------------------------------------------------------

func (a *API) listDatabases(w http.ResponseWriter, _ *http.Request) {
	a.mu.Lock()
	defer a.mu.Unlock()
	out := make([]*models.Database, 0, len(a.databases))
	for _, d := range a.databases {
		out = append(out, d)
	}
	sortByCreatedAt(out, func(d *models.Database) string { return strDeref(d.CreatedAt) })
	writeJSON(w, http.StatusOK, out)
}

func (a *API) createDatabase(w http.ResponseWriter, r *http.Request) {
	var in models.CreateDatabaseInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeAPIError(w, http.StatusBadRequest, "invalid body: "+err.Error())
		return
	}
	if in.Name == nil || *in.Name == "" {
		writeAPIError(w, http.StatusBadRequest, "name is required")
		return
	}
	if in.ClusterID == "" {
		writeAPIError(w, http.StatusBadRequest, "cluster_id is required")
		return
	}
	a.mu.Lock()
	cluster, ok := a.clusters[in.ClusterID.String()]
	a.mu.Unlock()
	if !ok {
		writeAPIError(w, http.StatusBadRequest, "cluster not found")
		return
	}

	// The real backend populates database nodes from the cluster's nodes;
	// mirror that so the provider's plan→state diff comes out clean.
	dbNodes := make([]*models.DatabaseNode, 0, len(cluster.Nodes))
	for _, cn := range cluster.Nodes {
		name := cn.Name
		dbNodes = append(dbNodes, &models.DatabaseNode{
			Name: &name,
			Connection: &models.Connection{
				Database: ptr("defaultdb"),
				Username: ptr("admin"),
				Password: ptr("fake-password"),
				Port:     ptr(int64(5432)),
				Host:     fmt.Sprintf("%s.db.fake.pgedge.io", name),
			},
			Location: &models.Location{
				Latitude:  ptr(0.0),
				Longitude: ptr(0.0),
			},
		})
	}

	d := &models.Database{
		Name:          in.Name,
		ClusterID:     &in.ClusterID,
		ConfigVersion: in.ConfigVersion,
		DisplayName:   in.DisplayName,
		Options:       in.Options,
		Backups:       in.Backups,
		Extensions:    in.Extensions,
		Nodes:         dbNodes,
		Status:        ptr("creating"),
	}
	a.SeedDatabase(d)
	a.recordTask(d.ID.String(), "database")
	writeJSON(w, http.StatusOK, d)
}

func (a *API) getDatabase(w http.ResponseWriter, _ *http.Request, id string) {
	a.mu.Lock()
	d, ok := a.databases[id]
	a.mu.Unlock()
	if !ok {
		writeAPIError(w, http.StatusNotFound, "database not found")
		return
	}
	writeJSON(w, http.StatusOK, d)
}

func (a *API) patchDatabase(w http.ResponseWriter, r *http.Request, id string) {
	a.mu.Lock()
	d, ok := a.databases[id]
	a.mu.Unlock()
	if !ok {
		writeAPIError(w, http.StatusNotFound, "database not found")
		return
	}
	var in models.UpdateDatabaseInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeAPIError(w, http.StatusBadRequest, "invalid body: "+err.Error())
		return
	}
	a.mu.Lock()
	if in.Options != nil {
		d.Options = in.Options
	}
	d.UpdatedAt = ptr(time.Now().UTC().Format(time.RFC3339))
	a.mu.Unlock()
	a.recordTask(id, "database")
	writeJSON(w, http.StatusOK, d)
}

func (a *API) deleteDatabase(w http.ResponseWriter, _ *http.Request, id string) {
	a.mu.Lock()
	if _, ok := a.databases[id]; !ok {
		a.mu.Unlock()
		writeAPIError(w, http.StatusNotFound, "database not found")
		return
	}
	delete(a.databases, id)
	a.mu.Unlock()
	a.recordTask(id, "database")
	w.WriteHeader(http.StatusNoContent)
}

// --- helpers --------------------------------------------------------------

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

// writeAPIError writes a body that the swagger client decodes into a
// runtime.APIError and that the provider's handleAPIError unwraps.
func writeAPIError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]any{"code": status, "message": msg})
}

func ptr[T any](v T) *T { return &v }

func strDeref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// fillID assigns a fresh UUID when the caller left the field nil.
func fillID(p **strfmt.UUID) {
	if *p != nil && (*p).String() != "" {
		return
	}
	id := newID()
	*p = &id
}

func fillTime(p **string) {
	if *p != nil && **p != "" {
		return
	}
	t := time.Now().UTC().Format(time.RFC3339)
	*p = &t
}

func newID() strfmt.UUID {
	var b [16]byte
	_, _ = rand.Read(b[:])
	// RFC 4122 variant + version 4.
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	s := fmt.Sprintf("%s-%s-%s-%s-%s",
		hex.EncodeToString(b[0:4]),
		hex.EncodeToString(b[4:6]),
		hex.EncodeToString(b[6:8]),
		hex.EncodeToString(b[8:10]),
		hex.EncodeToString(b[10:16]),
	)
	return strfmt.UUID(s)
}

// sortByCreatedAt sorts in place ascending by CreatedAt to give tests a
// deterministic order regardless of map iteration.
func sortByCreatedAt[T any](xs []T, key func(T) string) {
	// insertion sort — N is tiny (< 10 items per test).
	for i := 1; i < len(xs); i++ {
		for j := i; j > 0 && key(xs[j-1]) > key(xs[j]); j-- {
			xs[j-1], xs[j] = xs[j], xs[j-1]
		}
	}
}
