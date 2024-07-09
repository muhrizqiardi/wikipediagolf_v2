package create

type Repository interface {
	Insert(uid, username string) error
}
