package ip

/**
* @Author: super
* @Date: 2020-12-18 17:24
* @Description:
**/

import (
	"errors"
	"net"
)

func GetIntranceIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		//检查Ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("获取地址异常")
}
