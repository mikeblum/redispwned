package censys

import (
	"strings"
	"time"
)

// ISO 1601 ex. YYYY-MM-DD HH:MM:SS
const iso8601Timestamp = "2006-01-02 15:04:05"

type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	dt.Time, err = time.Parse(iso8601Timestamp, s)
	return
}

type Account struct {
	Login      string   `json:"login"`
	Email      string   `json:"email"`
	FirstLogin DateTime `json:"first_login"`
	LastLogin  DateTime `json:"last_login"`
	Quota      Quota    `json:"quota"`
}

type Quota struct {
	Used      int      `json:"used"`
	ResetsAt  DateTime `json:"resets_at"`
	Allowance int      `json:"allowance"`
}
