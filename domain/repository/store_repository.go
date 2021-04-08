package repository

import "multilingual_gurunavi_api/domain/models"

type StoreRepository interface {
	GnaviRequest(string) (*models.Store, error)
	GnaviRequestCh(string, chan []models.Response)
}
