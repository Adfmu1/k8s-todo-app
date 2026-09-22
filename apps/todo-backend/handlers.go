package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (application app) getTodos(w http.ResponseWriter, req *http.Request) {
	retArray, err := application.dbQuery.GetTodos(req.Context())
	checkErr(err, *application.logger)
	application.logger.Info("Retrieving todos from db")

	data, err := json.Marshal(retArray)
	checkErr(err, *application.logger)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

func (application app) postTodos(w http.ResponseWriter, req *http.Request) {
	type newTodo struct {
		Text string `json:"newTodo"`
	}
	decoder := json.NewDecoder(req.Body)
	nTodo := newTodo{}
	err := decoder.Decode(&nTodo)
	checkErr(err, *application.logger)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if strings.Trim(nTodo.Text, " ") == "" {
		application.logger.Error("Not a valid todo")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if len(nTodo.Text) > 140 {
		application.logger.Error("Todo too long")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	application.logger.Info("Writing data to db", "Todo text:", nTodo)

	application.dbQuery.CreateTodo(req.Context(), nTodo.Text)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(201)
}
