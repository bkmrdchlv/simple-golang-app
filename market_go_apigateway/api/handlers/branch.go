package handlers

import (
	"market_go_apigateway/api/http"
	"market_go_apigateway/genproto/branch"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateBranch(c *gin.Context) {

	var item branch.BranchCreate
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.BranchService().Create(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) GetAllBranch(c *gin.Context) {
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

	resp, err := h.services.BranchService().GetAll(c.Request.Context(), &branch.GetListRequest{
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

func (h Handler) DeleteBranch(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.BranchService().Delete(c.Request.Context(), &branch.Pkey{Id: id})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) GetByIdBranch(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.BranchService().GetById(c.Request.Context(), &branch.Pkey{Id: id})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) UpdateBranch(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.BranchService().Update(c.Request.Context(), &branch.Branch{
		Id: id,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
