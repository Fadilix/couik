package network

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net"
	"sync"
)

type Server struct {
	Clients map[net.Conn]string
	Mu      sync.Mutex
}

func NewServer() *Server {
	return &Server{
		Clients: make(map[net.Conn]string),
	}
}

func (s *Server) Start(port string) {
	listener, err := net.Listen("tcp", ":"+port)

	defer func() {
		log.Println("server has closed")
		if err := listener.Close(); err != nil {
			log.Println("Error while closing the server ", err)
		}
	}()

	if err != nil {
		log.Println("Error while listening to port : ", port, err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error while accepting new client", err)
		}

		go s.HandleJoin(conn)
	}
}

func (s *Server) HandleJoin(conn net.Conn) {
	defer func() {
		s.Mu.Lock()
		delete(s.Clients, conn)
		s.Mu.Unlock()
		if err := conn.Close(); err != nil {
			log.Println("Error while disconneting the client", err)
		}
	}()

	slog.Info("A new user entered the chat")
	s.Mu.Lock()
	s.Clients[conn] = "randomname"
	s.Mu.Unlock()

	decoder := json.NewDecoder(conn)

	for {
		var msg Message
		if err := decoder.Decode(&msg); err != nil {
			log.Println("Client disconnected on error", err)
			break
		}
		s.HandleMessage(conn, msg)
	}
}

func (s *Server) HandleMessage(sender net.Conn, msg Message) {
	switch msg.Type {
	case MsgJoin:
		var payload JoinPayload
		if err := json.Unmarshal(msg.Payload, &payload); err == nil {
			s.Mu.Lock()
			s.Clients[sender] = payload.PlayerName

			for conn, name := range s.Clients {
				if conn == sender {
					continue
				}
				p, _ := json.Marshal(JoinPayload{PlayerName: name})
				m := Message{Type: MsgJoin, Payload: p}
				data, _ := json.Marshal(m)
				fmt.Fprintln(sender, string(data))
			}
			s.Mu.Unlock()
		}
		s.Broadcast(msg)
	case MsgStart:
		s.Broadcast(msg)
	case MsgUpdate:
		s.Broadcast(msg)
	default:
		log.Printf("Unknown message %s\n", msg.Type)
	}
}

func (s *Server) Broadcast(msg Message) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	data, _ := json.Marshal(msg)
	for conn := range s.Clients {
		if _, err := fmt.Fprintln(conn, string(data)); err != nil {
			log.Println(err)
		}
	}
}
