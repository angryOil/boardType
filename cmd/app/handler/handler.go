package handler

import (
	"boardType/internal/controller"
	"boardType/internal/controller/req"
	"boardType/internal/controller/res"
	"boardType/internal/page"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	c controller.BoardTypeController
}

func NewHandler(c controller.BoardTypeController) http.Handler {
	m := mux.NewRouter()
	h := Handler{c: c}
	m.HandleFunc("/board-types/{cafeId:[0-9]+}", h.getList).Methods(http.MethodGet)
	m.HandleFunc("/board-types/{cafeId:[0-9]+}/{memberId:[0-9]+}", h.create).Methods(http.MethodPost)

	return m
}

func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}
	memberId, err := strconv.Atoi(vars["memberId"])
	if err != nil {
		http.Error(w, "invalid member id", http.StatusBadRequest)
		return
	}

	var d req.CreateBoardTypeDto
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		log.Println("create json.NewDecoder err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	err = h.c.Create(r.Context(), cafeId, memberId, d)
	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if strings.Contains(err.Error(), "duplicate") {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		log.Println("Create err: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) getList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, "invalid cafe id", http.StatusBadRequest)
		return
	}
	reqPage := page.GetPageReqByRequest(r)
	dtoList, total, err := h.c.GetListByCafe(r.Context(), cafeId, reqPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(res.NewListTotalDto(dtoList, total))
	if err != nil {
		log.Println("getList json.Marshal err: ", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}
