package database

import (
	"go-uaa/src/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserPasswordResetDbRepository struct {
	db *gorm.DB
}

func (repo *UserPasswordResetDbRepository) Save(userPasswordReset user.UserPasswordReset) error {
	result := repo.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&userPasswordReset)
	return result.Error
}

func (repo *UserPasswordResetDbRepository) FindByUserID(userID uuid.UUID) (*user.UserPasswordReset, error) {
	var foundReset user.UserPasswordReset
	result := repo.db.Where(user.UserPasswordReset{UserID: userID}).First(&foundReset)
	if result.Error != nil {
		return nil, result.Error
	}
	return &foundReset, nil
}

func (repo *UserPasswordResetDbRepository) Delete(userPasswordReset user.UserPasswordReset) error {
	result := repo.db.Delete(&userPasswordReset)
	return result.Error
}

func NewUserPasswordResetDbRepository(db *gorm.DB) *UserPasswordResetDbRepository {
	return &UserPasswordResetDbRepository{
		db: db,
	}
}
