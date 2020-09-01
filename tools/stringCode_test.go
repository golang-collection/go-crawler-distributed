package tools

import (
	"fmt"
	"testing"
)

/**
* @Author: super
* @Date: 2020-09-01 20:07
* @Description:
**/

func TestZipString(t *testing.T) {
	s, err := ZipString([]byte("helloworldasdafsdfasfsdgadfgadfweaweterteggdfsgdsbdfbvxvczxvfasdfasdfasdfsadfsadfsadfsd"))
	if err != nil{
		t.Error(err)
	}
	fmt.Println(s)
}

func TestUnzipString(t *testing.T) {
	s := UnzipString("H4sIAAAAAAAA/0zKMQ6AIBQD0LMW+/sdSEgoAeLpjbo4vO2dUWtbrVfChEzBMhPUYwVWjOgjMiknXagy9zyuPYW3f4wf3gAAAP//AQAA//9lk09DVwAAAA==")
	fmt.Println(s)
}