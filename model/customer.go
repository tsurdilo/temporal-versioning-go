package model

import "time"

type Customer struct {
	AccountNum       string           `json:"accountNum"`
	Name             string        `json:"name"`
	Email            string        `json:"email"`
	CustomerType     string        `json:"allDots"`
	DemoWaitDuration time.Duration `json:"demoWaitDuration"`
}
