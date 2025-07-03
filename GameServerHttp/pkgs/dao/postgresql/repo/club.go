package pgsql_repo

import (
	"gorm.io/gorm"
)

type clubDBRepo struct {
	db *gorm.DB
}

type ClubDBRepo interface {
}

func NewClubDBRepo(db *gorm.DB) ClubDBRepo {
	return &clubDBRepo{
		db: db,
	}
}
