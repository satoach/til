package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("form/login.gtpl")
		t.Execute(w, makeToken())
	} else {
		r.ParseForm()

		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("check token:", token)
		} else {
			fmt.Println("no token")
		}

		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		fmt.Println("fruit:", r.Form["fruit"],
			verifyMap(r, "fruit", []Element{"apple", "pear", "banane"}))
		fmt.Println("gender:", r.Form["gender"],
			verifyMap(r, "gender", []Element{"1", "2"}))
	}
}
