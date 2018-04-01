
package main

import ( "fmt"
"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,WWworld!",r.URL)
}

func user_handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"<h1>user<h1>",r.URL)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/user/",user_handler)

	http.ListenAndServe(":8080", nil)
}