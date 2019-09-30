// Code generated by goa v3.0.6, DO NOT EDIT.
//
// tasks HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wild-mouse/go-example-todo-application/design

package server

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// EncodeCountTasksResponse returns an encoder for responses returned by the
// tasks count_tasks endpoint.
func EncodeCountTasksResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(int)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}
