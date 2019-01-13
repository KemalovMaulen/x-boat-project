package endpoints

import (
	"github.com/salambayev/x-boat-project/server"
	"net/http"
	"github.com/salambayev/x-boat-project/domain"
	"github.com/salambayev/x-boat-project/utils"
	"strings"
	"fmt"
	"github.com/salambayev/x-boat-project/auth"
)

type ProfileEndpointFactory struct {

}

func NewProfileEndpointFactory() *ProfileEndpointFactory {
	return &ProfileEndpointFactory{}
}

var errProSys = &server.ErrorSystem{"x-boat-profile"}

func (fac *ProfileEndpointFactory) MakeCreateProfileEndpoint(services *server.Services) server.HttpEndpoint {

	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		profile := &domain.Profile{}
		err := utils.ParseJSON(r, profile)
		if err != nil {
			return errProSys.BadRequest(1, "Profile body error ", err.Error())
		}
		fmt.Println("MakeCreateProfileEndpoint email = ", profile.Email)
		profile.Uid = utils.GenerateId()
		userRecord , err := (&auth.FRAuth{}).SignUp(w, r, profile)
		if err != nil {
			return errProSys.BadRequest(11, "Profile SignUp ", err.Error())
		}
		fmt.Printf("userRecord: %+v\n", userRecord)
		fmt.Printf("userRecord.UserMetadata: %+v\n", userRecord.UserMetadata)
		fmt.Printf("userRecord.UserInfo: %+v\n", userRecord.UserInfo)
		token, err :=  (&auth.FRAuth{}).GetToken(w, r, profile)
		if err != nil {
			return errProSys.BadRequest(12, "Profile GetToken ", err.Error())
		}
		fmt.Println("token: ", token)
		if strings.TrimSpace(profile.Email) == "" {
			return errProSys.BadRequest(2, "Profile with no email coudl not be created")
		}
		err = services.Profiles.CreateProfile(profile)
		if err != nil {
			return errProSys.InternalServerError(3, "Profile already exists", err.Error())
		}
		return server.Created(profile)
	}
}

func (fac * ProfileEndpointFactory) MakeUpdateProfileEndpoint(services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		_ , err := (&auth.FRAuth{}).ValidateToken(w, r)
		if err != nil {
			return errProSys.BadRequest(12, "Profile ValidateToken: ", err.Error())
		}

		profile := &domain.Profile{}
		err = utils.ParseJSON(r, profile)
		if err != nil {
			return errProSys.BadRequest(4, "No profile", err.Error())
		}
		fmt.Println("MakeUpdateProfileEndpoint email = ", profile.Email)
		if strings.TrimSpace(profile.Email) == "" {
			return errProSys.BadRequest(5, "No email", err.Error())
		}
		err = services.Profiles.UpdateProfile(profile)
		if err != nil {
			return errProSys.InternalServerError(6, "Could Not found profile with email = " + profile.Email,  err.Error())
		}
		return server.OK(profile)
	}
}

func (fac *ProfileEndpointFactory) MakeGetProfileEndpoint(emailParam string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		_ , err := (&auth.FRAuth{}).ValidateToken(w, r)
		if err != nil {
			return errProSys.BadRequest(12, "Profile ValidateToken ", err.Error())
		}

		email := strings.TrimSpace(r.URL.Query().Get(emailParam))
		fmt.Println("MakeGetProfileEndpoint email = ", email)
		if email == "" {
			return errProSys.BadRequest(7,"No email")
		}
		d, err := services.Profiles.GetProfile(email)
		if err != nil {
			return errProSys.InternalServerError(8, "Could Not found profile with email = " + email, err.Error())
		}
		return server.OK(d)
	}
}

func (fac *ProfileEndpointFactory) MakeDeleteProfileEndpoint(emailParam string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {

		_ , err := (&auth.FRAuth{}).ValidateToken(w, r)
		if err != nil {
			return errProSys.BadRequest(12, "Profile ValidateToken ", err.Error())
		}

		email := strings.TrimSpace(r.URL.Query().Get(emailParam))
		fmt.Println("MakeDeleteProfileEndpoint email = ", email)
		if email == "" {
			return errProSys.BadRequest(9,"No email")
		}
		err = services.Profiles.DeleteProfile(email)
		if err != nil {
			return errProSys.InternalServerError(10, "Could Not found profile with email = " +email, err.Error())
		}
		return server.OK("ok")
	}
}
