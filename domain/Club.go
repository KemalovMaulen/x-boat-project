package domain

type Club struct {
	ClubId   string `json:"id,omitempty" bson:"_id,omitempty" firestore:"id,omitempty"`
	Name     string `json:"name,omitempty" bson:"owner" firestore:"name,omitempty"`
	Owner    string `json:"owner,omitempty" bson:"owner" firestore:"owner,omitempty"`
	IsActive *bool  `json:"is_active,omitempty" bson:"is_active" firestore:"is_active,omitempty"`
}
