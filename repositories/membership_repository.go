package repositories

import (
	"github.com/salambayev/x-boat-project/domain"
	"github.com/salambayev/x-boat-project/db"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	//"cloud.google.com/go/firestore"
	//"encoding/json"
)

type MembershipRepository interface {
	CreateMembership(membership *domain.Membership) error

	UpdateMembership(id string, membership *domain.Membership) error

	GetUserMemberships(email string) ([]*domain.Membership, error)

	GetMembershipById(id string) (*domain.Membership, error)

	DeleteMembership(id string) error

	GetClubMembers(clubId string) ([]*domain.Profile, error)
}

type membershipRepository struct {
}

func NewMembershipRepository() MembershipRepository {
	return &membershipRepository{}
}

func (mr *membershipRepository) CreateMembership(membership *domain.Membership) error {
	//var fireMap map[string]interface{}
	//tempByte, err := json.Marshal(membership)
	//if err != nil {
	//	return err
	//}
	//err = json.Unmarshal(tempByte, &fireMap)
	//if err != nil {
	//	return err
	//}
	_, err := db.MembershipCollection.Doc(membership.Id).Create(context.Background(), membership)

	return err
}

func (mr membershipRepository) UpdateMembership(id string, membership *domain.Membership) error {

	//var fireMap map[string]interface{}
	//tempByte, err := json.Marshal(membership)
	//if err != nil {
	//	return err
	//}
	//err = json.Unmarshal(tempByte, &fireMap)
	//if err != nil {
	//	return err
	//}
	//_, err = db.MembershipCollection.Doc(id).Set(context.Background(), fireMap, firestore.MergeAll)

	_, err := db.MembershipCollection.Doc(id).Set(context.Background(), membership)
	return err
}

func (mr *membershipRepository) GetUserMemberships(email string) ([]*domain.Membership, error) {
	result := []*domain.Membership{}
	iter := db.MembershipCollection.Where("Profile.Email", "==" , email).Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		membership := &domain.Membership{}
		err = doc.DataTo(membership)
		if err != nil {
			return nil, err
		}
		result = append(result, membership )
		fmt.Println( doc.Data() )
	}
	return result, nil
}

func (mr *membershipRepository) GetMembershipById(id string) (*domain.Membership, error) {
	dsnap, err := db.MembershipCollection.Doc(id).Get(context.Background())
	if err != nil {
		return nil, err
	}
	result := &domain.Membership{}
	err = dsnap.DataTo(result)
	return result, err
}

func (mr *membershipRepository) DeleteMembership(id string) error {
	_, err := db.MembershipCollection.Doc(id).Delete(context.Background())
	return err
}

func (mr *membershipRepository) GetClubMembers(clubId string) ([]*domain.Profile, error) {

	result := []*domain.Profile{}
	docSnapArr, err := db.MembershipCollection.Where("Club.ClubId", "==" , clubId).Documents(context.Background()).GetAll()
	if err != nil {
		return nil, err
	}
	for i := range docSnapArr {
		membership := &domain.Membership{}
		err = docSnapArr[i].DataTo(membership)
		if err != nil {
			return nil, err
		}
		result = append(result, membership.Profile )
		fmt.Println( docSnapArr[i].Data() )
	}

	return result, nil
	//var queryResult []*domain.Membership
	//err := db.MembershipCollection.Find(bson.M{"club._id": clubId}).All(&queryResult)
	//if err != nil {
	//	return nil, err
	//}
	//var result []*domain.Profile
	//for _, v := range queryResult {
	//	result = append(result, v.Profile)
	//}
	return result, nil
}
