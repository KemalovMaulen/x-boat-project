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


func (fac *MembershipEndpointFactory) MakeCreateMembershipEndpoint(services *server.Services) server.HttpEndpoint {

	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		membership := &domain.Membership{}
		err := utils.ParseJSON(r, membership)
		if err != nil {
			return &server.Response{Status: http.StatusBadRequest,
				Data:       server.Error{400, "2", "No membership body", ""},
				HeaderData: make(map[string]string) }
			//return errSys.BadRequest(110, err.Error())
		}
		err = services.Memberships.CreateMembership(membership)
		if err != nil {
			return &server.Response{Status: http.StatusInternalServerError,
				Data:       server.Error{500, "3", "Membership already exists",  err.Error()},
				HeaderData: make(map[string]string) }
			//return errSys.InternalServerError(119, err.Error())
		}
		return server.Created(membership)
	}
}

func (fac *MembershipEndpointFactory) MakeUpdateMembershipEndpoint(id string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		id, ok := vars[id]
		if !ok {
			return &server.Response{Status: http.StatusBadRequest,
				Data:       server.Error{400, "2", "No Id", ""},
				HeaderData: make(map[string]string) }
			//return errSys.BadRequest(123," ")
		}
		fmt.Println("MakeUpdateMembershipEndpoint id = ", id)
		membership := &domain.Membership{}
		err := utils.ParseJSON(r, membership)
		if err != nil {
			return &server.Response{Status: http.StatusBadRequest,
				Data:       server.Error{400, "2", "No Membership", err.Error()},
				HeaderData: make(map[string]string) }
			//return errSys.BadRequest(110, err.Error())
		}
		membership.Id = id
		err = services.Memberships.UpdateMembership(id, membership)
		if err != nil {
			return &server.Response{Status: http.StatusInternalServerError,
				Data:       server.Error{500, "3", "Could Not found Membership with id = " + id,  err.Error()},
				HeaderData: make(map[string]string) }
			//return errSys.InternalServerError(119, err.Error())
		}
		return server.OK(membership)
	}
}

func (fac *MembershipEndpointFactory) MakeGetMembershipByIdEndpoint(id string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		id, ok := vars[id]
		if !ok {
			return &server.Response{Status: http.StatusBadRequest,
				Data:       server.Error{400, "2", "No Id", ""},
				HeaderData: make(map[string]string) }
			//return errSys.BadRequest(123," ")
		}
		fmt.Println("MakeGetMembershipByIdEndpoint id = ", id)
		d, err := services.Memberships.GetMembershipById(id)
		if err != nil {
			return &server.Response{Status: http.StatusInternalServerError,
				Data:       server.Error{500, "3", "Could Not found Membership with id = " + id,  err.Error()},
				HeaderData: make(map[string]string) }
			//return errSys.InternalServerError(119, err.Error())
		}
		return server.OK(d)
	}
}

func (fac * MembershipEndpointFactory) MakeGetUserMembershipsEndpoint(email string, services *server.Services) server.HttpEndpoint {

	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		email, ok := vars[email]
		if !ok {
			return &server.Response{Status: http.StatusBadRequest,
				Data:       server.Error{400, "2", "No Id", ""},
				HeaderData: make(map[string]string) }
			//return errSys.BadRequest(123," ")
		}
		d, err := services.Memberships.GetUserMemberships(email)
		if err != nil {
			return &server.Response{Status: http.StatusInternalServerError,
				Data:       server.Error{500, "3", "Could Not found Membershipd with email = " + email,  err.Error()},
				HeaderData: make(map[string]string) }
			//return errSys.InternalServerError(119, err.Error())
		}
		return server.OK(d)
	}
}

func (fac * MembershipEndpointFactory) MakeGetClubMembersEndpoint(clubId string, services *server.Services) server.HttpEndpoint {

	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		clubId, ok := vars[clubId]
		if !ok {
			return &server.Response{Status: http.StatusBadRequest,
				Data:       server.Error{400, "2", "No Id", ""},
				HeaderData: make(map[string]string) }
			//return errSys.BadRequest(123," ")
		}
		fmt.Println("MakeGetClubMembersEndpoint id = ", clubId)
		d, err := services.Memberships.GetClubMembers(clubId)
		if err != nil {
			return &server.Response{Status: http.StatusInternalServerError,
				Data:       server.Error{500, "3", "Could Not found Membershipd with clubId = " + clubId,  err.Error()},
				HeaderData: make(map[string]string) }
			//return errSys.InternalServerError(119, err.Error())
		}
		return server.OK(d)
	}
}

func (fac * MembershipEndpointFactory) MakeDeleteMembershipEndpoint(id string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		id, ok := vars[id]
		if !ok {
			return &server.Response{Status: http.StatusBadRequest,
				Data:       server.Error{400, "2", "No Id", ""},
				HeaderData: make(map[string]string) }
			//return errSys.BadRequest(123," ")
		}
		fmt.Println("MakeDeleteMembershipEndpoint id = ", id)


		_, err := services.Memberships.GetMembershipById(id)
		if err != nil {
			return &server.Response{Status: http.StatusInternalServerError,
				Data:       server.Error{500, "3", "Could Not found membership with id = " + id,  err.Error()},
				HeaderData: make(map[string]string) }
			//return errSys.InternalServerError(119, err.Error())
		}

		err = services.Memberships.DeleteMembership(id)
		if err != nil {
			return &server.Response{Status: http.StatusInternalServerError,
				Data:       server.Error{500, "3", "Could Not found deleted membership with id = " + id,  err.Error()},
				HeaderData: make(map[string]string) }
			//return errSys.InternalServerError(119, err.Error())
		}

		return server.OK("ok")
	}
}