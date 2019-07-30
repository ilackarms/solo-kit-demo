package syncer

import (
	"context"
	v1 "github.com/ilackarms/solo-kit-demo/pkg/api/v1"
	"github.com/ilackarms/solo-kit-demo/pkg/googleapis"
	"github.com/solo-io/go-utils/errors"
	"log"
)

type PresentationSyncer struct{}

func (s *PresentationSyncer) Sync(ctx context.Context, snap *v1.PresentationsSnapshot) error {
	if len(snap.Googleapikeys) == 0 {
		return errors.Errorf("must provide a google api key")
	}
	key := snap.Googleapikeys[0]
	drv, srv, err := googleapis.CreateClients(key)
	if err != nil {
		return err
	}

	if err := googleapis.DeletePresentations(drv); err != nil {
		return err
	}

	for _, pres := range snap.Presentations {
		if err := googleapis.SyncPresentation(drv, srv, pres); err  != nil {
			return err
		}
	}

	log.Printf("synced %v presentations", len(snap.Presentations))

	return nil
}
