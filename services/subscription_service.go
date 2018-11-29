package services
//
//import (
//	"github.com/salambayev/x-boat-project/domain"
//	"github.com/salambayev/x-boat-project/repositories"
//)
//
//type SubscriptionService interface {
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
//type subscriptionService struct {
//	repo repositories.SubscriptionRepository
//}
//
//func NewSubscriptionService(repo repositories.SubscriptionRepository) SubscriptionService {
//	return &subscriptionService{ repo }
//}
//
//func (ss *subscriptionService) CreateSubscription(email string,
//	subscriptionType domain.SubscriptionType,
//	beginTime int64,
//	endTime int64) error {
//		return ss.repo.CreateSubscription(email, subscriptionType, beginTime, endTime)
//}
//
//func (ss *subscriptionService) ChangeSubscriptionType(email string,
//	subscriptionType domain.SubscriptionType) error {
//		return ss.repo.ChangeSubscriptionType(email, subscriptionType)
//}
//
//func (ss *subscriptionService) DeactivateSubscription(email string) error {
//	return ss.repo.DeactivateSubscription(email)
//}
//
//func (ss *subscriptionService) ActivateSubscription(email string, subscriptionType domain.SubscriptionType, beginTime int64,
//	endTime int64) error {
//		return ss.repo.ActivateSubscription(email, subscriptionType, beginTime, endTime)
//}
//
//func (ss *subscriptionService) GetSubscription(email string) (*domain.Subscription, error) {
//	return ss.repo.GetSubscription(email)
//}
