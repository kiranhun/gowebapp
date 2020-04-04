package main

import (
    // "html/template"
	// "log"
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"encoding/json"
)

// type PageVariables struct {
// 	Name		 string
// 	Time         string
// }

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/time", timeZonePage).Methods("GET")
	
	fmt.Println("Listening");
	http.ListenAndServe("0.0.0.0:3000", router)
}

func timeZonePage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	timeInfo := make(map[string]interface{})
	key := r.URL.Query()

	for _, value := range key {
		loc, err := time.LoadLocation(value[0])
		if err != nil {
				fmt.Println(err)
				timeInfo[value[0]] = nil
				break
		}
		timeInfo[value[0]] = time.Now().In(loc)
	}

	json.NewEncoder(w).Encode(timeInfo)
}

func homePage(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	Time := time.Now()
	json.NewEncoder(w).Encode(Time)
}	
    // HomePageVars := PageVariables{ //store the date and time in a struct
	//   Time: time.Now().Format(time.Stamp) ,
	//   Name: r.FormValue("name"),
	// }
	
	

	// http.Handle("/static/", //final url can be anything
    //   http.StripPrefix("/static/",
	// 	 http.FileServer(http.Dir("static"))))
		 
	// t, err := template.ParseFiles("templates/homepage.html")

    // if err != nil { // if there is an error
  	//   log.Print("template parsing error: ", err) // log it
  	// }
    // err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
    // if err != nil { // if there is an error
  	//   log.Print("template executing error: ", err) //log it
  	// }