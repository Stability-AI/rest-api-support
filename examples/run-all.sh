#!/bin/sh

# Exit on error
set -e

# Use the 'run-all.sh local' to run these examples against localhost
if [ "$1" = "local" ]; then
    export API_HOST=http://localhost:8080
fi

CWD=$(pwd)

# Go examples
cd "$CWD/go"
./run-all.sh

# Python examples
cd "$CWD/python"
./run-all.sh

# TypeScript examples
cd "$CWD/ts"
./run-all.sh

# cURL examples
cd "$CWD/curl"
./run-all.sh
