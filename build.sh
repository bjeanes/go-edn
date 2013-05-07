#!/usr/bin/env sh

set -e

ragel -Z -T0 -o ragel_parser.go parser.rl
ragel -V -Z -p -o ragel_parser.dot parser.rl
go test