#!/bin/sh

. ./env.sh

# Run the Go program with data race detection enabled and the CGO_ENABLED
# environment variable set to 1. This is done to enable the data race
# detection feature provided by the Go runtime.
#
# Args:
#   None

# Remove the data race detection feature on production builds

go run -race cmd/*.go
