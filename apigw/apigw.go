package apigw

type RouterHandler interface {
	Handle()
	GetPrefix() string
}
