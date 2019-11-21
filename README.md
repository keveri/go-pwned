# Go Pwned

Currently:

* Take a list of emails as input
* Check the email addresses for compromises in data leaks
* Output a report in different formats

Future ideas:

* Keep state of the results and only report the new information
* Plugins for collecting the email list from different sources
* Plugins for notifications
* Slack bot

## Setup

Get API key for the haveibeenpwned API from [here](https://haveibeenpwned.com/API/Key).

Emails should be in a file. One email per line. The file is provided as argument.

## Run

Execute: `GO_PWNED_API_KEY=$(cat api_key.txt) go run go-pwned.go -input emails.txt`

Default output format is `text`, but json can also be outputted with flag `-format json`.
