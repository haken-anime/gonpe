package pb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AdminHTTPConverter struct {
	srv AdminServer
}

func NewAdminHTTPConverter(srv AdminServer) *AdminHTTPConverter {
	return &AdminHTTPConverter{
		srv: srv,
	}
}

func (h *AdminHTTPConverter) GetAllUser(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error)) http.HandlerFunc {
	if cb == nil {
		cb = func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error) {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				p := status.New(codes.Unknown, err.Error()).Proto()
				switch r.Header.Get("Content-Type") {
				case "application/protobuf", "application/x-protobuf":
					buf, err := proto.Marshal(p)
					if err != nil {
						return
					}
					if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
						return
					}
				case "application/json":
					if err := json.NewEncoder(w).Encode(p); err != nil {
						return
					}
				default:
				}
			}
		}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			cb(ctx, w, r, nil, nil, err)
			return
		}

		arg := &GetAllUserReq{}

		contentType := r.Header.Get("Content-Type")
		switch contentType {
		case "application/protobuf", "application/x-protobuf":
			if err := proto.Unmarshal(body, arg); err != nil {
				cb(ctx, w, r, nil, nil, err)
				return
			}
		case "application/json":
			if err := jsonpb.Unmarshal(bytes.NewBuffer(body), arg); err != nil {
				cb(ctx, w, r, nil, nil, err)
				return
			}
		default:
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := fmt.Fprintf(w, "Unsupported Content-Type: %s", contentType)
			cb(ctx, w, r, nil, nil, err)
			return
		}

		ret, err := h.srv.GetAllUser(ctx, arg)
		if err != nil {
			cb(ctx, w, r, arg, nil, err)
			return
		}

		switch contentType {
		case "application/protobuf", "application/x-protobuf":
			buf, err := proto.Marshal(ret)
			if err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
			if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		case "application/json":
			if err := json.NewEncoder(w).Encode(ret); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		default:
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := fmt.Fprintf(w, "Unsupported Content-Type: %s", contentType)
			cb(ctx, w, r, arg, ret, err)
			return
		}
		cb(ctx, w, r, arg, ret, nil)
	})
}

func (h *AdminHTTPConverter) GetAllUserWithName(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error)) (string, string, http.HandlerFunc) {
	return "Admin", "GetAllUser", h.GetAllUser(cb)
}

func (h *AdminHTTPConverter) CreateGuestUser(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error)) http.HandlerFunc {
	if cb == nil {
		cb = func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error) {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				p := status.New(codes.Unknown, err.Error()).Proto()
				switch r.Header.Get("Content-Type") {
				case "application/protobuf", "application/x-protobuf":
					buf, err := proto.Marshal(p)
					if err != nil {
						return
					}
					if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
						return
					}
				case "application/json":
					if err := json.NewEncoder(w).Encode(p); err != nil {
						return
					}
				default:
				}
			}
		}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			cb(ctx, w, r, nil, nil, err)
			return
		}

		arg := &CreateGuestUserReq{}

		contentType := r.Header.Get("Content-Type")
		switch contentType {
		case "application/protobuf", "application/x-protobuf":
			if err := proto.Unmarshal(body, arg); err != nil {
				cb(ctx, w, r, nil, nil, err)
				return
			}
		case "application/json":
			if err := jsonpb.Unmarshal(bytes.NewBuffer(body), arg); err != nil {
				cb(ctx, w, r, nil, nil, err)
				return
			}
		default:
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := fmt.Fprintf(w, "Unsupported Content-Type: %s", contentType)
			cb(ctx, w, r, nil, nil, err)
			return
		}

		ret, err := h.srv.CreateGuestUser(ctx, arg)
		if err != nil {
			cb(ctx, w, r, arg, nil, err)
			return
		}

		switch contentType {
		case "application/protobuf", "application/x-protobuf":
			buf, err := proto.Marshal(ret)
			if err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
			if _, err := io.Copy(w, bytes.NewBuffer(buf)); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		case "application/json":
			if err := json.NewEncoder(w).Encode(ret); err != nil {
				cb(ctx, w, r, arg, ret, err)
				return
			}
		default:
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, err := fmt.Fprintf(w, "Unsupported Content-Type: %s", contentType)
			cb(ctx, w, r, arg, ret, err)
			return
		}
		cb(ctx, w, r, arg, ret, nil)
	})
}

func (h *AdminHTTPConverter) CreateGuestUserWithName(cb func(ctx context.Context, w http.ResponseWriter, r *http.Request, arg, ret proto.Message, err error)) (string, string, http.HandlerFunc) {
	return "Admin", "CreateGuestUser", h.CreateGuestUser(cb)
}