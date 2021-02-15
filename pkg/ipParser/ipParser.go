package ipParser

import (
	"errors"
	"github.com/kayon/iploc"
	"go-crawler-distributed/global"
	"go-crawler-distributed/pkg/setting"
)
/**
* @Author: super
* @Date: 2021-02-15 16:57
* @Description:
**/

func NewIpParser(ipParserSetting *setting.IpParserSettingS)(locator *iploc.Locator, err error){
	locator, err = iploc.Open(ipParserSetting.FilePath)
	if err != nil {
		return
	}
	return
}

// 查询IP所属位置
func GetIpLocationString(ip string) (string, error ){
	detail := global.IpParser.Find(ip)
	if detail != nil {
		return detail.String(), nil
	}else{
		return "", errors.New("can't find ip location")
	}
}