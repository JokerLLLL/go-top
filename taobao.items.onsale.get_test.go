package top_test

import (
	"fmt"
	"testing"

	"github.com/liuhengloveyou/go-top"
)

func TestTaobaoProductsGet(t *testing.T) {
	var topClient *top.TopClient = top.CreateTopClient("", "", "")

	var api = top.NewTaobaoItemsOnsaleGet()

	d, err := topClient.Run(api)

	fmt.Println(">>>>>>", string(d), err)

}
