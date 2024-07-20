package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	rows, err := s.db.Table("students").Select("*").Rows()
	if err != nil {
		return nil, err
	}

	var listStudent []model.Student
	for rows.Next() {
		s.db.ScanRows(rows, &listStudent)
	}

	return listStudent, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	result := s.db.Create(&student)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	result := s.db.Table("students").Where("id = ?", id).Updates(map[string]interface{}{
		"name":     student.Name,
		"address":  student.Address,
		"class_id": student.ClassId,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *studentRepoImpl) Delete(id int) error {
	student := model.Student{}
	result := s.db.Where("id = ?", id).Delete(&student)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var students model.Student
	result := s.db.Where("id = ?", id).First(&students)
	if result.Error != nil {
		return &model.Student{}, result.Error
	}

	return &students, nil
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	result := []model.StudentClass{}

	s.db.Table("students").
		Select("students.name, students.address, classes.name AS class_name, classes.professor AS professor, classes.room_number AS room_number").
		Joins("inner join classes on classes.id = students.class_id").
		Scan(&result)
	return &result, nil
}
