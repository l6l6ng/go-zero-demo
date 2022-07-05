package middleware

import (
	"net/http"
)

type RbacMiddleware struct {
}

func NewRbacMiddleware() *RbacMiddleware {
	return &RbacMiddleware{}
}

func (m *RbacMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next(w, r)
	}
}
