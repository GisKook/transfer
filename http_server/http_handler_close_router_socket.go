package http_server

import (
	"fmt"
	"github.com/giskook/transfer/conn"
	"log"
	"net/http"
	"strconv"
)

func CloseRouterSocket(clientid string) string {
	id, _ := strconv.ParseUint(clientid, 10, 64)
	log.Println("CloseRouterSocket0")
	c := conn.NewConnsDownstream().GetConn(id)
	log.Println("CloseRouterSocket1")
	if c != nil {
		c.Close2()
	}
	log.Println("CloseRouterSocket2")

	return EncodingGeneralResponse(0)
}

func CloseRouterSocketHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Fprint(w, EncodingGeneralResponse(HTTP_RESPONSE_RESULT_SERVER_FAILED))
		}
	}()

	r.ParseForm()
	id := r.Form.Get("id")
	if id == "" {
		fmt.Fprint(w, EncodingGeneralResponse(HTTP_RESPONSE_RESULT_PARAMTER_ERR))
		return
	}
	fmt.Fprint(w, CloseRouterSocket(id))
}
