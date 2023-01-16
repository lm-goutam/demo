package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func add_user(Username string, password string) bool {
	db, err := sql.Open("mysql", "root:admin@123@tcp(127.0.0.1)/gout")
	if err != nil {
		panic(err)
	}
	add, err := db.Query("Insert Into `Users`(`Name`, `Password`) VALUES (?, ?)", Username, password)
	//add,err := db.Query(("Insert Into Users(Name, Password) VALUES (?, ?)",(Username),(password)))
	if err != nil {
		panic(err)
	}
	fmt.Println(add)
	defer db.Close()
	return true
}

func check_user(Username string, password string) bool {
	db, err := sql.Open("mysql", "root:admin@123@tcp(127.0.0.1)/gout")
	if err != nil {
		panic(err)
	}
	var exists bool
	var query string
	query = fmt.Sprintf("SELECT EXISTS(SELECT Name FROM Users WHERE Name='%s' AND Password='%s')", (Username), (password))
	row := db.QueryRow(query).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	fmt.Println(row)
	defer db.Close()
	return exists
}

func login(w http.ResponseWriter, r *http.Request) {
	var temt = template.Must(template.ParseFiles("templates/login.html"))
	temt.Execute(w, nil)
}
func signup(w http.ResponseWriter, r *http.Request) {
	var temt = template.Must(template.ParseFiles("templates/signup.html"))
	temt.Execute(w, nil)
}

func signup_user(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var Username = r.Form["fname"]
	var Password = r.Form["Password"]
	fmt.Println(Username, " ", Password)
	if add_user(Username[0], Password[0]) {
		var temt = template.Must(template.ParseFiles("templates/index.html"))
		temt.Execute(w, nil)
	} else {
		var temt = template.Must(template.ParseFiles("templates/error.html"))
		temt.Execute(w, nil)
	}
}

func login_user(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var Username = r.Form["fname"]
	var Password = r.Form["Password"]
	fmt.Println(Username, " ", Password)
	if check_user(Username[0], Password[0]) {
		var temt = template.Must(template.ParseFiles("templates/index.html"))
		temt.Execute(w, nil)
	} else {
		var temt = template.Must(template.ParseFiles("templates/error.html"))
		temt.Execute(w, nil)
	}
}

func main() {

	http.HandleFunc("/", login)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login_user", login_user)
	http.HandleFunc("/signup_user", signup_user)

	http.ListenAndServe(":8000", nil)
}
