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

	err := file_checker.ReadData()
	if err != nil {
		log.Println(err)
	}
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
