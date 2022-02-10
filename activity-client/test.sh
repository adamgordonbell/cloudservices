#!/usr/bin/env sh
set -e

echo "=== Add Records ==="
./activity-client -add "overhead press: 70lbs"
./activity-client -add "20 minute walk"

echo "=== Retrieve Records ==="
./activity-client -get 1 | grep "overhead press"
./activity-client -get 2 | grep "20 minute walk"

echo "=== List Records ==="
./activity-client -list
./activity-client -list  | grep "overhead press"
./activity-client -list  | grep "20 minute walk"