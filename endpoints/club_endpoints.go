package endpoints

import (
	"github.com/salambayev/x-boat-project/server"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/salambayev/x-boat-project/domain"
	"github.com/salambayev/x-boat-project/utils"
)

type ClubsEndpointFactory struct {

}

func NewClubsEndpointFactory() *ClubsEndpointFactory {
	return &ClubsEndpointFactory{}
}

var errClSys = &server.ErrorSystem{"x-boat-club"}

func (fac *ClubsEndpointFactory) NotFoundEndpoint() server.HttpEndpoint {
	return func (w http.ResponseWriter, r *http.Request) server.HttpResponse {
		return errClSys.NotFound(1, "Whoops! Requested url not found")
	}
}

func (fac *ClubsEndpointFactory) MakeCreateClubEndpoint(services *server.Services) server.HttpEndpoint {

	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		club := &domain.Club{}
		err := utils.ParseJSON(r, club)
		if err != nil {
			return errClSys.BadRequest(2, "No Id")
		}
		if club.ClubId == "" {
			club.ClubId = utils.GenerateId()
		}
		err = services.Clubs.CreateClub(club)
		if err != nil {
			return errClSys.InternalServerError(3, "Club already exists",  err.Error())
		}
		return server.Created(club)
	}
}

func (fac * ClubsEndpointFactory) MakeUpdateClubEndpoint(id string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		id, ok := vars[id]
		if !ok {
			return errClSys.BadRequest(4,"No Id")
		}
		club := &domain.Club{}
		err := utils.ParseJSON(r, club)
		if err != nil {
			return errClSys.BadRequest(5, "No Club", err.Error())
		}
		club.ClubId = id
		err = services.Clubs.UpdateClub(id, club)
		if err != nil {
			return errClSys.InternalServerError(6, "Could Not found club with id = " + id,  err.Error())
		}
		return server.OK(club)
	}
}

func (fac *ClubsEndpointFactory) MakeGetClubByIdEndpoint(id string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		id, ok := vars[id]
		if !ok {
			return errClSys.BadRequest(7, "No id = " + id)
		}
		d, err := services.Clubs.GetClub(id)
		if err != nil {
			return errClSys.InternalServerError(8,"Could Not found club with id = " + id, err.Error())
		}
		return server.OK(d)
	}
}

func (fac *ClubsEndpointFactory) MakeDeleteClubEndpoint(id string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		id, ok := vars[id]
		if !ok {
			return errClSys.BadRequest(9,"No id")
		}
		_, err := services.Clubs.GetClub(id)
		if err != nil {
			return errClSys.InternalServerError(10, "Could Not found club with id = " + id,  err.Error())
		}

		err = services.Clubs.DeleteClub(id)
		if err != nil {
			return errClSys.InternalServerError(11, "Could Not found club with id = " + id,  err.Error())
		}
		return server.OK("ok")
	}
}

func (fac *ClubsEndpointFactory) MakeGetAllClubsEndpoint(services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		d, err := services.Clubs.GetAllClubs()
		if err != nil {
			return errClSys.InternalServerError(12, "Could Not found any club" , err.Error())
		}
		return server.OK(d)
	}
}