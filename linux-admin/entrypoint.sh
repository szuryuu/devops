#!/bin/bash
set -e

echo "Running startup scripts..."

# Run your scripts
# ./scripts/remove-user.sh
./scripts/create-user.sh

echo "All scripts completed successfully"

# Keep container running or start your main application
exec "$@"
