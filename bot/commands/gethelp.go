package commands

import (
	"bytes"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func Get_Help(url string) ([]byte, string) {
	var msg []byte
	var str string

	body := new(bytes.Buffer)
	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		zap.L().Error("Error", zap.Error(err))
	} else {
		if resp.StatusCode < 400 {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				zap.L().Error("Response Read Error", zap.Error(err))
			} else {
				msg := body
				str = ""
				return msg, str
			}
		} else {
			str = "Help Endpoint Not found Either! Report This to An Admin"
			return msg, str
		}
	}
	return msg, str
}
