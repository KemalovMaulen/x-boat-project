package db

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
)

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

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	ProfilesCollection = client.Collection("profiles")
	MembershipCollection = client.Collection("memberships")
	TimerecordCollection = client.Collection("timerecords")
	ClubCollection = client.Collection("clubs")
	SubscriptionCollection = client.Collection("subscriptions")

	return nil
}