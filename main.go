package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

type WordCount struct {
	word string
	count int
}

func main() {
	OpenFile,err := os.Open("Biodata.txt")
	if err != nil {
		fmt.Printf("Error: Could not find the file. %v\n",err)  // %v -->  Print Default format..
	}
	defer OpenFile.Close()

	StoreData := make(map[string]int)

	Scanner := bufio.NewScanner(OpenFile)  // Creates a reader configured to split input into line..
	for Scanner.Scan() {   //Scan reads the next line (token),Returns:true → successfully read something,false → EOF or error
		line := Scanner.Text() // Returns the line that was just read by Scan(),Does not read anything
		fmt.Println(line)
	}
	if err := Scanner.Err(); err != nil {
    log.Fatal(err)
	}

	
}