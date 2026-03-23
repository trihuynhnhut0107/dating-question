package service

import (
	"errors"
	"time"

	"dating-question-backend/internal/models"

	"gorm.io/gorm"
)

type CreateSessionInput struct {
	Name        string            `json:"name"`
	DrawCount   int               `json:"draw_count"`
	CategoryID  *uint             `json:"category_id"`
	Difficulty  *models.Difficulty `json:"difficulty"`
}

type SessionService struct {
	db *gorm.DB
}

func NewSessionService(db *gorm.DB) *SessionService {
	return &SessionService{db: db}
}

func (s *SessionService) GetAll() ([]models.Session, error) {
	var sessions []models.Session
	err := s.db.Find(&sessions).Error
	return sessions, err
}

// findByUUID resolves a UUID to a Session (no preloads – internal use only).
func (s *SessionService) findByUUID(uuid string) (*models.Session, error) {
	var session models.Session
	err := s.db.Where("uuid = ?", uuid).First(&session).Error
	return &session, err
}

func (s *SessionService) GetByUUID(uuid string) (*models.Session, error) {
	var session models.Session
	err := s.db.
		Preload("SessionQuestions.Question.Category").
		Where("uuid = ?", uuid).
		First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *SessionService) Create(input CreateSessionInput) (*models.Session, error) {
	session := models.Session{Name: input.Name}
	if err := s.db.Create(&session).Error; err != nil {
		return nil, err
	}

	drawCount := input.DrawCount
	if drawCount <= 0 {
		drawCount = 5
	}

	var questions []models.Question
	q := s.db.Where("is_active = ?", true)
	if input.CategoryID != nil {
		q = q.Where("category_id = ?", *input.CategoryID)
	}
	if input.Difficulty != nil {
		q = q.Where("difficulty = ?", *input.Difficulty)
	}
	if err := q.Order("RANDOM()").Limit(drawCount).Find(&questions).Error; err != nil {
		return nil, err
	}

	for _, question := range questions {
		sq := models.SessionQuestion{
			SessionID:  session.ID,
			QuestionID: question.ID,
		}
		if err := s.db.Create(&sq).Error; err != nil {
			return nil, err
		}
	}

	return s.GetByUUID(session.UUID)
}

func (s *SessionService) DrawMore(uuid string, count int, categoryID *uint, difficulty *models.Difficulty) (*models.Session, error) {
	session, err := s.findByUUID(uuid)
	if err != nil {
		return nil, err
	}
	sessionID := session.ID

	var alreadyDrawn []uint
	s.db.Model(&models.SessionQuestion{}).
		Where("session_id = ?", sessionID).
		Pluck("question_id", &alreadyDrawn)

	q := s.db.Where("is_active = ? AND id NOT IN ?", true, alreadyDrawn)
	if categoryID != nil {
		q = q.Where("category_id = ?", *categoryID)
	}
	if difficulty != nil {
		q = q.Where("difficulty = ?", *difficulty)
	}

	var questions []models.Question
	if err := q.Order("RANDOM()").Limit(count).Find(&questions).Error; err != nil {
		return nil, err
	}
	if len(questions) == 0 {
		return nil, errors.New("no more questions available to draw")
	}

	for _, question := range questions {
		sq := models.SessionQuestion{
			SessionID:  sessionID,
			QuestionID: question.ID,
		}
		if err := s.db.Create(&sq).Error; err != nil {
			return nil, err
		}
	}

	return s.GetByUUID(uuid)
}

func (s *SessionService) AnswerQuestion(uuid string, sqID uint) (*models.SessionQuestion, error) {
	session, err := s.findByUUID(uuid)
	if err != nil {
		return nil, err
	}
	var sq models.SessionQuestion
	if err := s.db.Where("id = ? AND session_id = ?", sqID, session.ID).First(&sq).Error; err != nil {
		return nil, err
	}
	now := time.Now()
	sq.IsAnswered = true
	sq.AnsweredAt = &now
	err = s.db.Save(&sq).Error
	return &sq, err
}

func (s *SessionService) Delete(uuid string) error {
	session, err := s.findByUUID(uuid)
	if err != nil {
		return err
	}
	if err := s.db.Where("session_id = ?", session.ID).Delete(&models.SessionQuestion{}).Error; err != nil {
		return err
	}
	return s.db.Delete(&models.Session{}, session.ID).Error
}
