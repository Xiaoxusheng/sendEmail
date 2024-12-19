package main

//go:generate go mod tidy
//go:generate rm -rf  email
//go:generate go build -o  email  "-w -s" -ldflags
//go:generate ./email
