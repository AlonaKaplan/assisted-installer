// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"

	"github.com/openshift/assisted-service/restapi/operations"
	"github.com/openshift/assisted-service/restapi/operations/events"
	"github.com/openshift/assisted-service/restapi/operations/installer"
	"github.com/openshift/assisted-service/restapi/operations/managed_domains"
	"github.com/openshift/assisted-service/restapi/operations/manifests"
	"github.com/openshift/assisted-service/restapi/operations/operators"
	"github.com/openshift/assisted-service/restapi/operations/versions"
)

type contextKey string

const AuthKey contextKey = "Auth"

//go:generate mockery -name EventsAPI -inpkg

/* EventsAPI  */
type EventsAPI interface {
	/* V2ListEvents Lists events for a cluster. */
	V2ListEvents(ctx context.Context, params events.V2ListEventsParams) middleware.Responder
}

//go:generate mockery -name InstallerAPI -inpkg

/* InstallerAPI  */
type InstallerAPI interface {
	/* BindHost Bind host to a cluster */
	BindHost(ctx context.Context, params installer.BindHostParams) middleware.Responder

	/* DeregisterInfraEnv Deletes an infra-env. */
	DeregisterInfraEnv(ctx context.Context, params installer.DeregisterInfraEnvParams) middleware.Responder

	/* DownloadMinimalInitrd Get the initial ramdisk for minimal ISO based installations.
	 */
	DownloadMinimalInitrd(ctx context.Context, params installer.DownloadMinimalInitrdParams) middleware.Responder

	/* GetClusterSupportedPlatforms A list of platforms that this cluster can support in its current configuration. */
	GetClusterSupportedPlatforms(ctx context.Context, params installer.GetClusterSupportedPlatformsParams) middleware.Responder

	/* GetInfraEnv Retrieves the details of the infra-env. */
	GetInfraEnv(ctx context.Context, params installer.GetInfraEnvParams) middleware.Responder

	/* GetInfraEnvDownloadURL Creates a new pre-signed image download URL for the infra-env. */
	GetInfraEnvDownloadURL(ctx context.Context, params installer.GetInfraEnvDownloadURLParams) middleware.Responder

	/* GetInfraEnvPresignedFileURL Creates a new pre-signed download URL for the infra-env. */
	GetInfraEnvPresignedFileURL(ctx context.Context, params installer.GetInfraEnvPresignedFileURLParams) middleware.Responder

	/* ListClusterHosts Get a list of cluster hosts according to supplied filters. */
	ListClusterHosts(ctx context.Context, params installer.ListClusterHostsParams) middleware.Responder

	/* ListInfraEnvs Retrieves the list of infra-envs. */
	ListInfraEnvs(ctx context.Context, params installer.ListInfraEnvsParams) middleware.Responder

	/* RegenerateInfraEnvSigningKey Regenerate InfraEnv token signing key. */
	RegenerateInfraEnvSigningKey(ctx context.Context, params installer.RegenerateInfraEnvSigningKeyParams) middleware.Responder

	/* RegisterInfraEnv Creates a new OpenShift Discovery ISO. */
	RegisterInfraEnv(ctx context.Context, params installer.RegisterInfraEnvParams) middleware.Responder

	/* TransformClusterToDay2 Transforming cluster to day2 and allowing adding hosts */
	TransformClusterToDay2(ctx context.Context, params installer.TransformClusterToDay2Params) middleware.Responder

	/* UnbindHost Unbind host to a cluster */
	UnbindHost(ctx context.Context, params installer.UnbindHostParams) middleware.Responder

	/* UpdateInfraEnv Updates an infra-env. */
	UpdateInfraEnv(ctx context.Context, params installer.UpdateInfraEnvParams) middleware.Responder

	/* V2CancelInstallation Cancels an ongoing installation. */
	V2CancelInstallation(ctx context.Context, params installer.V2CancelInstallationParams) middleware.Responder

	/* V2DownloadClusterCredentials Downloads credentials relating to the installed/installing cluster. */
	V2DownloadClusterCredentials(ctx context.Context, params installer.V2DownloadClusterCredentialsParams) middleware.Responder

	/* V2DownloadClusterFiles Downloads files relating to the installed/installing cluster. */
	V2DownloadClusterFiles(ctx context.Context, params installer.V2DownloadClusterFilesParams) middleware.Responder

	/* V2DownloadClusterLogs Download cluster logs. */
	V2DownloadClusterLogs(ctx context.Context, params installer.V2DownloadClusterLogsParams) middleware.Responder

	/* V2GetClusterDefaultConfig Get the default values for various cluster properties. */
	V2GetClusterDefaultConfig(ctx context.Context, params installer.V2GetClusterDefaultConfigParams) middleware.Responder

	/* V2GetCredentials Get the cluster admin credentials. */
	V2GetCredentials(ctx context.Context, params installer.V2GetCredentialsParams) middleware.Responder

	/* V2GetPresignedForClusterCredentials Get the cluster admin credentials. */
	V2GetPresignedForClusterCredentials(ctx context.Context, params installer.V2GetPresignedForClusterCredentialsParams) middleware.Responder

	/* V2GetPresignedForClusterFiles Retrieves a pre-signed S3 URL for downloading cluster files. */
	V2GetPresignedForClusterFiles(ctx context.Context, params installer.V2GetPresignedForClusterFilesParams) middleware.Responder

	/* V2UpdateCluster Updates an OpenShift cluster definition. */
	V2UpdateCluster(ctx context.Context, params installer.V2UpdateClusterParams) middleware.Responder

	/* V2UploadLogs Agent API to upload logs. */
	V2UploadLogs(ctx context.Context, params installer.V2UploadLogsParams) middleware.Responder

	/* V2CompleteInstallation Agent API to mark a finalizing installation as complete. */
	V2CompleteInstallation(ctx context.Context, params installer.V2CompleteInstallationParams) middleware.Responder

	/* V2DeregisterCluster Deletes an OpenShift cluster definition. */
	V2DeregisterCluster(ctx context.Context, params installer.V2DeregisterClusterParams) middleware.Responder

	/* V2DeregisterHost Deregisters an OpenShift host. */
	V2DeregisterHost(ctx context.Context, params installer.V2DeregisterHostParams) middleware.Responder

	/* V2DownloadHostIgnition Downloads the customized ignition file for this bound host, produces octet stream. For unbound host - error is returned */
	V2DownloadHostIgnition(ctx context.Context, params installer.V2DownloadHostIgnitionParams) middleware.Responder

	/* V2DownloadInfraEnvFiles Downloads the customized ignition file for this host */
	V2DownloadInfraEnvFiles(ctx context.Context, params installer.V2DownloadInfraEnvFilesParams) middleware.Responder

	/* V2GetCluster Retrieves the details of the OpenShift cluster. */
	V2GetCluster(ctx context.Context, params installer.V2GetClusterParams) middleware.Responder

	/* V2GetClusterInstallConfig Get the cluster's install config YAML. */
	V2GetClusterInstallConfig(ctx context.Context, params installer.V2GetClusterInstallConfigParams) middleware.Responder

	/* V2GetHost Retrieves the details of the OpenShift host. */
	V2GetHost(ctx context.Context, params installer.V2GetHostParams) middleware.Responder

	/* V2GetHostIgnition Fetch the ignition file for this host as a string. In case of unbound host produces an error */
	V2GetHostIgnition(ctx context.Context, params installer.V2GetHostIgnitionParams) middleware.Responder

	/* V2GetNextSteps Retrieves the next operations that the host agent needs to perform. */
	V2GetNextSteps(ctx context.Context, params installer.V2GetNextStepsParams) middleware.Responder

	/* V2GetPreflightRequirements Get preflight requirements for a cluster. */
	V2GetPreflightRequirements(ctx context.Context, params installer.V2GetPreflightRequirementsParams) middleware.Responder

	/* V2ImportCluster Import an AI cluster using minimal data assosiated with existing OCP cluster, in order to allow adding day2 hosts to that cluster */
	V2ImportCluster(ctx context.Context, params installer.V2ImportClusterParams) middleware.Responder

	/* V2InstallCluster Installs the OpenShift cluster. */
	V2InstallCluster(ctx context.Context, params installer.V2InstallClusterParams) middleware.Responder

	/* V2InstallHost install specific host for day2 cluster. */
	V2InstallHost(ctx context.Context, params installer.V2InstallHostParams) middleware.Responder

	/* V2ListClusters Retrieves the list of OpenShift clusters. */
	V2ListClusters(ctx context.Context, params installer.V2ListClustersParams) middleware.Responder

	/* V2ListFeatureSupportLevels Retrieves the support levels for features for each OpenShift version. */
	V2ListFeatureSupportLevels(ctx context.Context, params installer.V2ListFeatureSupportLevelsParams) middleware.Responder

	/* V2ListHosts Retrieves the list of OpenShift hosts that belong the infra-env. */
	V2ListHosts(ctx context.Context, params installer.V2ListHostsParams) middleware.Responder

	/* V2PostStepReply Posts the result of the operations from the host agent. */
	V2PostStepReply(ctx context.Context, params installer.V2PostStepReplyParams) middleware.Responder

	/* V2RegisterCluster Creates a new OpenShift cluster definition. */
	V2RegisterCluster(ctx context.Context, params installer.V2RegisterClusterParams) middleware.Responder

	/* V2RegisterHost Registers a new OpenShift agent. */
	V2RegisterHost(ctx context.Context, params installer.V2RegisterHostParams) middleware.Responder

	/* V2ResetCluster Resets a failed installation. */
	V2ResetCluster(ctx context.Context, params installer.V2ResetClusterParams) middleware.Responder

	/* V2ResetHost reset a failed host for day2 cluster. */
	V2ResetHost(ctx context.Context, params installer.V2ResetHostParams) middleware.Responder

	/* V2ResetHostValidation Reset failed host validation. */
	V2ResetHostValidation(ctx context.Context, params installer.V2ResetHostValidationParams) middleware.Responder

	/* V2UpdateClusterInstallConfig Override values in the install config. */
	V2UpdateClusterInstallConfig(ctx context.Context, params installer.V2UpdateClusterInstallConfigParams) middleware.Responder

	/* V2UpdateClusterLogsProgress Update log collection state and progress. */
	V2UpdateClusterLogsProgress(ctx context.Context, params installer.V2UpdateClusterLogsProgressParams) middleware.Responder

	/* V2UpdateHost Update an Openshift host */
	V2UpdateHost(ctx context.Context, params installer.V2UpdateHostParams) middleware.Responder

	/* V2UpdateHostIgnition Patch the ignition file for this host */
	V2UpdateHostIgnition(ctx context.Context, params installer.V2UpdateHostIgnitionParams) middleware.Responder

	/* V2UpdateHostInstallProgress Update installation progress. */
	V2UpdateHostInstallProgress(ctx context.Context, params installer.V2UpdateHostInstallProgressParams) middleware.Responder

	/* V2UpdateHostInstallerArgs Updates a host's installer arguments. */
	V2UpdateHostInstallerArgs(ctx context.Context, params installer.V2UpdateHostInstallerArgsParams) middleware.Responder

	/* V2UpdateHostLogsProgress Update log collection state and progress. */
	V2UpdateHostLogsProgress(ctx context.Context, params installer.V2UpdateHostLogsProgressParams) middleware.Responder

	/* V2UploadClusterIngressCert Transfer the ingress certificate for the cluster. */
	V2UploadClusterIngressCert(ctx context.Context, params installer.V2UploadClusterIngressCertParams) middleware.Responder
}

