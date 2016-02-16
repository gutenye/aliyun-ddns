package main

import (
	"fmt"
	"log"
	"net/http"
)

func Server(port string) {
	http.HandleFunc("/", homeHandler)
	fmt.Println(">> Listen on " + port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		rc.ACCESS_KEY_ID = req.FormValue("access_key_id")
		rc.ACCESS_KEY_SECRET = req.FormValue("access_key_secret")
		err := Update(req.FormValue("id"), req.FormValue("rr"), req.FormValue("value"))
		if err != nil {
			http.Error(res, err.Error(), 400)
		} else {
			fmt.Fprint(res, "Success")
		}
	}
}
