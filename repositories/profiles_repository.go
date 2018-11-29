package repositories
//
//import (
//	"github.com/salambayev/x-boat-project/domain"
//	"github.com/salambayev/x-boat-project/db"
//	"gopkg.in/mgo.v2/bson"
//)
//
//type ProfilesRepository interface {
//	CreateProfile(profile *domain.Profile) error
//
//	UpdateProfile(profile *domain.Profile) error
//
//	DeleteProfile(email string) error
//
//	GetProfile(email string) (*domain.Profile, error)
//}
//
//type profileRepository struct {
//}
//
//func NewProfileRepository() ProfilesRepository {
//	return &profileRepository{}
//}
//
//func (pr *profileRepository) CreateProfile(profile *domain.Profile) error {
//	return db.ProfilesCollection.Insert(&profile)
//}
//
//func (pr *profileRepository) UpdateProfile(profile *domain.Profile) error {
//	return db.ProfilesCollection.Update(bson.M{"email": profile.Email}, &profile)
//}
//
//func (pr *profileRepository) DeleteProfile(email string) error {
//	return db.ProfilesCollection.Remove(bson.M{"email": email})
//}
//
//func (pr *profileRepository) GetProfile(email string) (*domain.Profile, error) {
//	result := domain.Profile{}
//	err := db.ProfilesCollection.Find(bson.M{"email": email}).One(&result)
//	return &result, err
//}