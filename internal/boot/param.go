package boot

import (
	"Noah/api"
	"Noah/pkg/utils"
	"github.com/urfave/cli/v2"
	"os"
	"runtime"
	"strings"
)

func GetParam() ([]string, int) {
	/*
		获取cli的参数，返回由target组成的字符串和字典类型组成的字符串
	*/

	var targetString string
	var ports string
	var threads int

	cli.AppHelpTemplate = api.NewTemplate

	app := &cli.App{
		Name:    "Noah",
		Usage:   "A title collect",
		Version: api.Banner,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "targets",
				Usage:       "scan target (type: txt file or address string) `TARGET`",
				Destination: &targetString,
				//Required:    true,
			}, &cli.StringFlag{
				Name:        "ports",
				Usage:       "port, example: 80,8080",
				Destination: &ports,
				Value:       "80",
				//Required:    true,
			},
			&cli.IntFlag{
				Name:        "threads",
				Usage:       "threads count",
				Destination: &threads,
				Value:       runtime.NumCPU(),
				//Required:    true,
			},
		},
		Action: func(c *cli.Context) error {
			if len(os.Args) == 1 {
				if err := cli.ShowAppHelp(c); err != nil {
				}
				println("example: ")
				println("   noah --targets=http://127.0.0.1 --ports=80 --threads=12")
				println("   noah --targets=http://127.0.0.1")
				os.Exit(1)
			}

			if len(targetString) == 0 {
				println("You must enter target and threads")
				os.Exit(1)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		println(err.Error())
		os.Exit(1)
	}

	if len(targetString) == 0 {
		os.Exit(1)
	}

	var targetList []string

	targets := GetTargets(targetString)

	// 将ip进一步拆分，比如cidr
	for _, v := range targets {

		for _, k := range utils.AnalyseAddress(v) {
			for _, m := range strings.Split(ports, ",") {
				httpAddress := utils.ConvertAddress(k + ":" + m)
				targetList = append(targetList, httpAddress)
			}
		}

	}
	return targetList, threads
}
