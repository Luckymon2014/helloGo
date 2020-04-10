package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 此处根据Go的工作路径写相对路径
		t, _ := template.ParseFiles("Web\\login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		// FormValue函数可以获取提交的参数，并且会自动调用ParseForm，但只会获取同名参数的第一个，若不存在则返回空字符串
		fmt.Println("username is ", r.FormValue("username"))
		/*
			必须显示的调用ParseForm才能获取到完整的表单信息
			r.Form里面包含了所有请求的参数，比如URL中query-string、POST的数据、PUT的数据
			当URL中的query-string字段和POST冲突时，会保存成一个slice，里面存储了多个值
			Go官方文档中说在接下来的版本里面将会把POST、GET这些数据分离开来
		*/
		r.ParseForm()
		fmt.Println("username", r.Form["username"])                                // 这里会打印出post的username和url中的username合并成的slice
		fmt.Println("password", template.HTMLEscapeString(r.Form.Get("password"))) // Go自带了过滤html标签的防止XSS攻击的方法
	}
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9090", nil)
}
