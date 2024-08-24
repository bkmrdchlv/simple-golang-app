package handlers

import (
	"market_go_apigateway/api/http"
	"market_go_apigateway/genproto/employees"

	"github.com/gin-gonic/gin"
)

func (h Handler) Register(c *gin.Context) {
	var item employees.Regist
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.EmployeesService().Register(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) Login(c *gin.Context) {
	var item employees.Loginer
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.EmployeesService().Login(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
