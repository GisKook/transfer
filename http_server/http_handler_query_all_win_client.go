package http_server

import (
	"encoding/json"
	"fmt"
	"github.com/giskook/transfer/conn"
	"net/http"
)

type WinClientProperty struct {
	ClientID                   uint64 `json:"ClientID"`
	PeerID                     uint64 `json:"PeerID"`
	TransparentTransmissionKey uint32 `json:"Key"`
	EstablishedTime            string `json:"EstablishedTime"`
	RecvByteCount              uint32 `json:"RecvByteCount"`
	SendByteCount              uint32 `json:"SendByteCount"`
	Mode                       uint8  `json:"IsTTMode"`
}

func QueryAllWinClients() string {
	clients := make([]*WinClientProperty, 0)
	for _, v := range conn.NewConnsUpstream().Connsindex {
		clients = append(clients, &WinClientProperty{
			ClientID: v.ID,
			PeerID:   v.PeerID,
			TransparentTransmissionKey: v.TransparentTransmissionKey,
			EstablishedTime:            v.EstablishedTime,
			RecvByteCount:              v.RecvByteCount,
			SendByteCount:              v.SendByteCount,
			Mode:                       v.Mode,
		})
	}

	response, _ := json.Marshal(clients)

	return string(response)
}

func QueryAllWinClientHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Fprint(w, EncodingGeneralResponse(HTTP_RESPONSE_RESULT_SERVER_FAILED))
		}
	}()

	r.ParseForm()
	fmt.Fprint(w, QueryAllWinClients())
}
