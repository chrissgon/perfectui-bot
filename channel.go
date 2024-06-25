package main

import (
	"net/http"

	"github.com/chrissgon/lowbot"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type IServerChannel struct {
	*lowbot.Channel
	conn map[string]*websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func NewServerChannel() (lowbot.IChannel, error) {
	return &IServerChannel{
		Channel: &lowbot.Channel{
			ChannelID: uuid.New(),
			Name:      CHANNEL_SERVER_NAME,
		},

		conn: map[string]*websocket.Conn{},
	}, nil
}

func (channel *IServerChannel) GetChannel() *lowbot.Channel {
	return channel.Channel
}

func (channel *IServerChannel) Next(interaction chan *lowbot.Interaction) {
	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		var err error

		id := r.URL.Query().Get("id")

		if id == "" {
			id = uuid.NewString()
		}

		conn, exists := channel.conn[id]

		if !exists {
			conn, err = upgrader.Upgrade(w, r, nil)
		}

		if err != nil {
			return
		}

		channel.conn[id] = conn
		
		for {
			_, message, err := conn.ReadMessage()
			
			if err != nil {
				return
			}
			
			sender := lowbot.NewWho(id, "user")
			
			interaction <- lowbot.NewInteractionMessageText(channel, sender, string(message))
		}
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	http.ListenAndServe(":8090", nil)
}

func (channel *IServerChannel) SendAudio(*lowbot.Interaction) error {
	panic("unimplemented")
}

func (channel *IServerChannel) SendButton(*lowbot.Interaction) error {
	panic("unimplemented")
}

func (channel *IServerChannel) SendDocument(*lowbot.Interaction) error {
	panic("unimplemented")
}

func (channel *IServerChannel) SendImage(*lowbot.Interaction) error {
	panic("unimplemented")
}

func (channel *IServerChannel) SendText(interaction *lowbot.Interaction) error {
	sessionID := interaction.Sender.WhoID
	conn, exists := channel.conn[sessionID]

	if !exists {
		return ERR_UNKNOWN_CHANNEL_CONN
	}

	return conn.WriteMessage(1, []byte(interaction.Parameters.Text))
}

func (channel *IServerChannel) SendVideo(*lowbot.Interaction) error {
	panic("unimplemented")
}
