package http_server

import (
	"log"
	"net/http"
	"sync"
)

type HttpServer struct {
}

var once sync.Once
var h *HttpServer

func GetHttpServer() *HttpServer {
	once.Do(func() {
		h = &HttpServer{}
	})

	return h
}

func (server *HttpServer) Init() {
	http.HandleFunc(HTTP_QUERY_ALL_ROUTERS, QueryAllRoutersHandler)
	http.HandleFunc(HTTP_QUERY_ALL_WIN_CLIENT, QueryAllWinClientHandler)
	http.HandleFunc(HTTP_QUERY_WIN_CLIENT, QueryWinClientHandler)
	http.HandleFunc(HTTP_QUERY_ROUTER, QueryRouterHandler)
	http.HandleFunc(HTTP_CLOSE_ROUTER_SOCKET, CloseRouterSocketHandler)
}

func (server *HttpServer) Start(addr string) {
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe :", err)
	}
}
