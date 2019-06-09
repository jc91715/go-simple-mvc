package app

import (
	"contex"
	"controller"

	// "fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"routes"
	"strings"
)

type App struct {
	http.Handler
	routes []*routes.Route
}

func (a *App) AddRoutes() {
	routes.AddRoutes(a)
}
func (a *App) AddRoute(pattern string, m map[string]string, c controller.ControllerInterface) {
	parts := strings.Split(pattern, "/")

	j := 0
	params := make(map[int]string)
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			expr := "([^/]+)"

			// a user may choose to override the defult expression
			// similar to expressjs: ‘/user/:id([0-9]+)’

			if index := strings.Index(part, "("); index != -1 {
				expr = part[index:]
				part = part[1:index]
			} else {
				part = part[1:len(part)]
			}

			params[j] = part
			parts[i] = expr
			j++
		}
	}

	// recreate the url pattern, with parameters replaced
	// by regular expressions. then compile the regex

	pattern = strings.Join(parts, "/")
	regex, regexErr := regexp.Compile(pattern)
	if regexErr != nil {

		// TODO add error handling here to avoid panic
		panic(regexErr)
		return
	}

	// now create the Route
	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	route := &routes.Route{}
	route.Regex = regex
	route.Methods = m
	route.Params = params
	route.ControllerType = t

	a.routes = append(a.routes, route)

}
func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	started := false
	requestPath := r.URL.Path

	//find a matching Route
	for _, route := range app.routes {

		// check if Route pattern matches url
		if !route.Regex.MatchString(requestPath) {
			continue
		}

		// get submatches (params)
		matches := route.Regex.FindStringSubmatch(requestPath)

		// double check that the Route matches the URL pattern.
		if len(matches[0]) != len(requestPath) {
			continue
		}

		params := make(map[string]string)

		if len(route.Params) > 0 {
			// add url parameters to the query param map
			values := r.URL.Query()
			for i, match := range matches[1:] {
				values.Add(route.Params[i], match)
				params[route.Params[i]] = match
				// fmt.Println(route.Params[i])
				// fmt.Println(match)
			}

			// reassemble query params and add to RawQuery
			// r.URL.RawQuery = url.Values(values).Encode() + "&" + r.URL.RawQuery
			// r.URL.RawQuery = url.Values(values).Encode()
		}

		// Invoke the request handler
		vc := reflect.New(route.ControllerType)

		//init 调用
		init := vc.MethodByName("Init")
		args := make([]reflect.Value, 2)
		ct := &contex.Context{
			ResponseWriter: w,
			Request:        r,
			Params:         params,
		}
		args[0] = reflect.ValueOf(ct)
		args[1] = reflect.ValueOf(route.ControllerType.Name())
		init.Call(args)

		//Prepare 调用
		args = make([]reflect.Value, 0)
		method := vc.MethodByName("Prepare")
		method.Call(args)

		//请求方法调用
		if _, ok := route.Methods[r.Method]; !ok {
			http.NotFound(w, r)
		}
		method = vc.MethodByName(route.Methods[r.Method])
		if !method.IsValid() {
			http.NotFound(w, r)
		}
		method.Call(args)

		//请求完成调用
		method = vc.MethodByName("Finish")
		method.Call(args)
		started = true

	}

	// if no matches to url, throw a not found exception
	if started == false {
		http.NotFound(w, r)
	}
}
func Run() {

	app := &App{}
	app.AddRoutes()
	err := http.ListenAndServe(":9090", app) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
