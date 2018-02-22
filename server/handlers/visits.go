package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gauravbansal74/mlserver/pkg/jwt"
	"github.com/gauravbansal74/mlserver/pkg/response"
	"github.com/gauravbansal74/mlserver/pkg/router"
	"github.com/gauravbansal74/mlserver/server/exclusion"
	"github.com/gauravbansal74/mlserver/server/model/visits"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

func init() {
	router.Get("/visits/:selectedDate", visitsByDate)
	router.Get("/visits", visitOverAll)
}

func visitOverAll(w http.ResponseWriter, r *http.Request) {
	_, err := jwt.GetData(r.Header.Get("Authorization"))
	if err != nil {
		response.SendError(w, http.StatusUnauthorized, "Access denied")
		return
	}
	listExclusion, err := exclusion.GetList()
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	list, err := visits.GetOverAll(listExclusion)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}
	if len(list) > 0 {
		response.SendJSON(w, http.StatusOK, list)
		return
	}
	response.SendMessage(w, http.StatusOK, "no records")
	return
}

func visitsByDate(w http.ResponseWriter, r *http.Request) {
	_, err := jwt.GetData(r.Header.Get("Authorization"))
	if err != nil {
		response.SendError(w, http.StatusUnauthorized, "Access denied")
		return
	}
	p := context.Get(r, "params").(httprouter.Params)
	selectedDateString := p.ByName("selectedDate")
	if selectedDateString != "" {
		selectedDateSeconds, err := strconv.Atoi(selectedDateString)
		if err != nil {
			response.SendError(w, http.StatusBadRequest, err.Error())
			return
		}
		selectedDate := time.Unix(0, int64(selectedDateSeconds)*int64(time.Millisecond))
		listExclusion, err := exclusion.GetList()
		if err != nil {
			response.SendError(w, http.StatusBadRequest, err.Error())
			return
		}
		list, err := visits.GetByDate(time.Date(selectedDate.Year(), selectedDate.Month(), selectedDate.Day(), 0, 0, 0, 0, time.UTC), listExclusion)
		if err != nil {
			response.SendError(w, http.StatusBadRequest, err.Error())
			return
		}
		if len(list) > 0 {
			response.SendJSON(w, http.StatusOK, list)
			return
		}
		response.SendMessage(w, http.StatusOK, "no records")
		return
	}
	response.SendError(w, http.StatusBadRequest, "selected date can't be null or empty")
	return
}
