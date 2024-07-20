package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type ClassRepository interface {
	FetchAll() ([]model.Class, error)
}

type classRepoImpl struct {
	db *gorm.DB
}

func NewClassRepo(db *gorm.DB) *classRepoImpl {
	return &classRepoImpl{db}
}

func (s *classRepoImpl) FetchAll() ([]model.Class, error) {
	rows, err := s.db.Table("classes").Select("*").Rows()
	if err != nil {
		return nil, err
	}

	var listClass []model.Class
	for rows.Next() {
		s.db.ScanRows(rows, &listClass)
	}

	return listClass, nil
}
