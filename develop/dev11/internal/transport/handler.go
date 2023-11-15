package transport

import (
	"log"
	"net/http"
	"time"
)

type Handler interface {
	Init(router *http.ServeMux)
}

// CorsMiddleware - фунцкия которая перед срабатыванием любой API будет выполнять определенные действия
// В нашем случае, будет логировать каждый запрос
func CorsMiddleware(allowedHosts []string, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
