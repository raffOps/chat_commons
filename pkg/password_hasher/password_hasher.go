package password_hasher

//go:generate mockery --name PasswordHasher
type PasswordHasher interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
