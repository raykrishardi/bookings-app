#!/bin/bash

# go build will ignore *_test.go files (if you run with go run then it will include the *_test.go which will cause an error)
go build -o bookings cmd/web/*.go && ./bookings