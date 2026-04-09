package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
	"strings"
	"time"
	"unicode"
)

type WordCount struct {
	Word string
	Count int
}

func Cleanword(s string) string {
	var b strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}
// func OpenFile(filename string) (*os.File, error) {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Error: Please provide a filename.")
// 		fmt.Println("Usage: go run main.go <filename.txt>")
// 	}
// 	filename := os.Args[1]

// 	OpenFile,err := os.Open(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer OpenFile.Close()
// 	return &File,null
// }

func main() {
	start := time.Now()

	// OpenFile,err := os.Open("Biodata.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer OpenFile.Close()

	// Lets try to take argument from command line....
	// if len(os.Args) < 2 {  //In real CLI tools (like git, docker):Error → one line,Usage/help → separate line or block
	// 	fmt.Println("Error: Please provide a filename.")
	// 	fmt.Println("Usage: go run main.go <filename.txt>")
	// 	return
	// }
	
	filename := os.Args[1]

	OpenFile,err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer OpenFile.Close()

	StoreData := make(map[string]int)

	Scanner := bufio.NewScanner(OpenFile)  // Creates a reader configured to split input into line..
	for Scanner.Scan() {   //Scan reads the next line (token),Returns:true → successfully read something,false → EOF or error
		line := Scanner.Text() // Returns the line that was just read by Scan(),Does not read anything
		// fmt.Println(line)
		// "Cleaning" the line: split into words and remove junk...
		words := strings.Fields(line)  // Automatically  Handle tabs,multiple Space and Newline as a Seperator..
		for _, w := range words {
			// Clean "Go!!" into "go"
			cleanWord := strings.ToLower(strings.Trim(w, ".,!?:;\"()"))  // Normalize it Now go and Go are same...
			if cleanWord != "" {
				StoreData[cleanWord]++
			}
		}
	}
	if err := Scanner.Err(); err != nil {  // “Did scanning stop because of an error?”
    log.Fatal(err)
	}

	var SortedList []WordCount
	// Conversion of map to slice
	for word, count := range StoreData {
		SortedList = append(SortedList, WordCount{Word : word, Count : count})
	}
	// Lets sort the Slice....
	sort.Slice(SortedList,func(i, j int) bool {
		return SortedList[i].Count > SortedList[j].Count
	})

	// Lets Print the Result (but what we wants to print)..
	for i := 0; i < len(SortedList); i++ {
		fmt.Printf("%d %s: %d\n", i+1, SortedList[i].Word, SortedList[i].Count)
	}

	// Memory Check Status...
	var m runtime.MemStats 
	runtime.ReadMemStats(&m)
	fmt.Printf("\nMemory Used: %d KB\n", m.Alloc/1024)
	fmt.Printf("Execution Time: %v\n", time.Since(start))
}