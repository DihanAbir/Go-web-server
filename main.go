package main 

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseFrom() err: %v", err)
		return 
	}
	fmt.Fprintf(w, "Post request successfully")
	name := r.FormValue("name")
	
	fmt.Fprintf(w, " \n Home = %s \n", name)
}


func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path  != "/hello"{
		http.Error(w, "404 not found",http.StatusNotFound)
		return 
	}
	if r.Method  != "GET"{
		http.Error(w, "Method is not supported",http.StatusNotFound)
		return 
	}
	fmt.Fprintf(w, "hello Boss!")
}


func main(){
	fileServe := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServe)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting serving at port : 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err) 
	}

}