package model

import "time"

type TestMessage struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}
