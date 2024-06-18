package main

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Context is used to add a timeout or cancellation to a go routine

func main() {
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {

		timeoutContext, cancel := context.WithTimeout(ctx.Request.Context(), time.Second*2)
		defer cancel()

		req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://yahoo.com", nil)

		if err != nil {
			panic(err)
		}

		res, err := http.DefaultClient.Do(req)

		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)

		if err != nil {
			panic(err)
		}

		ctx.Data(200, "text/html", data)
	})

	r.Run()
}
