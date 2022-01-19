package data

import "time"

type LogInsert struct {
	LogID        string    `json:"log_id" db:"log_id"`
	Url          string    `json:"url" db:"url"`
	Action       string    `json:"action" db:"action"`
	HttpMethod   string    `json:"http_method" db:"http_method"`
	Request      string    `json:"request" db:"request"`
	Response     string    `json:"response" db:"response"`
	ResponseTime float64   `json:"response_time" db:"response_time"`
	Timestamp    time.Time `json:"timestamp" db:"timestamp"`
}

type MoviesData interface {
	InsertLog(payload LogInsert) error
}
