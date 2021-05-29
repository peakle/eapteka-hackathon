package server

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/pprofhandler"

	"github.com/peakle/eapteka-miniapp/internal"
)

type Handler struct {
	manager *internal.SQLManager
}

func NewHandler() *Handler {
	return &Handler{
		manager: internal.InitManager(),
	}
}

func (handler *Handler) Start(ctx context.Context) error {

	var requestHandler = func(ctx *fasthttp.RequestCtx) {
		path := strings.ToLower(string(ctx.Path()))

		if strings.HasPrefix(path, "/v1/speech/recognize") && string(ctx.Request.Header.Method()) == fasthttp.MethodPost {
			handler.SpeechRecognize(ctx)
		} else if strings.HasPrefix(path, "/v1/text/recognize") && string(ctx.Request.Header.Method()) == fasthttp.MethodPost {
			handler.TextRecognize(ctx)
		} else if strings.HasPrefix(path, "/v1/speech/callback/schedule/create") && string(ctx.Request.Header.Method()) == fasthttp.MethodPut {
		} else if strings.HasPrefix(path, "/v1/speech/callback/schedule/add") && string(ctx.Request.Header.Method()) == fasthttp.MethodPatch {
		} else if strings.HasPrefix(path, "/v1/speech/callback/drugs/create") && string(ctx.Request.Header.Method()) == fasthttp.MethodPut {
		} else if strings.HasPrefix(path, "/debug/pprof") {
			pprofhandler.PprofHandler(ctx)
		} else {
			ctx.SetConnectionClose()
		}
	}

	var server = fasthttp.Server{
		Handler:            requestHandler,
		IdleTimeout:        1 * time.Minute,
		TCPKeepalive:       true,
		CloseOnShutdown:    true,
		MaxRequestBodySize: 60 * 1024 * 1024,
	}

	go func() {
		if err := server.ListenAndServe(":80"); err != nil {
			log.Printf("on start api server: %s", err)
		}
	}()

	<-ctx.Done()

	return server.Shutdown()
}

func (*Handler) SpeechRecognize(ctx *fasthttp.RequestCtx) {
	// TODO
}

func (*Handler) TextRecognize(ctx *fasthttp.RequestCtx) {
	// TODO
}

func (*Handler) CallbackScheduleCreate(ctx *fasthttp.RequestCtx) {
	// TODO
}

func (*Handler) CallbackScheduleAdd(ctx *fasthttp.RequestCtx) {
	// TODO
}

func (*Handler) CallbackDrugsCreate(ctx *fasthttp.RequestCtx) {
	// TODO
}
