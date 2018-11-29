package services

import (
	"github.com/salambayev/x-boat-project/domain"
	"github.com/salambayev/x-boat-project/repositories"
)

type ClubService interface {
	CreateClub(club *domain.Club) error

	UpdateClub(clubCode string, club *domain.Club) error

	DeleteClub(code string) error

	GetClub(code string) (*domain.Club, error)

	GetAllClubs() ([]*domain.Club, error)
}

type clubService struct {
	repo repositories.ClubRepository
}

func NewClubService(repo repositories.ClubRepository) ClubService {
	return &clubService{ repo }
}

func (cs *clubService) CreateClub(club *domain.Club) error  {
	return cs.repo.CreateClub(club)
}

func (cs *clubService) UpdateClub(clubCode string, club *domain.Club) error {
	return cs.repo.UpdateClub(clubCode, club)
}

func (cs *clubService) DeleteClub(code string) error {
	return cs.repo.DeleteClub(code)
}

func (cs *clubService) GetClub(code string) (*domain.Club, error) {
	return cs.repo.GetClub(code)
}

func (cs *clubService) GetAllClubs() ([]*domain.Club, error) {
	return cs.repo.GetAllClubs()
}


