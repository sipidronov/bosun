package web

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"bosun.org/cmd/bosun/web/auth"
	"bosun.org/collect"
	"bosun.org/opentsdb"
	"github.com/MiniProfiler/go/miniprofiler"
	"github.com/NYTimes/gziphandler"
)

// MiddlewareFunc defines a function that returns a middleware. A Middleware can call next to continue the chain, or not if it is not appropriate
type MiddlewareFunc func(next http.HandlerFunc) http.Handler

// MiddlewareChain is a list of middlewares to be applied. The first element will be called first on a request.
type MiddlewareChain []MiddlewareFunc

//Extend takes a base middleware chain, copies its members, and adds the specified middlewares to the copy.
func (c MiddlewareChain) Extend(middlewares ...MiddlewareFunc) MiddlewareChain {
	newC := make(MiddlewareChain, 0, len(c)+len(middlewares))
	for _, m := range c {
		newC = append(newC, m)
	}
	for _, m := range middlewares {
		newC = append(newC, m)
	}
	return newC
}

// Build creates a single function that can be called to create concrete handlers from a middleware chain
func (c MiddlewareChain) Build() func(http.Handler) http.Handler {
	return func(root http.Handler) http.Handler {
		chain := root
		for i := len(c) - 1; i >= 0; i-- {
			chain = c[i](chain.ServeHTTP)
		}
		return chain
	}
}

func authMiddleware(required auth.PermissionLevel, provider auth.Provider) MiddlewareFunc {
	return func(next http.HandlerFunc) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := provider.GetUser(r)
			fmt.Println(user, err)
			if err != nil {
				//log maybe
			}
			if user == nil {
				user = &auth.User{
					Name:        "",
					AuthMethod:  "none",
					Permissions: auth.None,
				}
			}
			fmt.Println(user, user.Permissions, required)
			if user.Permissions >= required {
				newR := r.WithContext(context.WithValue(r.Context(), "user", user))
				next(w, newR)
				return
			}
			//auth failure. Redirect to login if html, or 403 for api
			if strings.Contains(r.Header.Get("Accept"), "html") {
				http.Redirect(w, r, "/login/?u="+url.QueryEscape(r.URL.String()), http.StatusFound)
				return
			}
			http.Error(w, "Not authorized", 403)
			return

		})
	}
}

// handle gzip with third-party package
func gzipMiddleware(next http.HandlerFunc) http.Handler {
	return gziphandler.GzipHandler(next)
}

//miniprofiler handler. Allows profiler to be pulled out of request context later
var miniprofileMiddleware = miniprofiler.NewContextHandler

// simple middleware to log if the request is http or https
func protocolLoggingMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proto := "http"
		if r.TLS != nil {
			proto = "https"
		}
		collect.Add("requests_by_protocol", opentsdb.TagSet{"proto": proto}, 1)
		next(w, r)
	})
}