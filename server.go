package transfer

import (
	"github.com/giskook/gotcp"
	"github.com/giskook/transfer/conf"
	"net"
	"time"
)

type ServerConfig struct {
	Listener      *net.TCPListener
	AcceptTimeout time.Duration
}

type Server struct {
	config           *ServerConfig
	srv              *gotcp.Server
	checkconnsticker *time.Ticker
}

var Gserver *Server

func SetServer(server *Server) {
	Gserver = server
}

func GetServer() *Server {
	return Gserver
}

func NewServer(srv *gotcp.Server, config *ServerConfig) *Server {
	serverstatistics := conf.GetConfiguration().ServerStatistics
	return &Server{
		config:           config,
		srv:              srv,
		checkconnsticker: time.NewTicker(time.Duration(serverstatistics) * time.Second),
	}
}

func (s *Server) Start() {
	//go s.checkStatistics()
	go s.srv.Start(s.config.Listener, s.config.AcceptTimeout)
}

func (s *Server) Stop() {
	s.srv.Stop()
	s.checkconnsticker.Stop()
}

//func (s *Server) checkStatistics() {
//	for {
//		<-s.checkconnsticker.C
//		log.Printf("--------------Upstream Connections : %d ----Downstream Connections : %d ---------------\n", conn.NewConnsUpstream().GetCount(), conn.NewConnsDownstream().GetCount())
//	}
//}
