package types

import (
	"log"

	pg "github.com/go-pg/pg/v10"
	"golang.org/x/crypto/bcrypt"
)

//Create create a new user and insert into db
func (u *User) Create(db *pg.DB) error {
	_, err := db.Model(u).Insert()
	return err
}

//Update update existing user
func (u *User) Update(db *pg.DB) error {
	_, err := db.Model(u).Update()
	return err
}

// CheckUserExists to check if user is present
func (u *User) CheckUserExists(db *pg.DB) (bool, error) {
	user := new(User)
	err := db.Model(user).Table("users").Where("users.email = ?", u.Email).Limit(1).Select()
	encryptionErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if encryptionErr == nil {
		u.FirstName = user.FirstName
		u.LastName = user.LastName
		return true, nil
	}
	if err != nil {
		log.Println(err)
	}
	return false, err
}

//Delete delete user
func (u *User) Delete(db *pg.DB) error {
	_, err := db.Model(u).Delete()
	return err
}
