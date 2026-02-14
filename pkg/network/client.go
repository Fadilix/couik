package network

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

type Client struct {
	conn net.Conn
	enc  *json.Encoder
	mu   sync.Mutex

	msgs chan Message
	errs chan error
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) NextMessage() <-chan Message {
	return c.msgs
}

func (c *Client) Errors() <-chan error {
	return c.errs
}

func (c *Client) ReadLoop() {
	defer c.conn.Close()

	decoder := json.NewDecoder(c.conn)

	for {
		var msg Message

		if err := decoder.Decode(&msg); err != nil {
			c.errs <- fmt.Errorf("connection error : %w", err)
			return
		}

		c.msgs <- msg
	}
}

func (c *Client) Send(m Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.enc.Encode(m)
}

func (c *Client) SendJoin(name string) error {
	payload, _ := json.Marshal(JoinPayload{PlayerName: name})
	return c.Send(Message{
		Type:    MsgJoin,
		Payload: payload,
	})
}

func (c *Client) SendUpdate(name string, wpm int, progress float64) error {
	payload, _ := json.Marshal(UpdatePayload{
		PlayerName: name,
		WPM:        wpm,
		Progress:   progress,
	})

	return c.Send(Message{
		Type:    MsgUpdate,
		Payload: payload,
	})
}

func NewClient(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server at %s: %w", addr, err)
	}

	c := &Client{
		conn: conn,
		enc:  json.NewEncoder(conn),
		msgs: make(chan Message, 100),
		errs: make(chan error, 1),
	}

	go c.ReadLoop()
	return c, nil
}
