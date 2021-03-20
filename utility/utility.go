package utility

import (
	"crypto/rand"
	"html"
	"io"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"bitbucket.org/cloud-platform/vnpt-sso-authentication/initial"
)

// GenerateUserID - GenerateUserID
func GenerateUserID(prefix, collection string) string {
	uid := strings.ReplaceAll(uuid.New().String(), "-", "")
	temUserID := prefix + uid
	if tempRecord := mgodb.FindOneByField(initial.MgoDBName, collection, "user_id", temUserID); tempRecord != nil {
		uid := strings.ReplaceAll(uuid.New().String(), "-", "")
		temUserID = prefix + uid
		if tempRecord := mgodb.FindOneByField(initial.MgoDBName, collection, "user_id", temUserID); tempRecord != nil {
			return ""
		}
	}

	return temUserID
}

// GenerateRandomWithDigit - GenerateRandomWithDigit
func GenerateRandomWithDigit(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// Hash - Hash
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

// CheckPasswordHash - CheckPasswordHash
func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Santize - Santize
func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
