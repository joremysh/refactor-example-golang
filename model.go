package main

type Play struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}
type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}
