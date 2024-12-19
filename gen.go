package main

//go:generate go mod tidy
//go:generate rm -rf  email
//go:generate go build -o  email  -ldflags "-w -s"
//go:generate  nohup  ./email &
