#!/bin/bash
swag fmt && swag init -g ./cmd/backend.go -o ./pkg/docs/ && go run ./cmd/
