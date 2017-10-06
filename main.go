package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"gopkg.in/jcelliott/turnpike.v2"
)

var (
	realm string
	port  int
	debug bool
)

func init() {
	flag.StringVar(&realm, "realm", "domain.app", "WAMP协议限定域名称")
	flag.IntVar(&port, "port", 8000, "服务器监听端口")
	flag.BoolVar(&debug, "debug", false, "是否开启Debug模式")
}

func main() {
	flag.Parse()
	if debug {
		turnpike.Debug()
	}
	s := turnpike.NewBasicWebsocketServer(realm)
	allowAllOrigin:= func(r *http.Request) bool {return true}
	s.Upgrader.CheckOrigin=allowAllOrigin
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	go func() {
		<-shutdown
		s.Close()
		log.Println("正在关闭服务器...")
		time.Sleep(time.Second)
		os.Exit(1)
	}()

	server := &http.Server{
		Handler: s,
		Addr:    fmt.Sprintf(":%d", port),
	}
	log.Printf("服务器已在[%d]端口运行,限定域名称为[%s] ...", port,realm)
	log.Fatal(server.ListenAndServe())
}