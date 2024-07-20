package repo

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TeacherRepo struct {
	db *gorm.DB
}

func NewTeacherRepo(db *gorm.DB) TeacherRepo {
	return TeacherRepo{db}
}

func (t TeacherRepo) Save(data model.Teacher) error {
	result := t.db.Create(&data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t TeacherRepo) Query() ([]model.Teacher, error) {
	rows, err := t.db.Table("teachers").Select("*").Rows()
	if err != nil {
		return nil, err
	}

	var listTeacher []model.Teacher
	for rows.Next() {
		t.db.ScanRows(rows, &listTeacher)
	}

	return listTeacher, nil
}

func (t TeacherRepo) Update(id uint, name string) error {
	result := t.db.Table("teachers").Where("id = ?", id).Update("name", name)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t TeacherRepo) Delete(id uint) error {
	teacher := model.Teacher{}
	result := t.db.Where("id = ?", id).Delete(&teacher)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
