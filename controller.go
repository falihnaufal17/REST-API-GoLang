package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arr_user []Users
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT userid, name, phone, location FROM user")
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err != nil {
			log.Print(err)
		}

		for rows.Next() {
			if err := rows.Scan(&users.UserId, &users.Name, &users.Phone, &users.Location); err != nil {
				log.Fatal(err.Error())
			} else {
				arr_user = append(arr_user, users)
			}
		}

		response.Status = 1
		response.Message = "Success"
		response.Data = arr_user

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func insertUser(w http.ResponseWriter, r *http.Request){
	var users Users
	var arr_user []Users
	var response Response

	db:=connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	type data struct {
		name string = r.FormValue("name")
		phone string = r.FormValue("phone")
		location string = r.FormValue("location")
	}

	_, err = db.Exec("INSERT INTO user SET ?", data)

	if err != nil{
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Add"
	response.Data = data

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}