package store

import (
	"github.com/jsdidierlaurent/echo-middleware/cache"

	coreConfig "github.com/monitoror/monitoror/config"
	pkgMonitorable "github.com/monitoror/monitoror/internal/pkg/monitorable"
	"github.com/monitoror/monitoror/registry"
	"github.com/monitoror/monitoror/service/router"
)

type (
	// Store is used to share Data in every monitorable
	Store struct {
		// Global CoreConfig
		CoreConfig *coreConfig.Config

		// CacheStore for every memory persistent data
		CacheStore cache.Store

		// Registry used to register Tile for verify / hydrate
		Registry registry.Registry

		// MonitorableRouter helper wrapping echo Router monitorable
		MonitorableRouter router.MonitorableRouter

		// MonitorableMetadata store data to print startup log
		MonitorableMetadata []pkgMonitorable.Metadata
	}
)
