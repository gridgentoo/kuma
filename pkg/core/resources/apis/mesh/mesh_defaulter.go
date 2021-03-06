package mesh

import (
	"github.com/pkg/errors"

	mesh_proto "github.com/Kong/kuma/api/mesh/v1alpha1"
	"github.com/Kong/kuma/pkg/util/proto"
)

func (mesh *MeshResource) Default() error {
	// default settings for Prometheus metrics
	for idx, backend := range mesh.Spec.GetMetrics().GetBackends() {
		if backend.GetType() == mesh_proto.MetricsPrometheusType {
			cfg := mesh_proto.PrometheusMetricsBackendConfig{}
			if err := proto.ToTyped(backend.GetConfig(), &cfg); err != nil {
				return errors.Wrap(err, "could not convert the backend")
			}

			if cfg.Port == 0 {
				cfg.Port = 5670
			}
			if cfg.Path == "" {
				cfg.Path = "/metrics"
			}

			str, err := proto.ToStruct(&cfg)
			if err != nil {
				return errors.Wrap(err, "could not convert the backend")
			}
			mesh.Spec.Metrics.Backends[idx].Config = &str
		}
	}
	return nil
}
