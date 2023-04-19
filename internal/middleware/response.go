package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

const HTTP_CODE_MD = "x-http-code"

// HTTPResponseModifier checks for the HTTP_CODE_MD metadata pair and explicitly sets the status code for the response if found.
func HTTPResponseModifier(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	// set http status code
	if vals := md.HeaderMD.Get(HTTP_CODE_MD); len(vals) > 0 {
		code, err := strconv.Atoi(vals[0])
		if err != nil {
			return err
		}
		// delete the headers to not expose any grpc-metadata in http response
		delete(md.HeaderMD, HTTP_CODE_MD)
		delete(w.Header(), "Grpc-Metadata-X-Http-Code")
		w.WriteHeader(code)
	}

	return nil
}

// HTTPStatusCode determines the status code from the request method.
func HTTPStatusCode(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	if vals := md.HeaderMD.Get(HTTP_CODE_MD); len(vals) == 0 {
		code := http.StatusOK
		if method, ok := httpMethod(ctx); ok {
			switch method {
			case "POST":
				code = http.StatusCreated
			case "PATCH", "DELETE":
				code = http.StatusNoContent
			}
			md.HeaderMD.Append(HTTP_CODE_MD, strconv.Itoa(code))
		}
	}

	return nil
}

func httpMethod(ctx context.Context) (string, bool) {
	m := ctx.Value(chi.RouteCtxKey)
	if m == nil {
		return "", false
	}
	chiCtx := chi.RouteContext(ctx)
	return chiCtx.RouteMethod, true
}

// SetHTTPStatusCode sets status code.
func SetHTTPStatusCode(ctx context.Context, code int) {
	_ = grpc.SetHeader(ctx, metadata.Pairs(HTTP_CODE_MD, strconv.Itoa(code)))
}