//go:generate mockery -name ManagedDomainsAPI -inpkg

/* ManagedDomainsAPI  */
type ManagedDomainsAPI interface {
	/* V2ListManagedDomains List of managed DNS domains. */
	V2ListManagedDomains(ctx context.Context, params managed_domains.V2ListManagedDomainsParams) middleware.Responder
}

//go:generate mockery -name ManifestsAPI -inpkg

/* ManifestsAPI  */
type ManifestsAPI interface {
	/* V2CreateClusterManifest Creates a manifest for customizing cluster installation. */
	V2CreateClusterManifest(ctx context.Context, params manifests.V2CreateClusterManifestParams) middleware.Responder

	/* V2DeleteClusterManifest Deletes a manifest from the cluster. */
	V2DeleteClusterManifest(ctx context.Context, params manifests.V2DeleteClusterManifestParams) middleware.Responder

	/* V2ListClusterManifests Lists manifests for customizing cluster installation. */
	V2ListClusterManifests(ctx context.Context, params manifests.V2ListClusterManifestsParams) middleware.Responder

	/* V2DownloadClusterManifest Downloads cluster manifest. */
	V2DownloadClusterManifest(ctx context.Context, params manifests.V2DownloadClusterManifestParams) middleware.Responder
}

//go:generate mockery -name OperatorsAPI -inpkg

