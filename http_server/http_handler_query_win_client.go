package http_server

import (
	"encoding/json"
	"fmt"
	"github.com/giskook/transfer/conn"
	"net/http"
	"strconv"
)

type SingleWinClientProperty struct {
	ClientID                   uint64 `json:"WinID"`
	PeerID                     uint64 `json:"WinPeerID"`
	TransparentTransmissionKey uint32 `json:"WinKey"`
	EstablishedTime            string `json:"WinEstablishedTime"`
	RecvByteCount              uint32 `json:"WinRecvByteCount"`
	SendByteCount              uint32 `json:"WinSendByteCount"`
	Mode                       uint8  `json:"WinIsTTMode"`

	RouterID                         uint64 `json:"RouterID"`
	RouterPeerID                     uint64 `json:"RouterPeerID"`
	RouterTransparentTransmissionKey uint32 `json:"RouterKey"`
	RouterEstablishedTime            string `json:"RouterEstablishedTime"`
	RouterRecvByteCount              uint32 `json:"RouterRecvByteCount"`
	RouterSendByteCount              uint32 `json:"RouterSendByteCount"`
	RouterMode                       uint8  `json:"RouterIsTTMode"`
}

func QueryWinClient(clientid string) string {
	id, _ := strconv.ParseUint(clientid, 10, 64)
	v := conn.NewConnsUpstream().GetConn(id)

	if v != nil {
		peer_v := conn.NewConnsDownstream().GetConn(v.PeerID)
		if peer_v != nil {
			clients := &SingleWinClientProperty{
				ClientID: v.ID,
				PeerID:   v.PeerID,
				TransparentTransmissionKey: v.TransparentTransmissionKey,
				EstablishedTime:            v.EstablishedTime,
				RecvByteCount:              v.RecvByteCount,
				SendByteCount:              v.SendByteCount,
				Mode:                       v.Mode,

				RouterID:                         peer_v.ID,
				RouterPeerID:                     peer_v.PeerID,
				RouterTransparentTransmissionKey: peer_v.TransparentTransmissionKey,
				RouterEstablishedTime:            peer_v.EstablishedTime,
				RouterRecvByteCount:              peer_v.RecvByteCount,
				RouterSendByteCount:              peer_v.SendByteCount,
				RouterMode:                       peer_v.Mode,
			}

			response, _ := json.Marshal(clients)

			return string(response)
		} else {
			clients := &SingleWinClientProperty{
				ClientID: v.ID,
				PeerID:   v.PeerID,
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

	response, _ := json.Marshal(&SingleWinClientProperty{
		ClientID: 0,
	})

	return string(response)
}

func QueryWinClientHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if x := recover(); x != nil {
			fmt.Fprint(w, EncodingGeneralResponse(HTTP_RESPONSE_RESULT_SERVER_FAILED))
		}
	}()

	r.ParseForm()
	clientid := r.Form.Get("id")
	fmt.Fprint(w, QueryWinClient(clientid))
}
