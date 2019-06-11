package router

import (
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	Handler http.HandlerFunc
	Path    string
}

type Routes []Route
