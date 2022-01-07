package router

import (
	"container/list"
	"github.com/gin-gonic/gin"
	"github.com/UseLiberty/swagin/security"
)

type Router struct {
	Handlers    *list.List
	Path        string
	Method      string
	Summary     string
	Description string
	Deprecated  bool
	ContentType string
	Tags        []string
	API         interface{}
	OperationID string
	Exclude     bool
	Securities  []security.ISecurity
	Response    Response
}

func (router *Router) GetHandlers() []gin.HandlerFunc {
	var handlers []gin.HandlerFunc
	for _, s := range router.Securities {
		handlers = append(handlers, s.Authorize)
	}
	for h := router.Handlers.Front(); h != nil; h = h.Next() {
		if f, ok := h.Value.(gin.HandlerFunc); ok {
			handlers = append(handlers, f)
		}
	}
	return handlers
}

func New(api interface{}, options ...Option) *Router {
	r := &Router{
		Handlers: list.New(),
		API:      api,
		Response: make(Response),
	}
	for _, option := range options {
		option(r)
	}
	return r
}
