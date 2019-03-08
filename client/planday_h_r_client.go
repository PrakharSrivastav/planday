// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/PrakharSrivastav/planday/client/contract_rules"
	"github.com/PrakharSrivastav/planday/client/contract_rules_employees"
	"github.com/PrakharSrivastav/planday/client/contract_rules_intervals"
	"github.com/PrakharSrivastav/planday/client/contract_rules_shift_types"
	"github.com/PrakharSrivastav/planday/client/custom_property_attachment_values"
	"github.com/PrakharSrivastav/planday/client/departments"
	"github.com/PrakharSrivastav/planday/client/device_tokens"
	"github.com/PrakharSrivastav/planday/client/employee_groups"
	"github.com/PrakharSrivastav/planday/client/employee_priorities"
	"github.com/PrakharSrivastav/planday/client/employee_types"
	"github.com/PrakharSrivastav/planday/client/employees"
	"github.com/PrakharSrivastav/planday/client/employees_activation"
	"github.com/PrakharSrivastav/planday/client/employees_columns"
	"github.com/PrakharSrivastav/planday/client/employees_contract_rules"
	"github.com/PrakharSrivastav/planday/client/employees_gdpr"
	"github.com/PrakharSrivastav/planday/client/employees_relations"
	"github.com/PrakharSrivastav/planday/client/employees_skills"
	"github.com/PrakharSrivastav/planday/client/employees_user_accounts"
	"github.com/PrakharSrivastav/planday/client/employees_utilities"
	"github.com/PrakharSrivastav/planday/client/gdpr"
	"github.com/PrakharSrivastav/planday/client/invites"
	"github.com/PrakharSrivastav/planday/client/pay_rates"
	"github.com/PrakharSrivastav/planday/client/portal_employee_form_definitions"
	"github.com/PrakharSrivastav/planday/client/portal_employee_form_schemas"
	"github.com/PrakharSrivastav/planday/client/portals"
	"github.com/PrakharSrivastav/planday/client/profile"
	"github.com/PrakharSrivastav/planday/client/roles"
	"github.com/PrakharSrivastav/planday/client/system"
	"github.com/PrakharSrivastav/planday/client/translations"
)

// Default planday h r HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new planday h r HTTP client.
func NewHTTPClient(formats strfmt.Registry) *PlandayHR {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new planday h r HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *PlandayHR {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new planday h r client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *PlandayHR {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(PlandayHR)
	cli.Transport = transport

	cli.ContractRules = contract_rules.New(transport, formats)

	cli.ContractRulesEmployees = contract_rules_employees.New(transport, formats)

	cli.ContractRulesIntervals = contract_rules_intervals.New(transport, formats)

	cli.ContractRulesShiftTypes = contract_rules_shift_types.New(transport, formats)

	cli.CustomPropertyAttachmentValues = custom_property_attachment_values.New(transport, formats)

	cli.Departments = departments.New(transport, formats)

	cli.DeviceTokens = device_tokens.New(transport, formats)

	cli.EmployeeGroups = employee_groups.New(transport, formats)

	cli.EmployeePriorities = employee_priorities.New(transport, formats)

	cli.EmployeeTypes = employee_types.New(transport, formats)

	cli.Employees = employees.New(transport, formats)

	cli.EmployeesActivation = employees_activation.New(transport, formats)

	cli.EmployeesColumns = employees_columns.New(transport, formats)

	cli.EmployeesContractRules = employees_contract_rules.New(transport, formats)

	cli.EmployeesGdpr = employees_gdpr.New(transport, formats)

	cli.EmployeesRelations = employees_relations.New(transport, formats)

	cli.EmployeesSkills = employees_skills.New(transport, formats)

	cli.EmployeesUserAccounts = employees_user_accounts.New(transport, formats)

	cli.EmployeesUtilities = employees_utilities.New(transport, formats)

	cli.Gdpr = gdpr.New(transport, formats)

	cli.Invites = invites.New(transport, formats)

	cli.PayRates = pay_rates.New(transport, formats)

	cli.PortalEmployeeFormDefinitions = portal_employee_form_definitions.New(transport, formats)

	cli.PortalEmployeeFormSchemas = portal_employee_form_schemas.New(transport, formats)

	cli.Portals = portals.New(transport, formats)

	cli.Profile = profile.New(transport, formats)

	cli.Roles = roles.New(transport, formats)

	cli.System = system.New(transport, formats)

	cli.Translations = translations.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// PlandayHR is a client for planday h r
type PlandayHR struct {
	ContractRules *contract_rules.Client

	ContractRulesEmployees *contract_rules_employees.Client

	ContractRulesIntervals *contract_rules_intervals.Client

	ContractRulesShiftTypes *contract_rules_shift_types.Client

	CustomPropertyAttachmentValues *custom_property_attachment_values.Client

	Departments *departments.Client

	DeviceTokens *device_tokens.Client

	EmployeeGroups *employee_groups.Client

	EmployeePriorities *employee_priorities.Client

	EmployeeTypes *employee_types.Client

	Employees *employees.Client

	EmployeesActivation *employees_activation.Client

	EmployeesColumns *employees_columns.Client

	EmployeesContractRules *employees_contract_rules.Client

	EmployeesGdpr *employees_gdpr.Client

	EmployeesRelations *employees_relations.Client

	EmployeesSkills *employees_skills.Client

	EmployeesUserAccounts *employees_user_accounts.Client

	EmployeesUtilities *employees_utilities.Client

	Gdpr *gdpr.Client

	Invites *invites.Client

	PayRates *pay_rates.Client

	PortalEmployeeFormDefinitions *portal_employee_form_definitions.Client

	PortalEmployeeFormSchemas *portal_employee_form_schemas.Client

	Portals *portals.Client

	Profile *profile.Client

	Roles *roles.Client

	System *system.Client

	Translations *translations.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *PlandayHR) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.ContractRules.SetTransport(transport)

	c.ContractRulesEmployees.SetTransport(transport)

	c.ContractRulesIntervals.SetTransport(transport)

	c.ContractRulesShiftTypes.SetTransport(transport)

	c.CustomPropertyAttachmentValues.SetTransport(transport)

	c.Departments.SetTransport(transport)

	c.DeviceTokens.SetTransport(transport)

	c.EmployeeGroups.SetTransport(transport)

	c.EmployeePriorities.SetTransport(transport)

	c.EmployeeTypes.SetTransport(transport)

	c.Employees.SetTransport(transport)

	c.EmployeesActivation.SetTransport(transport)

	c.EmployeesColumns.SetTransport(transport)

	c.EmployeesContractRules.SetTransport(transport)

	c.EmployeesGdpr.SetTransport(transport)

	c.EmployeesRelations.SetTransport(transport)

	c.EmployeesSkills.SetTransport(transport)

	c.EmployeesUserAccounts.SetTransport(transport)

	c.EmployeesUtilities.SetTransport(transport)

	c.Gdpr.SetTransport(transport)

	c.Invites.SetTransport(transport)

	c.PayRates.SetTransport(transport)

	c.PortalEmployeeFormDefinitions.SetTransport(transport)

	c.PortalEmployeeFormSchemas.SetTransport(transport)

	c.Portals.SetTransport(transport)

	c.Profile.SetTransport(transport)

	c.Roles.SetTransport(transport)

	c.System.SetTransport(transport)

	c.Translations.SetTransport(transport)

}
