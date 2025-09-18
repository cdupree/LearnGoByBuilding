package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	flag.Parse()
	filenames := flag.CommandLine.Args()

	var fh *os.File
	var err error

	if filenames[0] != "-" {
		fh, err = os.Open(filenames[0])
	} else {
		fh = os.Stdin
		err = nil
	}
	if err != nil {
		log.Fatal("file ", filenames[0], " couldn't be opened")
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
