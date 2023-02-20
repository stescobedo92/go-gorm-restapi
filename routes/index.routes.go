package routes

import "net/http"

func HomeHandler(writer http.ResponseWriter, rquest *http.Request) {
	writer.Write([]byte("Hello world"))
}
