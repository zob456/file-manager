package main

import (
	"github.com/zob456/file-manager/file_checker"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := file_checker.ReadData()
	if err != nil {
		log.Fatalln(err)
	}
}
