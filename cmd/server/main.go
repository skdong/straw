package main

import (
	"net/http"
	"fmt"
)

func IndexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello world")

}

func runServer()error{
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":80", nil)
	return nil
}

func main(){
	if err := runServer(); err != nil{
		fmt.Println(err)
	}
}