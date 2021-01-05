package cache

import (
	"fmt"
	"go-crawler-distributed/global"
	"go-crawler-distributed/pkg/setting"
	"strings"
	"testing"
)

/**
* @Author: super
* @Date: 2020-12-29 11:31
* @Description:
**/

func TestAddElementToSet(t *testing.T) {
	newSetting, err := setting.NewSetting(strings.Split("/Users/super/develop/superTools-frontground-backend/configs", ",")...)
	if err != nil {
		t.Error(err)
	}
	err = newSetting.ReadSection("Cache", &global.CacheSetting)
	if err != nil {
		t.Error(err)
	}
	global.RedisEngine, err = NewRedisEngine(global.CacheSetting)
	if err != nil {
		t.Error(err)
	}
	result, err := AddElementToSet("hello", "1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

func BenchmarkAddElementToSet(b *testing.B) {
	newSetting, err := setting.NewSetting(strings.Split("/Users/super/develop/superTools-frontground-backend/configs", ",")...)
	if err != nil {
		b.Error(err)
	}
	err = newSetting.ReadSection("Cache", &global.CacheSetting)
	if err != nil {
		b.Error(err)
	}
	global.RedisEngine, err = NewRedisEngine(global.CacheSetting)
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		_, err := AddElementToSet("hello", "1")
		if err != nil {
			b.Error(err)
		}
	}
}
