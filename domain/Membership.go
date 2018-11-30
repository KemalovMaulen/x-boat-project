package domain

type Membership struct {
	Id         string   `json:"id,omitempty" bson:"_id,omitempty" firestore:"id,omitempty"`
	Profile    *Profile `json:"profile,omitempty" bson:"profile" firestore:"profile,omitempty" `
	IsStuff    bool     `json:"is_stuff,omitempty" bson:"is_stuff" firestore:"is_stuff,omitempty"`
	BoatDriver string   `json:"boat_driver,omitempty" bson:"boat_driver" firestore:"boat_driver,omitempty"`
	Club       *Club    `json:"club,omitempty" bson:"club" firestore:"club,omitempty"`
}
