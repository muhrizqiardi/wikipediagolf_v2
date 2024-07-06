package join

type JoinRequest struct {
	DisplayName string `schema:"displayName"`
	RoomCode    string `schema:"roomCode"`
}
