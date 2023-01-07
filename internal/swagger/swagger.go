package swagger

import (
	"embed"
	"github.com/fastid/fastid/internal/logger"
	"github.com/labstack/echo/v4"
	"io/fs"
	"net/http"
)

import (
	"github.com/fastid/fastid/internal/config"
)

//go:embed assert/*
var embededFiles embed.FS

type Swagger interface {
	Register(e *echo.Group)
	getFileSystem() http.FileSystem
}

type swagger struct {
	cfg    *config.Config
	logger logger.Logger
}

func New(cfg *config.Config, logger logger.Logger) Swagger {
	return &swagger{cfg: cfg, logger: logger}
}

func (s *swagger) getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(embededFiles, "assert")
	if err != nil {
		panic(err)
	}
	return http.FS(fsys)
}

func (s *swagger) Register(e *echo.Group) {
	assetHandler := http.FileServer(s.getFileSystem())
	e.GET("/*", echo.WrapHandler(http.StripPrefix("/api/", assetHandler)))
}
