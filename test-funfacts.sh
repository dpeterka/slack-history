#!/bin/bash

# Test script to preview different fun facts without posting to Slack
# This is a simple wrapper around the test-funfacts-simple.go program

if [ $# -eq 0 ]; then
    # Run with default date range
    go run test-funfacts-simple.go
else
    # Run with specific date seed
    go run test-funfacts-simple.go "$1"
fi
