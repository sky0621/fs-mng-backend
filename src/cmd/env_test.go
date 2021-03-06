package main

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	tests := []struct {
		name string
		env  map[string]string
		want *env
	}{
		{
			name: "def",
			env:  nil,
			want: &env{
				Env:        "local",
				ServerPort: "5050",
				// DB接続関連
				DBDriverName: "postgres",
				DBHost:       "localhost",
				DBPort:       "19999",
				DBName:       "backenddb",
				DBUser:       "postgres",
				DBPassword:   "localpass",
				DBSSLMode:    "disable",
				// GCS関連
				MovieBucket: "bucket",
				// PubSub関連
				CreateMovieTopic: "create_movie_topic",
				// Auth関連
				AuthDebug:               true,
				AuthCredentialsOptional: true,
				Auth0Domain:             "localhost",
				Auth0ClientID:           "authid",
				Auth0ClientSecret:       "authsecret",
				Auth0Audience:           "http://api.localhost",
				Auth0Timeout:            "30s",
			},
		},
		{
			name: "set",
			env: map[string]string{
				"FS_ENV":         "dev",
				"FS_SERVER_PORT": "8080",
				// DB接続関連
				"FS_DB_DRIVER_NAME": "mysql",
				"FS_DB_HOST":        "example.com",
				"FS_DB_PORT":        "3456",
				"FS_DB_NAME":        "db01",
				"FS_DB_USER":        "testuser",
				"FS_DB_PASSWORD":    "passpass",
				"FS_DB_SSL_MODE":    "true",
				// GCS関連
				"FS_MOVIE_BUCKET": "bucket",
				// PubSub関連
				"FS_CREATE_MOVIE_TOPIC": "test_topic",
				// Auth関連
				"FS_AUTH_DEBUG":                "false",
				"FS_AUTH_CREDENTIALS_OPTIONAL": "false",
				"FS_AUTH0_DOMAIN":              "example2.com",
				"FS_AUTH0_CLIENT_ID":           "authidid",
				"FS_AUTH0_CLIENT_SECRET":       "authsecsec",
				"FS_AUTH0_AUDIENCE":            "https://example.com",
				"FS_AUTH0_TIMEOUT":             "0s",
			},
			want: &env{
				Env:        "dev",
				ServerPort: "8080",
				// DB接続関連
				DBDriverName: "mysql",
				DBHost:       "example.com",
				DBPort:       "3456",
				DBName:       "db01",
				DBUser:       "testuser",
				DBPassword:   "passpass",
				DBSSLMode:    "true",
				// GCS関連
				MovieBucket: "bucket",
				// PubSub関連
				CreateMovieTopic: "test_topic",
				// Auth関連
				AuthDebug:               false,
				AuthCredentialsOptional: false,
				Auth0Domain:             "example2.com",
				Auth0ClientID:           "authidid",
				Auth0ClientSecret:       "authsecsec",
				Auth0Audience:           "https://example.com",
				Auth0Timeout:            "0s",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.env {
				if err := os.Setenv(k, v); err != nil {
					t.Fail()
				}
			}
			if got := loadEnv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
