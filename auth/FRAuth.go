package auth

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/salambayev/x-boat-project/db"
	"github.com/salambayev/x-boat-project/domain"
	"github.com/salambayev/x-boat-project/utils"
	"log"
	"net/http"
	"strings"
)

type FRAuth struct {

}

func (frAuth *FRAuth) ValidateToken(w http.ResponseWriter, r *http.Request) (*auth.Token, error) {
	token := r.Header.Get("Authorization")
	splitedToken := strings.Split(token, "Bearer")
	trueToken := splitedToken[1]

	validatedToken, err := db.FRAuthClient.VerifyIDToken(context.Background(), trueToken)
	if err != nil {
		return nil, err
	}

	return validatedToken, nil
}

func (frAuth *FRAuth) SignUp(w http.ResponseWriter, r *http.Request) (*auth.UserRecord, error) {
	profile := &domain.Profile{}
	err := utils.ParseJSON(r, profile)
	if err != nil {
		return nil, err
	}
	params := (&auth.UserToCreate{}).
		Email(profile.Email).
		EmailVerified(false).
		Password(profile.Password).
		DisplayName(profile.FirstName + " " + profile.LastName).
		PhotoURL(profile.Image).
		Disabled(false)
	newUser, err := db.FRAuthClient.CreateUser(context.Background(), params)
	if err != nil {
		return nil, err
	}
	log.Printf("Successfully created user: %v/n", newUser)
	return newUser, nil
}