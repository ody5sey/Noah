package net

import (
	"Noah/configs"
	"net/http"
)

func GenerateGet(target string) (int, http.Header, []byte, error) {

	build := NewClientBuilder().SetUserAgent(configs.UserAgent).SetHeader(configs.Headers).SetTimeOut(configs.TimeOut).SetSkipVerify(true).Builder()
	return build.Get(target)

}
