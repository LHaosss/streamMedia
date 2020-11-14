package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"steamMedia/api/dbops"
	"steamMedia/api/defs"
	"steamMedia/api/session"

	"github.com/julienschmidt/httprouter"
)

//CreateUser ...创建用户函数
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCreadential{}

	//用Unmarshal对res获得的数据进行反序列化，并判断数据是否规范，将反序列化后的数据赋值给ubody，如果数据不符合规范或者添加数据到数据库报错，返回一个ErrorrResponse,如果都没毛病，使用GenerateNewSessionID函数生成一个sessionid
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}
//生成一个sessionID
	id := session.GenerateNewSessionID(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}

	//将su序列化赋值给resp，判断是否有问题，如果没问题，输出一个NomalResponse
	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

//Login ...
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
