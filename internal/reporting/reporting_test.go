package reporting_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	. "go-pwned/internal/reporting"
)

func Test_GenerateReport(t *testing.T) {
	emails := []string{"a@example.com", "b@example.com"}
	report := GenerateReport(emails)
	var emailsInReport []string
	for _, finding := range report.Findings {
		emailsInReport = append(emailsInReport, finding.Email)
	}
	assert.Equal(t, emails, emailsInReport, "Given emails should be in the report")
}
