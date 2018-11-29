package server

import "github.com/salambayev/x-boat-project/services"

type Services struct {
	Clubs services.ClubService
	Memberships services.MembershipService
	//Profiles services.ProfilesService
	//Timerecords services.TimerecordsService
	//Subscriptions services.SubscriptionService
}
