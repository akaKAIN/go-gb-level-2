package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/akaKAIN/go-gb-level-2/finder"
	"log"
	"os"
	"strings"
)

var (
	path *string
	file *string
	d    *bool
)

func init() {
	path = flag.String("path", ".", "dir for search")
	file = flag.String("file", "", "file name for search")
	d = flag.Bool("d", false, "Need to delete found copy files?")
	flag.Parse()
}

func main() {
	copyList, err := finder.FindCopy(*path, *file)
	if err != nil {
		log.Fatalf("Find copy error: %v", err)
	}
	copyCount := len(copyList)

	if copyCount > 0 {
		fmt.Printf("Found copies: %d\n", copyCount)
		for i, copyName := range copyList {
			fmt.Printf("%d. %s\n", i+1, copyName)
		}
	} else {
		fmt.Printf("No copy of %q in path: %q\n", *file, *path)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	if *d {
		fmt.Print("\nAre you sure you want to delete the copies? (for confirm enter: 'yes'):\t")
		scanner.Scan()
		userAnswer := scanner.Text()
		if strings.ToLower(userAnswer) == "yes" {
			for _, copyName := range copyList {
				fmt.Println("- deleted:", copyName)
			}
		}
	}
}
