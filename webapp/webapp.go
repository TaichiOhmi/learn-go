package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type ToDoList struct {
	ToDoCount int
	ToDos     []string
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// writerでブラウザへの書き込みができる
// メッセージを作成して、それをブラウザに反映させるレスポンスに加える。
func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	errorCheck(err)
}

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	errorCheck(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	errorCheck(scanner.Err())
	return lines
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello Internet")
}

func japaneseHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "こんにちは インターネット")
}

func spanishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hola Internet")
}

func interactHandler(writer http.ResponseWriter, request *http.Request) {
	// text を　[]string　として txtファイルから取得
	todoValues := getStrings("todos.txt")
	fmt.Printf("%#v\n", todoValues)
	tmpl, err := template.ParseFiles("view.html")
	errorCheck(err)
	todos := ToDoList{
		ToDoCount: len(todoValues),
		ToDos:     todoValues,
	}
	err = tmpl.Execute(writer, todos)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	errorCheck(err)
	err = tmpl.Execute(writer, nil)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	todo := request.FormValue("todo")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("todos.txt", options, os.FileMode(0600))
	errorCheck(err)
	_, err = fmt.Fprintln(file, todo)
	errorCheck(err)
	err = file.Close()
	errorCheck(err)
	http.Redirect(writer, request, "/interact", http.StatusFound)
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/hola", spanishHandler)
	http.HandleFunc("/こんにちは", japaneseHandler)
	http.HandleFunc("/interact", interactHandler)
	http.HandleFunc("/", newHandler)
	http.HandleFunc("/create", createHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
