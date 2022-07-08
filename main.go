package main

import "log"

func main() {
	log.Println(getMessage())
}

func getMessage() string {
	return "Welcome to deploybot with GitHub Actions!"
}
