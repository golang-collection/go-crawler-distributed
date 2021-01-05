package idGenerator

import "testing"

/**
* @Author: super
* @Date: 2020-11-24 13:57
* @Description:
**/

func TestGenerateSnowflake(t *testing.T) {
	err := InitSnowflake()
	if err != nil {
		t.Error(err)
	}
	id := GenerateID()
	t.Log(id)
}
