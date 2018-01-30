package data

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func UserByEmail(email string) (user User, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, password, created_at FROM users WHERE email=$1", email)
	if err != nil {
		return
	}

	for rows.Next() {

		err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Password, &user.CreatedAt)
		if err != nil {
			return
		}

	}
	rows.Close()
	return
}

func Encrypt(input string, salt *string) (encrypted string) {

	rand.Seed(time.Now().Unix())
	randInt := rand.Int()
	*salt = strconv.Itoa(randInt)
	fmt.Println(*salt)

	hasher := sha1.New()
	io.WriteString(hasher, input)
	io.WriteString(hasher, *salt)
	fmt.Println(hasher)
	fmt.Println(hasher.Sum(nil))
	return
}
