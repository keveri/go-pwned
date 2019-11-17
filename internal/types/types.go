package types

import (
	"time"
)

type Breach struct {
	Name         string   `json:"Name"`
	Title        string   `json:"Title"`
	Domain       string   `json:"Domain"`
	BreachDate   string   `json:"BreachDate"`   // date
	AddedDate    string   `json:"AddedDate"`    // datetime
	ModifiedDate string   `json:"ModifiedDate"` // datetime
	PwnCount     int      `json:"PwnCount"`
	Description  string   `json:"Description"`
	DataClasses  []string `json:"DataClasses"`
	IsVerified   bool     `json:"IsVerified"`
	IsFabricated bool     `json:"IsFabricated"`
	IsSensitive  bool     `json:"IsSensitive"`
	IsRetired    bool     `json:"IsRetired"`
	IsSpamList   bool     `json:"IsSpamList"`
	LogoPath     string   `json:"LogoPath"`
}

type Paste struct {
	Source     string `json:"Source"`
	Id         string `json:"Id"`
	Title      string `json:"Title"`
	Date       string `json:"Date"` // date
	EmailCount int    `json:"EmailCount"`
}

// NOTE: golang json package doesn't yet support empty arrays and converts them
//  	 to nil in json output. issue: https://github.com/golang/go/issues/31811
type Finding struct {
	Email    string   `json:"Email"`
	Breaches []Breach `json:"Breaches"`
	Pastes   []Paste  `json:"Pastes"`
}

type Report struct {
	Date     time.Time `json:"Date"`
	Findings []Finding `json:"Findings"`
}
