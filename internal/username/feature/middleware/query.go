package middleware

const (
	QueryFindUsernameByUID = `
		select uid, username from usernames 
			where uid = $1;
	`
)
