package report

import (
	"Noah/pkg/utils"
	"io/ioutil"
)

func ExportToJson(mapResults []map[string]string, savePath string) {

	// 最后面4个空格，让json格式更美观
	result, err := utils.CustomMarshal(mapResults)

	if err != nil {
		return
	}

	if err := ioutil.WriteFile(savePath, []byte(result), 0644); err != nil {
		return
	}
}
