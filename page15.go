<<<<<<< HEAD
package main

import (
	"fmt"
	"log"
	"net/http"
)
func main(){
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
=======
package main

import (
	"fmt"
	"log"
	"net/http"
)
func main(){
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
>>>>>>> 24181879451b8db8a6db4d439894d3ade71ca404
}