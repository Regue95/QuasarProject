package http

// RouterInterface is an interface
type RouterInterface interface {
	GRouter
}

// GRouter is an interface
type GRouter interface {
	GET(url string, i interface{})
	POST(url string, i interface{})
	SERVE(port string)
}
