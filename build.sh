#!/bin/bash

rm main
go build -o main cmd/main.go
./main