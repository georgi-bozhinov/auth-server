package users

import (
	"github.com/georgi-bozhinov/auth-server/server/router"
	"net/http"
)

type loginDTO struct {
	Username string
	Password string
}

func (c *Controller) Routes() router.Routes {
	return []router.Route{
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: c.GetUsers,
			Name:    "GET /users",
		},
		{
			Method:  http.MethodPost,
			Path:    "/users/register",
			Handler: c.Register,
			Name:    "POST /users/register",
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{username}",
			Handler: c.GetUserByUsername,
			Name:    "GET /users/{username}",
		},
		{
			Method:  http.MethodPost,
			Path:    "/users/login",
			Handler: c.Login,
			Name:    "POST /users/login",
		},
	}
}
