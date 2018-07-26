package main

import (
	"log"
	"net/http"
)

type account map[string]string

func (ac *account) auth(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	log.Println("req.Form")
	if err != nil {
		log.Fatal("Can't parse post data.")
	}
	username := req.Form.Get("username")
	passwd := req.Form.Get("password")
	log.Println(username)
	log.Println(passwd)
	if username == "" || passwd == "" {
		w.WriteHeader(http.StatusNotFound)
		log.Println("username or passwd is blank.")
		return
	}
	pwd, ok := (*ac)[username]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		log.Println("username is not exist.")
		return
	}
	if passwd != pwd {
		w.WriteHeader(http.StatusNotFound)
		log.Println("passwd is wrong.")
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	ac := &account{"tom": "tompasswd"}
	http.HandleFunc("/auth", ac.auth)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
