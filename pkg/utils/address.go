package utils

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)

func AnalyseAddress(ips string) []string {
	/*
		支持以下三种格式:
		192.168.1.1
		192.168.1.1-256
		192.168.1.1/24

	*/

	var results []string

	// 首先判断是否符合上面三个ip格式
	regStr := `^(([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){2}([1-9/-]{1,10})$`
	if match, _ := regexp.MatchString(regStr, ips); !match {
		results = append(results, ips)
		return results
	}

	if strings.Contains(ips, "-") {

		var ipRange []string

		ip := strings.Split(ips, ".")

		rangeIP := strings.Split(ip[3], "-")

		startNum, _ := strconv.Atoi(rangeIP[0])
		endNum, _ := strconv.Atoi(rangeIP[1])

		for i := startNum; i <= endNum; i++ {
			ipRange = append(ipRange, ip[0], ip[1], ip[2], strconv.Itoa(i))
			results = append(results, strings.Join(ipRange, "."))

		}

		return results

	} else if strings.Contains(ips, "/") {
		ip, ipNet, err := net.ParseCIDR(ips)
		if err != nil {
			return results
		}

		for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
			results = append(results, ip.String())
		}

		if results == nil {
			return []string{}
		}

		return results[1 : len(results)-1]

	} else {
		return append(results, ips)
	}

}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
