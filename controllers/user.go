package controllers

import (
	"addVWeb/models"
	"encoding/json"
	"net/http"
	"regexp"
)

type userConstroller struct {
	userIDPattern *regexp.Regexp
}

func (uc userConstroller) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello form the user controller"))
}


func (uc * userConstroller ) getAll(w http.ResponseWriter, r *http.Request){
	encodeResponseAsJSON(model.GetUsers(), w)
}


func (uc *userConstroller) get(id int, w http.ResponseWriter){

	u, err := models.GetUserByID((id))

	if err !=nil{

		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	encodeResponseAsJSON(u, w)
}



func (uc *userConstroller) post(w http.ResponseWriter, r *http.Request){

	u, err := uc.parseRequest(r)

	if err != nil {

		w.WriteHeader((http.StatusInternalServerError))
		w.Write([]byte("Could not parse User Object"))
		return 
	}

	u, err = models.AddUser((u))
	if err !=nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	encodeResponseAsJSON(u, w)
}


func (uc *userConstroller) put(id int, w http.ResponseWriter, r *http.Request){
u, err := uc.parseRequest(r)
if err != nil{

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Could not parse User object"))
	return
}

if id != u.ID{

	w.WriteHeader(http.StatusBadRequest)
	w.Write([] byte ("ID of submitted user must match ID in  URL"))
	return 
}

u, err = models.updateUser(u)
if err != nil{

	w.WriteHeader((http.StatusInternalServerError))
	w.Write([]byte(err.Error()))
	return 
}


encodeResponseAsJSON(u, w)
}


func (uc *userConstroller) parseRequest(r *http.Request)(models.User, error){

	dec := json.NewDecoder(r, r.Body)
	var u models.User
	err := dec.Decode((&u))
	if err!= nil {
		return models.User{}, err
	}

	return u, nil
}





func newUserController() *userConstroller {

	return &userConstroller{

		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?$`),
	}

}
