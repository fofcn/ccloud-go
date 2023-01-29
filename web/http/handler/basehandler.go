package handler

import (
	"ccloud/web/constant"
	"ccloud/web/entity"
	"ccloud/web/http/interceptor"
	"ccloud/web/http/util"
	"net/http"
)

type BaseHandler struct {
	Handle HttpHandler
}

// Handler implements HttpHandler
func (base BaseHandler) Handler() http.Handler {
	return base
}

// Pattern implements HttpHandler
func (base BaseHandler) Pattern() string {
	return base.Handle.Pattern()
}

func (base BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result := interceptor.GetInstance().Intecept(w, r)
	if !result {
		util.WriteJson(w, entity.Fail(constant.InterceptorFailed))
		return
	}
	base.Handle.Handler().ServeHTTP(w, r)
}
