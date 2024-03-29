package datacenterhttphandler

import (
	"RetroPGF-Hub/RetroPGF-Hub-Backend-Go/modules/datacenter"
	datacenterusecase "RetroPGF-Hub/RetroPGF-Hub-Backend-Go/modules/datacenter/datacenterUsecase"
	"RetroPGF-Hub/RetroPGF-Hub-Backend-Go/pkg/request"
	"RetroPGF-Hub/RetroPGF-Hub-Backend-Go/pkg/response"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	DatacenterHttpHandlerService interface {
		InsertUrlCache(c echo.Context) error
		DeleteUrlCache(c echo.Context) error
		FindManyUrlCache(c echo.Context) error
		FindCacheData(c echo.Context) error
		CronJobUpdateCache() error
		TriggerUpdateCache(c echo.Context) error
	}

	datacenterHttpHandler struct {
		datacenterUsecase datacenterusecase.DatacenterUsecaseService
	}
)

func NewDatacenterHttpHandler(datacenterUsecase datacenterusecase.DatacenterUsecaseService) DatacenterHttpHandlerService {
	return &datacenterHttpHandler{
		datacenterUsecase: datacenterUsecase,
	}
}

func (h *datacenterHttpHandler) InsertUrlCache(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	req := new(datacenter.CacheModel)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.datacenterUsecase.InsertUrlCache(ctx, req.Url)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, map[string]string{
		"msg":   "insert url success",
		"urlId": id,
	})
}

func (h *datacenterHttpHandler) DeleteUrlCache(c echo.Context) error {
	ctx := context.Background()

	urlId := c.Param("urlId")
	if len(urlId) < 5 {
		return response.ErrResponse(c, http.StatusBadRequest, "urlId is required")
	}

	if err := h.datacenterUsecase.DeleteUrlCache(ctx, urlId); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, map[string]string{
		"msg": "delete url success",
	})
}

func (h *datacenterHttpHandler) FindManyUrlCache(c echo.Context) error {
	ctx := context.Background()

	data, err := h.datacenterUsecase.FindManyUrlsCache(ctx)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, map[string]any{
		"msg":  "ok",
		"urls": data,
	})

}

func (h *datacenterHttpHandler) FindCacheData(c echo.Context) error {
	ctx := context.Background()
	cacheId := c.Param("cacheId")
	if len(cacheId) < 5 {
		return response.ErrResponse(c, http.StatusBadRequest, "cacheId is required")
	}

	data, err := h.datacenterUsecase.FindCacheData(ctx, cacheId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, map[string]any{
		"msg":       "ok",
		"cacheData": data,
	})
}

func (h *datacenterHttpHandler) TriggerUpdateCache(c echo.Context) error {
	ctx := context.Background()
	cacheId := c.Param("cacheId")
	if len(cacheId) < 5 {
		return response.ErrResponse(c, http.StatusBadRequest, "cacheId is required")
	}

	data, err := h.datacenterUsecase.TriggerUpdateCache(ctx, cacheId)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}
	return response.SuccessResponse(c, http.StatusOK, map[string]any{
		"msg":       "ok",
		"cacheData": data,
	})
}

func (h *datacenterHttpHandler) CronJobUpdateCache() error {
	ctx := context.Background()
	return h.datacenterUsecase.CronJobUpdateCache(ctx)
}
