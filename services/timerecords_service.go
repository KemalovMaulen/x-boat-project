package services

import (
	"github.com/salambayev/x-boat-project/domain"
	"github.com/salambayev/x-boat-project/repositories"
)

type TimerecordsService interface {
	CreateTimerecord(timerecord *domain.Timerecord) error

	UpdateTimerecord(timestamp int64, timerecord *domain.Timerecord) error

	DeleteTimerecord(timestamp int64) error

	GetTimerecord(timestamp int64) (*domain.Timerecord, error)

	GetAllTimerecords() ([]*domain.Timerecord, error)
}

type timerecordsService struct {
	repo repositories.TimerecordsRepository
}

func NewTimerecordsService(repo repositories.TimerecordsRepository) TimerecordsService {
	return &timerecordsService{ repo }
}

func (ts *timerecordsService) CreateTimerecord(timerecord *domain.Timerecord) error {
	fullTimerecord := &domain.Timerecord{ timerecord.Timestamp,
	timerecord.Email, timerecord.StartTime,
	timerecord.EndTime, timerecord.EndTime - timerecord.StartTime,
	timerecord.BoatDriver}
	return ts.repo.CreateTimerecord(fullTimerecord)
}

func (ts *timerecordsService) UpdateTimerecord(timestamp int64, timerecord *domain.Timerecord) error {
	return ts.repo.UpdateTimerecord(timestamp, timerecord)
}

func (ts *timerecordsService) DeleteTimerecord(timestamp int64) error {
	return ts.repo.DeleteTimerecord(timestamp)
}

func (ts *timerecordsService) GetTimerecord(timestamp int64) (*domain.Timerecord, error) {
	return ts.repo.GetTimerecord(timestamp)
}

func (ts *timerecordsService) GetAllTimerecords() ([]*domain.Timerecord, error) {
	return ts.repo.GetAllTimerecords()
}