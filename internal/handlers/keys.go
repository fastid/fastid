package handlers

import (
	"github.com/fastid/fastid/internal/config"
	"github.com/fastid/fastid/internal/services"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type KeyHandler interface {
	Register(router *echo.Group)
	post() echo.HandlerFunc
}

type keyHandler struct {
	cfg *config.Config
	log *log.Logger
	srv services.Services
}

func NewKeyHandler(cfg *config.Config, log *log.Logger, srv services.Services) KeyHandler {
	return &keyHandler{cfg: cfg, log: log, srv: srv}
}

func (h *keyHandler) Register(router *echo.Group) {
	router.Add("POST", "/keys/", h.post())
}

func (h *keyHandler) post() echo.HandlerFunc {

	type Request struct {
		Count int `json:"count" validate:"required,min=1,max=10"`
	}

	return func(e echo.Context) error {
		//u := new(Request)
		//if err := e.Bind(u); err != nil {
		//	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		//}
		//
		//if err := e.Validate(u); err != nil {
		//	var errs []Errors
		//	var errMessage string
		//
		//	for _, err := range err.(validator.ValidationErrors) {
		//		if err.Field() == "Count" && err.Tag() == "required" {
		//			errMessage = "Field not filled \"count\""
		//		}
		//
		//		errs = append(errs, Errors{
		//			errMessage,
		//			err.Field(),
		//			err.Tag(),
		//			err.ActualTag(),
		//		})
		//	}
		//	return echo.NewHTTPError(http.StatusBadRequest, &Error{Message: "", Errors: errs})
		//}
		//
		//fmt.Println(u.Count)

		//cipher, err := h.srv.Keys().GenerateKey()
		//if err != nil {
		//	return err
		//}
		h.log.Traceln("Start")

		err := h.srv.Keys().Key(e.Request().Context())
		if err != nil {
			return err
		}

		json := make(map[string]string)
		json["key"] = "test"
		return e.JSON(http.StatusOK, json)
	}
}
