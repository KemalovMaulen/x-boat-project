package repositories

import (
	"cloud.google.com/go/firestore"
	"github.com/salambayev/x-boat-project/domain"
	"golang.org/x/net/context"
	"github.com/salambayev/x-boat-project/db"
	"google.golang.org/api/iterator"
	"fmt"
	"github.com/salambayev/x-boat-project/utils"
)

type ClubRepository interface {

	CreateClub(club *domain.Club) error

	UpdateClub(clubCode string, club *domain.Club) error

	DeleteClub(code string) error

	GetClub(code string) (*domain.Club, error)

	GetAllClubs() ([]*domain.Club, error)
}

type clubRepository struct {
}

func NewClubRepository() ClubRepository {
	return &clubRepository{}
}

func (cr *clubRepository) CreateClub(club *domain.Club) error {
	//_, err := db.ClubCollection.Doc(club.ClubId).Set(context.Background(), club)
	_, err := db.ClubCollection.Doc(club.ClubId).Create(context.Background(), club)
	return err
}

func (cr *clubRepository) UpdateClub(clubCode string, club *domain.Club) error {
	fmt.Printf("club %+v", club)
	fireMap, err := utils.GetMap(club)
	if err != nil {
		return err
	}
	_, err = db.MembershipCollection.Doc(clubCode).Set(context.Background(), fireMap, firestore.MergeAll)
	return err
}

func (cr *clubRepository) DeleteClub(code string) error {
	isActive := false
	_, err := db.ClubCollection.Doc(code).Set(context.Background(), &domain.Club{IsActive: &isActive}, firestore.Merge([]string{"is_active"}))
	return err
}

func (cr *clubRepository) GetClub(code string) (*domain.Club, error) {
	dsnap, err := db.ClubCollection.Doc(code).Get(context.Background())
	if err != nil {
		return nil, err
	}
	result := &domain.Club{}
	err = dsnap.DataTo(result)
	return result, err
}

func (cr *clubRepository) GetAllClubs() ([]*domain.Club, error) {
	var result []*domain.Club
	iter := db.ClubCollection.Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		club := &domain.Club{}
		doc.DataTo(club)
		result = append(result, club)
	}
	return result, nil
}