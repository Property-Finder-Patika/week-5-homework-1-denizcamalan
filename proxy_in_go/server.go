package main

type server interface {
	clientRequest(packageName string) string
}