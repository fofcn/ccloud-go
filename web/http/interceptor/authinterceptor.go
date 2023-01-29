package interceptor

import (
	"ccloud/web/log"
	"ccloud/web/token"
	"net/http"
	"time"
)

type AuthInterceptor struct {
	paths []string
}

func NewAuthInterceptor(paths ...string) Interceptor {
	return &AuthInterceptor{
		paths: paths,
	}
}

func (auth AuthInterceptor) preRequest(w http.ResponseWriter, r *http.Request) bool {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		log.Logger.Error("No value in Authorization header")
		return false
	}

	tokenService := token.NewAuthTokenService()
	payload, err := tokenService.GetToken(tokenHeader)
	if err != nil {
		log.Logger.Error("get token from token store error, ", err)
		return false
	}

	// 检查过期时间
	expire := payload.Expire
	if expire.Before(time.Now()) {
		log.Logger.Info("token expired, token expire time: %v, now time: ", expire, time.Now())
		return false
	}

	// 添加自定义协议头
	r.Header.Set("x-user-id", payload.Payload["userId"])
	return true
}

func (auth AuthInterceptor) afterRequest(w http.ResponseWriter, r *http.Request) {

}
