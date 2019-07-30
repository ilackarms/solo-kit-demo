package googleapis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"google.golang.org/api/slides/v1"
	"log"
	"net/http"
	"os"

	v1 "github.com/ilackarms/solo-kit-demo/pkg/api/v1"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func CreateClients(key *v1.GoogleApiKey) (*drive.Service, *slides.Service, error) {
	b, err := base64.StdEncoding.DecodeString(key.CredentialsBase64)
	if err != nil {
		return nil, nil, err
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, slides.PresentationsScope, drive.DriveScope)
	if err != nil {
		return nil, nil, err
	}
	client := getClient(config)

	drv, err := drive.New(client)
	if err != nil {
		return nil, nil, err
	}

	sld, err := slides.New(client)
	if err != nil {
		return nil, nil, err
	}

	return drv, sld, nil
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
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
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
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func filterFiles(files []*drive.File) []*drive.File {
	var filtered []*drive.File
	for _, f := range files {
		if f.MimeType == "application/vnd.google-apps.presentation" {
			filtered = append(filtered, f)
		}
	}
	return filtered
}

const basePresentationId = "1lqAT2UgIw3SwM_fSVUHhLlOGT_JyvW1e03CeRH0mA2M"

var demoLabels = map[string]string{"demo": "only"}

func DeletePresentations(drv *drive.Service) error {
	fileList, err := drv.Files.List().Do()
	if err != nil {
		return err
	}

	presentations := filterFiles(fileList.Files)

	for _, pres := range presentations {
		if pres.Properties["demo"] == "only" {
			if err := drv.Files.Delete(pres.Id).Do(); err != nil {
				return err
			}
		}
	}

	return nil
}

func SyncPresentation(drv *drive.Service, sld *slides.Service, pres *v1.Presentation) error {
	fileList, err := drv.Files.List().Do()
	if err != nil {
		return err
	}

	presentations := filterFiles(fileList.Files)

	for _, presentation := range presentations {
		if presentation.Name == pres.Title {
			if err := drv.Files.Delete(presentation.Id).Do(); err != nil {
				return err
			}
		}
	}

	// copy base presentation
	file := &drive.File{
		Name:       pres.Title,
		Properties: demoLabels,
	}
	presentationCopy, err := drv.Files.Copy(basePresentationId, file).Do()
	if err != nil {
		return err
	}

	updates := &slides.BatchUpdatePresentationRequest{
		Requests: []*slides.Request{
			{
				ReplaceAllText: &slides.ReplaceAllTextRequest{
					ContainsText: &slides.SubstringMatchCriteria{
						Text: "PLACE-TITLE-HERE",
					},
					ReplaceText: pres.Title,
				},
			},
			{
				ReplaceAllText: &slides.ReplaceAllTextRequest{
					ContainsText: &slides.SubstringMatchCriteria{
						Text: "CONTENT-SLIDE-NAME",
					},
					ReplaceText: pres.ContentSlideTitle,
				},
			},
			{
				ReplaceAllText: &slides.ReplaceAllTextRequest{
					ContainsText: &slides.SubstringMatchCriteria{
						Text: "CONTENT-SLIDE-BODY",
					},
					ReplaceText: pres.ContentSlideContent,
				},
			},
		},
	}

	_, err = sld.Presentations.BatchUpdate(presentationCopy.Id, updates).Do()
	return err
}
