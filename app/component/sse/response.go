package sse

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Sse SSE /**
type Sse struct {
	ResponseWriter gin.ResponseWriter
}

func New(writer gin.ResponseWriter) *Sse {
	return &Sse{ResponseWriter: writer}
}

// Send 推送/**
func (sse *Sse) Send(data string) (len int, err error) {
	flusher, ok := sse.ResponseWriter.(http.Flusher)
	if !ok {
		log.Panic("server not support")
	}

	len, err = sse.ResponseWriter.Write([]byte("data: " + data + "\n\n"))

	flusher.Flush()

	return len, err
}

// Close 关闭 /**
func (sse *Sse) Close(data string) (len int, err error) {
	flusher, ok := sse.ResponseWriter.(http.Flusher)
	if !ok {
		log.Panic("server not support")
	}
	len, err = sse.ResponseWriter.Write([]byte("event: close\ndata:" + data + " close\n\n"))
	flusher.Flush()

	return len, err
}
