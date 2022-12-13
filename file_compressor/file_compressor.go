package file_compressor

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func Compressor(rc io.ReadCloser) io.Reader {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	//var compressedBuf bytes.Buffer
	var buf bytes.Buffer
	_, err = buf.ReadFrom(rc)
	if err != nil {
		return nil
	}
	zt, err := os.Create(fmt.Sprintf("%s/testdata/file.gz", pwd))
	w, err := gzip.NewWriterLevel(zt, gzip.BestCompression)
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(buf.Bytes())
	if err != nil {
		log.Println(err)
	}
	err = w.Close()
	if err != nil {
		log.Println(err)
	}
	fi, _ := os.Stat(fmt.Sprintf("%s/testdata/file.gz", pwd))
	fmt.Println("After: ", fi.Size())
	r := bytes.NewReader(buf.Bytes())
	return r
}
