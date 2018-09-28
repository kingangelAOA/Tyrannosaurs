package samples

import (
	"testing"
	"encoding/json"
	"fmt"
)

func TestHttp(t *testing.T) {
	j := "{\"name\":\"aaa\",\"annotation\":null,\"protocol\":\"http\",\"host\":\"www.baidu.com\",\"port\":null,\"method\":\"get\",\"path\":\"/\",\"codeing\":\"utf-8\",\"consuming\":-1,\"error\":null,\"Header\":{\"aa\":\"bb\",\"cc\":\"dd\"},\"body\":\"aaaa\",\"form\":{\"ee\":\"hh\",\"jj\":\"gg\"}}"
	var http Http
	json.Unmarshal([]byte(j), &http)
	fmt.Print(http.Consuming)
}