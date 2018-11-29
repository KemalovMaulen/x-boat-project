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

func (fac *ClubsEndpointFactory) NotFoundEndpoint() server.HttpEndpoint {
	return func (w http.ResponseWriter, r *http.Request) server.HttpResponse {
		return &server.Response{Status: http.StatusNotFound,
		Data:       server.Error{404, "1", "Whoops! Requested url not found", ""},
		HeaderData: make(map[string]string) }
	}
}

//var errSys = &server.ErrorSystem{"x-boat", 10}


func (fac *ClubsEndpointFactory) MakeCreateClubEndpoint(services *server.Services) server.HttpEndpoint {

	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		club := &domain.Club{}
		err := utils.ParseJSON(r, club)
		if err != nil {
			return &server.Response{Status: http.StatusBadRequest,
				Data:       server.Error{400, "2", "No Id", ""},
				HeaderData: make(map[string]string) }
			//return errSys.BadRequest(110, err.Error())
		}
		err = services.Clubs.CreateClub(club)
		if err != nil {
			return &server.Response{Status: http.StatusInternalServerError,
				Data:       server.Error{500, "3", "Club already exists",  err.Error()},
				HeaderData: make(map[string]string) }
			//return errSys.InternalServerError(119, err.Error())
		}
		return server.Created(club)
	}
}

func (fac * ClubsEndpointFactory) MakeUpdateClubEndpoint(id string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		id, ok := vars[id]
		if !ok {
			return &server.Response{Status: http.StatusBadRequest,
				Data:       server.Error{400, "2", "No Id", ""},
				HeaderData: make(map[string]string) }
			//return errSys.BadRequest(123," ")
		}
		club := &domain.Club{}
		err := utils.ParseJSON(r, club)
		if err != nil {
			return &server.Response{Status: http.StatusBadRequest,
				Data:       server.Error{400, "2", "No Club", err.Error()},
				HeaderData: make(map[string]string) }
			//return errSys.BadRequest(110, err.Error())
		}
		club.ClubId = id
		err = services.Clubs.UpdateClub(id, club)
		if err != nil {
			return &server.Response{Status: http.StatusInternalServerError,
				Data:       server.Error{500, "3", "Could Not found club with id = " + id,  err.Error()},
				HeaderData: make(map[string]string) }
			//return errSys.InternalServerError(119, err.Error())
		}
		return server.OK(club)
	}
}

func (fac *ClubsEndpointFactory) MakeGetClubByIdEndpoint(id string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		id, ok := vars[id]
		if !ok {
			return &server.Response{Status: http.StatusBadRequest,
				Data:       server.Error{400, "2", "No Id", ""},
				HeaderData: make(map[string]string) }
			//return errSys.BadRequest(123," ")
		}
		d, err := services.Clubs.GetClub(id)
		if err != nil {
			return &server.Response{Status: http.StatusInternalServerError,
				Data:       server.Error{500, "3", "Could Not found club with id = " + id,  ""},
				HeaderData: make(map[string]string) }
			//return errSys.InternalServerError(119, err.Error())
		}
		return server.OK(d)
	}
}

func (fac *ClubsEndpointFactory) MakeDeleteClubEndpoint(id string, services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		vars := mux.Vars(r)
		id, ok := vars[id]
		if !ok {
			return &server.Response{Status: http.StatusBadRequest,
				Data: server.Error{400, "2", "No Id", ""},
				HeaderData: make(map[string]string)}
			//return errSys.BadRequest(123," ")
		}
		_, err := services.Clubs.GetClub(id)
		if err != nil {
			return &server.Response{Status: http.StatusInternalServerError,
				Data:       server.Error{500, "3", "Could Not found club with id = " + id,  err.Error()},
				HeaderData: make(map[string]string) }
			//return errSys.InternalServerError(119, err.Error())
		}

		err = services.Clubs.DeleteClub(id)
		if err != nil {
			return &server.Response{Status: http.StatusInternalServerError,
				Data:       server.Error{500, "3", "Could Not found club with id = " + id,  err.Error()},
				HeaderData: make(map[string]string) }
			//return errSys.InternalServerError(119, err.Error())
		}
		return server.OK("ok")
	}
}

func (fac *ClubsEndpointFactory) MakeGetAllClubsEndpoint(services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {

		d, err := services.Clubs.GetAllClubs()
		if err != nil {
			return &server.Response{Status: http.StatusInternalServerError,
				Data:       server.Error{500, "3", "Could Not found club with id = " ,  ""},
				HeaderData: make(map[string]string) }
			//return errSys.InternalServerError(119, err.Error())
		}
		return server.OK(d)
	}
}