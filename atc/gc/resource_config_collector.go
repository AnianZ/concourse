package gc

import (
	"context"

	"code.cloudfoundry.org/lager/lagerctx"
	"github.com/concourse/concourse/atc/db"
)

type resourceConfigCollector struct {
	configFactory db.ResourceConfigFactory
}

func NewResourceConfigCollector(configFactory db.ResourceConfigFactory) *resourceConfigCollector {
	return &resourceConfigCollector{
		configFactory: configFactory,
	}
}

func (rcuc *resourceConfigCollector) Run(ctx context.Context) error {
	logger := lagerctx.FromContext(ctx).Session("resource-config-collector")

	logger.Debug("start")
	defer logger.Debug("done")

	return rcuc.configFactory.CleanUnreferencedConfigs()
}
