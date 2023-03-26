#!/bin/sh

set -e

echo "\033[0;34mcURL examples:\033[0m"

# Use the 'run-all.sh local' to run these examples against localhost
if [ "$1" = "local" ]; then
    export API_HOST=http://localhost:8080
fi

# Clear output dir
rm -rf out
mkdir out

# Customize the format that 'time' prints
TIMEFORMAT="%3R seconds"

# Execute every .sh file in the code directory recursively
for file in $(find code -name '*.sh'); do
    printf "\t %s..." "${file#code/}"

    # Run the script and silence stdout but keep stderr
    time "./$file" 1> /dev/null
done
