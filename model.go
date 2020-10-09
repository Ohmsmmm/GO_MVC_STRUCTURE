package main

type DataCleansingRequest struct {
	InputString string `json:"input_string"`
}

type DataCleansingResponse struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}