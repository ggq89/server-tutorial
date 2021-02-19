// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "strconv"

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from the clients.
	unregister chan *Client

	started	bool

	broadcast chan ClientMsg
}

type ClientMsg struct {
	client	*Client
	msg		[]byte
}

func newHub() *Hub {
	return &Hub {
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		started:    false,
		broadcast:  make(chan ClientMsg),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <- h.register:
			h.registerClient(client)
		case client := <-h.unregister:
			h.unregisterClient(client)
		case clmsg := <-h.broadcast:
			h.broadcastMsg(clmsg)
		}
	}
}

func (h *Hub) registerClient(client *Client) {
	h.clients[client] = true
	if len(h.clients) == 0 {
		client.isRoomOwner = true
	}
}

func (h *Hub) unregisterClient(client *Client) {
	if _, ok := h.clients[client]; ok {

		delete(h.clients, client)

		if client.isRoomOwner {
			for c, _ := range h.clients {
				if c != client {
					c.isRoomOwner = true
					break
				}
			}
		}

		close(client.send)
	}
}

func (h *Hub) broadcastMsg(clmsg ClientMsg) {
	// if message is from room owner and start is the
	// message given, when the room status is not started,
	// start the game
	cl, message := clmsg.client, clmsg.msg

	for client := range h.clients {
		select {
		case client.send <- []byte("Client " + strconv.Itoa(cl.clientNo) + ")"):
			if cl.isRoomOwner {
				client.send <- []byte("<<message from roomOwner>>")
			}
			client.send <- message
		default:
			close(client.send)
			delete(h.clients, client)
		}
	}
}

