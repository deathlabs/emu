package models

import "time"

type ResponseData struct {
	Status     string              `json:"status"`
	StatusCode int                 `json:"status_code"`
	Headers    map[string][]string `json:"headers"`
	Body       interface{}         `json:"body"`
	Timestamp  time.Time           `json:"timestamp"`
	Duration   time.Duration       `json:"duration"`
}
