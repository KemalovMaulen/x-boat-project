package services
//
//import (
//	"github.com/salambayev/x-boat-project/domain"
//	"github.com/salambayev/x-boat-project/repositories"
//)
//
//type ProfilesService interface {
//	CreateProfile(profile *domain.Profile) error
//
//	UpdateProfile(profile *domain.Profile) error
//
//	DeleteProfile(email string) error
//
//	GetProfile(email string) (*domain.Profile, error)
//}
//
//type profileService struct {
//	repo repositories.ProfilesRepository
//}
//
//func NewProfileService(repo repositories.ProfilesRepository) ProfilesService {
//	return &profileService{ repo }
//}
//
//func (ps *profileService) CreateProfile(profile *domain.Profile) error {
//	return ps.repo.CreateProfile(profile)
//}
//
//func (ps *profileService) UpdateProfile(profile *domain.Profile) error {
//	return ps.repo.UpdateProfile(profile)
//}
//
//func (ps *profileService) DeleteProfile(email string) error {
//	return ps.repo.DeleteProfile(email)
//}
//
//func (ps *profileService) GetProfile(email string) (*domain.Profile, error) {
//	return ps.repo.GetProfile(email)
//}