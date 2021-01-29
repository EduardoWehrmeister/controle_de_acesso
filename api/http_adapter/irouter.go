package http_adapter

import "net/http"

type IRouter interface {
	GET(uri string, function func(response http.ResponseWriter, request *http.Request))
	POST(uri string, function func(response http.ResponseWriter, request *http.Request))
	SERVE(port string)
}
