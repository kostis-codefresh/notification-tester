package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/callme", list_notifications)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Listening now at port 8080")
	http.ListenAndServe(":8080", nil)
}

func list_notifications(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<div class=\"entry\"><span id=\"dfdf\">No ArgoCD Notifications yet</span></div>")
}

// func create_notification(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "<div class=\"entry\"><span id=\"dfdf\">No ArgoCD Notifications yet</span></div>")
// }
