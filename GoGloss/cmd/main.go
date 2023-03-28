package main

import (
	"html/template"
	"net/http"
)

type Term struct {
	Name        string
	Description string
}

var terms = []Term{
	{Name: "Язык программирования", Description: "формальный язык, используемый для формирования структур данных и алгоритмов, то есть вычислительных правил которые могут быть выполнены компьютером."},
	{Name: "Rust", Description: "мультипарадигменный компилируемый язык программирования общего назначения, сочетающий парадигмы функционального и процедурного программирования с объектной системой, основанной на типажах."},
	{Name: "Go", Description: "компилируемый многопоточный язык программирования от Google с открытым исходным кодом. Считается языком общего назначения, но основное применение — разработка веб-сервисов и клиент-серверных приложений."},
	{Name: "Сервер", Description: "компьютер, на котором хранятся данные, или который выполняет определенные служебные функции для других компьютеров сети."},
	{Name: "Клиент", Description: "компьютер, запрашивающий некоторую функцию или данные у сервера."},
	{Name: "Приложение", Description: "программа, ориентированная на решение конкретных задач, рассчитанная на взаимодействие с пользователем."},
	{Name: "«Клиент-сервер»", Description: "вычислительная или сетевая архитектура, в которой задания или сетевая нагрузка распределены между поставщиками услуг, называемыми серверами, и заказчиками услуг, называемыми клиентами."},
	{Name: "Backend", Description: "внутренняя часть продукта, которая находится на сервере и скрыта от пользователей."},
}

func renderTerms(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("terms.html"))
	err := tmpl.Execute(w, terms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderMindMap(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("mindMap.html"))
	err := tmpl.Execute(w, terms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	http.HandleFunc("/termsGo", renderTerms)
	http.HandleFunc("/mindMap", renderMindMap)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}
