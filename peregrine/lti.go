package peregrine

import (
	"github.com/google/uuid"
	"time"
)

// ToolDataRepo is intended to be a storage (e.g. DB) service for an LTI Tools registration and launch
type ToolDataRepo interface {
	// GetPlatformInstance should return a PlatformInstance by ID
	GetPlatformInstance(id uuid.UUID) (PlatformInstance, error)
	// GetPlatformInstanceByGUID should return a PlatformInstance by GUID
	GetPlatformInstanceByGUID(guid string) (PlatformInstance, error)
	// CreatePlatformInstance should create a PlatformInstance returning PlatformInstance with ID
	CreatePlatformInstance(platIns PlatformInstance) (PlatformInstance, error)
	// UpdatePlatformInstance should update a PlatformInstance by ID
	UpdatePlatformInstance(platIns PlatformInstance) (PlatformInstance, error)
	// DeletePlatformInstance should delete a PlatformInstance by ID
	DeletePlatformInstance(id uuid.UUID) error
	// GetRegistration should return a Registration by ID
	GetRegistration(id uuid.UUID) (Registration, error)
	// GetRegistrationByClientID should return a Registration by ClientID
	GetRegistrationByClientID(clientId string) (Registration, error)
	// CreateRegistration should create a Registration returning Registration with ID
	CreateRegistration(reg Registration) (Registration, error)
	// UpdateRegistration should update a Registration by ID
	UpdateRegistration(reg Registration) (Registration, error)
	// DeleteRegistration should delete a Registration by ID
	DeleteRegistration(id uuid.UUID) (Registration, error)
	// GetDeployment should return a Deployment by ID
	GetDeployment(id uuid.UUID) (Deployment, error)
	// GetDeploymentByPlatformDeploymentID should return a Deployment by PlatformDeploymentID
	GetDeploymentByPlatformDeploymentID(deploymentId string) (Deployment, error)
	// CreateDeployment should create a Deployment returning Deployment with ID
	CreateDeployment(dep Deployment) (Deployment, error)
	// UpdateDeployment should update a Deployment by ID
	UpdateDeployment(dep Deployment) (Deployment, error)
	// DeleteDeployment should delete a Deployment by ID
	DeleteDeployment(id uuid.UUID) (Deployment, error)
	// GetLaunch should return a Launch by ID
	GetLaunch(id uuid.UUID) (Launch, error)
	// CreateLaunch should create a Launch returning Launch with ID
	CreateLaunch(launch Launch) (Launch, error)
	// UpdateLaunch should update a Launch by ID
	UpdateLaunch(launch Launch) (Launch, error)
	// DeleteLaunch should delete a Launch by ID
	DeleteLaunch(id uuid.UUID) error
}

// PlatformInstance composes properties associated with the platform instance initiating the launch
// optional https://purl.imsglobal.org/spec/lti/claim/tool_platform claim
// In a multi-tenancy case, a single platform (iss) will host multiple instances,
// but each LTI message is originating from a single instance identified by its guid.
type PlatformInstance struct {
	// ID (REQUIRED) is the tools UUID for the Platform Instance
	ID uuid.UUID
	// GUID (REQUIRED)
	// A stable locally unique to the iss identifier for an instance of the tool platform.
	// The value of guid is a case-sensitive string that MUST NOT exceed 255 ASCII characters in length.
	// The use of Universally Unique IDentifier (UUID) defined in [RFC4122] is recommended.
	GUID string
	// ContactEmail (OPTIONAL). Administrative contact email for the platform instance.
	ContactEmail string
	// Description (OPTIONAL). Descriptive phrase for the platform instance.
	Description string
	// Name (OPTIONAL). Name for the platform instance.
	Name string
	// URL (OPTIONAL). Home HTTPS URL endpoint for the platform instance.
	URL string
	// ProductFamilyCode (OPTIONAL). Vendor product family code for the type of platform.
	ProductFamilyCode string
	// Version (OPTIONAL). Vendor product version for the platform.
	Version string
}

// Registration is the installation of an LTI in a platform instance
type Registration struct {
	// ID is the tools UUID for the Tool Registration
	ID uuid.UUID
	// PlatformInstanceID (REQUIRED) is the Platform Instance in which the Tool is Registered
	PlatformInstanceID uuid.UUID
	// ClientID (REQUIRED) or client_id is the tool registration ID in the Platform Instance
	ClientID string
	// Issuer (REQUIRED) is the id_token JWT issuer
	Issuer string
	// KeySetURL (REQUIRED) or JWK URL for the PlatformInstance
	KeySetURL string
	// AuthLoginURL (REQUIRED) is the redirect_uri for the PlatformInstance launch
	AuthLoginURL string
}

// Deployment
// A deployment of a tool defines the scope of contexts under which a tool is made available.
// For example, a tool may be deployed by the instructor into a single course,
// or the institution may deploy a tool across the whole institution, available to all institution's contexts,
// present and future. For further details see https://www.imsglobal.org/spec/lti/v1p3#tool-deployment
type Deployment struct {
	// ID (REQUIRED) is the tools UUID for the Tool Deployment
	ID uuid.UUID
	// PlatformDeploymentID (REQUIRED)
	// When a user deploys a tool within their tool platform, the platform MUST generate an immutable deployment_id
	// identifier to identify the integration.
	PlatformDeploymentID string
	// RegistrationID (REQUIRED) associates the Deployment back to the Registration
	// of the Tool in the Platform Instance
	RegistrationID uuid.UUID
	// Name (OPTIONAL) for the Deployment location e.g. Global or Sub account
	Name string
	// Description (OPTIONAL) of the Deployment e.g. In Course Navigation
	Description string
}

// Launch is the individual user launch event in the Platform
type Launch struct {
	// ID (REQUIRED) is the tools UUID for the Launch event
	ID uuid.UUID
	// Nonce (REQUIRED) is the tools UNIQUE Nonce for the Launch event
	Nonce string
	// RegistrationID (REQUIRED) is the ID of the Registration in which the Launch event occurred
	RegistrationID uuid.UUID
	// DeploymentID (OPTIONAL) is the ID of the Deployment in which the Launch event occurred
	DeploymentID uuid.NullUUID
	// Used (OPTIONAL) is the timestamp of when the Launch was completed, upon completion the
	// Launch should not be reusable (whether querying by Nonce or ID)
	Used time.Time
}
