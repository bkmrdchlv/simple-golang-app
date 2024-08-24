package handlers

import (
	"market_go_apigateway/api/http"
	"market_go_apigateway/genproto/category"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateCategory(c *gin.Context) {

	var item category.CategoryCreate
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.CategoryService().Create(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) GetAllCategory(c *gin.Context) {
	resp, err := h.services.CategoryService().GetAll(c.Request.Context(), &category.GetListRequest{
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

func (h Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.CategoryService().Delete(c.Request.Context(), &category.Pkey{Id: id})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) GetByIdCategory(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.CategoryService().GetById(c.Request.Context(), &category.Pkey{Id: id})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.CategoryService().Update(c.Request.Context(), &category.Category{
		Id: id,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
