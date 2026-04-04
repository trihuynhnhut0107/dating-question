package services

import (
	"dating-question/config"
	"dating-question/models"
)

func CreateSession(name *string, clientIp *string) (*models.Session, error) {
	var tempName string

	if name != nil {
		tempName = *name
	}

	if clientIp != nil {
		tempName += *clientIp
	}

	if tempName == "" {
		tempName = "Anonymous Session"
	}
	session := &models.Session{
		Name: &tempName,
	}

	if err := config.DB.Create(session).Error; err != nil {
		return nil, err
	}

	return session, nil
}