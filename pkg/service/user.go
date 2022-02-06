package service

import (
	"grpc-streams/pkg/pb"
)

type Connection struct {
	stream pb.MyService_GetStreamServer
	error  chan error
}

type UserId string
type User struct {
	id         UserId
	chunk      Chunk
	connection Connection
}

var (
	users map[UserId]*User
)

func init() {
	users = make(map[UserId]*User)
}

func FindUser(stream *pb.MyService_GetStreamServer) *User {
	userId := UserId((*stream).Context().Value("user").(string))
	if _, ok := users[userId]; !ok {
		conn := Connection{stream: *stream}
		user := User{
			id:         userId,
			connection: conn,
		}
		users[userId] = &user
	}
	return users[userId]
}

func (u *User) Logout() {
	for _, world := range worlds {
		world.Unsubscribe(u)
	}
	delete(users, u.id)
}
