package httpserver

import (
	"net/http"
	"os"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
)

const (
	docsPrefix = "/swagger-ui"
	apiPrefix  = "/api"
)

type Config struct {
	Addr string
}

type Docs interface {
	GetAsset() func(name string) ([]byte, error)
	GetAssetInfo() func(name string) (os.FileInfo, error)
	GetAssetDir() func(name string) ([]string, error)
}

func New(cfg Config, mux *runtime.ServeMux, docs Docs, logger *zap.Logger) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:     docs.GetAsset(),
		AssetDir:  docs.GetAssetDir(),
		AssetInfo: docs.GetAssetInfo(),
		Prefix:    "/",
	})
	r.HandleFunc(docsPrefix+"/*", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/swagger.json") {
			w.Header().Set("Cache-Control", "max-age=60, must-revalidate")
		}
		http.StripPrefix(docsPrefix+"/", fileServer).ServeHTTP(w, r)
	})

	r.HandleFunc(apiPrefix+"/*", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.Replace(r.URL.Path, apiPrefix, "", -1)
		mux.ServeHTTP(w, r)
	})

	server := &http.Server{Addr: cfg.Addr, Handler: r}
	go func(s *http.Server) {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}

	}(server)
	logger.Sugar().Infof("http server listening at %v", server.Addr)
	return server
}
