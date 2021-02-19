// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	flag.Parse()
	hubs := make([]*Hub, 0)
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		handleWs(w, r, &hubs)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleWs(w http.ResponseWriter, r * http.Request, h *[]*Hub) {
	removeEmptyHubs(h)
	success := assignExistingHub(w, r, h)
	if !success {
		assignNewHub(w, r, h)
	}
}

func assignNewHub(w http.ResponseWriter, r *http.Request, h *[]*Hub) {
	hub := newHub()
	*h = append(*h, hub)

	go hub.run()
	serveWs(hub, w, r)
}

func assignExistingHub(w http.ResponseWriter, r *http.Request, h *[]*Hub) bool {
	assigned := false
	for _, hub := range *h {
		if !hub.started && len(hub.clients) < 5 {
			serveWs(hub, w, r)
			assigned = true
			break
		}
	}
	return assigned
}

func removeEmptyHubs(h *[]*Hub) {
	nextHub := make([]*Hub, 0)
	for _, hub := range *h {
		if len(hub.clients) == 0 {
			continue
		}

		nextHub = append(nextHub, hub)
	}
	h = &nextHub
}