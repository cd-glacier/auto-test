package main

import (
	"log"
	"strings"
)

// これを検知するtestを追加する
func ReDefinication(s string) string {
	v := s
	if strings.Contains(v, "hoge") {
		v = "arg contains hoge"
		log.Println(v)
	} else if len(s) >= 10 {
		v = "arg is more than 10"
		log.Println(v)
	} else {
		// trap
		v := "arg is other one"
		log.Println(v)
	}
	return v
}
