package socket

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/lib/pq"
)

var upgrader = websocket.Upgrader{}

type Socket struct {
	sockets map[string]Connection
	DB      *sql.DB
}

func (s *Socket) SetUpSocket(c *gin.Context) {
	type SocketConnID struct {
		ConnID string `form:"conn_id"`
	}
	var socketId SocketConnID

	if c.ShouldBindQuery(&socketId) != nil {
		c.JSON(500, gin.H{
			"response": "Invalid Query Params",
		})
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ws.Close()

	connection := &Connection{ws: ws}

	if s.sockets == nil {
		s.sockets = make(map[string]Connection)
	}

	s.sockets[socketId.ConnID] = *connection

	// why use go multi threading ? because it is blocking for loop
	connection.readMessages(s)

}

func (s *Socket) GetConnectionFromId(receiverId string) []Connection {
	var connections []Connection
	recTuple := strings.Split(receiverId, "_")

	if recTuple[0] == "p" {
		connections = append(connections, s.sockets[recTuple[1]])
		return connections
	}

	getGroupUsersQuery := "SELECT users FROM groups WHERE id=$1"
	var users []int64

	if err := s.DB.QueryRow(getGroupUsersQuery, recTuple[1]).Scan(pq.Array(&users)); err != nil {
		fmt.Println("query", err)
		log.Fatal(err)
	}

	for _, user := range users {
		stringifiedUserId := strconv.FormatInt(user, 10)
		connections = append(connections, s.sockets[stringifiedUserId])
	}

	return connections
}
