---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "pgedge Provider"
subcategory: ""
description: |-
  Interface with the pgEdge service API.
---

# pgedge Provider

Interface with the pgEdge service API.

## Example Usage

```terraform
provider "pgedge" {
  base_url = "https://api.pgedge.com" // (Optional) The base URL of the pgedge API.
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `base_url` (String) Base Url to use when connecting to the PgEdge service.
