package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/pprofhandler"

	"github.com/peakle/eapteka-miniapp/generated/client/client"
	"github.com/peakle/eapteka-miniapp/generated/client/client/operations"
	serverApi "github.com/peakle/eapteka-miniapp/generated/server/client/operations"
	serverModels "github.com/peakle/eapteka-miniapp/generated/server/models"
	"github.com/peakle/eapteka-miniapp/internal"
)

type Handler struct {
	manager   *internal.SQLManager
	config    *Config
	botClient operations.ClientService
}

func NewHandler() *Handler {
	return &Handler{
		manager: internal.InitManager(),
		config: &Config{
			ApiKey: "jgSAiwzYGRgVX2ei5eU03W9QIKSmNlab",
		},
	}
}

func (h *Handler) Start(ctx context.Context) error {
	h.botClient = client.Default.Operations

	var requestHandler = func(ctx *fasthttp.RequestCtx) {
		path := strings.ToLower(string(ctx.Path()))

		if strings.HasPrefix(path, "/v1/speech/recognize") && string(ctx.Request.Header.Method()) == fasthttp.MethodPost {
			h.SpeechRecognize(ctx)
		} else if strings.HasPrefix(path, "/v1/text/recognize") && string(ctx.Request.Header.Method()) == fasthttp.MethodPost {
			h.TextRecognize(ctx)
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
		if err := server.ListenAndServe(":8081"); err != nil {
			log.Printf("on start api server: %s", err)
		}
	}()

	<-ctx.Done()

	return server.Shutdown()
}

func (*Handler) SpeechRecognize(ctx *fasthttp.RequestCtx) {
	// TODO
}

func (h *Handler) TextRecognize(ctx *fasthttp.RequestCtx) {
	const handler = "TextRecognize"
	defer ctx.Response.Header.Set("Content-Type", "application/json")

	fail := func(err error) {
		log.Printf("[%s]: %s \n", handler, err)
		_, _ = fmt.Fprint(ctx, "{\"status\": \"failure\"}")
	}

	params := serverApi.TextRecognizeBody{}
	err := json.Unmarshal(ctx.PostBody(), &params)
	if err != nil {
		fail(err)
		return
	}

	if err = params.Validate(strfmt.Default); err != nil {
		fail(err)
		return
	}

	b := operations.BotTextRequestParams{
		Context: ctx,
		Body: operations.BotTextRequestBody{
			Key:   &h.config.ApiKey,
			Query: params.Text,
			Unit:  params.UserID,
		},
	}

	resp, err := h.botClient.BotTextRequest(&b)
	if err != nil {
		fail(err)
		return
	}

	res, _ := (&serverModels.RecognizeResponse{
		Status: "success",
		Text:   resp.Payload.Text,
		URI:    "",
	}).MarshalBinary()

	_, _ = fmt.Fprint(ctx, string(res))
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
