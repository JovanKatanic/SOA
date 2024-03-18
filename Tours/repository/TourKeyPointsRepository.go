package repository

import "gorm.io/gorm"

type TourKeyPointsRepository struct {
	DatabaseConnection *gorm.DB
}
