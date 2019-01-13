package main

import (
	"github.com/salambayev/x-boat-project/repositories"
	"github.com/salambayev/x-boat-project/services"
	"golang.org/x/net/context"
	"strconv"
	"os"
	"github.com/urfave/cli"
	"github.com/joho/godotenv"
	"fmt"
	"github.com/salambayev/x-boat-project/db"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/salambayev/x-boat-project/logger"
	"github.com/salambayev/x-boat-project/server"
	"github.com/salambayev/x-boat-project/endpoints"
)

var (
	configPath  = ".env"
	port string
	version = "1.0.0"
)

var log_level, _ = strconv.Atoi(os.Getenv("LOG_LEVEL"))

var flags = []cli.Flag{
	cli.StringFlag{
		Name:        "config, c",
		Value:       configPath,
		Usage:       "path to .env config file",
		Destination: &configPath,
	},
}


func parseEnvFile() error {
	if configPath != "" {
		if err := godotenv.Load(configPath); err != nil {
			return err
		}
	}

	port = os.Getenv("PORT")
	log_level, _ = strconv.Atoi(os.Getenv("LOG_LEVEL"))
	version = os.Getenv("VERSION")

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "X-Boat API"
	app.Usage = "some test usage"
	app.UsageText = "x-boat-project [global options]"
	app.Version = version
	app.Flags = flags
	app.Action = run

	fmt.Println(app.Run(os.Args))
}

func printConfig() {
	fmt.Println("PORT="+port)
	fmt.Println("LOG_LEVEL="+strconv.Itoa(log_level))
	fmt.Println("VERSION="+version)
}

