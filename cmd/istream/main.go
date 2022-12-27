package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	if err := run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	return nil
}
