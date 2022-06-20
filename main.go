package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	name  string
	age   uint16 //Целое положительное число
	money float32
}

func (u User) getAllInfo() string {
	return fmt.Sprint("User name is: ", u.name, ". He is ", u.age, ". And he has money ", u.money, "\n")
}

func (u *User) setNewName(newName string) {
	u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{
		name: "Bob", age: 24, money: 24.5,
	}
	templ, _ := template.ParseFiles("templates/index.html")
	templ.Execute(w, bob)
}

func profile(w http.ResponseWriter, r *http.Request) {
	bob := User{
		name: "Bob", age: 24, money: 24.5,
	}
	fmt.Fprintf(w, bob.getAllInfo())
	bob.setNewName("Ivan")
	fmt.Fprintf(w, bob.getAllInfo())
}

func handleRequest() {
	// Вывод главной страницы
	http.HandleFunc("/", home_page)
	//Ссылка на профиль
	http.HandleFunc("/profile/", profile)
	//Инициализация localhost с портом 8000
	http.ListenAndServe(":8000", nil)
}

func main() {
	handleRequest()
}
