package main

import (
	"fmt"
	"github.com/giskook/gotcp"
	"github.com/giskook/transfer"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// read configuration
	conf, err := transfer.ReadConfig("./conf.json")
	configuration := conf.Configure
	transfer.SetConfiguration(configuration)

	checkError(err)

	// creates a tcp listener upstream
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":"+configuration.UpstreamPort)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// creates a tcp server
	config := &gotcp.Config{
		PacketSendChanLimit:    20,
		PacketReceiveChanLimit: 20,
	}
	srv := gotcp.NewServer(config, &transfer.UpstreamCallback{}, &transfer.UpstreamProtocol{})

	// create transfer server
	upstreamconf := &transfer.ServerConfig{
		Listener:      listener,
		AcceptTimeout: time.Duration(configuration.ConnTimeout) * time.Second,
	}
	server_upstream := transfer.NewServer(srv, upstreamconf)
	// starts service
	fmt.Println("upstream listening:", listener.Addr())
	server_upstream.Start()

	// creates a tcp listener downstream
	dtcpAddr, derr := net.ResolveTCPAddr("tcp4", ":"+configuration.DownstreamPort)
	checkError(derr)
	dlistener, derr := net.ListenTCP("tcp", dtcpAddr)
	checkError(derr)

	// creates a tcp server
	dconfig := &gotcp.Config{
		PacketSendChanLimit:    20,
		PacketReceiveChanLimit: 20,
	}
	dsrv := gotcp.NewServer(dconfig, &transfer.DownstreamCallback{}, &transfer.DownstreamProtocol{})

	// create transfer server
	downstreamconf := &transfer.ServerConfig{
		Listener:      dlistener,
		AcceptTimeout: time.Duration(configuration.ConnTimeout) * time.Second,
	}
	server_downstream := transfer.NewServer(dsrv, downstreamconf)
	// starts service
	fmt.Println("downstream listening:", dlistener.Addr())
	server_downstream.Start()

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)

	// stops service
	server_upstream.Stop()
	server_downstream.Stop()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
