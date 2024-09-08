package model

import (
	known "github.com/costa92/krm/internal/pkg/known/usercenter"
	"github.com/costa92/krm/internal/pkg/zid"
	"github.com/costa92/krm/pkg/authn"
	"gorm.io/gorm"
)

// AfterCreate hook
func (u *UserM) AfterCreate(tx *gorm.DB) (err error) {
	u.UserID = zid.User.New(uint64(u.ID)) // Generate and set a new user ID.
	return tx.Save(u).Error               // Save the updated user record to the database.
}

// BeforeCreate runs before creating a UserM database record and initializes various fields.
func (u *UserM) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, err = authn.Encrypt(u.Password) // Encrypt the user password.
	if err != nil {
		return err // Return error if there's a problem with encryption.
	}

	u.Status = known.UserStatusActived // Set the default status for the user as active.

	return nil
}
