package main

import (
	"net/http"
	"fmt"
	"log"
)

type Hello struct {

}

func (h Hello) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprint(w,"Hello!")
}


func main() {

	var h Hello

	err := http.ListenAndServe("localhost:4000",h)

	if err != nil {
		log.Fatal(err)
	}


}
