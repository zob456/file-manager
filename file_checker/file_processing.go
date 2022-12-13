package file_checker

import (
	"encoding/csv"
	"fmt"
	"github.com/zob456/file-manager/models"
	"io"
	"log"
	"os"
	"strings"
)

func ReadData() error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	f, err := os.Open(fmt.Sprintf("%s/testdata", pwd))
	if err != nil {
		return err
	}
	fileInfo, err := f.ReadDir(-1)
	if err != nil {
		return err
	}
	ch := make(chan bool)
	for _, file := range fileInfo {
		fileData, err := os.Open(pwd + "/testdata/" + file.Name())
		if err != nil {
			return err
		}
		cReader := csv.NewReader(fileData)
		cReader.FieldsPerRecord = -1
		go func() {
			CheckCode(fileData, ch)
			if err != nil {
				log.Println(err)
			}
		}()
		b := <-ch
		if b {
			return nil
		}
	}
	close(ch)
	return nil
}

func CheckCode(data io.Reader, ch chan bool) {
	cReader := csv.NewReader(data)
	cReader.FieldsPerRecord = -1
	records, err := cReader.ReadAll()
	if err != nil {
		log.Println(err)
	}
	code := ""
	for _, record := range records {
		file := models.File{}
		file.Barcode = record[0]
		file.Code = record[1]
		file.YearWeek = record[2]
		if strings.Contains(code, file.Code) {
			log.Println("Duplicate code found. Stopping execution at code: ", file.Code)
			ch <- true
		}
		code += file.Code
	}
}
