package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/serin0837/go-templ/view/user"
)

type UserHandler struct {
}

func (h UserHandler) HandlerUserShow(c echo.Context) error {
	return render(c, user.Show())
}
