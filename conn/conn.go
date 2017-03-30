package conn

import (
	"bytes"
	"github.com/giskook/gotcp"
	//	"github.com/giskook/transfer/base"
	"log"
	"time"
)

var ConnSuccess uint8 = 0
var ConnUnauth uint8 = 1

type ConnConfig struct {
	ConnCheckInterval uint16
	ReadLimit         uint16
	WriteLimit        uint16
}

type Conn struct {
	conn                       *gotcp.Conn
	config                     *ConnConfig
	recieveBuffer              *bytes.Buffer
	ticker                     *time.Ticker
	readflag                   int64
	writeflag                  int64
	closeChan                  chan bool
	index                      uint32
	ID                         uint64
	Status                     uint8
	PeerID                     uint64
	TransparentTransmissionKey uint32
	ReadMore                   bool
}

func NewConn(conn *gotcp.Conn, config *ConnConfig) *Conn {
	return &Conn{
		conn:          conn,
		recieveBuffer: bytes.NewBuffer([]byte{}),
		config:        config,
		readflag:      time.Now().Unix(),
		writeflag:     time.Now().Unix(),
		ticker:        time.NewTicker(time.Duration(config.ConnCheckInterval) * time.Second),
		closeChan:     make(chan bool),
		index:         0,
		Status:        ConnUnauth,
		ReadMore:      true,
	}
}

func (c *Conn) Close() {
	c.closeChan <- true
	c.ticker.Stop()
	c.recieveBuffer.Reset()
	close(c.closeChan)
}

func (c *Conn) GetBuffer() *bytes.Buffer {
	return c.recieveBuffer
}

func (c *Conn) SendToTerm(p gotcp.Packet) {
	//c.packetNsqReceiveChan <- p
	c.conn.AsyncWritePacket(p, time.Second)
}

func (c *Conn) UpdateReadflag() {
	c.readflag = time.Now().Unix()
}

func (c *Conn) UpdateWriteflag() {
	c.writeflag = time.Now().Unix()
}

func (c *Conn) checkHeart() {
	defer func() {
		c.conn.Close()
	}()

	var now int64
	for {
		select {
		case <-c.ticker.C:
			now = time.Now().Unix()
			log.Printf("%x mac %x check now %d read flag %d\n", &c, c.ID, now, c.readflag)
			if now-c.readflag > int64(c.config.ReadLimit) {
				log.Printf("read limit %x\n", c.ID)
				return
			}
			//			if now-c.writeflag > int64(c.config.WriteLimit) {
			//				log.Println("write limit")
			//				return
			//			}
			//			if c.Status == ConnUnauth {
			//				log.Printf("unauth's gateway gatewayid %x\n", c.ID)
			//				return
			//			}
		case <-c.closeChan:
			return
		}
	}
}

func (c *Conn) Do() {
	go c.checkHeart()
}
