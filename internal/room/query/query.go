package query

const (
	queryInsertRoom = `
		insert into rooms (code, state) values ($1, $2)
			returning id, code, state, created_at, updated_at;
	`
	queryInsertRoomMember = `
		insert into room_members (room_id, user_uid, is_owner) values ($1, $2, $3)
			returning id, is_owner, room_id, user_uid, created_at, updated_at;
	`
	queryDeleteRoomMember = `
		delete from room_members
			where room_id = $1 and user_uid = $2;
	`
	queryGetRoomByCode = `
		select id, code, state, created_at, updated_at
			from rooms
			where code = $1;
	`
	queryGetRoomByID = `
		select id, code, state, created_at, updated_at
			from rooms
			where id = $1;
	`
	queryGetRoomMembers = `
		select 
				rm.id as id, 
				rm.owner_id as owner_id,
				rm.room_id as room_id,
				rm.user_uid as user_uid,
				rm.created_at as created_at, 
				rm.updated_at as updated_at
			from rooms as r
			inner join room_members as rm
				on rm.room_id = r.id
			where r.id = $1;
	`
	queryGetRoomBelongToMember = `
		select 
				r.id as id, 
				r.code as code, 
				r.state as state, 
				r.created_at as created_at, 
				r.updated_at as updated_at
			from rooms as r
			inner join room_members as rm
				on rm.room_id = r.id
			where rm.user_uid = $1;
	`
	queryUpdateRoomState = `
		update rooms
			set 
				status = $2,
				updated_at = current_timestamp
			where id = $1
			returning id, code, state, created_at, updated_at;
	`
	queryDelete = `
		delete from rooms
			where id = $1 and user_uid = $2;
	`
)
