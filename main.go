package main

import (
	"fmt"
	"net/http"
	"sync"
)

type notificationHandler struct {
	mu sync.Mutex // guards n
	n  int
}

func main() {

	nf := new(notificationHandler)

	http.HandleFunc("/callme", nf.listNotifications)
	http.HandleFunc("/notify", nf.createNotification)
	http.HandleFunc("/clear", nf.clearNotifications)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	fmt.Println("Listening now at port 8080")
	http.ListenAndServe(":8080", nil)
}

func (nh *notificationHandler) listNotifications(w http.ResponseWriter, req *http.Request) {
	nh.mu.Lock()
	defer nh.mu.Unlock()
	fmt.Fprintf(w, "<div class=\"entry\"><span id=\"dfdf\">%d</span></div>", nh.n)
}

func (nh *notificationHandler) createNotification(w http.ResponseWriter, req *http.Request) {
	nh.mu.Lock()
	defer nh.mu.Unlock()
	nh.n++
	fmt.Fprintf(w, "Created")
}

func (nh *notificationHandler) clearNotifications(w http.ResponseWriter, req *http.Request) {
	nh.mu.Lock()
	defer nh.mu.Unlock()
	nh.n = 0
	fmt.Fprintf(w, "<div class=\"entry\"><span id=\"dfdf\">No Argo CD Notifications yet</span></div>")
}
