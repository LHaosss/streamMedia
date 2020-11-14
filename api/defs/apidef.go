package defs

//UserCreadential ...
type UserCreadential struct {
	//打tip
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

//SignedUp ...
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

//VideoInfo ...
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

//Comments ...
type Comments struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

//SimpleSession ...
type SimpleSession struct {
	Username string
	TTL      int64
}
