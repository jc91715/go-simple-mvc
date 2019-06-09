package routes

import (
	"controller"
	"reflect"
	"regexp"
)

type Route struct {
	Regex          *regexp.Regexp
	Methods        map[string]string
	Params         map[int]string
	ControllerType reflect.Type
}

type app interface {
	AddRoute(pattern string, m map[string]string, c controller.ControllerInterface)
}

func AddRoutes(a app) {

	a.AddRoute("/", map[string]string{
		"GET": "Index",
	}, &controller.PostController{})
	a.AddRoute("/posts/:post_id([0-9]+)", map[string]string{
		"GET": "Show",
	}, &controller.PostController{})
}
