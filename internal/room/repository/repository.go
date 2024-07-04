package repository

type Repository interface {
	InsertRoom()
	InsertRoomMember()
	DeleteRoomMember()
	GetRoomByCode()
	GetRoomByID()
	GetRoomMembers()
	GetRoomBelongToMember()
	UpdateRoomStatus()
	Delete()
}
