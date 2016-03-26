package controllers

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/Pruthvik-n/Go-sample/api/v1/models"
	"database/sql"
	_"github.com/lib/pq"
	"strconv"
	"github.com/gorilla/mux"
	"log"
	)

type RegisterController struct{}

var Register RegisterController

func (r RegisterController) Create(rw http.ResponseWriter, req *http.Request) {

	var u models.User

	req_body, err := ioutil.ReadAll(req.Body)
	if err != nil{
		panic(err)
	}

	body := json.Unmarshal(req_body, &u)
	if body != nil{
		panic(err)
	}

	db, err := sql.Open("postgres", "user=pruthvik password=password host=localhost dbname=go_sample sslmode=disable")
	if err !=nil {
		panic(err)
	}

	insert_query := "INSERT INTO users (id, name) VALUES ($1, $2)"
	if err != nil {
		panic(err)
	}

	prepare_insert_query, err := db.Prepare(insert_query)
	if err != nil{
		panic(err)
		
	}
	defer prepare_insert_query.Close()

	execute_prepare_query, err := prepare_insert_query.Exec(u.Id, u.Name)
	if err != nil || execute_prepare_query == nil{
		panic(err)
	}

    b, err := json.Marshal(models.UserSignUpResponse{
      Message: "User Created Successfully",
    })
    rw.Header().Set("Content-type","application/json")
    rw.Write(b)
	db.Close()

}

func (r RegisterController) Fetch(rw http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	id := vars["id"]

	tmp, err := strconv.Atoi(id)
	user_id := tmp
	if err != nil || tmp == 0 {
		panic(err)
	}
	db, err := sql.Open("postgres", "user=pruthvik password=password host=localhost dbname=go_sample sslmode=disable")
	if err != nil {
		panic(err)
	}

	get_details, err := db.Query("SELECT * FROM USERS WHERE id=$1", user_id)
  if err != nil {
    panic(err)
  }

  var u_details []models.User
  for get_details.Next() {
    var id int
    var name string
    err := get_details.Scan(&id, &name)
    if err != nil {
      panic(err)
    }
    user_details := models.User{id, name}
    u_details = append(u_details, user_details)
    b, err := json.Marshal(models.UserFetchResponse{
      Message: "All Users",
      User_details: u_details,
    })
    rw.Header().Set("Content-type","application/json")
    rw.Write(b)
  }
  db.Close()	
}	

func (r RegisterController) FetchAll(rw http.ResponseWriter, req *http.Request) {

	db, err := sql.Open("postgres", "user=pruthvik password=password host=localhost dbname=go_sample sslmode=disable")
		if err != nil {
			panic(err)
	}

	get_details, err := db.Query("SELECT * FROM USERS")
  if err != nil {
    panic(err)
  }

   var u_details []models.User
  	

  	for get_details.Next() {
    var id int
    var name string
    err := get_details.Scan(&id, &name)
    if err != nil {
      panic(err)
    }
    user_details := models.User{id, name}
    u_details = append(u_details, user_details)
}
    b, err := json.Marshal(models.UserFetchResponse{
      Message: "All Users",
      User_details: u_details,
    })
    if err != nil {
    	panic(err)
    }
    rw.Header().Set("Content-type","application/json")
    log.Println(u_details)
    rw.Write(b)
  db.Close()


}

func (r RegisterController) Delete(rw http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	id := vars["id"]

	tmp, err := strconv.Atoi(id)
	user_id := tmp
	if err != nil || tmp == 0 {
		panic(err)
	}
	db, err := sql.Open("postgres", "user=pruthvik password=password host=localhost dbname=go_sample sslmode=disable")
	if err != nil {
		panic(err)
	}

	// check_user, err := db.Query("SELECT id FROM USERS WHERE id=$1", user_id)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(check_user)
	// for check_user.Next(){
	// 	var id int
	// 	err := check_user.Scan(&id)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	
	get_details, err := db.Query("DELETE FROM USERS WHERE id=$1", user_id)
  	if err != nil || get_details == nil{
    panic(err)
  }
    b, err := json.Marshal(models.UserSignUpResponse{
      Message: "Boom shaka laka ! User deleted Successfully !",
    })
    rw.Header().Set("Content-type","application/json")
    rw.Write(b)
  db.Close()	
}