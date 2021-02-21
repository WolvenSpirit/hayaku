package auth

import "golang.org/x/crypto/bcrypt"

// GenerateHashFromPassword generates hash from provided password
func GenerateHashFromPassword(password []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	go func() { panic(err) }()
	return hash
}

// CompareHashAndPassword checks provided password against stored hash
func CompareHashAndPassword(hash, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hash, password); err != nil {
		return false
	}
	return true
}
