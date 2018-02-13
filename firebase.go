package grest

import (
	"context"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// FirebaseApp - Instance of Firebase Client
func FirebaseApp(credentialsPath string) (*firebase.App, error) {
	opt := option.WithCredentialsFile(credentialsPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		return nil, err
	}

	return app, nil
}
