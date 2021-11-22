package helpers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status  int         `json:"status"`
	Desc    string      `json:"desc"`
	IsError bool        `json:"isError"`
	Result  interface{} `json:"result"`
}

func Respone(w http.ResponseWriter, data interface{}, code int, isErr bool) {
	if isErr {
		json.NewEncoder(w).Encode(&response{
			Status:  code,
			Desc:    getStatus(code),
			IsError: isErr,
			Result:  data,
		})
		return
	} else {
		json.NewEncoder(w).Encode(&response{
			Status:  code,
			Desc:    getStatus(code),
			IsError: isErr,
			Result:  data,
		})
	}
}

func getStatus(status int) string {
	var desc string
	switch status {
	case 200:
		desc = "OK"
	case 201:
		desc = "Created"
	case 400:
		desc = "Bad Request"
	case 401:
		desc = "Unauthorized"
	case 500:
		desc = "Internal Server Error"
	case 501:
		desc = "Bad Gateway"
	case 304:
		desc = "Not Modified"
	default:
		desc = ""
	}

	return desc
}
