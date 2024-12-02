"""
Accepts a list of OpenAPI spec paths and merged them into a single spec.
This allows more modular and incremental development of OpenAPI specs.

Usage:
    python specmerge.py -s spec1.yaml spec2.yaml -o combined.yaml

Specifically this is meant to help with incremental additions to
components/schemas/* and paths/*. 

The top-level attributes of `info`, `openapi`, and `servers` should
be just defined in one of the input specs or be the same in all of them.

It's fine if one input spec references a schema defined in another input spec.
"""
import argparse
import yaml
from typing import Dict, List, Any


def read_yaml(path: str):
    """
    Reads and parses a yaml file located at the given path.
    """
    with open(path) as f:
        return yaml.safe_load(f)


def merge(specs: List[Dict[str, Any]]):
    """
    Merges the list of given OpenAPI specs into a single output spec.
    """
    combined = {}
    for spec in specs:
        for key, value in spec.items():
            if key not in combined:
                combined[key] = value
            elif key == "components":
                combined[key].setdefault("schemas", {})
                combined[key].setdefault("securitySchemes", {})
                for schema_name, schema in value.get("schemas", {}).items():
                    combined[key]["schemas"][schema_name] = schema
                for schema_name, schema in value.get("securitySchemes", {}).items():
                    combined[key]["securitySchemes"][schema_name] = schema
            elif key == "paths":
                combined[key].update(value)
    return combined


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "-s",
        "--spec",
        nargs="+",
        type=str,
        required=True,
        help="The spec files to merge",
    )
    parser.add_argument("-o", "--output", help="The output file")
    args = parser.parse_args()

    specs = [read_yaml(path) for path in args.spec]
    result = yaml.dump(merge(specs), sort_keys=True)
    if args.output:
        with open(args.output, "w") as f:
            f.write(result)
    else:
        print(result)
