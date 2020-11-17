package net

import (
	"Noah/pkg/utils"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type clientBuilder struct {
	//HTTP请求中的header信息
	header map[string]string

	//HTTP请求中,携带的cookies
	cookies []*http.Cookie

	//发起请求的client(go 自带的client)
	client *http.Client

	// 连接超时设置
	timeOut time.Duration

	//是否跳过HTTPS证书校验(默认跳过)
	skipVerify bool

	//http代理
	Proxy string
	// 用户代理
	UserAgent string
}

func (cb *clientBuilder) SetSkipVerify(skip bool) *clientBuilder {
	cb.skipVerify = skip
	return cb

}

// 设置超时
func (cb *clientBuilder) SetTimeOut(t time.Duration) *clientBuilder {
	cb.timeOut = t
	return cb
}

// 设置代理
func (cb *clientBuilder) SetProxy(u string) *clientBuilder {
	if len(u) > 0 {
		cb.Proxy = u
	} else {
		cb.Proxy = ""
	}

	return cb
}

// cookie设置
func (cb *clientBuilder) SetCookie(cookies string) *clientBuilder {

	if len(cookies) > 0 {
		if utils.MatchCookie(cookies) {
			nameList := strings.Split(cookies, "; ")

			var cookieSlice []*http.Cookie

			for _, c := range nameList {
				cookie := strings.Split(c, "=")
				newCookie := &http.Cookie{Name: cookie[0], Value: cookie[1]}
				cookieSlice = append(cookieSlice, newCookie)
			}
			cb.cookies = cookieSlice

		} else {
			cb.cookies = nil
		}
	} else {
		cb.cookies = nil
	}

	return cb
}

// user-agent设置
func (cb *clientBuilder) SetUserAgent(ua string) *clientBuilder {
	cb.UserAgent = ua
	return cb
}

// 设置http请求头文件
func (cb *clientBuilder) SetHeader(header map[string]string) *clientBuilder {

	cb.header = header
	return cb
}

func (cb *clientBuilder) Builder() *client {

	tlsConfig := &tls.Config{
		InsecureSkipVerify: cb.skipVerify,
	}

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   cb.timeOut * time.Second,
			KeepAlive: 10 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       10 * time.Second,
		TLSHandshakeTimeout:   3 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
	}

	if len(cb.Proxy) > 0 {

		if utils.MatchProxy(cb.Proxy) {
			proxy, err := url.Parse(cb.Proxy)
			if err == nil {
				transport.Proxy = http.ProxyURL(proxy)
			}

		}

	}

	c := &client{
		client: &http.Client{
			Transport: transport,
			Timeout:   cb.timeOut * time.Second,
		},
		header:    cb.header,
		cookies:   cb.cookies,
		userAgent: cb.UserAgent,
	}
	return c
}

//初始化 clientBuilder
func NewClientBuilder() *clientBuilder {
	return &clientBuilder{
		skipVerify: true,
	}
}
