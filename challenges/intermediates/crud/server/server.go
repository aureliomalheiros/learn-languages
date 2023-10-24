package server

import(
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"crud/db"
	"strconv"
	"github.com/gorilla/mux"
)

type user struct {
	ID uint32 `json:"id"`
	Name string `json:"name"`
	Email string `json: "email"`
}
// Create user in the DB
func CreateUser(w http.ResponseWriter, r *http.Request){

	Requesting, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("ERROR! FAILD READ REQUEST"))
	}

	var user user

	if erro = json.Unmarshal(Requesting, &user); erro != nil {
		w.Write([]byte("Erro to convert user"))
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		w.Write([]byte("ERROR! Connect in the DB"))
		return
	}
	statement, erro := db.Prepare("insert into user (name, email) values (?,?)")

	if erro != nil {
		w.Write([]byte("Erro! Create statement"))
		return
	}
	defer statement.Close()

	insert, erro := statement.Exec(user.Name, user.Email)
	if erro != nil {
		w.Write([]byte("Error! Execute statement"))
		return
	}

	idInsert, erro := insert.LastInsertId()
	if erro != nil {
		w.Write([]byte("Error! Get id insert"))
		return
	}

	//STATUS CODE

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User insert succecss! id: %d", idInsert)))
}

//Find all users
func FindUsers(w http.ResponseWriter, r *http.Request){
	db, erro := db.Connect()

	if erro != nil {
		w.Write([]byte("Error! Faild connection db"))
	}
	defer db.Close()

	lines, erro := db.Query("select * from user")

	if erro !=  nil {
		w.Write([]byte("ERRO! Find the users"))
	}

	defer lines.Close()

	var users []user
	for lines.Next() {
		var user user

		if erro := lines.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.Write([]byte("ERROR! Scan user"))
			return
		}

		users = append(users, user)

		w.WriteHeader(http.StatusOK)

		if erro := json.NewEncoder(w).Encode(users); erro != nil {
			w.Write([]byte("ERROR! Convert user to JSON"))
			return
		}
	}
}

//Find specified user
func FindUser(w http.ResponseWriter, r *http.Request){
	parametrs := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametrs["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Error! Converter parametrs to int"))
		return
	}

	db, erro := db.Connect()

	if erro != nil {
		w.Write([]byte("Error! Connect DB"))
		return
	}
	line, erro := db.Query("select * from user where id = ?", ID)
	
	if erro != nil {
		w.Write([]byte("Error! Query"))
		return
	}

	var user user 

	if line.Next(){
		if erro := line.Scan(&user.ID, &user.Name, &user.Email); erro != nil {
			w.Write([]byte("ERROR! Scan user"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		w.Write([]byte("ERROR! Converter user to JSON"))
		return
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	parametrs := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametrs["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("Erro! Converter parametrs to int"))
		return
	}

	Requesting, erro := ioutil.ReadAll(r.Body)

	var user user

	if erro := json.Unmarshal(Requesting, &user); erro != nil {
		w.Write([]byte("ERROR, convert user to struct"))
		return
	}

	db, erro := db.Connect()

	if erro != nil {
		w.Write([]byte("ERRO! Connect DB"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update user set name = ?, email = ? where id = ?")

	if erro != nil {
		w.Write([]byte("Erro! Create of statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(user.Name, user.Email, ID); erro != nil {
		w.Write([]byte("Erro! Update user"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}


//Delete users
func DeleteUser(w http.ResponseWriter, r *http.Request){
	parametrs := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametrs["id"], 10, 32)

	if erro != nil {
		w.Write([]byte("ERROR! Converter parametrs to INT"))
		return
	}

	db, erro := db.Connect()

	if erro != nil {
		w.Write([]byte("ERRO! Connect DB"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("delete from user where id = ?")

	if erro != nil {
		w.Write([]byte("ERRO! Create statement"))
		return
	}

	defer statement.Close()

	if  _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("ERROR! Delete user"))
		return
	}
}