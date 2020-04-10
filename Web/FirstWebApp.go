package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// 路由器实现思路
type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHello(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHello(response http.ResponseWriter, request *http.Request) {
	fmt.Println("=========== start ==========")
	// 解析参数，默认不解析
	request.ParseForm()
	fmt.Println(request.Form)
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["test"]) // have a test with test param: http://localhost:9090/?test=111&test=222
	for key, value := range request.Form {
		fmt.Println("key: ", key)
		fmt.Println("value: ", strings.Join(value, ","))
	}
	fmt.Println("===========  end  ==========")
	// 写入到response，输出到客户端
	fmt.Fprintf(response, "hello, world")
}

func main() {
	// 设置访问的路由
	// 同时强制类型转换函数成为HandlerFunc类型，从而实现底层的ServeHTTP接口
	// ServeHTTP接口会将具体的请求分发至路由的Handler
	http.HandleFunc("/", sayHello)
	// 设置监听端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}

	// 测试自定义的路由器
	//http.ListenAndServe(":9090", &MyMux{})

	/*
	   首先调用Http.HandleFunc
	   按顺序做了几件事：
	   1 调用了DefaultServeMux的HandleFunc
	   2 调用了DefaultServeMux的Handle
	   3 往DefaultServeMux的map[string]muxEntry中增加对应的handler和路由规则

	   其次调用http.ListenAndServe(":9090", nil)
	   按顺序做了几件事情：
	   1 实例化Server
	   2 调用Server的ListenAndServe()
	   3 调用net.Listen("tcp", addr)监听端口
	   4 启动一个for循环，在循环体中Accept请求
	   5 对每个请求实例化一个Conn，并且开启一个goroutine为这个请求进行服务go c.serve()
	   6 读取每个请求的内容w, err := c.readRequest()
	   7 判断handler是否为空，如果没有设置handler（这个例子就没有设置handler），handler就设置为DefaultServeMux
	   8 调用handler的ServeHttp
	   9 在这个例子中，下面就进入到DefaultServeMux.ServeHttp
	   10 根据request选择handler，并且进入到这个handler的ServeHTTP
	     mux.handler(r).ServeHTTP(w, r)
	   11 选择handler：
	   A 判断是否有路由能满足这个request（循环遍历ServeMux的muxEntry）
	   B 如果有路由满足，调用这个路由handler的ServeHTTP
	   C 如果没有路由满足，调用NotFoundHandler的ServeHTTP
	*/

	/*
		Go底层用TCP协议搭建了一个服务，监听设置的端口
		func (srv *Server) Serve(l net.Listener) error {
			defer l.Close()
			var tempDelay time.Duration // how long to sleep on accept failure
			for {
				rw, e := l.Accept()
				if e != nil {
					if ne, ok := e.(net.Error); ok && ne.Temporary() {
						if tempDelay == 0 {
							tempDelay = 5 * time.Millisecond
						} else {
							tempDelay *= 2
						}
						if max := 1 * time.Second; tempDelay > max {
							tempDelay = max
						}
						log.Printf("http: Accept error: %v; retrying in %v", e, tempDelay)
						time.Sleep(tempDelay)
						continue
					}
					return e
				}
				tempDelay = 0
				c, err := srv.newConn(rw)
				if err != nil {
					continue
				}
				go c.serve()
			}
		}
	*/
}
