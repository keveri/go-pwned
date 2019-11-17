# Go Pwned

Check email addresses for compromises and pastes.

## Setup

Get API key for the haveibeenpwned API from [here](https://haveibeenpwned.com/API/Key).

Emails should be in a file. One email per line. The file is provided as argument.

## Run

Execute: `GO_PWNED_API_KEY=$(cat api_key.txt) go run go-pwned.go emails.txt`
