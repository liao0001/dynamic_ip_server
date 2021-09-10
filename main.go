package main

import (
	"flag"
	"fmt"
	"github.com/liao0001/dynamic_ip_server/socket"
	"github.com/liao0001/dynamic_ip_server/utils"
	"log"
	"net/http"
)

var port = flag.String("port", "9210", "http service address")

func Start(w http.ResponseWriter, r *http.Request) {
	port := r.FormValue("port")
	//获取到真实ip  然后建立socket连接
	realIP := utils.RealIP(r)
	//建立连接
	go func() {
		socket.Connect(realIP, port)
	}()

	//数据返回
	w.Write([]byte(realIP))
}

func RealIP(w http.ResponseWriter, r *http.Request) {
	//获取到真实ip  然后建立socket连接
	realIP := utils.RealIP(r)
	//数据返回
	w.Write([]byte(realIP))
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/start", Start)
	http.HandleFunc("/realIP", RealIP)

	fmt.Println("start http: 0.0.0.0:", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", *port), nil))
}
