#!/bin/bash
set -e

echo "Running startup scripts..."

./scripts/create-user.sh
./scripts/disk-monitor.sh

echo "Cron started"
cron -f

exec "$@"
