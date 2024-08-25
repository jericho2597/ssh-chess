package main

import (
	"log"
	"net/http"
	"ssh-server/config"
)

func setupHTTPServer() {

	handler := &sshHandler{addr: config.SshServerAddress, keyfile: config.SshClientKeyPath}

	http.Handle("/", http.FileServer(http.Dir("./web/")))
	http.HandleFunc("/web-socket/ssh", handler.webSocket)
}

func startHTTPServer() {
	log.Printf("Starting HTTP server on %s", config.HttpServerAddress)
	if err := http.ListenAndServe(config.HttpServerAddress, nil); err != nil {
		log.Fatalf("Error starting HTTP server: %s", err)
	}
}

type sshHandler struct {
	addr    string
	keyfile string
}

// webSocket handles WebSocket requests for SSH from the clients.
func (h *sshHandler) webSocket(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("upgrader.Upgrade:", err)
		return
	}

	sshCli := &sshClient{
		conn:     conn,
		addr:     h.addr,
		keyfile:  h.keyfile,
		closeSig: make(chan struct{}, 1),
	}
	go sshCli.bridgeWSAndSSH()
}
