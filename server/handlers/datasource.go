package handlers

import (
	"net/http"

	"github.com/gauravbansal74/mlserver/pkg/jwt"
	"github.com/gauravbansal74/mlserver/pkg/response"
	"github.com/gauravbansal74/mlserver/pkg/router"
	"github.com/gauravbansal74/mlserver/server/model/dataSource"
)

func init() {
	router.Get("/data-sources", AllDataSourcesHandler)
}

// AllDataSourcesHandler - get all data source file processing logs
func AllDataSourcesHandler(w http.ResponseWriter, r *http.Request) {
	_, err := jwt.GetData(r.Header.Get("Authorization"))
	if err != nil {
		response.SendError(w, http.StatusUnauthorized, "Access denied")
		return
	}
	list, err := dataSource.GetAllRecords()
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}
	response.SendJSON(w, http.StatusOK, list)
	return
}
