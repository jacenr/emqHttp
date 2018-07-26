package main

import (
	"log"
	"net/http"
)

type account map[string]string
type acl map[string]map[string][]string

func (ac *account) auth(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	log.Println(req.Form)
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

func (al *acl) aclcheck(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	log.Println(req.Form)
	if err != nil {
		log.Fatal("Can't parse post data.")
	}
	username := req.Form.Get("username")
	topic := req.Form.Get("topic")
	access := req.Form.Get("access")
	if username == "" || topic == "" || access == "" {
		w.WriteHeader(http.StatusNotFound)
		log.Println("All param is blank.")
		return
	}
	accessMap, ok := (*al)[username]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Username deny.")
		return
	}
	accessList, ok := accessMap[access]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Access deny.")
		return
	}
	for _, t := range accessList {
		if t == topic {
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	log.Println("Topic deny.")
}

func main() {
	ac := &account{"tom": "tompasswd"}
	al := &acl{
		"tom": {
			"1": {"test"},
			"2": {"test"},
		},
	}
	http.HandleFunc("/auth", ac.auth)
	http.HandleFunc("/acl", al.aclcheck)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
