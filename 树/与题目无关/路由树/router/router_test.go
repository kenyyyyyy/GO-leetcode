package router

import (
	"fmt"
	"testing"
)

func newTestRouter() *Router {
	Router := NewRouter()
	return Router
}

func TestGetRouteGet(t *testing.T) {
	r := newTestRouter()
	r.AddRoute("POST", "hello/:name/test")
	node, err := r.GetRoute("GET", "/hello/yyy/test")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(node)
}

func TestGetRouteFail(t *testing.T) {
	r := newTestRouter()
	r.AddRoute("GET", "hello/:name")
	node, err := r.GetRoute("GET", "/hello/yyy/test")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(node)
}

func TestGetRouteSuccess(t *testing.T) {
	r := newTestRouter()
	r.AddRoute("GET", "hello/:name/test")
	node, err := r.GetRoute("GET", "/hello/yyy/test")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(node)
}
