#!/bin/sh

set -e

echo "\033[0;34mGo examples:\033[0m"

# Use the 'run-all.sh local' to run these examples against localhost
if [ "$1" = "local" ]; then
    export API_HOST=http://localhost:8080
fi

# Clear output dir
rm -rf out
mkdir out

# Customize the format that 'time' prints
TIMEFORMAT="%3R seconds"

# Execute every .go file in the code directory recursively
for file in $(find code -name '*.go'); do
    printf "\t %s..." "${file#code/}"
    time go run "./$file"
done
