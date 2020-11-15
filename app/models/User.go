package models

import (
	"errors"
	"time"

	helper "github.com/somkiet073/Golang-api-for-testskill/app/helpers"

	"github.com/jinzhu/gorm"
)

// User = user
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Fristname string    `gorm:"type:varchar(255);" json:"fristname"`
	Lastname  string    `gorm:"type:varchar(255);" json:"lastname"`
	Nickname  string    `gorm:"type:varchar(255);" json:"nickname"`
	Email     string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password  string    `gorm:"type:varchar(100);" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// CreateUser = createUser
func (u *User) CreateUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

// FindAllUser = findAllUser
func (u *User) FindAllUser(db *gorm.DB) (*[]User, error) {
	var err error
	var users []User
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, nil
}

// FindUserByID = findUserByID
func (u *User) FindUserByID(db *gorm.DB, uid uint64) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id=?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("RecordNotFound")
	}
	return u, nil
}

// UpdateUser = updateUser
func (u *User) UpdateUser(db *gorm.DB, uid uint64) (*User, error) {
	var err error
	hashedpassword, err := helper.Hash(u.Password)
	if err != nil {
		return &User{}, err
	}

	// encypt password
	u.Password = string(hashedpassword)

	// map data
	dataUpdate := map[string]interface{}{
		"fristname":  u.Fristname,
		"lastname":   u.Lastname,
		"nickname":   u.Nickname,
		"email":      u.Email,
		"password":   u.Password,
		"updated_at": time.Now(),
	}

	// update data
	db = db.Debug().Model(&User{}).Where("id=?", uid).Take(&User{}).Update(dataUpdate)
	if db.Error != nil {
		return &User{}, db.Error
	}

	err = db.Debug().Model(&User{}).Where("id=?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, nil
}

// DeleteUser = deleteUser
func (u *User) DeleteUser(db *gorm.DB, uid uint64) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id=?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
