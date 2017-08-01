package http_server

import (
	"encoding/json"
	"fmt"
	"github.com/giskook/transfer/conn"
	"net/http"
	"strconv"
)

type SingleRouterProperty struct {
	ID                         uint64 `json:"RouterID"`
	PeerID                     uint64 `json:"RouterPeerID"`
	TransparentTransmissionKey uint32 `json:"RouterKey"`
	EstablishedTime            string `json:"RouterEstablishedTime"`
	RecvByteCount              uint32 `json:"RouterRecvByteCount"`
	SendByteCount              uint32 `json:"RouterSendByteCount"`
	Mode                       uint8  `json:"RouterIsTTMode"`

	WinID                         uint64 `json:"WinID"`
	WinPeerID                     uint64 `json:"WinPeerID"`
	WinTransparentTransmissionKey uint32 `json:"WinKey"`
	WinEstablishedTime            string `json:"WinEstablishedTime"`
	WinRecvByteCount              uint32 `json:"WinRecvByteCount"`
	WinSendByteCount              uint32 `json:"WinSendByteCount"`
	WinMode                       uint8  `json:"WinIsTTMode"`
}

func QueryRouter(clientid string) string {
	id, _ := strconv.ParseUint(clientid, 10, 64)
	v := conn.NewConnsDownstream().GetConn(id)

	if v != nil {
		peer_v := conn.NewConnsUpstream().GetConn(v.PeerID)
		if peer_v != nil {
			clients := &SingleRouterProperty{
				ID:     v.ID,
				PeerID: v.PeerID,
				TransparentTransmissionKey: v.TransparentTransmissionKey,
				EstablishedTime:            v.EstablishedTime,
				RecvByteCount:              v.RecvByteCount,
				SendByteCount:              v.SendByteCount,
				Mode:                       v.Mode,

				WinID:                         peer_v.ID,
				WinPeerID:                     peer_v.PeerID,
				WinTransparentTransmissionKey: peer_v.TransparentTransmissionKey,
				WinEstablishedTime:            peer_v.EstablishedTime,
				WinRecvByteCount:              peer_v.RecvByteCount,
				WinSendByteCount:              peer_v.SendByteCount,
				WinMode:                       peer_v.Mode,
			}

			response, _ := json.Marshal(clients)

			return string(response)
		} else {
			clients := &SingleRouterProperty{
				ID:     v.ID,
				PeerID: v.PeerID,
				TransparentTransmissionKey: v.TransparentTransmissionKey,
				EstablishedTime:            v.EstablishedTime,
				RecvByteCount:              v.RecvByteCount,
				SendByteCount:              v.SendByteCount,
				Mode:                       v.Mode,
			}

			response, _ := json.Marshal(clients)

			return string(response)
		}
	}

	response, _ := json.Marshal(&SingleRouterProperty{
		ID: 0,
	})

	return string(response)
}

func QueryRouterHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Fprint(w, EncodingGeneralResponse(HTTP_RESPONSE_RESULT_SERVER_FAILED))
		}
	}()

	r.ParseForm()
	id := r.Form.Get("id")
	fmt.Fprint(w, QueryRouter(id))
}
