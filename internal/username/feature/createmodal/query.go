package createmodal

const (
	QueryFindUsernameByUID = `
		select uid, username from usernames 
			where uid = $1;
	`
)
