package main

import (
"fmt"
	"log"
"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query();
	a := v.Get("a")
	fmt.Printf("param a value is [%s]\n", a);
	b := v.Get("b")
	fmt.Printf("param a value is [%s]\n", b);
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}