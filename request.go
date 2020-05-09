package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// ResultStruct for server
type ResultStruct struct {
	Result string
}

// HomePage for server
func HomePage(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.ParseFloat(r.FormValue("a"), 10)
	b, err := strconv.ParseFloat(r.FormValue("b"), 10)
	operator := r.FormValue("operator")
	var arithmetic float64
	var result string
	switch operator {
	case "plus":
		arithmetic = a + b
		result = fmt.Sprintf("%.2f + %.2f = %.2f", a, b, arithmetic)
	case "minus":
		arithmetic = a - b
		result = fmt.Sprintf("%.2f - %.2f = %.2f", a, b, arithmetic)
	case "multiply":
		arithmetic = a * b
		result = fmt.Sprintf("%.2f * %.2f = %.2f", a, b, arithmetic)
	case "divide":
		arithmetic = a / b
		result = fmt.Sprintf("%.2f / %.2f = %.2f", a, b, arithmetic)
	}

	Result := ResultStruct{
		Result: result,
	}
	t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
	if err != nil {                                // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, Result) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {            // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func main() {

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/calc", HomePage)

	http.ListenAndServe(":8090", nil)
}
