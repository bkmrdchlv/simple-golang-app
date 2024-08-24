package handlers

import (
	"market_go_apigateway/api/http"
	"market_go_apigateway/genproto/coming"
	"market_go_apigateway/genproto/remainder"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateRemainder(c *gin.Context) {
	var item remainder.RemainderCreate
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.RemainderService().Create(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
func (h Handler) GetAllRemainder(c *gin.Context) {
	resp, err := h.services.RemainderService().GetAll(c.Request.Context(), &remainder.GetListRequest{
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
func (h Handler) DeleteRemainder(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.RemainderService().Delete(c.Request.Context(), &remainder.Pkey{Id: id})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
func (h Handler) GetByIdRemainder(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.RemainderService().GetById(c.Request.Context(), &remainder.Pkey{Id: id})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
func (h Handler) UpdateRemainder(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.RemainderService().Update(c.Request.Context(), &remainder.Remainder{
		Id: id,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) CreateComing(c *gin.Context) {
	var item coming.ComingCreate
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ComingService().Create(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
func (h Handler) GetAllComing(c *gin.Context) {
	resp, err := h.services.ComingService().GetAll(c.Request.Context(), &coming.GetListRequest{
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
func (h Handler) DeleteComing(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.ComingService().Delete(c.Request.Context(), &coming.Pkey{Id: id})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
func (h Handler) GetByIdComing(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.ComingService().GetById(c.Request.Context(), &coming.Pkey{Id: id})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
func (h Handler) UpdateComing(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.ComingService().Update(c.Request.Context(), &coming.Coming{
		Id: id,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) SdelatPrihod(c *gin.Context) {
	var item coming.Zavershit
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ComingService().SdelatPrihod(c.Request.Context(), &item)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) CreateComingProducts(c *gin.Context) {
	var item coming.ComingProductsCreate
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ComingProductsService().Create(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
func (h Handler) GetAllComingProducts(c *gin.Context) {
	resp, err := h.services.ComingProductsService().GetAll(c.Request.Context(), &coming.GetListRequest{
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
func (h Handler) DeleteComingProducts(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.ComingProductsService().Delete(c.Request.Context(), &coming.Pkey{Id: id})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
func (h Handler) GetByIdComingProducts(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.ComingProductsService().GetById(c.Request.Context(), &coming.Pkey{Id: id})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) UpdateComingProducts(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.ComingProductsService().Update(c.Request.Context(), &coming.ComingProducts{
		Id: id,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
