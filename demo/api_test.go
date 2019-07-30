package main

import (
	"encoding/base64"
	v1 "github.com/ilackarms/solo-kit-demo/pkg/api/v1"
	"github.com/ilackarms/solo-kit-demo/pkg/googleapis"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/utils/protoutils"
	"io/ioutil"
	"log"
)

func createKey() (*v1.GoogleApiKey, error) {
	b, err := ioutil.ReadFile("slides-creds.json")
	if err != nil {
		return nil, err
	}

	key := &v1.GoogleApiKey{
		Metadata:          core.Metadata{Name: "my-key", Namespace: "default"},
		CredentialsBase64: base64.StdEncoding.EncodeToString(b)}
	b, err = protoutils.MarshalBytes(key)
	if err != nil {
		return nil, err
	}

	return key, ioutil.WriteFile("google-api-key.json", b, 0644)
}

func run() error {
	key, err := createKey()
	if err != nil {
		return err
	}
	drv, sld, err := googleapis.CreateClients(key)
	if err != nil {
		return err
	}

	if err := googleapis.DeletePresentations(drv); err != nil {
		return err
	}

	if err := googleapis.SyncPresentation(drv, sld, &v1.Presentation{
		Title:               "test4",
		ContentSlideTitle:   "aaaaa,",
		ContentSlideContent: "test2",
	}); err != nil {
		return err
	}

	log.Printf("finished")
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
