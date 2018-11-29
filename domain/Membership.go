package domain

type Membership struct {
	Id         string `json:"id,omitempty" bson:"_id,omitempty"`
	Profile    *Profile      `json:"profile,omitempty" bson:"profile"`
	IsStuff    bool          `json:"is_stuff,omitempty" bson:"is_stuff"`
	BoatDriver string        `json:"boat_driver,omitempty" bson:"boat_driver"`
	Club       *Club         `json:"club,omitempty" bson:"club"`
}
