package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Element interface{}

func makeToken() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func verifyMap(r *http.Request, name string, values []Element) bool {
	for _, v := range values {
		fmt.Println(name, v)
		if v == r.Form.Get(name) {
			return true
		}
	}
	return false
}
