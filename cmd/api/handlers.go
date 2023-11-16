package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(writer http.ResponseWriter, request *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "LIVE",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(out)
}

func (app *application) Lists(w http.ResponseWriter, r *http.Request) {
	var lists []models.List

	testList := models.List{
		ID:          0,
		Name:        "The List",
		Description: "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	lists = append(lists, testList)

	testList2 := models.List{
		ID:          1,
		Name:        "The (censored) List",
		Description: "Something about Hell's Kitchen",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	lists = append(lists, testList2)

	out, err := json.Marshal(lists)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