func run(c *cli.Context) error {
	if err := parseEnvFile(); err != nil {
		return err
	}
	printConfig()
	log := logger.New(uint8(log_level))

	if err := db.Connect("newtest-8692c", context.Background(), "newtest-8692c-4d1f7d82ea6d.json"); err != nil {
		return errors.New("Failed to connect to database: "+err.Error())
	}

	clubsFac := endpoints.NewClubsEndpointFactory()
	r := mux.NewRouter()

	repoServices := initializeServices()

	// Club Endpoints
	r.HandleFunc("/clubs",
		server.Json(
			server.Logging(log,
				clubsFac.MakeCreateClubEndpoint(repoServices))),
	).Methods("POST")

	r.HandleFunc("/clubs/{id}",
		server.Json(
			server.Logging(log,
				clubsFac.MakeUpdateClubEndpoint("id", repoServices))),
	).Methods("PUT")

	r.HandleFunc("/clubs/{id}",
		server.Json(
			server.Logging(log,
				clubsFac.MakeGetClubByIdEndpoint("id", repoServices))),
	).Methods("GET")

	r.HandleFunc("/clubs/{id}",
		server.Json(
			server.Logging(log,
				clubsFac.MakeDeleteClubEndpoint("id", repoServices))),
	).Methods("DELETE")

	r.HandleFunc("/all/clubs",
		server.Json(
			server.Logging(log,
				clubsFac.MakeGetAllClubsEndpoint(repoServices))),
	).Methods("GET")


	// Membership Endpoints

	membershipFac := endpoints.NewMembershipEndpointFactory()

	r.HandleFunc("/membership",
		server.Json(
			server.Logging(log,
				membershipFac.MakeCreateMembershipEndpoint(repoServices))),
	).Methods("POST")

	r.HandleFunc("/membership/{id}",
		server.Json(
			server.Logging(log,
				membershipFac.MakeUpdateMembershipEndpoint("id", repoServices))),
	).Methods("PUT")

	r.HandleFunc("/membership/{id}",
		server.Json(
			server.Logging(log,
				membershipFac.MakeGetMembershipByIdEndpoint("id", repoServices))),
	).Methods("GET")


	r.HandleFunc("/memberships/byemail/{email}",
		server.Json(
			server.Logging(log,
				membershipFac.MakeGetUserMembershipsEndpoint("email", repoServices))),
	).Methods("GET")

	r.HandleFunc("/memberships/byclub/{id}",
		server.Json(
			server.Logging(log,
				membershipFac.MakeGetClubMembersEndpoint("id", repoServices))),
	).Methods("GET")

	r.HandleFunc("/memberships/{id}",
		server.Json(
			server.Logging(log,
				membershipFac.MakeDeleteMembershipEndpoint("id", repoServices))),
	).Methods("DELETE")

	// Profile Endpoints
	profileFac := endpoints.NewProfileEndpointFactory()

	r.HandleFunc("/profile",
		server.Json(
			server.Logging(log,
				profileFac.MakeCreateProfileEndpoint(repoServices))),
	).Methods("POST")

	r.HandleFunc("/profile",
		server.Json(
			server.Logging(log,
				profileFac.MakeUpdateProfileEndpoint(repoServices))),
	).Methods("PUT")

	r.HandleFunc("/profile",
		server.Json(
			server.Logging(log,
				profileFac.MakeGetProfileEndpoint  ("email", repoServices))),
	).Methods("GET")

	r.HandleFunc("/profile",
		server.Json(
			server.Logging(log,
				profileFac.MakeDeleteProfileEndpoint("email", repoServices))),
	).Methods("DELETE")


	// Subscription Endpoints
	subscriptionFac := endpoints.NewSubscriptionEndpointFactory()
	r.HandleFunc("/subscription",
		server.Json(
			server.Logging(log,
				subscriptionFac.MakeCreateSubscriptionEndpoint(repoServices))),
				).Methods("POST")

	r.HandleFunc("/subscription/activate",
		server.Json(
			server.Logging(log,
				subscriptionFac.MakeActivateSubscriptionEndpoint(repoServices))),
				).Methods("PUT")

	r.HandleFunc("/subscription/type",
		server.Json(
			server.Logging(log,
				subscriptionFac.MakeChangeSubscriptionTypeEndpoint(repoServices))),
				).Methods("PUT")

	r.HandleFunc("/subscription/deactivate",
		server.Json(
			server.Logging(log,
				subscriptionFac.MakeDeactivateSubscriptionEndpoint(repoServices))),
				).Methods("PUT")

	r.HandleFunc("/subscription",
		server.Json(
			server.Logging(log,
				subscriptionFac.MakeGetSubscriptionEndpoint("email", repoServices))),
				).Methods("GET")


	// Timerecord Endpoints
	timerecordFac := endpoints.NewTimerecordsEndpointFactory()

	r.HandleFunc("/timerecord",
		server.Json(
			server.Logging(log,
				timerecordFac.MakeCreateTimerecordEndpoint(repoServices))),
				).Methods("POST")

	r.HandleFunc("/timerecord",
		server.Json(
			server.Logging(log,
				timerecordFac.MakeUpdateTimerecordEndpoint(repoServices))),
				).Methods("PUT")

	r.HandleFunc("/timerecord",
		server.Json(
			server.Logging(log,
				timerecordFac.MakeDeleteTimerecordEndpoint(repoServices))),
				).Methods("DELETE")

	r.HandleFunc("/timerecord/all",
		server.Json(
			server.Logging(log,
				timerecordFac.MakeGetAllTimerecordsEndpoint(repoServices))),
				).Methods("GET")

	r.HandleFunc("/timerecord",
		server.Json(
			server.Logging(log,
				timerecordFac.MakeGetTimerecordEndpoint(repoServices))),
				).Methods("GET")


	r.NotFoundHandler = http.HandlerFunc( server.Json( server.Logging( log, clubsFac.NotFoundEndpoint())))
	return http.ListenAndServe(":"+port, r)
}

func initializeServices() *server.Services {
	clubs := services.NewClubService(repositories.NewClubRepository())
	memberships := services.NewMembershipService(repositories.NewMembershipRepository())
	profiles := services.NewProfileService(repositories.NewProfileRepository())
	timerecords := services.NewTimerecordsService(repositories.NewTimerecordsRepository())
	subscriptions := services.NewSubscriptionService(repositories.NewSubscriptionRepository())
	return &server.Services {
		Clubs: clubs,
		Memberships: memberships,
		Profiles: profiles,
		Subscriptions: subscriptions,
		Timerecords: timerecords,
	}
}
