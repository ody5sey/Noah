package utils

import (
	"Noah/api"
	"bufio"
	"bytes"
	"encoding/json"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"regexp"
	"strings"
)

func ConvertAddress(target string) string {
	// 检查目标是否是 http://target的格式
	/*
		@param: http://127.0.0.1/ || 127.0.0.1
		@return: http://127.0.0.1
	*/
	NotLine := "^(http://|https://).*"
	match, _ := regexp.MatchString(NotLine, target)

	if !match {
		target = "http://" + target
	}

	return strings.TrimSuffix(target, "/")
}

// 探测网页编码
func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytesHtml, err := r.Peek(1024)

	if err != nil {
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytesHtml, "")

	return e

}

func CustomMarshal(message interface{}) (string, error) {
	/*
		自定义序列化函数，解决 "&"被转译的问题
	*/

	bf := bytes.NewBuffer([]byte{})

	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", "    ")

	if err := jsonEncoder.Encode(message); err != nil {
		return api.ErrorFlag, err
	}

	return bf.String(), nil
}
