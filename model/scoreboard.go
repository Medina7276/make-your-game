package model

type scoreboard struct {
	name  string `json:"String"`
	rank  int    `json:"Rank"`
	score int    `json:"Score"`
	time  string `json:"Time"`
}
