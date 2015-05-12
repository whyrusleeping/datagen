package main

import (
	"flag"
	"fmt"
	"github.com/dustin/randbo"
	"io"
	"math/rand"
	"os"
)

func writeRandFile(path string, size int64) error {
	fi, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = io.CopyN(fi, randbo.New(), size)
	if err != nil {
		return err
	}

	return fi.Close()
}

func main() {
	maxsize := flag.Int64("maxsize", 0, "max size for file generation")
	minsize := flag.Int64("minsize", 0, "min size for file generation")
	total := flag.Int64("total", 0, "total amount of data to write")
	dirname := flag.String("dirname", "", "name of dir to write into")
	flag.Parse()

	if *maxsize == 0 || *minsize == 0 || *total == 0 || *dirname == "" {
		fmt.Println("please specify all arguments")
		return
	}

	err := os.Mkdir(*dirname, 0775)
	if err != nil && !os.IsExist(err) {
		fmt.Println(err)
		return
	}

	written := int64(0)
	for i := 0; *total > 0; i++ {
		nextSize := rand.Int63n(*maxsize-*minsize) + *minsize
		if nextSize > *total {
			nextSize = *total
		}

		err := writeRandFile(fmt.Sprintf("%s/file%d", *dirname, i), nextSize)
		if err != nil {
			fmt.Printf("wrote %d bytes before encountering error: %s", written, err)
			return
		}

		written += nextSize
		*total -= nextSize

	}
}
