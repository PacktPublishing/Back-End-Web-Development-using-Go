package authenticate

import (
	"log"
	"net/http"
	"os"

	"github.com/EngineerKamesh/gofullstack/volume2/section7/gopherfaceq/models"

	"github.com/gorilla/sessions"
)

var SessionStore *sessions.FilesystemStore

func CreateUserSession(u *models.User, sessionID string, w http.ResponseWriter, r *http.Request) error {

	gfSession, err := SessionStore.Get(r, "gopherface-session")

	if err != nil {
		log.Print(err)
	}

	gfSession.Values["sessionID"] = sessionID
	gfSession.Values["username"] = u.Username
	gfSession.Values["firstName"] = u.FirstName
	gfSession.Values["lastName"] = u.LastName
	gfSession.Values["email"] = u.Email

	err = gfSession.Save(r, w)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func ExpireUserSession(w http.ResponseWriter, r *http.Request) {
	gfSession, err := SessionStore.Get(r, "gopherface-session")

	if err != nil {
		log.Print(err)
	}

	gfSession.Options.MaxAge = -1
	gfSession.Save(r, w)
}

func init() {

	SessionStore = sessions.NewFilesystemStore("/tmp/gopherface-sessions", []byte(os.Getenv("GOPHERFACE_HASH_KEY")))

}
