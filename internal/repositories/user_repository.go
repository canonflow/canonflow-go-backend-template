package repositories

import (
	"canonflow-golang-backend-template/internal/models/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[domain.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *Repository[T]) FindByUsername(db *gorm.DB, entity *T, username string) error {
	return db.Where("username = ?", username).Take(entity).Error
}
