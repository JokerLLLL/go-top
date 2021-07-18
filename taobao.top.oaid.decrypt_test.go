package top_test

import (
	"fmt"
	"testing"

	"github.com/liuhengloveyou/go-top"
)

func TestTaboTopOaidDecrypt(t *testing.T) {
	var topClient *top.TopClient = top.CreateTopClient("", "", "")

	var api = top.NewTaobaoTopAoidDecrypt()
	api.SetParam("query_list", []map[string]string{
		{"oaid":"xxx","tid":"111"},
		{"oaid":"xxx","tid":"111"},
		{"oaid":"xxx","tid":"111"},
	})

	d, err := topClient.Run(api)
	fmt.Println(">>>>>>", string(d), err)
}
