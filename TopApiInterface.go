package top

type TopApi interface {
	ApiName() string
	getParam() map[string]interface{}
	SetParam(k string, v interface{})
	CheckParam() (msg string, ok bool)
}
