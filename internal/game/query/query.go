package query

const (
	QueryCreateGame = `
		insert into games (room_id, index, language, from_title, to_title)
			values ($1, $2, $3, $4, $5)
			returning id, room_id, index, is_finished, language, from_title, to_title, created_at, updated_at; 
	`
	QueryGetLatestGame = `
		select id, room_id, index, is_finished, language, from_title, to_title, created_at, updated_at
			from games
			where 
				room_id = $1
			order by index desc
			limit 1;
	`
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
