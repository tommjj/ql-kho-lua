package domain

import (
	"errors"
)

var (
	// ErrInternal is an error for when an internal service fails to process the request
	ErrInternal = errors.New("internal error")
	// ErrDataNotFound is an error for when requested data is not found
	ErrDataNotFound = errors.New("data not found")
	// ErrNoUpdatedData is an error for when no data is provided to update
	ErrNoUpdatedData = errors.New("no data to update")
	// ErrDataConflict is an error for data conflict
	ErrDataConflict = errors.New("data conflict error")
	// ErrConflictingData is an error for when data conflicts with existing data
	ErrConflictingData = errors.New("data conflicts with existing data in unique column")
	// ErrTokenDuration is an error for when the token duration format is invalid
	ErrTokenDuration = errors.New("invalid token duration format")
	// ErrTokenCreation is an error for when the token creation fails
	ErrTokenCreation = errors.New("error creating token")
	// ErrExpiredToken is an error for when the access token is expired
	ErrExpiredToken = errors.New("access token has expired")
	// ErrInvalidToken is an error for when the access token is invalid
	ErrInvalidToken = errors.New("access token is invalid")
	// ErrInvalidCredentials is an error for when the credentials are invalid
	ErrInvalidCredentials = errors.New("invalid email or password")
	// ErrEmptyAuthorizationHeader is an error for when the authorization header is empty
	ErrEmptyAuthorizationHeader = errors.New("authorization header is not provided")
	// ErrInvalidAuthorizationHeader is an error for when the authorization header is invalid
	ErrInvalidAuthorizationHeader = errors.New("authorization header format is invalid")
	// ErrInvalidAuthorizationType is an error for when the authorization type is invalid
	ErrInvalidAuthorizationType = errors.New("authorization type is not supported")
	// ErrUnauthorized is an error for when the user is unauthorized
	ErrUnauthorized = errors.New("user is unauthorized to access the resource")
	// ErrForbidden is an error for when the user is forbidden to access the resource
	ErrForbidden = errors.New("user is forbidden to access the resource")
	// ErrInvalidLocation is an error for the location is not valid
	ErrInvalidLocation = errors.New("location is not valid")
)

// File storage
var (
	// ErrConflictDir is an error for base dir and temp dir conflicting
	ErrConflictingDirectory = errors.New("directory conflicting err")
	// ErrCanNotCreateTemp is an error for can not create base directory
	ErrCreateBaseDirectory = errors.New("can not create base directory")
	// ErrCanNotCreateTemp is an error for can not create temp directory
	ErrCreateTempDirectory = errors.New("can not create temp directory")
	// ErrFileIsNotExist is an error for file is not exist
	ErrFileIsNotExist = errors.New("file is not exist")
	// ErrCreateFile is an for create file err
	ErrCreateFile = errors.New("create file err")
	// ErrSaveFile is an error for save file err
	ErrSaveFile = errors.New("save file err")
	// ErrOpenFile is an error for open file err
	ErrOpenFile = errors.New("open file err")
)
