package templates

import (
	"fmt"
	. "github.com/keveri/go-pwned/internal/types"
	"strings"
)

func Header(report Report) string {
	template := `Breaches and Pastes searched for %d emails
Report generated %s
==========`
	return fmt.Sprintf(
		template,
		len(report.Findings),
		report.Date,
	)
}

// TODO: include more information and simplify
func FindingContent(finding Finding) string {
	headerTemplate := `

%s

`
	breachTemplate := `Breach:  %s
Classes: %s

`
	footer := "----------"
	var output strings.Builder
	headerStr := fmt.Sprintf(headerTemplate, finding.Email)
	output.WriteString(headerStr)
	for _, breach := range finding.Breaches {
		breachStr := fmt.Sprintf(
			breachTemplate,
			breach.Title,
			breach.DataClasses,
		)
		output.WriteString(breachStr)
	}
	output.WriteString(footer)
	return output.String()
}
