package utils

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Frontend string = "layout/front.html"

var sessionKey string = "test-session-key"
var sessionName string = "session-alert"
var Store *sessions.CookieStore = sessions.NewCookieStore([]byte(sessionKey))

func ReturnAlertFlash(response http.ResponseWriter, request *http.Request) (string, string) {
	session, err := Store.Get(request, sessionName)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return "", ""
	}

	var flashes []string

	fm := session.Flashes("css")
	if len(fm) > 0 {
		session.Save(request, response)

		for _, css := range fm {
			flashes = append(flashes, css.(string))
		}

	}

	fm2 := session.Flashes("message")

	if len(fm2) > 0 {
		session.Save(request, response)

		for _, message := range fm2 {
			flashes = append(flashes, message.(string))
		}
	}

	if len(flashes) == 0 {
		return "", ""
	}

	return flashes[0], flashes[1]
}

func CreateAlertFlash(response http.ResponseWriter, request *http.Request, css string, message string) {

	session, err := Store.Get(request, sessionName)

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	session.AddFlash(css, "css")
	session.AddFlash(message, "message")
	err = session.Save(request, response)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

}
