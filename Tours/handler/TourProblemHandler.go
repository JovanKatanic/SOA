package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tours_service/service"

	"github.com/gorilla/mux"
)

type TourProblemHandler struct {
	TourProblemService *service.TourProblemService
}

func (handler *TourProblemHandler) GetByAuthorId(resp http.ResponseWriter, req *http.Request) {
	authorIdInt64, _ := strconv.ParseInt(mux.Vars(req)["authorId"], 10, 64)

	authorId := int(authorIdInt64)

	tourProblems, err := handler.TourProblemService.GetByAuthorId(&authorId)
	resp.Header().Set("Content-Type", "application/json")
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(tourProblems)
}
