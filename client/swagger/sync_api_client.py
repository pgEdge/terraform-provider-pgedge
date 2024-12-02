import os
import shutil
import argparse
import sys
from pathlib import Path

def setup_argparse():
    parser = argparse.ArgumentParser(description='Sync OpenAPI specs from SAAS to provider')
    parser.add_argument('--saas-dir', required=True, help='Path to the SAAS repository root')
    parser.add_argument('--provider-dir', required=True, help='Path to the Terraform provider repository root')
    return parser

def ensure_directory(path):
    if not os.path.exists(path):
        os.makedirs(path)

def copy_oapi_files(saas_dir, provider_dir):
    source_dir = os.path.abspath(os.path.join(saas_dir, 'internal', 'starfleet', 'oapi'))
    target_dir = os.path.abspath(os.path.join(provider_dir))
    
    print(f"Copying from: {source_dir}")
    print(f"Copying to: {target_dir}")
    
    ensure_directory(target_dir)
    
    files_to_copy = [
        'specmerge.py',
        'pgedge_vacuum_rules.yaml',
        'pgedge.yaml',
        'api_common.yaml',
        'api_ssh_keys.yaml',
        'api_backup_stores.yaml',
        'api_cloud_accounts.yaml',
        'api_clusters.yaml',
        'api_databases.yaml',
        'api_tasks.yaml',
        'api_oauth.yaml'
    ]
    
    for file in files_to_copy:
        source = os.path.join(source_dir, file)
        target = os.path.join(target_dir, file)
        if os.path.exists(source):
            shutil.copy2(source, target)
            print(f"Copied {file}")
        else:
            print(f"Warning: File {file} not found in {source_dir}")

def main():
    parser = setup_argparse()
    args = parser.parse_args()
    
    saas_dir = os.path.abspath(args.saas_dir)
    provider_dir = os.path.abspath(args.provider_dir)
    
    print(f"SAAS directory: {saas_dir}")
    print(f"Provider directory: {provider_dir}")
    
    copy_oapi_files(saas_dir, provider_dir)
    
    print("\nSpec sync completed!")

if __name__ == "__main__":
    main()