package main

import (
	"context"
	"github.com/hashicorp/vault/api"
	v1 "github.com/ilackarms/solo-kit-demo/pkg/api/v1"
	"github.com/ilackarms/solo-kit-demo/pkg/syncer"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"log"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	log.Printf("started")
	ctx := context.Background()

	cfg, err := kubeutils.GetConfig("", "")
	if err != nil {
		return err
	}
	cache := kube.NewKubeCache(ctx)

	// use kubernetes as our presentation storage
	presentationClient, err := v1.NewPresentationClient(&factory.KubeResourceClientFactory{
		Crd:         v1.PresentationCrd,
		Cfg:         cfg,
		SharedCache: cache,
	})
	if err != nil {
		return err
	}

	if err := presentationClient.Register(); err != nil {
		return err
	}

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

	// initialize and run the event loop
	emitter := v1.NewPresentationsEmitter(presentationClient, credentialsClient)
	eventLoop := v1.NewPresentationsEventLoop(emitter, &syncer.PresentationSyncer{})

	errs, err := eventLoop.Run(nil, clients.WatchOpts{Ctx:ctx, RefreshRate: time.Second * 2})
	if err != nil {
		return err
	}

	for err := range errs {
		log.Printf("error: %v", err)
	}

	return nil
}


