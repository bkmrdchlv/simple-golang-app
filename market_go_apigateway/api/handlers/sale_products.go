package handlers

import (
	"fmt"
	"market_go_apigateway/api/http"
	"market_go_apigateway/genproto/sale"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetAllSaleProducts(c *gin.Context) {
	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.SaleProductService().GetAll(c.Request.Context(), &sale.GetListRequest{
		Offset: int64(page),
		Limit:  int64(limit),
		Search: "",
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) CreateSaleProducts(c *gin.Context) {
	var item sale.SaleProductsCreate
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleProductService().Create(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) DeleteSaleProducts(c *gin.Context) {

	id := c.Param("id")

	fmt.Println(id)

	resp, err := h.services.SaleProductService().Delete(c.Request.Context(), &sale.Pkey{Id: id})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) GetByIdSaleProducts(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.SaleProductService().GetById(c.Request.Context(), &sale.Pkey{Id: id})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) UpdateSaleProducts(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.SaleProductService().Update(c.Request.Context(), &sale.SaleProducts{
		Id: id,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h Handler) Paytype(c *gin.Context) {

	var item sale.PayTypes
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleProductService().PayType(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h Handler) Prodaja(c *gin.Context) {
	var item sale.Sold
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleProductService().Prodaja(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
