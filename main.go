package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a search term")
		return
	}

	input := strings.Join(os.Args[1:], " ")

	re := regexp.MustCompile(`-\w+`)
	flagsFound := re.FindAllString(input, -1)
	query := re.ReplaceAllString(input, "")

	flags := parseFlagsString(flagsFound)

	data, err := getWikipediaSummary(query, flags)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println(data.title)
	fmt.Println()
	fmt.Println(data.extract)
	fmt.Println()
}
