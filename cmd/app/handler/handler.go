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
	m.HandleFunc("/board-types/{cafeId:[0-9]+}/{id:[0-9]+}", h.getDetail).Methods(http.MethodGet)
	m.HandleFunc("/board-types/{cafeId:[0-9]+}/{typeId:[0-9]+}", h.delete).Methods(http.MethodDelete)
	m.HandleFunc("/board-types/{cafeId:[0-9]+}/{typeId:[0-9]+}", h.patch).Methods(http.MethodPatch)
	m.HandleFunc("/board-types/{cafeId:[0-9]+}/{memberId:[0-9]+}", h.create).Methods(http.MethodPost)

	return m
}

const (
	InvalidCafeId       = "invalid cafe id"
	InvalidMemberId     = "invalid member id"
	InvalidTypeId       = "invalid type id"
	InternalServerError = "internal server error"
	InvalidId           = "invalid board type id"
)

func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	memberId, err := strconv.Atoi(vars["memberId"])
	if err != nil {
		http.Error(w, InvalidMemberId, http.StatusBadRequest)
		return
	}

	var d req.CreateBoardTypeDto
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		log.Println("create json.NewDecoder err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) getList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
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
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (h Handler) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// 사실 카페 아이디 까지 필요하진 않지만 한번더 확인
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	typeId, err := strconv.Atoi(vars["typeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	err = h.c.Delete(r.Context(), cafeId, typeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h Handler) patch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// 사실 카페 아이디 까지 필요하진 않지만 한번더 확인 실제론 cafeAPI 측에서 해야함
	cafeId, err := strconv.Atoi(vars["cafeId"])
	if err != nil {
		http.Error(w, InvalidCafeId, http.StatusBadRequest)
		return
	}
	typeId, err := strconv.Atoi(vars["typeId"])
	if err != nil {
		http.Error(w, InvalidTypeId, http.StatusBadRequest)
		return
	}
	var d req.PatchBoardDto
	err = json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		log.Println("patch json.NewDecoder err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}

	err = h.c.Patch(r.Context(), cafeId, typeId, d)
	if err != nil {
		if strings.Contains(err.Error(), "no row") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h Handler) getDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, InvalidId, http.StatusBadRequest)
		return
	}
	dto, err := h.c.GetDetail(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	date, err := json.Marshal(dto)
	if err != nil {
		log.Println("getDetail json.Marshal err: ", err)
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(date)
}
