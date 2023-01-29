package interceptor

import (
	"net/http"
	"sync"
)

type InterceptorChain interface {
	RegisterInterceptor(Interceptor)
	Intecept(w http.ResponseWriter, r *http.Request) bool
}

type DefaultInterceptorChain struct {
	list []Interceptor
}

var instance InterceptorChain
var once sync.Once

func GetInstance() InterceptorChain {
	once.Do(func() {
		instance = &DefaultInterceptorChain{
			list: make([]Interceptor, 0),
		}
	})

	return instance
}

func (chain *DefaultInterceptorChain) RegisterInterceptor(interceptor Interceptor) {
	chain.list = append(chain.list, interceptor)
}

func (chain *DefaultInterceptorChain) Intecept(w http.ResponseWriter, r *http.Request) bool {
	for _, interceptor := range chain.list {
		result := interceptor.preRequest(w, r)
		if !result {
			return false
		}
	}

	return true
}
