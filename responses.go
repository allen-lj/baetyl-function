package main

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

const (
	jsonContentTypeHeader = "application/json"
)

// ErrorResponse is an HTTP response message sent back to calling clients by the Dapr Runtime HTTP API
type ErrorResponse struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

// NewErrorResponse returns a new ErrorResponse
func NewErrorResponse(errorCode, message string) ErrorResponse {
	return ErrorResponse{
		ErrorCode: errorCode,
		Message:   message,
	}
}

// respondWithJSON overrides the content-type with application/json
func respondWithJSON(ctx *fasthttp.RequestCtx, code int, obj []byte) {
	respond(ctx, code, obj)
	ctx.Response.Header.SetContentType(jsonContentTypeHeader)
}

// respond sets a default application/json content type if content type is not present
func respond(ctx *fasthttp.RequestCtx, code int, obj []byte) {
	ctx.Response.SetStatusCode(code)
	ctx.Response.SetBody(obj)

	if len(ctx.Response.Header.ContentType()) == 0 {
		ctx.Response.Header.SetContentType(jsonContentTypeHeader)
	}
}

func respondWithError(ctx *fasthttp.RequestCtx, code int, resp ErrorResponse) {
	b, _ := json.Marshal(&resp)
	respondWithJSON(ctx, code, b)
}

func respondEmpty(ctx *fasthttp.RequestCtx, code int) {
	ctx.Response.SetBody(nil)
	ctx.Response.SetStatusCode(code)
}
