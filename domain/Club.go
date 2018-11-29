package domain

type Club struct {
	ClubId string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"owner"`
	Owner string `json:"owner,omitempty" bson:"owner"`
	IsActive *bool `json:"is_active,omitempty" bson:"is_active"`
}
