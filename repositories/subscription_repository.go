package repositories
//
//import (
//	"github.com/salambayev/x-boat-project/db"
//	"github.com/salambayev/x-boat-project/domain"
//	"gopkg.in/mgo.v2/bson"
//)
//
//type SubscriptionRepository interface {
//	CreateSubscription(email string,
//		subscriptionType domain.SubscriptionType,
//		beginTime int64,
//		endTime int64) error
//
//	ChangeSubscriptionType(email string,
//		subscriptionType domain.SubscriptionType) error
//
//	DeactivateSubscription(email string) error
//
//	ActivateSubscription(email string, subscriptionType domain.SubscriptionType, beginTime int64,
//		endTime int64) error
//
//	GetSubscription(email string) (*domain.Subscription, error)
//}
//
//type subscriptionRepository struct {
//
//}
//
//func NewSubscriptionRepository() SubscriptionRepository {
//	return &subscriptionRepository{}
//}
//
//func (sr *subscriptionRepository) CreateSubscription(email string,
//	subscriptionType domain.SubscriptionType,
//	beginTime int64,
//	endTime int64) error {
//		return db.SubscriptionCollection.Insert(&domain.Subscription{Email:email, Type:subscriptionType,
//		BeginTime: beginTime, EndTime: endTime, IsActive: true})
//}
//
//func (sr *subscriptionRepository) ChangeSubscriptionType(email string,
//	subscriptionType domain.SubscriptionType) error {
//		return db.SubscriptionCollection.Update(bson.M{"email": email},
//			bson.M{"$set": bson.M{"type": subscriptionType}})
//}
//
//func (sr *subscriptionRepository) DeactivateSubscription(email string) error {
//	return db.SubscriptionCollection.Update(bson.M{"email": email}, bson.M{"$set": bson.M{ "is_active": false }})
//}
//
//func (sr *subscriptionRepository) ActivateSubscription(email string,
//	subscriptionType domain.SubscriptionType, beginTime int64,
//	endTime int64) error {
//		return db.SubscriptionCollection.Update(bson.M{"email": email},
//		bson.M{ "$set": bson.M{ "type": subscriptionType, "begin_time": beginTime, "end_time": endTime, "is_active": true }})
//}
//
//func (sr *subscriptionRepository) GetSubscription(email string) (*domain.Subscription, error) {
//	result := domain.Subscription{}
//	err := db.SubscriptionCollection.Find(bson.M{"email": email}).One(&result)
//	return &result, err
//}