package handlers

import (
	"market_go_apigateway/api/http"
	"market_go_apigateway/genproto/change"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateChange(c *gin.Context) {
	var item change.ChangeCreate
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ChangeService().Create(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
func (h Handler) GetAllChange(c *gin.Context) {
	resp, err := h.services.ChangeService().GetAll(c.Request.Context(), &change.GetListRequest{
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
func (h Handler) DeleteChange(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.ChangeService().Delete(c.Request.Context(), &change.Pkey{Id: id})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
func (h Handler) GetByIdChange(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.ChangeService().GetById(c.Request.Context(), &change.Pkey{Id: id})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
func (h Handler) UpdateChange(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.ChangeService().Update(c.Request.Context(), &change.Change{
		Id: id,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
func (h Handler) OpenChange(c *gin.Context) {
	var item change.Pkey
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ChangeService().OpenChange(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
func (h Handler) CloseChange(c *gin.Context) {
	var item change.Pkey
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ChangeService().CloseChange(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
