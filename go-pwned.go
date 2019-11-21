package main

import (
	"flag"
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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func validArguments(filepath string, format string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		fmt.Println("ERROR: Given filepath not found!")
		os.Exit(2)
	}
	validFormats := []string{"text", "json"}
	return (filepath != "") && stringInSlice(format, validFormats)
}

func main() {
	filepath := flag.String("input", "", "filepath to email list")
	format := flag.String("format", "text", "output format. possible values: json, text")
	flag.Parse()

	if !(validArguments(*filepath, *format)) {
		flag.Usage()
		os.Exit(1)
	}

	emails := readEmails(*filepath)
	report := reporting.GenerateReport(emails)

	switch *format {
	case "json":
		jsonReport := reporting.JsonReport(report)
		fmt.Println(jsonReport)
	default:
		textReport := reporting.PrintableReport(report)
		fmt.Println(textReport)
	}
}
