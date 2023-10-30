#!/bin/bash
rm -rf go.mod go.sum
go mod init main
go mod tidy