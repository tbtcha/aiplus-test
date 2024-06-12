package controller

import (
	"aiplus/internal/app/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary	createStaff
// @Tags		staff
// @Accept		json
// @Produce	json
// @Param		RequestBody	body		domain.StaffRequest	true	"body"
// @Success	200		{object}	domain.StaffResponse				"Created"
// @Router		/staff [post]
func (c Controller) createStaff(ec echo.Context) error {
	req := new(domain.StaffRequest)
	err := ec.Bind(req)
	if err != nil {
		return err
	}
	response, err := c.s.CreateStaff(ec.Request().Context(), *req)
	if err != nil {
		return err

	}
	return ec.JSON(http.StatusCreated, response)
}
