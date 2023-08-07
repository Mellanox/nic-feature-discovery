package daemon

import (
	"context"
	"time"

	"github.com/go-logr/logr"

	"github.com/Mellanox/nic-feature-discovery/pkg/feature"
	"github.com/Mellanox/nic-feature-discovery/pkg/label"
	"github.com/Mellanox/nic-feature-discovery/pkg/writer"
)

// New creates a new Daemon
func New(scanInterval time.Duration, labelWriter writer.LabelWriter, sources []feature.Source) *Daemon {
	return &Daemon{
		scanInterval: scanInterval,
		writer:       labelWriter,
		sources:      sources,
	}
}

// Daemon periodically scans for features and writes labels
type Daemon struct {
	scanInterval time.Duration
	writer       writer.LabelWriter
	sources      []feature.Source
}

// Run daemon control loop
func (d *Daemon) Run(ctx context.Context) {
	log := logr.FromContextOrDiscard(ctx)
	d.discover(ctx)
OUTER:
	for {
		select {
		case <-ctx.Done():
			log.Info("context closed exiting daemon")

			break OUTER
		case <-time.After(d.scanInterval):
			d.discover(ctx)
		}
	}
}

func (d *Daemon) discover(ctx context.Context) {
	log := logr.FromContextOrDiscard(ctx)
	log.WithName("discovery-daemon")

	log.Info("discovering features")

	features := make([]feature.Feature, 0)
	for _, s := range d.sources {
		log.V(2).Info("discovering features from source", "name", s.Name())
		sourceFeatures, err := s.Discover(ctx)
		if err != nil {
			log.Error(err, "failed to discover features from source", "name", s.Name())

			continue
		}
		features = append(features, sourceFeatures...)
	}

	labels := make([]label.Label, 0)
	for _, f := range features {
		labels = append(labels, f.Labels()...)
	}

	log.Info("conditionally updating features file")
	err := d.writer.Write(labels)
	if err != nil {
		log.Error(err, "failed to write feature labels")
	}
	log.Info("discovery complete")
}
