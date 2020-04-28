package monitorable

import (
	coreModels "github.com/monitoror/monitoror/models"
)

type (
	Metadata struct {
		MonitorableName        string
		EnabledVariantNames    []coreModels.VariantName
		ErroredVariantMetadata []ErroredVariantMetadata
	}

	ErroredVariantMetadata struct {
		VariantName coreModels.VariantName
		Errors      []error
	}
)

func NewMonitorableMetadata(monitorableName string) *Metadata {
	return &Metadata{MonitorableName: monitorableName}
}

func (m *Metadata) IsDisabled() bool {
	return len(m.EnabledVariantNames) == 0 && len(m.ErroredVariantMetadata) == 0
}

func (m *Metadata) AddEnabledVariant(variantName coreModels.VariantName) {
	m.EnabledVariantNames = append(m.EnabledVariantNames, variantName)
}

func (m *Metadata) AddErroredVariant(variantName coreModels.VariantName, errors []error) {
	m.ErroredVariantMetadata = append(m.ErroredVariantMetadata, ErroredVariantMetadata{
		VariantName: variantName,
		Errors:      errors,
	})
}
