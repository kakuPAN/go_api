package handler

import (
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World\n")
}

func postArticleHundler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

func ListArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List\n")
}
func ShowArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article No.1\n")
}
func NiceArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}
func CommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
