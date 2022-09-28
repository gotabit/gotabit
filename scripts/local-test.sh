#!/bin/sh

BASE_DIR=$(dirname "$0")

$BASE_DIR/docker/chain_init "$@"
$BASE_DIR/docker/chain_start "--log_level warn"
