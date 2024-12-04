#!/bin/bash
set -e 

# Remove the old documentation
rm -rf docs

# Run godoc
godoc -http=:6060 &
PROCESS_ID=$!

# Wait for godoc to start
sleep 2

# Get all the documentation
! wget -r -np -N -E -p -k http://localhost:6060/pkg/SeedTradingSystem/

# Rename it
echo "Renaming the directory"
mv localhost\:6060/ docs

# kill the process
kill $PROCESS_ID

echo "Documentation generated in docs/"
