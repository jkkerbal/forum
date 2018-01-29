package main

import (
	"github.com/jkkerbal/forum/data"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sees.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}