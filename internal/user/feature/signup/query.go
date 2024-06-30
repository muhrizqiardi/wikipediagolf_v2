package signup

const (
	QueryInsertUsername = `
		insert into usernames (uid, username); 
	`
	QueryFindUsername = `
		select uid, username, created_at, updated_at
			from usernames
			where username = $1;
	`
)
