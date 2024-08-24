package handlers

import (
	"market_go_apigateway/api/http"
	"market_go_apigateway/genproto/employees"

	"github.com/gin-gonic/gin"
)

func (h Handler) CreateEmployees(c *gin.Context) {
	var item employees.EmployeesCreate
	err := c.ShouldBindJSON(&item)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.EmployeesService().Create(c.Request.Context(), &item)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) GetAllEmployees(c *gin.Context) {
	resp, err := h.services.EmployeesService().GetAll(c.Request.Context(), &employees.GetListRequest{
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

func (h Handler) DeleteEmployees(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.EmployeesService().Delete(c.Request.Context(), &employees.Pkey{Id: id})
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) GetByIdEmployees(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.EmployeesService().GetById(c.Request.Context(), &employees.Pkey{Id: id})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}

func (h Handler) UpdateEmployees(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.services.EmployeesService().Update(c.Request.Context(), &employees.Employees{
		Id: id,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)

}