/* OperatorsAPI  */
type OperatorsAPI interface {
	/* V2ListOfClusterOperators Lists operators to be monitored for a cluster. */
	V2ListOfClusterOperators(ctx context.Context, params operators.V2ListOfClusterOperatorsParams) middleware.Responder

	/* V2ListOperatorProperties Lists properties for an operator. */
	V2ListOperatorProperties(ctx context.Context, params operators.V2ListOperatorPropertiesParams) middleware.Responder

	/* V2ListSupportedOperators Retrieves the list of supported operators. */
	V2ListSupportedOperators(ctx context.Context, params operators.V2ListSupportedOperatorsParams) middleware.Responder

	/* V2ReportMonitoredOperatorStatus Controller API to report of monitored operators. */
	V2ReportMonitoredOperatorStatus(ctx context.Context, params operators.V2ReportMonitoredOperatorStatusParams) middleware.Responder
}

//go:generate mockery -name VersionsAPI -inpkg

/* VersionsAPI  */
type VersionsAPI interface {
	/* V2ListComponentVersions List of component versions. */
	V2ListComponentVersions(ctx context.Context, params versions.V2ListComponentVersionsParams) middleware.Responder

	/* V2ListSupportedOpenshiftVersions Retrieves the list of OpenShift supported versions. */
	V2ListSupportedOpenshiftVersions(ctx context.Context, params versions.V2ListSupportedOpenshiftVersionsParams) middleware.Responder
}

