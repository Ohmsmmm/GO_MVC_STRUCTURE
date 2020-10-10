package main

type Hashtag struct {
	Hashtag1 string `json:"hashtag1"`
	Hashtag2 string `json:"hashtag2"`
	Hashtag3 string `json:"hashtag3"`
	Hashtag4 string `json:"hashtag4"`
	Hashtag5 string `json:"hashtag5"`
}

type MsgAccGen struct {
	Msg string `json:"msg"`
}

type AccountIO struct {
	AccountNo int `json:"account_no"`
	Msg string `json:"msg"`
	CountIO int `json:"count_io"`
}
