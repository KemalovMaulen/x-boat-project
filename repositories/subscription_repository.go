package repositories

import (
	"github.com/salambayev/x-boat-project/db"
	"github.com/salambayev/x-boat-project/domain"
	//"gopkg.in/mgo.v2/bson"
	"context"
	"cloud.google.com/go/firestore"
	"github.com/salambayev/x-boat-project/utils"
)

type SubscriptionRepository interface {
	CreateSubscription(email string,
		subscriptionType domain.SubscriptionType,
		beginTime int64,
		endTime int64) error

	ChangeSubscriptionType(email string,
		subscriptionType domain.SubscriptionType) error

	DeactivateSubscription(email string) error

	ActivateSubscription(email string, subscriptionType domain.SubscriptionType, beginTime int64,
		endTime int64) error

	GetSubscription(email string) (*domain.Subscription, error)
}

type subscriptionRepository struct {

}

func NewSubscriptionRepository() SubscriptionRepository {
	return &subscriptionRepository{}
}

func (sr *subscriptionRepository) CreateSubscription(email string, subscriptionType domain.SubscriptionType, beginTime int64,
	endTime int64) error {

	isActive := true
	subscription := &domain.Subscription{email, subscriptionType, beginTime, endTime, &isActive}
	_, err := db.SubscriptionCollection.Doc(email).Create(context.Background(), subscription)
	return err
}

func (sr *subscriptionRepository) ChangeSubscriptionType(email string, subscriptionType domain.SubscriptionType) error {

	subscription := &domain.Subscription{Email: email, Type:subscriptionType}
	_, err := db.SubscriptionCollection.Doc(email).Set(context.Background(), subscription ,firestore.Merge([]string{"email","type"})) //, firestore.MergeAll)
	return err
}

func (sr *subscriptionRepository) DeactivateSubscription(email string) error {
	isActive := false
	_, err := db.SubscriptionCollection.Doc(email).Set(context.Background(), &domain.Subscription{IsActive: &isActive}, firestore.Merge([]string{"is_active"}))
	return err
}

func (sr *subscriptionRepository) ActivateSubscription(email string, subscriptionType domain.SubscriptionType, beginTime int64,
	endTime int64) error {
	isActive := true
	subscription := &domain.Subscription{IsActive: &isActive, Type: subscriptionType, BeginTime: beginTime, EndTime: endTime}
	fireMap, err := utils.GetMap(subscription)
	if err != nil {
		return err
	}
	_, err = db.SubscriptionCollection.Doc(email).Set(context.Background(), fireMap, firestore.MergeAll)
//		firestore.Merge([]string{"is_active","begin_time", "type", "end_time"})) //TODO normolize it
	return err
}

func (sr *subscriptionRepository) GetSubscription(email string) (*domain.Subscription, error) {
	dsnap, err := db.SubscriptionCollection.Doc(email).Get(context.Background())
	if err != nil {
		return nil, err
	}
	result := &domain.Subscription{}
	err = dsnap.DataTo(result)
	return result, err
}