package services

import (
	"github.com/salambayev/x-boat-project/domain"
	"github.com/salambayev/x-boat-project/repositories"
	//"gopkg.in/mgo.v2/bson"
)

type MembershipService interface {
	CreateMembership(membership *domain.Membership) error

	UpdateMembership(id string, membership *domain.Membership) error

	GetUserMemberships(email string) ([]*domain.Membership, error)

	GetMembershipById(id string) (*domain.Membership, error)

	DeleteMembership(id string) error

	GetClubMembers(clubId string) ([]*domain.Profile, error)
}

type membershipService struct {
	repo repositories.MembershipRepository
}

func NewMembershipService(repo repositories.MembershipRepository) MembershipService {
	return &membershipService{ repo }
}

func (ms *membershipService) CreateMembership(membership *domain.Membership) error {
	return ms.repo.CreateMembership(membership)
}

func (ms *membershipService) UpdateMembership(id string, membership *domain.Membership) error {
	return ms.repo.UpdateMembership(id, membership)
}

func (ms *membershipService) GetUserMemberships(email string) ([]*domain.Membership, error) {
	return ms.repo.GetUserMemberships(email)
}

func (ms *membershipService) GetMembershipById(id string) (*domain.Membership, error) {
	return ms.repo.GetMembershipById(id)
}

func (ms *membershipService) DeleteMembership(id string) error {
	return ms.repo.DeleteMembership(id)
}

func (ms *membershipService) GetClubMembers(clubId string) ([]*domain.Profile, error) {
	return ms.repo.GetClubMembers(clubId)
}