package endpoints

import (
	"github.com/salambayev/x-boat-project/server"
	"net/http"
	"github.com/salambayev/x-boat-project/domain"
	"github.com/salambayev/x-boat-project/utils"
	"fmt"
	"strconv"
)

type TimerecordsEndpointFactory struct {

}

func NewTimerecordsEndpointFactory() *TimerecordsEndpointFactory {
	return &TimerecordsEndpointFactory{}
}

var errTimeSys = &server.ErrorSystem{"x-boat-timerecord"}

func (fac *TimerecordsEndpointFactory) MakeCreateTimerecordEndpoint(services *server.Services) server.HttpEndpoint {

	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		timerecord := &domain.Timerecord{}
		err := utils.ParseJSON(r, timerecord)
		if err != nil {
			return errTimeSys.BadRequest(1, "timerecord body error ", err.Error())
		}
		fmt.Println("MakeCreateTimerecordEndpoint email = ", timerecord.Email)
		if timerecord.Timestamp == 0 {
			return errTimeSys.BadRequest(2, "subscription with no Timestamp could not be created")
		}
		err = services.Timerecords.CreateTimerecord(timerecord)
		if err != nil {
			return errTimeSys.InternalServerError(3, "timerecord already exists",  err.Error())
		}
		return server.Created(timerecord)
	}
}

func (fac * TimerecordsEndpointFactory) MakeUpdateTimerecordEndpoint(services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		timerecord := &domain.Timerecord{}
		err := utils.ParseJSON(r, timerecord)
		if err != nil {
			return errTimeSys.BadRequest(4, "timerecord body error ", err.Error())
		}
		fmt.Println("MakeUpdateTimerecordEndpoint email = ", timerecord.Email)
		if timerecord.Timestamp == 0 {
			return errTimeSys.BadRequest(5,"no timestamp")
		}
		err = services.Timerecords.UpdateTimerecord(timerecord.Timestamp, timerecord)
		if err != nil {
			return errTimeSys.InternalServerError(6, "Could Not found timerecord with Timestamp = " + strconv.FormatInt(timerecord.Timestamp, 10),  err.Error())
		}
		return server.OK(timerecord)
	}
}

func (fac *TimerecordsEndpointFactory) MakeDeleteTimerecordEndpoint(services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		timerecord := &domain.Timerecord{}
		err := utils.ParseJSON(r, timerecord)
		if err != nil {
			return errTimeSys.BadRequest(7, "timerecord body error ", err.Error())
		}
		fmt.Println("MakeDeleteTimerecordEndpoint Timestamp = ", timerecord.Timestamp)
		if timerecord.Timestamp == 0 {
			return errTimeSys.BadRequest(8, "no timestamp")
		}
		err = services.Timerecords.DeleteTimerecord(timerecord.Timestamp)
		if err != nil {
			return errTimeSys.InternalServerError(9, "Could Not found timerecord with Timestamp = " + strconv.FormatInt(timerecord.Timestamp, 10),  err.Error())
		}
		return server.OK("ok")
	}
}

func (fac *TimerecordsEndpointFactory) MakeGetTimerecordEndpoint(services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		timerecord := &domain.Timerecord{}
		err := utils.ParseJSON(r, timerecord)
		if err != nil {
			return errTimeSys.BadRequest(10, "timerecord body error ", err.Error())
		}
		fmt.Println("MakeGetTimerecordEndpoint email = ", timerecord.Email)
		if timerecord.Timestamp == 0 {
			return errTimeSys.BadRequest(11, "no timestamp")
		}

		d, err := services.Timerecords.GetTimerecord(timerecord.Timestamp)
		if err != nil {
			return errTimeSys.InternalServerError(12, "Could Not found timerecord with Timestamp = " + strconv.FormatInt(timerecord.Timestamp, 10),  err.Error())
		}
		return server.OK(d)
	}
}

func (fac *TimerecordsEndpointFactory) MakeGetAllTimerecordsEndpoint( services *server.Services) server.HttpEndpoint {
	return func(w http.ResponseWriter, r *http.Request) server.HttpResponse {
		fmt.Println("MakeGetAllTimerecordsEndpoint ")

		d, err := services.Timerecords.GetAllTimerecords()
		if err != nil {
			return errTimeSys.BadRequest(13, "no timerecords", err.Error())
		}
		return server.OK(d)
	}
}

