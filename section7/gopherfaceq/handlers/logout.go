package handlers

import (
	"net/http"

	"github.com/EngineerKamesh/gofullstack/volume2/section7/gopherfaceq/common/authenticate"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)
}
