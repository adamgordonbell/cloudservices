#!/usr/bin/env sh
set -e

echo "=== Add Records ==="
./activityclient --add "overhead press: 70lbs"
./activityclient --add "20 minute walk"

echo "=== Retrieve Records ==="
./activityclient --get 0 | grep "overhead press"
./activityclient --get 1 | grep "20 minute walk"