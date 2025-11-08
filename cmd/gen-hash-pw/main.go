package main

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: gen-hash-pw <password>")
	}
	pw := []byte(os.Args[1])

	hashed, err := bcrypt.GenerateFromPassword(pw, 10)
	if err != nil {
		log.Fatal(err)
	}

	println(string(hashed))
}
