package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

const MinBody = 10
const MaxBody = 200

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomAlphanumeric(length int) []byte {
	b := make([]byte, length)
	charsetSize := len(charset)
	for i := range b {
		b[i] = charset[rand.Intn(charsetSize)]
	}
	return b
}

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // default fallback
	}
	// Create a TCP listener manually
	addr := ":" + port
	log.Printf("listening addr = %d", addr)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to listen: %v\n", err)
		os.Exit(1)
	}

	// Wrap it so we can set TCP-level options
	tcpListener := ln.(*net.TCPListener)

	// Set TCP_NODELAY manually for every new connection
	listener := &tcpKeepAliveListener{TCPListener: tcpListener}

	server := &fasthttp.Server{
		Handler: requestHandler,
		// Optional: set any fasthttp tuning params here
	}

	// Serve using our tuned listener
	if err := server.Serve(listener); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
	}
}

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	conn, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}

	// Turn off Nagle’s algorithm (TCP_NODELAY)
	conn.SetNoDelay(true)

	// Keep connection alive
	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(3e9) // 3 seconds

	// You can even get lower-level if you want to do syscall.SetsockoptInt

	return conn, nil
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	nano := time.Now().UnixNano()
	var str []byte
	body := generateRandomAlphanumeric(rand.Intn((MaxBody - MinBody + 1) + MinBody))
	if str = ctx.Request.Header.Peek("X-Arrived-time"); str != nil {
		ctx.Response.Header.Set("X-Arrived-time", string(str))
	}
	if str = ctx.Request.Header.Peek("X-Start-Time"); str != nil {
		ctx.Response.Header.Set("X-Start-Time", string(str))
	}
	ctx.Response.Header.Set("X-App-time", strconv.FormatInt(nano, 10))
	ctx.Response.Header.Set("Content-Type", "application/text")
	ctx.Response.Header.Set("Server", "OUR.TEST.SERVER/28")
	ctx.SetBodyString(string(body))
	log.Printf("take %d", (time.Now().UnixNano()-nano)/1000000)
}