// Config is configuration for Handler
type Config struct {
	EventsAPI
	InstallerAPI
	ManagedDomainsAPI
	ManifestsAPI
	OperatorsAPI
	VersionsAPI
	Logger func(string, ...interface{})
	// InnerMiddleware is for the handler executors. These do not apply to the swagger.json document.
	// The middleware executes after routing but before authentication, binding and validation
	InnerMiddleware func(http.Handler) http.Handler

	// Authorizer is used to authorize a request after the Auth function was called using the "Auth*" functions
	// and the principal was stored in the context in the "AuthKey" context value.
	Authorizer func(*http.Request) error

	// AuthAgentAuth Applies when the "X-Secret-Key" header is set
	AuthAgentAuth func(token string) (interface{}, error)

	// AuthImageAuth Applies when the "Image-Token" header is set
	AuthImageAuth func(token string) (interface{}, error)

	// AuthImageURLAuth Applies when the "image_token" query is set
	AuthImageURLAuth func(token string) (interface{}, error)

	// AuthURLAuth Applies when the "api_key" query is set
	AuthURLAuth func(token string) (interface{}, error)

	// AuthUserAuth Applies when the "Authorization" header is set
	AuthUserAuth func(token string) (interface{}, error)

	// Authenticator to use for all APIKey authentication
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// Authenticator to use for all Bearer authentication
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// Authenticator to use for all Basic authentication
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator
}

// Handler returns an http.Handler given the handler configuration
// It mounts all the business logic implementers in the right routing.
func Handler(c Config) (http.Handler, error) {
	h, _, err := HandlerAPI(c)
	return h, err
}

