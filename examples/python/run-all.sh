#!/bin/sh

# Exit on error
set -e

echo "\033[0;34mPython examples:\033[0m"

# Use the 'run-all.sh local' to run these examples against localhost
if [ "$1" = "local" ]; then
    export API_HOST=http://localhost:8080
fi

# Clear output dir
rm -rf out
mkdir out

# Customize the format that 'time' prints
TIMEFORMAT="%3R seconds"

# Execute every .py file in the code directory recursively (ignoring files starting with _)
for file in $(find code -name '*.py' | grep -v '/_'); do
    # Remove /code from $file
    printf "\t %s..." "${file#code/}"
    time poetry run python "./$file"
done
