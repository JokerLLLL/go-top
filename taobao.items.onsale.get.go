package top

type TaobaoItemsOnsaleGet struct {
	param map[string]interface{}
}

func NewTaobaoItemsOnsaleGet() *TaobaoItemsOnsaleGet {
	return &TaobaoItemsOnsaleGet{
		param: map[string]interface{}{
			"fields":"num_iid,num,outer_id,title,pic_url,price",
			"page_no":1,
			"page_size":5,
		}}

}

func (t *TaobaoItemsOnsaleGet) ApiName() (s string) {
	return "taobao.items.onsale.get"
}

func (t *TaobaoItemsOnsaleGet) SetParam(k string, v interface{}) {
	t.param[k] = v
}

func (t *TaobaoItemsOnsaleGet) getParam() map[string]interface{} {
	return t.param
}

func (t *TaobaoItemsOnsaleGet) CheckParam() (msg string, ok bool) {
	if _, ok := t.param["fields"]; !ok {
		return "Missing required arguments:fields", false
	}
	if _, ok := t.param["page_no"]; !ok {
		return "Missing required arguments:page_no", false
	}
	if _, ok := t.param["page_size"]; !ok {
		return "Missing required arguments:page_size", false
	}

	return "", true
}
