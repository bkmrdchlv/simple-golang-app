package handlers

import (
	"market_go_apigateway/api/http"
	"market_go_apigateway/config"
	"market_go_apigateway/grpc/client"
	"market_go_apigateway/pkg/logger"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	cfg      config.Config
	log      logger.LoggerI
	services client.ServiceManagerI
}

func NewHandler(cfg config.Config, log logger.LoggerI, svcs client.ServiceManagerI) Handler {
	return Handler{
		cfg:      cfg,
		log:      log,
		services: svcs,
	}
}

func (h *Handler) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			// logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

func (h *Handler) getOffsetParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("offset", h.cfg.DefaultOffset)
	return strconv.Atoi(offsetStr)
}

func (h *Handler) getLimitParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)
	return strconv.Atoi(offsetStr)
}

func (h *Handler) getPageParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("page", "1")
	return strconv.Atoi(offsetStr)
}

func (h *Handler) IsUUID(c *gin.Context, id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.BadRequest.Code, gin.H{
			"error": err.Error(),
		})
		return err
	}
	return nil
}

func (h *Handler) ParseUUID(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return nil
}
