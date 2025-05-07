package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"math/rand"
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
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		nano := time.Now().UnixNano()
		var str string
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
		startTime, err := strconv.ParseInt(str, 10, 64) // base 10, 64-bit
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		log.Printf("take %d", (nano-startTime)/1000000)
	}

	if err := fasthttp.ListenAndServe(":8080", requestHandler); err != nil {
		panic(err)
	}
}
