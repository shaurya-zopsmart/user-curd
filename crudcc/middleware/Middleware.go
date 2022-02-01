package Middlewares

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

func Logger(inner http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		userName, pwd, ok := r.BasicAuth()
		if ok {
			userNameHash := sha256.Sum256([]byte(userName))
			userPwdHash := sha256.Sum256([]byte(pwd))
			expecName := sha256.Sum256([]byte("anusri18101"))
			expecPwd := sha256.Sum256([]byte("anusri@18101"))

			userNameMatch := (subtle.ConstantTimeCompare(userNameHash[:], expecName[:]) == 1)
			pwdMatch := (subtle.ConstantTimeCompare(userPwdHash[:], expecPwd[:]) == 1)

			if userNameMatch && pwdMatch {
				inner.ServeHTTP(rw, r)
				return
			}
		}
		// rw.Header().Set("WWW-Authenticate", Basic realm="restricted", charset="UTF-8")
		// rw.WriteHeader(http.StatusUnauthorized)
		http.Error(rw, "Unauthorized", http.StatusUnauthorized)
	})
}
