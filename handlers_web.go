package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Debug("Serving template with url = " + r.URL.Path)
	url := vars["path"]

	// For root, show the index page
	if len(url) == 0 {
		url = "index"
	}

	if err := templates.ExecuteTemplate(w, url, staticPath); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func rememberMeHandler(w http.ResponseWriter, r *http.Request) {
	session, err := api.store.Get(r, "session-login")
	if err != nil {
		log.Fatal(err.Error())
	}

	if session.Values["user"] == nil {
		if err = WriteJson(w, map[string]interface{}{"remember": false}); err != nil {
			Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		WriteJson(w, map[string]interface{}{"remember": true, "user": session.Values["user"]})
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	user := Farmacias{}
	if err := DecodeJson(r, &user); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := api.DB.Where("nik = ? AND pass = ?", user.Nik, user.Pass).Find(&user).Error; err != nil {
		Error(w, "User not found", http.StatusNotFound)
		return
	}

	session, err := api.store.Get(r, "session-login")
	if err != nil {
		log.Fatal(err.Error())
	}

	session.Values["user"] = user
	if err = session.Save(r, w); err != nil {
		log.Fatal(err.Error())
	}

	WriteJson(w, map[string]interface{}{"user": &user})
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := api.store.Get(r, "session-login")
	if err != nil {
		log.Fatal(err.Error())
	}

	session.Values["user"] = nil
	if err = session.Save(r, w); err != nil {
		log.Fatal(err.Error())
	}

	WriteJson(w, map[string]interface{}{"user": nil})
}

/* Call a go-json-rest handler,
check if it's in the store here,
	if not lookup in database,
		if it is
			save in store
			return ok in json,
		if not
			return user not found
	if it it's alreay in the store, show logged in.
*/
