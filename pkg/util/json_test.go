package util

import (
	"testing"
)

/**
* @Author: super
* @Date: 2020-12-06 17:45
* @Description:
**/

type LoginUser struct {
	ID        string `json:"id"`
	UserName  string `json:"user_name"`
	IPAddress string `json:"ip_address"`
}

func TestEncodeToJson(t *testing.T) {
	loginUser := LoginUser{
		ID:        "123",
		UserName:  "username",
		IPAddress: "192.1.1.222",
	}
	result, err := EncodeToJson(loginUser)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func BenchmarkEncodeToJson(b *testing.B) {
	loginUser := LoginUser{
		ID:        "123",
		UserName:  "username",
		IPAddress: "192.1.1.222",
	}
	for i := 0; i < b.N; i++ {
		_, err := EncodeToJson(loginUser)
		if err != nil {
			b.Error(err)
		}
	}
}

func TestDecodeToStruct(t *testing.T) {
	input := `{"id":"123","user_name":"username","ip_address":"192.1.1.222"}`
	result, err := DecodeToStruct(input)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
