package boot

import (
	"bufio"
	"os"
	"strings"
)

func GetTargets(targetString string) []string {
	/*
		获取目标列表
	*/
	var targetList []string
	if strings.HasSuffix(targetString, ".txt") {

		file, err := os.Open(targetString)
		if err != nil {
			println("can not open file " + targetString)
			os.Exit(1)
		}

		defer func() { _ = file.Close() }()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lineText := scanner.Text()
			targetList = append(targetList, lineText)
		}

	} else {
		targetList = append(targetList, targetString)

	}

	return targetList

}
