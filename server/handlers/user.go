package handlers

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gauravbansal74/mlserver/pkg/jwt"
	"github.com/gauravbansal74/mlserver/pkg/logger"
	"github.com/gauravbansal74/mlserver/pkg/response"
	"github.com/gauravbansal74/mlserver/pkg/router"
	"github.com/gauravbansal74/mlserver/pkg/utils"
	"github.com/gauravbansal74/mlserver/server/model/user"
)

var (
	errFriendly = errors.New("internal server error")
)

const (
	userRegisteredSuccessfully = "Successfully registered."
)

func init() {
	router.Get("/user", profileHandler)
	router.Post("/login", loginHandler)
	router.Post("/register", registerHandler)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	userData, err := jwt.GetData(r.Header.Get("Authorization"))
	if err != nil {
		response.SendError(w, http.StatusUnauthorized, "Access denied")
		return
	}
	userInfo, err := user.GetByEmail(userData.Email)
	if err != nil {
		logger.Error(err, "Error while reading body Register")
		response.SendError(w, http.StatusBadRequest, errFriendly.Error())
		return
	}
	response.SendJSON(w, http.StatusOK, userInfo)
	return
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err, "Error while reading body Register")
		response.SendError(w, http.StatusBadRequest, errFriendly.Error())
		return
	}
	userData, err := user.New(body)
	if err != nil {
		logger.Error(err, "error while getting user data into struct value")
		response.SendError(w, http.StatusBadRequest, errFriendly.Error())
		return
	}
	err = userData.Create()
	if err != nil {
		logger.Error(err, "error while creating user")
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}
	response.SendMessage(w, http.StatusOK, userRegisteredSuccessfully)
	return
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err, "error while reading body Register")
		response.SendError(w, http.StatusBadRequest, errFriendly.Error())
		return
	}
	userData, err := user.Body(body)
	if err != nil {
		logger.Error(err, "error while getting user data into struct value")
		response.SendError(w, http.StatusBadRequest, errFriendly.Error())
		return
	}

	userInfo, err := user.GetByEmail(userData.Email)
	if err != nil {
		if err.Error() != "not found" {
			logger.Error(err, "error while fetching user data")
			response.SendError(w, http.StatusBadRequest, "error while fetching user data")
			return
		}
		response.SendError(w, http.StatusBadRequest, "not registered")
		return

	}
	userHash, err := utils.GetHash(userData.Password, userInfo.Salt)
	if err != nil {
		logger.Error(err, "error while getting user hash value")
		response.SendError(w, http.StatusBadRequest, "error while validation")
		return
	}
	if userHash == userInfo.Hash {
		token := jwt.GetToken(userInfo.Email, userInfo.ID.Hex())
		w.Header().Add("Authorization", token)
		response.SendMessage(w, http.StatusOK, "success")
		return
	}
	response.SendError(w, http.StatusUnauthorized, "email or password invalid")
	return
}
