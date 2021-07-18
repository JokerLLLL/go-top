package top

type TaobaoTopAoidDecrypt struct {
	param map[string]interface{}
}

func NewTaobaoTopAoidDecrypt() *TaobaoTopAoidDecrypt {
	return &TaobaoTopAoidDecrypt{
		param: map[string]interface{}{},
	}
}

func (t *TaobaoTopAoidDecrypt) ApiName() (s string) {
	return "taobao.top.oaid.decrypt"
}

func (t *TaobaoTopAoidDecrypt) SetParam(k string, v interface{}) {
	t.param[k] = v
}

func (t *TaobaoTopAoidDecrypt) getParam() map[string]interface{} {
	return t.param
}

func (t *TaobaoTopAoidDecrypt) CheckParam() (msg string, ok bool) {
	if _, ok := t.param["query_list"]; !ok {
		return "Missing required arguments:query_list", false
	}
	return "", true
}
