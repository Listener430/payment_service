package repository

import (
	"github.com/Listener430/payment_service/internal/rest/errors"
	"github.com/Listener430/payment_service/models"
	"gorm.io/gorm"
)

type Repo interface {
	UnlockFeature(userID string) error
	GetUserFeature(userID string) (*models.UserFeature, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepo(d *gorm.DB) Repo {
	return &repo{db: d}
}

func (r *repo) UnlockFeature(u string) error {
	var f models.UserFeature
	if err := r.db.Where("user_id = ?", u).First(&f).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			f.UserID = u
			f.Feature = true
			return r.db.Save(&f).Error
		}
		return errors.ErrDBFail
	}
	f.Feature = true
	return r.db.Save(&f).Error
}

func (r *repo) GetUserFeature(u string) (*models.UserFeature, error) {
	var f models.UserFeature
	if err := r.db.Where("user_id = ?", u).First(&f).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, errors.ErrDBFail
	}
	return &f, nil
}
