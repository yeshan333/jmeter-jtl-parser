#!/usr/bin/env bash

set -euo pipefail

if [[ "$#" -ne 1 ]]; then
  echo "Usage: $0 FILE" >&2
  exit 1
fi

xml_file="$1"

parsed=$(./bin/jtl-parser-linux -o xml "$xml_file" | xmllint --c14n --pretty 1 -)
original=$(xmllint --c14n --pretty 1 "$xml_file")

# diff and ignore white space
cmp -b <(echo ${original}) <(echo ${parsed})