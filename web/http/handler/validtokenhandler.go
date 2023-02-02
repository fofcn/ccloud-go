package handler

import (
	"ccloud/web/entity"
	"ccloud/web/http/util"
	"ccloud/web/service"
	"net/http"
)

type validtokenhandler struct {
	pattern      string
	handler      http.Handler
	loginservice service.LoginService
}

func NewViliadTokenHandler() (HttpHandler, error) {
	loginservice, err := service.NewLoginService()
	if err != nil {
		return nil, err
	}

	login := &validtokenhandler{
		pattern:      "/account/token/valid",
		loginservice: loginservice,
	}

	login.handler = login
	return login, nil
}

func (login validtokenhandler) Pattern() string {
	return login.pattern
}

func (login validtokenhandler) Handler() http.Handler {
	return login.handler
}

func (login *validtokenhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	util.WriteJson(w, entity.OK())
}
