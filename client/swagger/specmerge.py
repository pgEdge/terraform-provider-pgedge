import yaml
import sys
import os
import glob

def specmerge(main_file):
    with open(main_file, 'r') as f:
        main_data = yaml.safe_load(f)

    # Get the directory of the main file
    dir_path = os.path.dirname(main_file)
    
    # Find all yaml files in the directory, excluding the main file
    yaml_files = [f for f in glob.glob(os.path.join(dir_path, "*.yaml")) if f != main_file]
    
    for file in yaml_files:
        with open(file, 'r') as f:
            data = yaml.safe_load(f)
            if data:  # Check if the file is not empty
                main_data['paths'].update(data.get('paths', {}))
                main_data['definitions'].update(data.get('definitions', {}))

    return main_data

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python merge_yaml.py <main_file>")
        sys.exit(1)

    main_file = sys.argv[1]
    merged = specmerge(main_file)
    yaml.dump(merged, sys.stdout)