package delivery

import "github.com/labstack/echo/v4"

// Register routes for our api.
func (h *Handler) Register(e *echo.Echo) {
	e.GET("/recs/:id", h.Fetch)
	e.GET("/recs", h.FetchAll)
	e.POST("/rec", h.Update)
}
