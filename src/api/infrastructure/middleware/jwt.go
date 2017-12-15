package middleware

import (
	"github.com/dtrinh100/Music-Playlist/src/api/interfaces"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"context"
)

type JWTMiddleware struct {
	MiddlewareHandler
	JWTHandler interfaces.JWTHandler
}

func (middleware *JWTMiddleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	_, claims, extractErr := middleware.JWTHandler.ExtractJWTFromRequest(req)

	if extractErr != nil {
		if extractErr.Error() == "http: named cookie not present" {
			middleware.Logger.Log("No JWT cookie present")
			middleware.Responder.Unauthorized(rw)
			return
		}

		switch extractErr.(type) {
		case *jwt.ValidationError:
			switch extractErr.(*jwt.ValidationError).Errors {
			case jwt.ValidationErrorExpired:
				middleware.Logger.Log("JWT has expired")
				middleware.Responder.Unauthorized(rw)
			default:
				middleware.Logger.Log("JWT claim-validation error")
				middleware.Responder.InternalServerError(rw)
			}
		default:
			middleware.Logger.Log("JWT is not valid")
			middleware.Responder.InternalServerError(rw)
		}

		return
	}

	if claims.RefreshAt < time.Now().Unix() {
		middleware.Logger.Log("Refreshing JWT for email: " + claims.UserEmail)
		middleware.JWTHandler.ValidateUserEmail(rw, claims.UserEmail)
	}

	ctx := context.WithValue(req.Context(), "mpEmailKey", claims.UserEmail)

	middleware.next.ServeHTTP(rw, req.WithContext(ctx))
}

func (middleware *JWTMiddleware) Handle(next http.Handler) http.Handler {
	middleware.next = next
	return middleware
}
