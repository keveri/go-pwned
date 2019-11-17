package types

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
