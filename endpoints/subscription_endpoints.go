package endpoints

import (
	"github.com/salambayev/x-boat-project/server"
	"net/http"
	"github.com/salambayev/x-boat-project/domain"
	"github.com/salambayev/x-boat-project/utils"
	"strings"
	"fmt"
)

type SubscriptionEndpointFactory struct {

}

func NewSubscriptionEndpointFactory() *SubscriptionEndpointFactory {
	return &SubscriptionEndpointFactory{}
}

var errSubSys = &server.ErrorSystem{"x-boat-subscription"}

func (fac *SubscriptionEndpointFactory) MakeCreateSubscriptionEndpoint(services *server.Services) server.HttpEndpoint {

	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		subscription := &domain.Subscription{}
		err := utils.ParseJSON(r, subscription)
		if err != nil {
			return errSubSys.BadRequest(1, "subscription body error ", err.Error())
		}
		fmt.Println("MakeCreateSubscriptionEndpoint email = ", subscription.Email)
		if strings.TrimSpace(subscription.Email) == "" {
			return errSubSys.BadRequest(2, "subscription with no email coudl not be created")
		}
		err = services.Subscriptions.CreateSubscription(subscription.Email, subscription.Type, subscription.BeginTime, subscription.EndTime)
		if err != nil {
			return errSubSys.InternalServerError(3, "subscription already exists", err.Error())
		}
		return server.Created(subscription)
	}
}

func (fac * SubscriptionEndpointFactory) MakeChangeSubscriptionTypeEndpoint(services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		subscription := &domain.Subscription{}
		err := utils.ParseJSON(r, subscription)
		if err != nil {
			return errSubSys.BadRequest(4, "subscription body error ", err.Error())
		}
		fmt.Println("MakeChangeSubscriptionTypeEndpoint email = ", subscription.Email)
		if strings.TrimSpace(subscription.Email) == "" {
			return errSubSys.BadRequest(5, "No email")
		}
		err = services.Subscriptions.ChangeSubscriptionType(subscription.Email, subscription.Type)
		if err != nil {
			return errSubSys.InternalServerError(6, "Could Not found subscription with email = " + subscription.Email, err.Error())
		}
		return server.OK(subscription)
	}
}

func (fac *SubscriptionEndpointFactory) MakeDeactivateSubscriptionEndpoint(services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		subscription := &domain.Subscription{}
		err := utils.ParseJSON(r, subscription)
		if err != nil {
			return errSubSys.BadRequest(7, "subscription body error ", err.Error())
		}
		fmt.Println("MakeDeactivateSubscriptionEndpoint email = ", subscription.Email)
		if strings.TrimSpace(subscription.Email) == "" {
			return errSubSys.BadRequest(8 ,"No email", err.Error())
		}
		err = services.Subscriptions.DeactivateSubscription(subscription.Email)
		if err != nil {
			return errSubSys.BadRequest(9, "Could Not found subscription with email = " + subscription.Email, err.Error())
		}
		return server.OK("ok")
	}
}

func (fac *SubscriptionEndpointFactory) MakeActivateSubscriptionEndpoint(services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		subscription := &domain.Subscription{}
		err := utils.ParseJSON(r, subscription)
		if err != nil {
			return errSubSys.BadRequest(10, "subscription body error ", err.Error())
		}
		fmt.Println("MakeActivateSubscriptionEndpoint email = ", subscription.Email)
		if strings.TrimSpace(subscription.Email) == "" {
			return errSubSys.BadRequest(11, "No email = " )
		}
		if subscription.BeginTime == 0 && subscription.EndTime == 0 { //TODO

			return errSubSys.BadRequest(12, "subscription time is 0" )
		}
		err = services.Subscriptions.ActivateSubscription(subscription.Email, subscription.Type, subscription.BeginTime, subscription.EndTime)
		if err != nil {
			return errSubSys.BadRequest(13, "Could Not found subscription with email = " + subscription.Email, err.Error())
		}
		return server.OK("ok")
	}
}

func (fac *SubscriptionEndpointFactory) MakeGetSubscriptionEndpoint(emailParam string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		email := strings.TrimSpace(r.URL.Query().Get(emailParam))
		fmt.Println("MakeGetSubscriptionEndpoint email = ", email)
		if email == "" {
			return errSubSys.BadRequest(4, "No email")
		}
		d, err := services.Subscriptions.GetSubscription(email)
		if err != nil {
			return errSubSys.BadRequest(4, "Could Not found profile with email = " + email, err.Error())
		}
		return server.OK(d)
	}
}

