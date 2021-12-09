#!/usr/bin/env sh
set -e

echo "=== Insert Test Data ==="

curl -X POST localhost:8080 -d \
'{"activity": {"description": "christmas eve bike class", "time":"2021-12-09T16:34:04Z"}}'

curl -X POST localhost:8080 -d \
'{"activity": {"description": "cross country skiing is horrible and cold", "time":"2021-12-09T16:56:12Z"}}'

curl -X POST localhost:8080 -d \
'{"activity": {"description": "sledding with nephew", "time":"2021-12-09T16:56:23Z"}}'

echo "=== Test Descriptions ==="

curl -X GET localhost:8080 -d '{"id": 0}' | rg -q 'christmas eve bike class'
curl -X GET localhost:8080 -d '{"id": 1}' | rg -q 'cross country skiing'
curl -X GET localhost:8080 -d '{"id": 2}' | rg -q 'sledding'

echo "=== Test Other Fields ==="
curl -X GET localhost:8080 -d '{"id": 0}' | rg -q '2021-12-09T16:34:04Z'
curl -X GET localhost:8080 -d '{"id": 1}' | rg -q '"id":1'

echo "Success"