package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Todo struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoReport struct {
	Count int
	Todos []Todo
}

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var todos []Todo

	err = json.Unmarshal(data, &todos)

	if err != nil {
		log.Fatal(err)
	}

	reportData := TodoReport{len(todos), todos}

	report, err := template.New("index.tmpl").Funcs(template.FuncMap{"upper": Upper}).ParseFiles("index.tmpl")

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("index.html")

	if err != nil {
		log.Fatal(err)
	}

	err = report.Execute(f, reportData)

	if err != nil {
		log.Fatal(err)
	}
}

func Upper(s string) string {
	return strings.ToUpper(s)
}
