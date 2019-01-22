package sessions

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

//Store the cookie store which is going to store mysession data in the cookie

var Store = sessions.NewCookieStore([]byte("secret-password"))
var mysession *sessions.Session

//IsLoggedIn will check if the user has an active mysession and return True
func IsLogined(r *http.Request) bool {

	mysession, err := Store.Get(r, "mysession")
	fmt.Println("=========", mysession, Store)
	if err == nil && (mysession.Values["loggedin"] == "true") {
		return true
	}
	return false
}

//GetCurrentUserName returns the username of the logged in user
func GetCurrentUserName(r *http.Request) string {
	mysession, err := Store.Get(r, "mysession")
	if err == nil {
		return mysession.Values["username"].(string)
	}
	return ""
}

//LogoutFunc Implements the logout functionality. WIll delete the mysession information from the cookie store
func LogoutFunc(w http.ResponseWriter, r *http.Request) {
	mysession, err := Store.Get(r, "mysession")
	if err == nil { //If there is no error, then remove mysession
		if mysession.Values["loggedin"] != "false" {
			mysession.Values["loggedin"] = "false"
			mysession.Save(r, w)
		}
	}
	http.Redirect(w, r, "/gologin", 302) //redirect to login irrespective of error or not
}

//LoginFunc implements the login functionality, will add a cookie to the cookie store for managing authentication
func SetLogined(w http.ResponseWriter, r *http.Request) {
	Store.MaxAge(0)
	mysession, err := Store.Get(r, "mysession")
	if err != nil {
		fmt.Println("error identifying mysession")
	}
	mysession.Values["loggedin"] = "true"
	mysession.Save(r, w)
}

/*

 */
func MustLogin(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !IsLogined(r) {
			fmt.Println("you must login...")
			http.Redirect(w, r, "/gologin", 302)
			return
		}
		handler(w, r)
	}
}
