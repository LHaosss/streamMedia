package defs

type UserCreadential struct {
	//打tip
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

type Comments struct {
	Id      string
	VideoId string
	Author  string
	Content string
}


type SimpleSession struct {
	Username string
	TTL int64
}
