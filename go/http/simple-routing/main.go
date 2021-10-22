package main

import (
	"fmt"
	"net/http"
)

type patternStr string

type methodStr string

type route func(writer http.ResponseWriter, request *http.Request)

var (
	router = map[patternStr]map[methodStr]route{
		"/hello": {
			http.MethodGet: getHello,
		},
	}
)

type handler struct {}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	res, ok := router[patternStr(request.RequestURI)]
	if !ok{
		writer.Write([]byte("not found"))
		writer.WriteHeader(http.StatusNotFound)
	} else if r, ok := res[methodStr(request.Method)]; ok {
		r(writer, request)
	} else {
		writer.Write([]byte("method not allowed"))
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	hd := &handler{}
	fmt.Println(http.ListenAndServe(":8000", hd))
}

func getHello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
}
