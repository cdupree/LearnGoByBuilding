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

    fh, err := os.Open(filenames[0])

    if err != nil {
        // fmt.Println("file", filenames[0], "couldn't be opened")
        log.Fatal("file", filenames[0], " couldn't be opened")
    }

    scanner := bufio.NewScanner(fh)

    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    fh.Close()
}
