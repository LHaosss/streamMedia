package main

import (
	"net/http"
	"steamMedia/api/defs"
	"steamMedia/api/session"
)

//HeaderFieldSession ...
var HeaderFieldSession = "X-Session-Id"

//HeaderFieldUname ...
var HeaderFieldUname = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HeaderFieldSession)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	r.Header.Add(HeaderFieldUname, uname)
	return true
}

//ValidateUser ...
func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HeaderFieldUname)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
