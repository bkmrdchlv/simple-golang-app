package handlers

import (
	"market_go_apigateway/api/http"
	"market_go_apigateway/genproto/provider"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateProvider(c *gin.Context) {
	var item provider.ProviderCreate
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ProvicderService().Create(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) GetAllProvider(c *gin.Context) {
	resp, err := h.services.ProvicderService().GetAll(c.Request.Context(), &provider.GetListRequest{
		Offset: 0,
		Limit:  10,
		Search: "",
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) DeleteProvider(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.ProvicderService().Delete(c.Request.Context(), &provider.Pkey{Id: id})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) GetByIdProvider(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.ProvicderService().GetById(c.Request.Context(), &provider.Pkey{Id: id})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) UpdateProvider(c *gin.Context) {

}
