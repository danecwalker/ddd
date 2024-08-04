package account

func HashPassword(password string) string {
	return password
}

func VerifyPassword(hashedPassword, password string) bool {
	return hashedPassword == password
}
