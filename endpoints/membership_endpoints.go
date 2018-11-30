package endpoints

import (
	"github.com/salambayev/x-boat-project/server"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/salambayev/x-boat-project/domain"
	"github.com/salambayev/x-boat-project/utils"
	"fmt"
)

type MembershipEndpointFactory struct {

}

func NewMembershipEndpointFactory() *MembershipEndpointFactory {
	return &MembershipEndpointFactory{}
}

var errMemSys = &server.ErrorSystem{"x-boat-membership"}

func (fac *MembershipEndpointFactory) MakeCreateMembershipEndpoint(services *server.Services) server.HttpEndpoint {

	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		membership := &domain.Membership{}
		err := utils.ParseJSON(r, membership)
		if err != nil {
			return errMemSys.BadRequest(1, "No membership body")
		}
		if membership.Id == "" {
			membership.Id = utils.GenerateId()
		}
		err = services.Memberships.CreateMembership(membership)
		if err != nil {
			return errMemSys.InternalServerError(2,"Membership already exists",  err.Error())
		}
		return server.Created(membership)
	}
}

func (fac *MembershipEndpointFactory) MakeUpdateMembershipEndpoint(id string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		id, ok := vars[id]
		if !ok {
			return errMemSys.InternalServerError(3,"No id")
		}
		fmt.Println("MakeUpdateMembershipEndpoint id = ", id)
		membership := &domain.Membership{}
		err := utils.ParseJSON(r, membership)
		if err != nil {
			return errMemSys.BadRequest(4, "No Membership", err.Error())
		}
		membership.Id = id
		err = services.Memberships.UpdateMembership(id, membership)
		if err != nil {
			return errMemSys.InternalServerError(5, "Could Not found Membership with id = " + id,  err.Error())
		}
		return server.OK(membership)
	}
}

func (fac *MembershipEndpointFactory) MakeGetMembershipByIdEndpoint(id string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		id, ok := vars[id]
		if !ok {
			return errMemSys.BadRequest(6,"No id")
		}
		fmt.Println("MakeGetMembershipByIdEndpoint id = ", id)
		d, err := services.Memberships.GetMembershipById(id)
		if err != nil {
			return errMemSys.InternalServerError(7, "Could Not found Membership with id = " + id,  err.Error())
		}
		return server.OK(d)
	}
}

func (fac * MembershipEndpointFactory) MakeGetUserMembershipsEndpoint(email string, services *server.Services) server.HttpEndpoint {

	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		email, ok := vars[email]
		if !ok {
			return errMemSys.BadRequest(8,"No Id")
		}
		d, err := services.Memberships.GetUserMemberships(email)
		if err != nil {
			return errMemSys.InternalServerError(9, "Could Not found Membershipd with email = " + email,  err.Error())
		}
		return server.OK(d)
	}
}

func (fac * MembershipEndpointFactory) MakeGetClubMembersEndpoint(clubId string, services *server.Services) server.HttpEndpoint {

	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		clubId, ok := vars[clubId]
		if !ok {
			return errMemSys.BadRequest(10,"No id")
		}
		fmt.Println("MakeGetClubMembersEndpoint id = ", clubId)
		d, err := services.Memberships.GetClubMembers(clubId)
		if err != nil {
			return errMemSys.InternalServerError(11, "Could Not found Membershipd with clubId = " + clubId,  err.Error())
		}
		return server.OK(d)
	}
}

func (fac * MembershipEndpointFactory) MakeDeleteMembershipEndpoint(id string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		id, ok := vars[id]
		if !ok {
			return errMemSys.BadRequest(12,"No Id")
		}
		fmt.Println("MakeDeleteMembershipEndpoint id = ", id)

		err := services.Memberships.DeleteMembership(id)
		if err != nil {
			return errMemSys.InternalServerError(12, "Could Not found deleted membership with id = " + id,  err.Error())
		}

		return server.OK("ok")
	}
}