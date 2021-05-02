package main

import "time"

type Profile struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Weight     int     `json:"weight"`
	LowTarget  float64 `json:"low"`
	HighTarget float64 `json:"high"`
}

type Data struct {
	ID      string    `json:"id"`
	Time    time.Time `json:"time"`
	BGLevel float64   `json:"bg"`
}
