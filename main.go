package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/iam/v1"

	"google.golang.org/api/oauth2/v2"
)

func main() {
	ctx := context.Background()

	// Use default credentials
	creds, err := google.FindDefaultCredentials(ctx, iam.CloudPlatformScope)
	if err != nil {
		log.Fatalf("Failed to find default credentials: %v", err)
	}

	fmt.Printf("Application Default Credentials:\n%s\n", string(creds.JSON))
	fmt.Println()

	// Extract the email of the service account or user
	tokenSource := creds.TokenSource
	token, err := tokenSource.Token()
	if err != nil {
		slog.Warn("Failed to get token:", "error", err)
	}

	ok := false
	idToken := ""
	if err == nil {
		idToken, ok = token.Extra("id_token").(string)
		fmt.Printf("id_token:\n%+v\n", token)
	}

	if ok && idToken != "" {
		fmt.Println("Running as a user account")

		segments := strings.Split(idToken, ".")
		if len(segments) < 2 {
			log.Fatalf("Invalid ID token format")
		}

		payload, err := base64.RawURLEncoding.DecodeString(segments[1])
		if err != nil {
			log.Fatalf("Failed to decode ID token payload: %v", err)
		}

		var claims struct {
			Email string `json:"email"`
		}
		if err := json.Unmarshal(payload, &claims); err != nil {
			log.Fatalf("Failed to unmarshal ID token payload: %v", err)
		}

		fmt.Printf("Authenticated as: %s \n", claims.Email)
	} else {
		fmt.Println("Most likely as a service account")
		// It's likely a service account, so let's fetch user info using the OAuth2 API
		oauth2Service, err := oauth2.NewService(ctx)
		if err != nil {
			log.Fatalf("Failed to create OAuth2 service: %v", err)
		}

		userInfoService := oauth2.NewUserinfoV2MeService(oauth2Service)
		userInfo, err := userInfoService.Get().Do()
		if err != nil {
			log.Fatalf("Failed to get user info: %v", err)
		}

		fmt.Printf("Authenticated as: %s\n", userInfo.Email)
	}
}
