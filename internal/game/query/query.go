package query

const (
	// args: `room_id`, `index`, `language`, `from_title`, `to_title`
	QueryCreateGame = `
		insert into games (room_id, index, language, from_title, to_title)
			values ($1, $2, $3, $4, $5)
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
