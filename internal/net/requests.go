package net

import (
	"Noah/api"
	"Noah/pkg/utils"
	"bufio"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

type client struct {

	//HTTP请求中的header信息
	header map[string]string

	//HTTP请求中,携带的cookies
	cookies []*http.Cookie

	//发起请求的client(go 自带的client)
	client *http.Client

	//设置UserAgent
	userAgent string
}

//初始化一个 http.Request, 并填充属性
func (c *client) getRequest(method, url string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range c.header {
		request.Header.Set(k, v)
	}

	for _, v := range c.cookies {
		request.AddCookie(v)
	}

	if len(c.userAgent) > 0 {
		request.Header["User-Agent"] = []string{c.userAgent}
	}

	return request, nil
}

func (c *client) Get(address string) (int, http.Header, []byte, error) {
	request, err := c.getRequest(http.MethodGet, address, nil)
	if err != nil {
		return api.ErrorCode, nil, nil, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return api.ErrorCode, nil, nil, err
	}

	defer func() { _ = response.Body.Close() }()

	//  开始探测网页编码
	bodyReader := bufio.NewReader(response.Body)
	e := utils.DetermineEncoding(bodyReader)

	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	body, err := ioutil.ReadAll(utf8Reader)
	return response.StatusCode, response.Header, body, err

}
