package main

import (
	"encoding/base64"
	"github.com/hashicorp/vault/api"
	v1 "github.com/ilackarms/solo-kit-demo/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"io/ioutil"
	"log"
	"sigs.k8s.io/yaml"
)

func writeKey() error {
	vaultCfg := api.DefaultConfig()
	vaultCfg.Address = "http://127.0.0.1:8200"
	vaultClient, err := api.NewClient(vaultCfg)
	if err != nil {
		return err
	}
	vaultClient.SetToken("root")

	// use vault as our credentials store
	credentialsClient, err := v1.NewGoogleApiKeyClient(&factory.VaultSecretClientFactory{
		Vault:   vaultClient,
		RootKey: "apikeys",
	})

	b, err := ioutil.ReadFile("slides-creds.json")
	if err != nil {
		return err
	}

	key := &v1.GoogleApiKey{
		Metadata:          core.Metadata{Name: "my-key", Namespace: "default"},
		CredentialsBase64: base64.StdEncoding.EncodeToString(b)}

	_, err = credentialsClient.Write(key, clients.WriteOpts{})
	return err
}

func main() {
	if err := writePres(); err != nil {
		log.Fatal(err)
	}
	if err := writeKey(); err != nil {
		log.Fatal(err)
	}
}

func writePres() error {
	kr := v1.PresentationCrd.KubeResource(&v1.Presentation{
		Metadata: core.Metadata{
			Name:      "presentation-1",
			Namespace: "default",
		},
		Title:               "Demo Ideas",
		ContentSlideTitle:   "Why Coming Up with Demo Ideas is Difficult",
		ContentSlideContent: "* You might think that coming up with demo ideas is easy.\n* But it's a lot harder than you think.\n",
	})
	b, err := yaml.Marshal(kr)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("presentation-1.yaml", b, 0644)
}
