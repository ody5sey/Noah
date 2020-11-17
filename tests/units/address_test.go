package units

import (
	"Noah/pkg/utils"
	"testing"
)

func TestAddress(t *testing.T) {

	ips := []string{"192.168.1.1/24"}
	for _, v := range ips {
		info := utils.AnalyseAddress(v)
		for i := range info {
			println(i)
		}
	}
}
