package handlers

import (
	"market_go_apigateway/api/http"
	"market_go_apigateway/genproto/branch"
	"market_go_apigateway/genproto/product"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateProduct(c *gin.Context) {
	var item product.ProductCreate
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ProductService().Create(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) GetAllProduct(c *gin.Context) {
	resp, err := h.services.BranchService().GetAll(c.Request.Context(), &branch.GetListRequest{
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

func (h Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.ProductService().Delete(c.Request.Context(), &product.Pkey{Id: id})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) GetByIdProduct(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.ProductService().GetById(c.Request.Context(), &product.Pkey{Id: id})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.ProductService().Update(c.Request.Context(), &product.Product{
		Id: id,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
