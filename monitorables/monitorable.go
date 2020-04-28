package monitorables

import (
	pkgMonitorable "github.com/monitoror/monitoror/internal/pkg/monitorable"
	coreModels "github.com/monitoror/monitoror/models"
	"github.com/monitoror/monitoror/store"
)

type Monitorable interface {
	//GetDisplayName return monitorable name display in console
	GetDisplayName() string

	//GetVariantsNames return variant list extract from config
	GetVariantsNames() []coreModels.VariantName

	//Validate test if config variant is valid
	// return false if empty and error if config have an error (ex: wrong url format)
	Validate(variantName coreModels.VariantName) (bool, []error)

	//Enable monitorable variant (add route to echo and enable tile for config verify / hydrate)
	Enable(variantName coreModels.VariantName)
}

type (
	Manager struct {
		store *store.Store

		monitorables []Monitorable
	}
)

func NewMonitorableManager(store *store.Store) *Manager {
	return &Manager{store: store}
}

func (m *Manager) register(monitorable Monitorable) {
	m.monitorables = append(m.monitorables, monitorable)
}

func (m *Manager) EnableMonitorables() {
	for _, monitorable := range m.monitorables {
		monitorableMetadata := *pkgMonitorable.NewMonitorableMetadata(monitorable.GetDisplayName())

		for _, variantName := range monitorable.GetVariantsNames() {
			valid, errors := monitorable.Validate(variantName)
			if errors != nil {
				monitorableMetadata.AddErroredVariant(variantName, errors)
			}

			if valid {
				monitorable.Enable(variantName)
				monitorableMetadata.AddEnabledVariant(variantName)
			}
		}

		m.store.MonitorableMetadata = append(m.store.MonitorableMetadata, monitorableMetadata)
	}
}
