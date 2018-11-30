package repositories

import (
	"github.com/salambayev/x-boat-project/domain"
	"github.com/salambayev/x-boat-project/db"
	"context"
	"github.com/salambayev/x-boat-project/utils"
	"cloud.google.com/go/firestore"
)

type ProfilesRepository interface {
	CreateProfile(profile *domain.Profile) error

	UpdateProfile(profile *domain.Profile) error

	DeleteProfile(email string) error

	GetProfile(email string) (*domain.Profile, error)
}

type profileRepository struct {
}

func NewProfileRepository() ProfilesRepository {
	return &profileRepository{}
}

func (pr *profileRepository) CreateProfile(profile *domain.Profile) error {
	_, err := db.ProfilesCollection.Doc(profile.Email).Create(context.Background(), profile)
	return err
}

func (pr *profileRepository) UpdateProfile(profile *domain.Profile) error {
	fireMap, err := utils.GetMap(profile)
	if err != nil {
		return err
	}
	_, err = db.ProfilesCollection.Doc(profile.Email).Set(context.Background(), fireMap, firestore.MergeAll)
	return err
}

func (pr *profileRepository) DeleteProfile(email string) error {
	_, err := db.ProfilesCollection.Doc(email).Delete(context.Background())
	return err
}

func (pr *profileRepository) GetProfile(email string) (*domain.Profile, error) {
	dsnap, err := db.ProfilesCollection.Doc(email).Get(context.Background())
	if err != nil {
		return nil, err
	}
	result := &domain.Profile{}
	err = dsnap.DataTo(result)
	return result, err
}