// HandlerAPI returns an http.Handler given the handler configuration
// and the corresponding *AssistedInstall instance.
// It mounts all the business logic implementers in the right routing.
func HandlerAPI(c Config) (http.Handler, *operations.AssistedInstallAPI, error) {
	spec, err := loads.Analyzed(swaggerCopy(SwaggerJSON), "")
	if err != nil {
		return nil, nil, fmt.Errorf("analyze swagger: %v", err)
	}
	api := operations.NewAssistedInstallAPI(spec)
	api.ServeError = errors.ServeError
	api.Logger = c.Logger

	if c.APIKeyAuthenticator != nil {
		api.APIKeyAuthenticator = c.APIKeyAuthenticator
	}
	if c.BasicAuthenticator != nil {
		api.BasicAuthenticator = c.BasicAuthenticator
	}
	if c.BearerAuthenticator != nil {
		api.BearerAuthenticator = c.BearerAuthenticator
	}

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer
	api.BinProducer = runtime.ByteStreamProducer()
	api.JSONProducer = runtime.JSONProducer()
	api.AgentAuthAuth = func(token string) (interface{}, error) {
		if c.AuthAgentAuth == nil {
			return token, nil
		}
		return c.AuthAgentAuth(token)
	}

	api.ImageAuthAuth = func(token string) (interface{}, error) {
		if c.AuthImageAuth == nil {
			return token, nil
		}
		return c.AuthImageAuth(token)
	}

	api.ImageURLAuthAuth = func(token string) (interface{}, error) {
		if c.AuthImageURLAuth == nil {
			return token, nil
		}
		return c.AuthImageURLAuth(token)
	}

	api.URLAuthAuth = func(token string) (interface{}, error) {
		if c.AuthURLAuth == nil {
			return token, nil
		}
		return c.AuthURLAuth(token)
	}

	api.UserAuthAuth = func(token string) (interface{}, error) {
		if c.AuthUserAuth == nil {
			return token, nil
		}
		return c.AuthUserAuth(token)
	}

	api.APIAuthorizer = authorizer(c.Authorizer)
	api.InstallerBindHostHandler = installer.BindHostHandlerFunc(func(params installer.BindHostParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.BindHost(ctx, params)
	})
	api.InstallerDeregisterInfraEnvHandler = installer.DeregisterInfraEnvHandlerFunc(func(params installer.DeregisterInfraEnvParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.DeregisterInfraEnv(ctx, params)
	})
	api.InstallerDownloadMinimalInitrdHandler = installer.DownloadMinimalInitrdHandlerFunc(func(params installer.DownloadMinimalInitrdParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.DownloadMinimalInitrd(ctx, params)
	})
	api.InstallerGetClusterSupportedPlatformsHandler = installer.GetClusterSupportedPlatformsHandlerFunc(func(params installer.GetClusterSupportedPlatformsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.GetClusterSupportedPlatforms(ctx, params)
	})
	api.InstallerGetInfraEnvHandler = installer.GetInfraEnvHandlerFunc(func(params installer.GetInfraEnvParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.GetInfraEnv(ctx, params)
	})
	api.InstallerGetInfraEnvDownloadURLHandler = installer.GetInfraEnvDownloadURLHandlerFunc(func(params installer.GetInfraEnvDownloadURLParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.GetInfraEnvDownloadURL(ctx, params)
	})
	api.InstallerGetInfraEnvPresignedFileURLHandler = installer.GetInfraEnvPresignedFileURLHandlerFunc(func(params installer.GetInfraEnvPresignedFileURLParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.GetInfraEnvPresignedFileURL(ctx, params)
	})
	api.InstallerListClusterHostsHandler = installer.ListClusterHostsHandlerFunc(func(params installer.ListClusterHostsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.ListClusterHosts(ctx, params)
	})
	api.InstallerListInfraEnvsHandler = installer.ListInfraEnvsHandlerFunc(func(params installer.ListInfraEnvsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.ListInfraEnvs(ctx, params)
	})
	api.InstallerRegenerateInfraEnvSigningKeyHandler = installer.RegenerateInfraEnvSigningKeyHandlerFunc(func(params installer.RegenerateInfraEnvSigningKeyParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.RegenerateInfraEnvSigningKey(ctx, params)
	})
	api.InstallerRegisterInfraEnvHandler = installer.RegisterInfraEnvHandlerFunc(func(params installer.RegisterInfraEnvParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.RegisterInfraEnv(ctx, params)
	})
	api.InstallerTransformClusterToDay2Handler = installer.TransformClusterToDay2HandlerFunc(func(params installer.TransformClusterToDay2Params, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.TransformClusterToDay2(ctx, params)
	})
	api.InstallerUnbindHostHandler = installer.UnbindHostHandlerFunc(func(params installer.UnbindHostParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.UnbindHost(ctx, params)
	})
	api.InstallerUpdateInfraEnvHandler = installer.UpdateInfraEnvHandlerFunc(func(params installer.UpdateInfraEnvParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.UpdateInfraEnv(ctx, params)
	})
	api.InstallerV2CancelInstallationHandler = installer.V2CancelInstallationHandlerFunc(func(params installer.V2CancelInstallationParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2CancelInstallation(ctx, params)
	})
	api.ManifestsV2CreateClusterManifestHandler = manifests.V2CreateClusterManifestHandlerFunc(func(params manifests.V2CreateClusterManifestParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.ManifestsAPI.V2CreateClusterManifest(ctx, params)
	})
	api.ManifestsV2DeleteClusterManifestHandler = manifests.V2DeleteClusterManifestHandlerFunc(func(params manifests.V2DeleteClusterManifestParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.ManifestsAPI.V2DeleteClusterManifest(ctx, params)
	})
	api.InstallerV2DownloadClusterCredentialsHandler = installer.V2DownloadClusterCredentialsHandlerFunc(func(params installer.V2DownloadClusterCredentialsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2DownloadClusterCredentials(ctx, params)
	})
	api.InstallerV2DownloadClusterFilesHandler = installer.V2DownloadClusterFilesHandlerFunc(func(params installer.V2DownloadClusterFilesParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2DownloadClusterFiles(ctx, params)
	})
	api.InstallerV2DownloadClusterLogsHandler = installer.V2DownloadClusterLogsHandlerFunc(func(params installer.V2DownloadClusterLogsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2DownloadClusterLogs(ctx, params)
	})
	api.InstallerV2GetClusterDefaultConfigHandler = installer.V2GetClusterDefaultConfigHandlerFunc(func(params installer.V2GetClusterDefaultConfigParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2GetClusterDefaultConfig(ctx, params)
	})
	api.InstallerV2GetCredentialsHandler = installer.V2GetCredentialsHandlerFunc(func(params installer.V2GetCredentialsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2GetCredentials(ctx, params)
	})
	api.InstallerV2GetPresignedForClusterCredentialsHandler = installer.V2GetPresignedForClusterCredentialsHandlerFunc(func(params installer.V2GetPresignedForClusterCredentialsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2GetPresignedForClusterCredentials(ctx, params)
	})
	api.InstallerV2GetPresignedForClusterFilesHandler = installer.V2GetPresignedForClusterFilesHandlerFunc(func(params installer.V2GetPresignedForClusterFilesParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2GetPresignedForClusterFiles(ctx, params)
	})
	api.ManifestsV2ListClusterManifestsHandler = manifests.V2ListClusterManifestsHandlerFunc(func(params manifests.V2ListClusterManifestsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.ManifestsAPI.V2ListClusterManifests(ctx, params)
	})
	api.ManagedDomainsV2ListManagedDomainsHandler = managed_domains.V2ListManagedDomainsHandlerFunc(func(params managed_domains.V2ListManagedDomainsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.ManagedDomainsAPI.V2ListManagedDomains(ctx, params)
	})
	api.OperatorsV2ListOfClusterOperatorsHandler = operators.V2ListOfClusterOperatorsHandlerFunc(func(params operators.V2ListOfClusterOperatorsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.OperatorsAPI.V2ListOfClusterOperators(ctx, params)
	})
	api.OperatorsV2ListOperatorPropertiesHandler = operators.V2ListOperatorPropertiesHandlerFunc(func(params operators.V2ListOperatorPropertiesParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.OperatorsAPI.V2ListOperatorProperties(ctx, params)
	})
	api.OperatorsV2ListSupportedOperatorsHandler = operators.V2ListSupportedOperatorsHandlerFunc(func(params operators.V2ListSupportedOperatorsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.OperatorsAPI.V2ListSupportedOperators(ctx, params)
	})
	api.InstallerV2UpdateClusterHandler = installer.V2UpdateClusterHandlerFunc(func(params installer.V2UpdateClusterParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2UpdateCluster(ctx, params)
	})
	api.InstallerV2UploadLogsHandler = installer.V2UploadLogsHandlerFunc(func(params installer.V2UploadLogsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2UploadLogs(ctx, params)
	})
	api.InstallerV2CompleteInstallationHandler = installer.V2CompleteInstallationHandlerFunc(func(params installer.V2CompleteInstallationParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2CompleteInstallation(ctx, params)
	})
	api.InstallerV2DeregisterClusterHandler = installer.V2DeregisterClusterHandlerFunc(func(params installer.V2DeregisterClusterParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2DeregisterCluster(ctx, params)
	})
	api.InstallerV2DeregisterHostHandler = installer.V2DeregisterHostHandlerFunc(func(params installer.V2DeregisterHostParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2DeregisterHost(ctx, params)
	})
	api.ManifestsV2DownloadClusterManifestHandler = manifests.V2DownloadClusterManifestHandlerFunc(func(params manifests.V2DownloadClusterManifestParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.ManifestsAPI.V2DownloadClusterManifest(ctx, params)
	})
	api.InstallerV2DownloadHostIgnitionHandler = installer.V2DownloadHostIgnitionHandlerFunc(func(params installer.V2DownloadHostIgnitionParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2DownloadHostIgnition(ctx, params)
	})
	api.InstallerV2DownloadInfraEnvFilesHandler = installer.V2DownloadInfraEnvFilesHandlerFunc(func(params installer.V2DownloadInfraEnvFilesParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2DownloadInfraEnvFiles(ctx, params)
	})
	api.InstallerV2GetClusterHandler = installer.V2GetClusterHandlerFunc(func(params installer.V2GetClusterParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2GetCluster(ctx, params)
	})
	api.InstallerV2GetClusterInstallConfigHandler = installer.V2GetClusterInstallConfigHandlerFunc(func(params installer.V2GetClusterInstallConfigParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2GetClusterInstallConfig(ctx, params)
	})
	api.InstallerV2GetHostHandler = installer.V2GetHostHandlerFunc(func(params installer.V2GetHostParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2GetHost(ctx, params)
	})
	api.InstallerV2GetHostIgnitionHandler = installer.V2GetHostIgnitionHandlerFunc(func(params installer.V2GetHostIgnitionParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2GetHostIgnition(ctx, params)
	})
	api.InstallerV2GetNextStepsHandler = installer.V2GetNextStepsHandlerFunc(func(params installer.V2GetNextStepsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2GetNextSteps(ctx, params)
	})
	api.InstallerV2GetPreflightRequirementsHandler = installer.V2GetPreflightRequirementsHandlerFunc(func(params installer.V2GetPreflightRequirementsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2GetPreflightRequirements(ctx, params)
	})
	api.InstallerV2ImportClusterHandler = installer.V2ImportClusterHandlerFunc(func(params installer.V2ImportClusterParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2ImportCluster(ctx, params)
	})
	api.InstallerV2InstallClusterHandler = installer.V2InstallClusterHandlerFunc(func(params installer.V2InstallClusterParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2InstallCluster(ctx, params)
	})
	api.InstallerV2InstallHostHandler = installer.V2InstallHostHandlerFunc(func(params installer.V2InstallHostParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2InstallHost(ctx, params)
	})
	api.InstallerV2ListClustersHandler = installer.V2ListClustersHandlerFunc(func(params installer.V2ListClustersParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2ListClusters(ctx, params)
	})
	api.VersionsV2ListComponentVersionsHandler = versions.V2ListComponentVersionsHandlerFunc(func(params versions.V2ListComponentVersionsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.VersionsAPI.V2ListComponentVersions(ctx, params)
	})
	api.EventsV2ListEventsHandler = events.V2ListEventsHandlerFunc(func(params events.V2ListEventsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.EventsAPI.V2ListEvents(ctx, params)
	})
	api.InstallerV2ListFeatureSupportLevelsHandler = installer.V2ListFeatureSupportLevelsHandlerFunc(func(params installer.V2ListFeatureSupportLevelsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2ListFeatureSupportLevels(ctx, params)
	})
	api.InstallerV2ListHostsHandler = installer.V2ListHostsHandlerFunc(func(params installer.V2ListHostsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2ListHosts(ctx, params)
	})
	api.VersionsV2ListSupportedOpenshiftVersionsHandler = versions.V2ListSupportedOpenshiftVersionsHandlerFunc(func(params versions.V2ListSupportedOpenshiftVersionsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.VersionsAPI.V2ListSupportedOpenshiftVersions(ctx, params)
	})
	api.InstallerV2PostStepReplyHandler = installer.V2PostStepReplyHandlerFunc(func(params installer.V2PostStepReplyParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2PostStepReply(ctx, params)
	})
	api.InstallerV2RegisterClusterHandler = installer.V2RegisterClusterHandlerFunc(func(params installer.V2RegisterClusterParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2RegisterCluster(ctx, params)
	})
	api.InstallerV2RegisterHostHandler = installer.V2RegisterHostHandlerFunc(func(params installer.V2RegisterHostParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2RegisterHost(ctx, params)
	})
	api.OperatorsV2ReportMonitoredOperatorStatusHandler = operators.V2ReportMonitoredOperatorStatusHandlerFunc(func(params operators.V2ReportMonitoredOperatorStatusParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.OperatorsAPI.V2ReportMonitoredOperatorStatus(ctx, params)
	})
	api.InstallerV2ResetClusterHandler = installer.V2ResetClusterHandlerFunc(func(params installer.V2ResetClusterParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2ResetCluster(ctx, params)
	})
	api.InstallerV2ResetHostHandler = installer.V2ResetHostHandlerFunc(func(params installer.V2ResetHostParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2ResetHost(ctx, params)
	})
	api.InstallerV2ResetHostValidationHandler = installer.V2ResetHostValidationHandlerFunc(func(params installer.V2ResetHostValidationParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2ResetHostValidation(ctx, params)
	})
	api.InstallerV2UpdateClusterInstallConfigHandler = installer.V2UpdateClusterInstallConfigHandlerFunc(func(params installer.V2UpdateClusterInstallConfigParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2UpdateClusterInstallConfig(ctx, params)
	})
	api.InstallerV2UpdateClusterLogsProgressHandler = installer.V2UpdateClusterLogsProgressHandlerFunc(func(params installer.V2UpdateClusterLogsProgressParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2UpdateClusterLogsProgress(ctx, params)
	})
	api.InstallerV2UpdateHostHandler = installer.V2UpdateHostHandlerFunc(func(params installer.V2UpdateHostParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2UpdateHost(ctx, params)
	})
	api.InstallerV2UpdateHostIgnitionHandler = installer.V2UpdateHostIgnitionHandlerFunc(func(params installer.V2UpdateHostIgnitionParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2UpdateHostIgnition(ctx, params)
	})
	api.InstallerV2UpdateHostInstallProgressHandler = installer.V2UpdateHostInstallProgressHandlerFunc(func(params installer.V2UpdateHostInstallProgressParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2UpdateHostInstallProgress(ctx, params)
	})
	api.InstallerV2UpdateHostInstallerArgsHandler = installer.V2UpdateHostInstallerArgsHandlerFunc(func(params installer.V2UpdateHostInstallerArgsParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2UpdateHostInstallerArgs(ctx, params)
	})
	api.InstallerV2UpdateHostLogsProgressHandler = installer.V2UpdateHostLogsProgressHandlerFunc(func(params installer.V2UpdateHostLogsProgressParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2UpdateHostLogsProgress(ctx, params)
	})
	api.InstallerV2UploadClusterIngressCertHandler = installer.V2UploadClusterIngressCertHandlerFunc(func(params installer.V2UploadClusterIngressCertParams, principal interface{}) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		ctx = storeAuth(ctx, principal)
		return c.InstallerAPI.V2UploadClusterIngressCert(ctx, params)
	})
	api.ServerShutdown = func() {}
	return api.Serve(c.InnerMiddleware), api, nil
}

// swaggerCopy copies the swagger json to prevent data races in runtime
func swaggerCopy(orig json.RawMessage) json.RawMessage {
	c := make(json.RawMessage, len(orig))
	copy(c, orig)
	return c
}

// authorizer is a helper function to implement the runtime.Authorizer interface.
type authorizer func(*http.Request) error

func (a authorizer) Authorize(req *http.Request, principal interface{}) error {
	if a == nil {
		return nil
	}
	ctx := storeAuth(req.Context(), principal)
	return a(req.WithContext(ctx))
}

func storeAuth(ctx context.Context, principal interface{}) context.Context {
	return context.WithValue(ctx, AuthKey, principal)
}
