package domain

type SubscriptionType int

const (
	Monthly SubscriptionType = iota
)

type Subscription struct {
	Email     string           `json:"email,omitempty" bson:"email" firestore:"email,omitempty"`
	Type      SubscriptionType `json:"type,omitempty" bson:"type" firestore:"type,omitempty"`
	BeginTime int64            `json:"begin_time,omitempty" bson:"begin_time" firestore:"begin_time,omitempty"`
	EndTime   int64            `json:"end_time,omitempty" bson:"end_time" firestore:"end_time,omitempty"`
	IsActive  *bool            `json:"is_active,omitempty" bson:"is_active" firestore:"is_active,omitempty"`
}
