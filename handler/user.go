package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/serin0837/go-templ/model"
	"github.com/serin0837/go-templ/view/user"
)

type UserHandler struct {
}

func (h UserHandler) HandlerUserShow(c echo.Context) error {
	u := model.User{
		Email: "serin@serin.com",
	}
	return render(c, user.Show(u))
}
