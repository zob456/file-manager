package main

import (
	"bytes"
	"fmt"
	"github.com/zob456/file-manager/file_checker"
	"github.com/zob456/file-manager/file_compressor"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var arg string

	if len(os.Args) == 2 {
		arg = os.Args[1]
	}

	if len(os.Args) > 2 {
		log.Fatal("You provided too many arguments to execute a function")
	}

	if len(os.Args) < 2 {
		log.Fatal("FAILED to provide an argument for function execution")
	}

	if arg == "checker" {
		// Execute file checker
		err := file_checker.ReadData()
		if err != nil {
			log.Println(err)
		}

	}
	if arg == "compressor" {
		pwd, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		f, err := os.ReadFile(fmt.Sprintf("%s/testdata/TestProcessEligibleChannel2_0_TestProcessEligibleChannel2_0_CODES.csv", pwd))
		if err != nil {
			log.Println(err)
		}
		fi, _ := os.Stat(fmt.Sprintf("%s/testdata/TestProcessEligibleChannel2_0_TestProcessEligibleChannel2_0_CODES.csv", pwd))
		fmt.Println("Before: ", fi.Size())

		r := bytes.NewReader(f)
		rc := io.NopCloser(r)
		file_compressor.Compressor(rc)
	}

	if arg != "compressor" && arg != "checker" {
		fmt.Println("You passed an incorrect command so no function was executed")
	}
}
