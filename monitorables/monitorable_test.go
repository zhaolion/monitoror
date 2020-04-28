package monitorables

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/monitoror/monitoror/config"
	coreModels "github.com/monitoror/monitoror/models"
	"github.com/monitoror/monitoror/store"
)

type monitorableMock struct {
	displayName    string
	variants       []coreModels.VariantName
	validateBool   bool
	validateErrors []error
}

func (m *monitorableMock) GetDisplayName() string                     { return m.displayName }
func (m *monitorableMock) GetVariantsNames() []coreModels.VariantName { return m.variants }
func (m *monitorableMock) Validate(_ coreModels.VariantName) (bool, []error) {
	return m.validateBool, m.validateErrors
}
func (m *monitorableMock) Enable(_ coreModels.VariantName) {}

func TestManager_EnableMonitorables(t *testing.T) {
	mockMonitorable1 := &monitorableMock{
		displayName:    "Monitorable mock 1",
		variants:       []coreModels.VariantName{coreModels.DefaultVariant},
		validateBool:   true,
		validateErrors: nil,
	}
	mockMonitorable2 := &monitorableMock{
		displayName:    "Monitorable mock 2",
		variants:       []coreModels.VariantName{coreModels.DefaultVariant},
		validateBool:   false,
		validateErrors: []error{errors.New("boom"), errors.New("boom2")},
	}

	manager := NewMonitorableManager(&store.Store{
		CoreConfig: &config.Config{
			Env: "production",
		},
	})

	manager.register(mockMonitorable1)
	assert.Len(t, manager.monitorables, 1)
	manager.register(mockMonitorable2)
	assert.Len(t, manager.monitorables, 2)

	manager.EnableMonitorables()

	// Count non-enable monitorables
	nonEnabledMonitorable := 0
	for _, monitorable := range manager.store.MonitorableMetadata {
		if monitorable.IsDisabled() {
			nonEnabledMonitorable++
		}
	}

	assert.Equal(t, 0, nonEnabledMonitorable)

	mockMonitorable3 := &monitorableMock{
		displayName:    "Monitorable mock 3",
		variants:       []coreModels.VariantName{},
		validateBool:   false,
		validateErrors: nil,
	}
	manager.register(mockMonitorable3)
	assert.Len(t, manager.monitorables, 3)

	manager.EnableMonitorables()

	// Count non-enable monitorables
	nonEnabledMonitorable = 0
	for _, monitorable := range manager.store.MonitorableMetadata {
		if monitorable.IsDisabled() {
			nonEnabledMonitorable++
		}
	}

	assert.Equal(t, 1, nonEnabledMonitorable)
}
