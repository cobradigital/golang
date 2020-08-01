package services

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"../crypto"
	"../flags"
	"../repositories"
	"../response"
	"../util"
	"github.com/dgrijalva/jwt-go"
	"github.com/prometheus/common/log"
)

const emailVerificationDelimiter = "-"

type AuthService interface {
	// ExtractToken takes value of Authorization Header, validate expected token type and extract token value
	ExtractToken(authString string, expectedType string) (string, error)
	NewToken(id, key string) (*response.Success, error)
	ValidateAccessToken(r *http.Request, purpose string) error
	ValidateJWT(authString string) (*crypto.JWTClaim, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func (s *authService) ExtractToken(authString string, expectedType string) (string, error) {
	if authString == "" {
		log.Error("Authorization Header is empty")
		return "", util.NewError("401")
	}
	// Extract token
	splitToken := strings.Split(authString, " ")
	if len(splitToken) != 2 {
		log.Error("Bearer token is malformed")
		return "", util.NewError("401")
	}
	tokenType := splitToken[0]
	if tokenType != expectedType {
		log.Errorf("Unmatch token type. Type: %s", tokenType)
		return "", util.NewError("401")
	}
	return splitToken[1], nil
}

func (s *authService) ValidateAccessToken(r *http.Request, purpose string) (err error) {
	// Get authentication string
	authString := r.Header.Get(flags.HeaderKeyCOBRAAuthorization)
	// If client secret

	// Validate jwt token and get claim
	claim, err := Auth.ValidateJWT(authString)
	if err != nil {
		// If error is a jwt validation error, converts error to standard error
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors == jwt.ValidationErrorExpired {
				err = util.NewError("499")
			}
		} else {
			err = util.NewError("401")
		}
		return err
	}
	// Check purpose if exist
	if purpose != "" && claim.Purpose != purpose {
		log.Errorf("Invalid Purpose: %s", claim.Purpose)
		return util.NewError("401")
	} else if purpose == flags.ACLAuthenticatedUser {
		// Check if blacklisted or not
		isBlacklisted, err := false, err
		if err != nil {
			return err
		}
		if isBlacklisted {
			return util.NewError("401")
		}
	}
	// if subject is not empty, apply subject
	if claim.Subject != "" {
		r.Header.Set(flags.HeaderKeyCOBRASubject, claim.Subject)
		// If purpose is authenticated user, validate user
		if purpose == flags.ACLAuthenticatedUser {
			// Get user from repository
			u, _, err := repositories.User.FindById("id = ?", claim.Subject)
			if err != nil {
				// If user not found, change error
				if err == sql.ErrNoRows {
					err = util.NewError("498")
				}
				// Log Error
				log.Errorf("Error: %s, UserId: %s", err, claim.Subject)
				return err
			}
			// If user is banned, return error
			if u.Status == flags.UserStatusBanned {
				err = util.NewError("497")
				log.Errorf("Error: %s, UserId: %s", err, claim.Subject)
				return err
			}
		}
	}
	// Get expired at
	r.Header.Set(flags.HeaderKeyCOBRATokenExpired, strconv.FormatInt(claim.ExpiresAt, 10))
	// If client secret
	return nil
}

func (s *authService) NewToken(id, key string) (*response.Success, error) {
	// Get User
	user, len, _ := repositories.User.Find("secret_id = ? and secret_key = ? and status = ?", id, key, 1)
	if len < 1 {
		return nil, util.NewError("498")
	}
	fmt.Println(user)
	token, _ := authTokenGenerator(strconv.Itoa(int(user[0].Id)))
	// Compose success message
	success := response.Success{
		Result: "Success",
		Header: token,
	}
	return &success, nil
}

func authTokenGenerator(userId string) (map[string]string, error) {
	// Initiate token
	token, expiredAt, err := crypto.NewJWT(userId)
	if err != nil {
		return nil, err
	}
	// Generate token header
	return response.NewToken(token, expiredAt), nil
}

func (s *authService) ValidateJWT(authString string) (*crypto.JWTClaim, error) {
	// Extract token
	token, err := Auth.ExtractToken(authString, crypto.TokenTypeBearer)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return crypto.VerifyJWT(token)
}
