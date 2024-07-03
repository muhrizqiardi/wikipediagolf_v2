package signup

const (
	QueryInsertUsername = `
		insert into usernames (uid, username) values ($1, $2); 
	`
)
