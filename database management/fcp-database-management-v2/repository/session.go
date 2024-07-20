package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	result := s.db.Create(&session)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	session := model.Session{}
	result := s.db.Where("token = ?", token).Delete(&session)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	result := s.db.Table("sessions").Where("username = ?", session.Username).Updates(map[string]interface{}{
		"token":  session.Token,
		"expiry": session.Expiry,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	var sessions model.Session
	result := s.db.Where("username = ?", name).First(&sessions)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	var sessions model.Session
	result := s.db.Where("token = ?", token).First(&sessions)
	if result.Error != nil {
		return model.Session{}, result.Error
	}

	return sessions, nil
}
