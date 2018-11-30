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
	mongoDbUrl string
	mongoDbHost string
	mongoDbPort string
	mongoDbName string
	configPath  = ".env"
	port string
	version = "0.0.1"
)

var log_level, _ = strconv.Atoi(os.Getenv("LOG_LEVEL"))

var flags = []cli.Flag{
	cli.StringFlag{
		Name:        "config, c",
		Value:       configPath,
		Usage:       "path to .env config file",
		Destination: &configPath,
	},
	cli.StringFlag{
		Name: "mongo_url, mgu",
		Value: "",
		Usage: "Mongo DB Url",
		Destination: &mongoDbUrl,
	},
	cli.StringFlag{
		Name: "mongo_name, mgn",
		Value: "",
		Usage: "Mongo DB Name",
		Destination: &mongoDbName,
	},
	cli.StringFlag{
		Name: "mongo_host, mgh",
		Value: "",
		Usage: "Mongo DB Host",
		Destination: &mongoDbHost,
	},
	cli.StringFlag{
		Name: "mongo_port, mgp",
		Value: "",
		Usage: "Mongo DB Port",
		Destination: &mongoDbPort,
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
	mongoDbUrl = os.Getenv("MONGODB_URL")
	mongoDbName = os.Getenv("MONGODB_NAME")
	mongoDbHost = os.Getenv("MONGODB_HOST")
	mongoDbPort = os.Getenv("MONGODB_PORT")
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
	fmt.Println("MONGODB_URL="+mongoDbUrl)
	fmt.Println("MONGODB_NAME="+mongoDbName)
	fmt.Println("MONGODB_HOST="+mongoDbHost)
	fmt.Println("MONGODB_PORT="+mongoDbPort)
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

	services := initializeServices()


	//Here should be the handlers -->
	// Club Endpoints
	r.HandleFunc("/clubs",
		server.Json(
			server.Logging(log,
				clubsFac.MakeCreateClubEndpoint( services))),
	).Methods("POST")

	r.HandleFunc("/clubs/{id}",
		server.Json(
			server.Logging(log,
				clubsFac.MakeUpdateClubEndpoint("id", services))),
	).Methods("PUT")

	r.HandleFunc("/clubs/{id}",
		server.Json(
			server.Logging(log,
				clubsFac.MakeGetClubByIdEndpoint("id", services))),
	).Methods("GET")

	r.HandleFunc("/clubs/{id}",
		server.Json(
			server.Logging(log,
				clubsFac.MakeDeleteClubEndpoint("id", services))),
	).Methods("DELETE")

	r.HandleFunc("/all/clubs",
		server.Json(
			server.Logging(log,
				clubsFac.MakeGetAllClubsEndpoint(services))),
	).Methods("GET")


	// Membership Endpoints

	membershipFac := endpoints.NewMembershipEndpointFactory()

	r.HandleFunc("/membership",
		server.Json(
			server.Logging(log,
				membershipFac.MakeCreateMembershipEndpoint( services))),
	).Methods("POST")

	r.HandleFunc("/membership/{id}",
		server.Json(
			server.Logging(log,
				membershipFac.MakeUpdateMembershipEndpoint("id", services))),
	).Methods("PUT")

	r.HandleFunc("/membership/{id}",
		server.Json(
			server.Logging(log,
				membershipFac.MakeGetMembershipByIdEndpoint("id", services))),
	).Methods("GET")


	r.HandleFunc("/memberships/byemail/{email}",
		server.Json(
			server.Logging(log,
				membershipFac.MakeGetUserMembershipsEndpoint("email", services))),
	).Methods("GET")

	r.HandleFunc("/memberships/byclub/{id}",
		server.Json(
			server.Logging(log,
				membershipFac.MakeGetClubMembersEndpoint("id", services))),
	).Methods("GET")

	r.HandleFunc("/memberships/{id}",
		server.Json(
			server.Logging(log,
				membershipFac.MakeDeleteMembershipEndpoint("id", services))),
	).Methods("DELETE")

	// Profile Endpoints
	profileFac := endpoints.NewProfileEndpointFactory()

	r.HandleFunc("/profile",
		server.Json(
			server.Logging(log,
				profileFac.MakeCreateProfileEndpoint( services))),
	).Methods("POST")

	r.HandleFunc("/profile",
		server.Json(
			server.Logging(log,
				profileFac.MakeUpdateProfileEndpoint( services))),
	).Methods("PUT")

	r.HandleFunc("/profile",
		server.Json(
			server.Logging(log,
				profileFac.MakeGetProfileEndpoint  ("email", services))),
	).Methods("GET")

	r.HandleFunc("/profile",
		server.Json(
			server.Logging(log,
				profileFac.MakeDeleteProfileEndpoint("email", services))),
	).Methods("DELETE")


	// Subscription Endpoints
	//subscriptionFac := endpoints.NewSubscriptionEndpointFactory()





	// Timerecord Endpoints
	//timerecordFac := endpoints.NewTimerecordsEndpointFactory()


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
