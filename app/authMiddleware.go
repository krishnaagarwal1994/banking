package app

import (
	"banking/domain"
	"banking/errs"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type AuthMiddleware struct {
	repository domain.AuthRepository
}

func (m AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			auth_header := r.Header.Get("Authorization")

			if auth_header != "" {
				auth_token := getTokenFromHeader(auth_header)
				isAuthorized := m.repository.IsAuthorized(auth_token, currentRoute.GetName(), currentRouteVars)
				if isAuthorized {
					h.ServeHTTP(w, r)
				} else {
					appError := errs.AppError{Message: "unauthorised", Code: http.StatusForbidden}
					writeResponse(w, appError.Code, appError.AsMessage())
				}
			} else {
				writeResponse(w, http.StatusUnauthorized, errs.AppError{Message: "missing token"}.AsMessage())
			}
		})
	}
}

func getTokenFromHeader(value string) string {
	words := strings.Split(value, " ")
	if len(words) <= 1 {
		return value
	} else {
		return words[1]
	}
}
