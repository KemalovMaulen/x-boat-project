package domain

type SubscriptionType int

const (
	Monthly SubscriptionType = iota
)


type Subscription struct {
	Email string `json:"email,omitempty" bson:"email"`
	Type SubscriptionType `json:"type,omitempty" bson:"type"`
	BeginTime int64 `json:"begin_time,omitempty" bson:"begin_time"`
	EndTime int64 `json:"end_time,omitempty" bson:"end_time"`
	IsActive *bool `json:"is_active,omitempty" bson:"is_active"`
}
