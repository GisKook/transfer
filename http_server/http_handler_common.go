package http_server

import (
	"encoding/json"
	"net/http"
)

func CheckParamters(r *http.Request, keys ...string) bool {
	for _, key := range keys {
		value := r.Form.Get(key)
		if value == "" {
			return false
		}
	}

	return true
}

type GeneralResponse struct {
	Result uint8  `json:"result"`
	Desc   string `json:"desc"`
}

func EncodingGeneralResponse(result uint8) string {
	general_response := &GeneralResponse{
		Result: result,
		Desc:   HTTP_RESULT[result],
	}

	response, _ := json.Marshal(general_response)

	return string(response)
}
