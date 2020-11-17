package library

import (
	"Noah/internal/boot"
	"Noah/internal/net"
	"Noah/pkg/report"
	"Noah/pkg/utils"
	"strconv"
	"strings"
	"sync"
)

func Scanner() {
	// 获取参数
	targets, threads := boot.GetParam()

	var wg = &sync.WaitGroup{}
	var ThreadsChan chan struct{}
	//var length = len(targets)

	var threadSlice = NewSlice()

	// 使用强制模式，多线程并发
	ThreadsChan = make(chan struct{}, threads)

	for _, v := range targets {
		ThreadsChan <- struct{}{}
		wg.Add(1)

		go func(t string) {

			code, header, body, err := net.GenerateGet(t)
			if err != nil {
				wg.Done()
				<-ThreadsChan
				return
			}
			title := utils.MatchTitle(string(body))
			headers := strings.Join(header["Server"], "")
			mess := ResultPath{
				Code:    code,
				Address: t,
				Title:   title,
				Header:  headers,
			}

			threadSlice.Add(mess)

			wg.Done()
			<-ThreadsChan

		}(v)
	}

	wg.Wait()

	var mapResults []map[string]string

	for _, v := range threadSlice.Get() {
		var result = make(map[string]string)
		result["address"] = v.Address
		result["header"] = v.Header
		result["code"] = strconv.Itoa(v.Code)
		result["title"] = v.Title
		mapResults = append(mapResults, result)
	}

	report.ExportToJson(mapResults, "results.json")
}
