package main

import (
	"fmt"
	"go-pwned/internal/reporting"
	"io/ioutil"
	"os"
	"strings"
)

func filterNonEmpty(lines []string) []string {
	var result []string
	for _, line := range lines {
		if len(line) > 0 {
			result = append(result, line)
		}
	}
	return result
}

func readEmails(fileName string) []string {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := strings.Split(string(fileBytes), "\n")
	return filterNonEmpty(lines)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <file_with_emails> \n", os.Args[0])
		os.Exit(0)
	}

	fileName := os.Args[1]
	emails := readEmails(fileName)
	report := reporting.GenerateReport(emails)

	textReport := reporting.PrintableReport(report)
	fmt.Println(textReport)

	jsonReport := reporting.JsonReport(report)
	fmt.Println(jsonReport)
}
