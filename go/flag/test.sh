#!/bin/bash

set -x
go run flag.go
go run flag.go -h
go run flag.go -b
go run flag.go -b=true --number 1,3,5
go run flag.go --number 1,3,5,ng
