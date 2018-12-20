package db

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
)

var FRAuthClient *auth.Client

var ProfilesCollection *firestore.CollectionRef = nil
var MembershipCollection *firestore.CollectionRef = nil
var TimerecordCollection *firestore.CollectionRef = nil
var ClubCollection *firestore.CollectionRef = nil
var SubscriptionCollection *firestore.CollectionRef = nil

func Connect(projectId string, ctx context.Context, pathToAccount string) error {
	conf := &firebase.Config{ProjectID: projectId}
	app, err := firebase.NewApp(ctx, conf, option.WithCredentialsFile(pathToAccount))
	if err != nil {
		log.Fatalln(err)
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	ProfilesCollection = client.Collection("profiles")
	MembershipCollection = client.Collection("memberships")
	TimerecordCollection = client.Collection("timerecords")
	ClubCollection = client.Collection("clubs")
	SubscriptionCollection = client.Collection("subscriptions")

	FRAuthClient = auth

	return nil
}