package server

import (
	"errors"
	"io"
	"net/http"

	"github.com/no-src/nscache"
	_ "github.com/no-src/nscache/all"
	"github.com/no-src/nscache/proxy/request"
	"github.com/no-src/nscache/proxy/response"
	"github.com/no-src/nsgo/jsonutil"
)

// Start start the proxy server
func Start(addr string, conn string) error {
	cache, cacheErr := nscache.NewCache(conn)
	if cacheErr != nil {
		return cacheErr
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{key}", func(w http.ResponseWriter, req *http.Request) {
		key := req.PathValue("key")
		var v any
		err := cache.Get(key, &v)
		if errors.Is(err, nscache.ErrNil) {
			w.Write(response.NewNilErrorResponse().Bytes())
		} else if err != nil {
			w.Write(response.NewErrorResponse(err.Error()).Bytes())
		} else {
			w.Write(response.NewSuccessResponse(v).Bytes())
		}
	})

	mux.HandleFunc("PUT /{key}", func(w http.ResponseWriter, req *http.Request) {
		key := req.PathValue("key")
		data, err := io.ReadAll(req.Body)
		if err != nil {
			w.Write(response.NewErrorResponse(err.Error()).Bytes())
			return
		}
		var setReq request.SetRequest
		err = jsonutil.Unmarshal(data, &setReq)
		if err != nil {
			w.Write(response.NewErrorResponse(err.Error()).Bytes())
			return
		}

		err = cache.Set(key, setReq.Value, setReq.Expiration)
		if err != nil {
			w.Write(response.NewErrorResponse(err.Error()).Bytes())
		} else {
			w.Write(response.NewSuccessResponse("").Bytes())
		}
	})

	mux.HandleFunc("DELETE /{key}", func(w http.ResponseWriter, req *http.Request) {
		key := req.PathValue("key")
		err := cache.Remove(key)
		if err != nil {
			w.Write(response.NewErrorResponse(err.Error()).Bytes())
		} else {
			w.Write(response.NewSuccessResponse("").Bytes())
		}
	})
	return http.ListenAndServe(addr, mux)
}
