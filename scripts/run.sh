#!/usr/bin/env bash
# Do not use this script manually, Use makefile

set -e

source ./scripts/setup-variables.sh

###############################################
# This script is used to start monitoror core #
###############################################

# Set environment (default: development)
export MO_ENV=${MO_ENV:-$MB_ENVIRONMENT}
export MO_DISABLEUI=true

target=$MB_BINARIES_PATH/monitoror-run
go build -o "$target" --ldflags "$MB_GO_LDFLAGS" --tags "$MB_GO_TAGS" "${MB_SOURCE_PATH}"
$target "$@"
