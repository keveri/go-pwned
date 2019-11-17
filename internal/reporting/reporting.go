package reporting

import (
	"encoding/json"
	"fmt"
	"go-pwned/internal/api"
	"go-pwned/internal/reporting/templates"
	. "go-pwned/internal/types"
	"os"
	"strings"
	"time"
)

// TODO: include errors to responses instead of exiting
func emailToFinding(email string) Finding {
	breaches, err := api.GetAllBreachesForEmail(email)
	if err != nil {
		fmt.Printf("API call failed with error %s\n", err)
		os.Exit(1)
	}
	pastes, err := api.GetAllPastesForEmail(email)
	if err != nil {
		fmt.Printf("API call failed with error %s\n", err)
		os.Exit(1)
	}
	return Finding{Email: email, Breaches: breaches, Pastes: pastes}
}

func GenerateReport(emails []string) Report {
	timestamp := time.Now().UTC()
	var findings []Finding
	for _, email := range emails {
		findings = append(findings, emailToFinding(email))
	}
	return Report{Date: timestamp, Findings: findings}
}

func PrintableReport(report Report) string {
	var output strings.Builder
	headerStr := templates.Header(report)
	output.WriteString(headerStr)
	for _, finding := range report.Findings {
		findingStr := templates.FindingContent(finding)
		output.WriteString(findingStr)
	}
	return output.String()
}

func JsonReport(report Report) string {
	jsonReport, _ := json.Marshal(report)
	return string(jsonReport)
}
