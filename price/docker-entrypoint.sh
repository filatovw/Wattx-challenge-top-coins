#!/bin/bash
set -e
set -o pipefail

echo "ENTRYPOINT"

exec /usr/bin/app "$@"