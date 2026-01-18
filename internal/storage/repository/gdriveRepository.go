package storageRepository

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	storageModule "github.com/hifat/mallow-sale-api/internal/storage"
	"github.com/hifat/mallow-sale-api/pkg/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

//go:generate mockgen -source=./gdriveRepository.go -destination=./mock/repository.go -package=mockStorageRepository
type IRepository interface {
	Upload(ctx context.Context, req *storageModule.UploadRequest) (*storageModule.UploadResponse, error)
}

type gdriveRepository struct {
	driveService *drive.Service
}

func NewGDrive(cfg *config.Config) (IRepository, error) {
	ctx := context.Background()

	var credentialsJSON []byte
	var err error

	// Try to use the value as JSON string first
	clientSecret := strings.TrimSpace(cfg.GoogleDrive.ServiceAccount)
	if clientSecret == "" {
		return nil, errors.New("GDRIVE_SERVICE_ACCOUNT is not configured (must provide Client Secret JSON)")
	}

	// Check if it's a JSON string (starts with '{')
	if len(clientSecret) > 0 && clientSecret[0] == '{' {
		// Use as JSON string directly
		credentialsJSON = []byte(clientSecret)
	} else {
		// Try to read as file path
		credentialsJSON, err = os.ReadFile(clientSecret)
		if err != nil {
			return nil, fmt.Errorf("failed to read client secret file: %w", err)
		}
	}

	// --- User OAuth Flow (Quickstart) ---
	config, err := google.ConfigFromJSON(credentialsJSON, drive.DriveFileScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %w", err)
	}

	client := getClient(config)
	driveService, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Drive client: %w", err)
	}

	return &gdriveRepository{
		driveService: driveService,
	}, nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		panic(fmt.Sprintf("Unable to read authorization code: %v", err))
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		panic(fmt.Sprintf("Unable to retrieve token from web: %v", err))
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(fmt.Sprintf("Unable to cache oauth token: %v", err))
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func (r *gdriveRepository) Upload(ctx context.Context, req *storageModule.UploadRequest) (*storageModule.UploadResponse, error) {
	// Create file metadata
	file := &drive.File{
		Name:     req.FileName,
		MimeType: req.MimeType,
	}

	// Set parent folder if provided
	if req.FolderID != "" {
		file.Parents = []string{req.FolderID}
	}

	// Create reader from file bytes
	reader := io.NopCloser(bytes.NewReader(req.File))

	// Upload file to Google Drive
	uploadedFile, err := r.driveService.Files.Create(file).
		Context(ctx).
		Media(reader).
		Fields("id, name, webViewLink").
		Do()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &storageModule.UploadResponse{
		FileID:      uploadedFile.Id,
		FileName:    uploadedFile.Name,
		WebViewLink: uploadedFile.WebViewLink,
		UploadedAt:  &now,
	}, nil
}
