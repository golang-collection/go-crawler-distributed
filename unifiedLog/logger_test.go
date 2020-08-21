package unifiedLog

import "testing"

/**
* @Author: super
* @Date: 2020-08-21 09:20
* @Description:
**/

func TestGetLogger(t *testing.T) {
	logger := GetLogger()
	logger.Info("hello")
	logger.Error("hello")
}