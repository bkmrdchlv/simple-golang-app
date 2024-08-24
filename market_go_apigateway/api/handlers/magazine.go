package handlers

import (
	"market_go_apigateway/api/http"
	"market_go_apigateway/genproto/magazine"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateMagazine(c *gin.Context) {
	var item magazine.MagazineCreate
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.MagazineService().Create(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) GetAllMagazine(c *gin.Context) {
	resp, err := h.services.MagazineService().GetAll(c.Request.Context(), &magazine.GetListRequest{
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

func (h Handler) DeleteMagazine(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.MagazineService().Delete(c.Request.Context(), &magazine.Pkey{Id: id})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) GetByIdMagazine(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.MagazineService().GetById(c.Request.Context(), &magazine.Pkey{Id: id})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) UpdateMagazine(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.MagazineService().Update(c.Request.Context(), &magazine.Magazine{
		Id: id,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
