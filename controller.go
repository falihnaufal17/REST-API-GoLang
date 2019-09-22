package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arrUser []Users
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
				arrUser = append(arrUser, users)
			}
		}

		response.Status = 1
		response.Message = "Success"
		response.Data = arrUser

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func insertUser(w http.ResponseWriter, r *http.Request) {
	var user Users
	var arrUser []Users
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	user.Name = r.FormValue("name")
	user.Phone = r.FormValue("phone")
	user.Location = r.FormValue("location")

	_, err = db.Exec(`INSERT INTO user (name, phone, location) value (?, ?, ?)`, user.Name, user.Phone, user.Location)

	if err != nil {
		log.Print(err)
	}

	arrUser = append(arrUser, user)

	response.Status = 1
	response.Message = "Success Add"
	response.Data = arrUser
	log.Println("insert data : ", user, " Sukses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
