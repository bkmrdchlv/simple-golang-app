package handlers

import (
	"fmt"
	"market_go_apigateway/api/http"
	"market_go_apigateway/genproto/sale"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetAllSale(c *gin.Context) {
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

	resp, err := h.services.SaleService().GetAll(c.Request.Context(), &sale.GetListRequest{
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

func (h Handler) GetById(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.SaleService().GetById(c.Request.Context(), &sale.Pkey{Id: id})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h Handler) DeleteSale(c *gin.Context) {

	var resp http.Empty

	id := c.Param("id")

	fmt.Println(id)

	_, err := h.services.SaleService().Delete(c.Request.Context(), &sale.Pkey{Id: id})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h Handler) CreateSale(c *gin.Context) {

	var item sale.SalectCreate
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleService().Create(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

func (h Handler) UpdateSale(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.SaleService().Update(c.Request.Context(), &sale.Sale{
		Id: id,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) CodeProd(c *gin.Context) {
	var item sale.SaleCode
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SaleProductService().Code(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
