package query

const (
	QueryInsertUsername = `
		insert into usernames (uid, username) values ($1, $2); 
	`
	QueryFindUsernameByUID = `
		select uid, username from usernames 
			where uid = $1;
	`
)
