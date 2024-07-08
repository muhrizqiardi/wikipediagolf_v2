package query

const (
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
	// args: `room_id`, `index`, `language`, `from_title`, `to_title`
	QueryCreateGame = `
		insert into games (room_id, index, language, from_title, to_title, is_finished)
			values ($1, $2, $3, $4, $5, false)
			returning id, room_id, index, is_finished, language, from_title, to_title, created_at, updated_at; 
	`
	// args = `room_id`
	QueryGetLatestGame = `
		select id, room_id, index, is_finished, language, from_title, to_title, created_at, updated_at
			from games
			where 
				room_id = $1
			order by index desc
			limit 1;
	`
	// args = `id`, `room_id`, `is_finished`
	QueryUpdateGame = `
		update games
			set
				is_finished = $3,
				updated_at = current_timestamp
			where
				id = $1,
				room_id = $2;
	`
)
