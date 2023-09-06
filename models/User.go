package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null;" json:"name"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null;" json:"password"`
}

type APIUser struct {
	gorm.Model
	Name  string
	Email string
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) beforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		// if err := checkmail.ValidateFormat(u.Email); err != nil {
		// 	return errors.New("Invalid Email")
		// }

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		// if err := checkmail.ValidateFormat(u.Email); err != nil {
		// 	return errors.New("Invalid Email")
		// }
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		// if err := checkmail.ValidateFormat(u.Email); err != nil {
		// 	return errors.New("Invalid Email")
		// }
		return nil
	}
}

func (u *User) SaveUser() (*User, error) {

	err := u.beforeSave()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("user:: ", u)
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	u.Password = ""
	return u, nil
}

func (u *User) FindAllUsers(page int, limit int) (*[]APIUser, error) {
	var err error
	users := []APIUser{}
	err = DB.Model(&User{}).Offset((page - 1) * limit).Limit(limit).Find(&users).Error
	if err != nil {
		return &[]APIUser{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(uid uint32) (*APIUser, error) {
	var err error
	user := APIUser{}
	err = DB.Model(User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &APIUser{}, err
	}
	return &user, err
}

func (u *User) FindUserByEmail(email string) (*User, error) {
	var err error
	user := User{}
	err = DB.Model(User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	return &user, err
}

func (u *User) UpdateAUser(uid uint32) (*User, error) {
	// To hash the password
	err := u.beforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db := DB.Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":   u.Password,
			"name":       u.Name,
			"email":      u.Email,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(uid uint32) (int64, error) {

	db := DB.Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (u *User) Count() (int64, error) {
	var count int64
	err := DB.Model(&User{}).Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}
