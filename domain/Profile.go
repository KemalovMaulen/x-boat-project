package domain

type Profile struct {
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	FirstName string `json:"first_name,omitempty" bson:"first_name,omitempty"`
	LastName string `json:"last_name,omitempty" bson:"last_name"`
	LivingAddress *LivingAddress `json:"living_address,omitempty"`
	Image string `json:"image,omitempty" bson:"image"`
	IsAdmin *bool `json:"is_admin,omitempty" bson:"is_admin"`
	IsVerified *bool `json:"is_verified,omitempty" bson:"is_verified"`
	IsActive *bool `json:"is_active,omitempty" bson:"is_active"`
}

type LivingAddress struct {
	Street string `json:"living_address_street,omitempty" bson:"living_address_street"`
	StreetName string `json:"living_address_street_name,omitempty" bson:"living_address_street_name"`
	PostalCode string `json:"living_address_postal_code,omitempty" bson:"living_address_postal_code"`
	City string `json:"living_address_city,omitempty" bson:"living_address_city"`
	Country string `json:"living_address_country,omitempty" bson:"living_address_country"`
}
