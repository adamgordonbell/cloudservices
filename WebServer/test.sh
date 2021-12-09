#!/usr/bin/env sh

set -e
set -x

echo "=== Test GET ==="

curl -X GET -s localhost:8080 | rg -q "get" 

echo "=== Test POST ==="

curl -X POST -s localhost:8080 | rg -q "post"

echo "Success"