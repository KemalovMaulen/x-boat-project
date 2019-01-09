package domain

type Profile struct {
	Uid			  string         `json:"uid,omitempty" bson:"uid,omitempty" firestore:"uid,omitempty"`
	Email         string         `json:"email,omitempty" bson:"email,omitempty" firestore:"email,omitempty"`
	Password      string         `json:"password,omitempty" bson:"password,omitempty" firestore:"password,omitempty"`
	FirstName     string         `json:"first_name,omitempty" bson:"first_name,omitempty" firestore:"first_name,omitempty"`
	LastName      string         `json:"last_name,omitempty" bson:"last_name" firestore:"last_name,omitempty"`
	LivingAddress *LivingAddress `json:"living_address,omitempty" firestore:"living_address,omitempty"`
	Image         string         `json:"image,omitempty" bson:"image" firestore:"image,omitempty"`
	IsAdmin       *bool          `json:"is_admin,omitempty" bson:"is_admin" firestore:"is_admin,omitempty"`
	IsVerified    *bool          `json:"is_verified,omitempty" bson:"is_verified" firestore:"is_verified,omitempty"`
	IsActive      *bool          `json:"is_active,omitempty" bson:"is_active" firestore:"is_active,omitempty"`
}

type LivingAddress struct {
	Street     string `json:"living_address_street,omitempty" bson:"living_address_street" firestore:"living_address_street,omitempty"`
	StreetName string `json:"living_address_street_name,omitempty" bson:"living_address_street_name" firestore:"living_address_street_name,omitempty" `
	PostalCode string `json:"living_address_postal_code,omitempty" bson:"living_address_postal_code" firestore:"living_address_postal_code,omitempty"`
	City       string `json:"living_address_city,omitempty" bson:"living_address_city" firestore:"living_address_city,omitempty"`
	Country    string `json:"living_address_country,omitempty" bson:"living_address_country" firestore:"living_address_country,omitempty"`
}
