package query

const (
	// args: `code`, `status`
	QueryInsertRoom = `
		insert into rooms (code, status) values ($1, $2)
			returning id, code, status, created_at, updated_at;
	`

	// args: `room_id`, `user_id`, `is_owner`
	QueryInsertRoomMember = `
		insert into room_members (room_id, user_uid, is_owner) values ($1, $2, $3)
			returning id, is_owner, room_id, user_uid, created_at, updated_at;
	`

	// args: `room_id`, `user_uid`
	QueryDeleteRoomMember = `
		delete from room_members
			where room_id = $1 and user_uid = $2;
	`

	// args: `code`
	QueryGetRoomByCode = `
		select id, code, status, created_at, updated_at
			from rooms
			where code = $1;
	`

	// args: `id`
	QueryGetRoomByID = `
		select id, code, status, created_at, updated_at
			from rooms
			where id = $1;
	`

	// args: `id`
	QueryGetRoomMembers = `
		select 
				rm.id as id, 
				rm.room_id as room_id,
				rm.user_uid as user_uid,
				coalesce(u.username, '') as username,
				rm.created_at as created_at, 
				rm.updated_at as updated_at
			from rooms as r
			inner join room_members as rm
				on rm.room_id = r.id
			left join usernames as u
				on u.uid = rm.user_uid
			where r.id = $1;
	`

	// args: `user_uid`
	QueryGetRoomBelongToMember = `
		select 
				r.id as id, 
				r.code as code, 
				r.status as status, 
				r.created_at as created_at, 
				r.updated_at as updated_at
			from rooms as r
			inner join room_members as rm
				on rm.room_id = r.id
			where 
				rm.user_uid = $1 and
				r.status = 'open';
	`

	// args: `id`, `status`
	QueryUpdateRoomstatus = `
		update rooms
			set 
				status = $2,
				updated_at = current_timestamp
			where id = $1
			returning id, code, status, created_at, updated_at;
	`

	// args: `id`, `user_uid`
	QueryDelete = `
		delete from rooms
			where id = $1 and user_uid = $2;
	`
)
