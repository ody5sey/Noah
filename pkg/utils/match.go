package utils

import (
	"regexp"
)

/*
	正则校验socks5/http/https代理
*/
func MatchProxy(address string) bool {

	matched, err := regexp.MatchString(`^(https?|socks5)://.*?:[0-9]{1,5}$`, address)
	if err != nil {
		return false
	}
	if matched {
		return true

	} else {
		return false
	}

}

/*
	匹配cookie的规则，比如: username=admin; userid=1; PHPSESSID=9d1q9o4927a42p2thki1ql82p7
*/
func MatchCookie(cookies string) bool {
	matched, err := regexp.MatchString(`^([\w]*?=[\w]*?; )+([\w]*?=[\w]*?)$`, cookies)
	if err != nil {
		return false
	}
	if matched {
		return true

	} else {
		return false
	}
}

/*
正则提取网站标题
*/
func MatchTitle(html string) string {
	matched := regexp.MustCompile(`<title>([\S\s]*?)</title>`)
	results := matched.FindStringSubmatch(html)

	if len(results) > 1 {
		return results[1]
	}

	return " "
}
