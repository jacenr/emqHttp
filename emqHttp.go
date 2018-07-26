package main

import (
	"log"
	"net/http"
)

type account map[string]string

func (ac *account) auth(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatal("Can't parse post data.")
	}
	username := req.Form.Get("u")
	passwd := req.Form.Get("P")
	if username == "" || passwd == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	pwd, ok := (*ac)[username]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if passwd != pwd {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	ac := &account{"tom": "tompasswd"}
	http.HandleFunc("/auth", ac.auth)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
