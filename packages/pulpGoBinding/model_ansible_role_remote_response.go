/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpGoBinding

import (
	"encoding/json"
	"time"
)

// checks if the AnsibleRoleRemoteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnsibleRoleRemoteResponse{}

// AnsibleRoleRemoteResponse A serializer for Ansible Remotes.
type AnsibleRoleRemoteResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// A unique name for this remote.
	Name string `json:"name"`
	// The URL of an external content source.
	Url string `json:"url"`
	// A PEM encoded CA certificate used to validate the server certificate presented by the remote server.
	CaCert NullableString `json:"ca_cert,omitempty"`
	// A PEM encoded client certificate used for authentication.
	ClientCert NullableString `json:"client_cert,omitempty"`
	// If True, TLS peer validation must be performed.
	TlsValidation *bool `json:"tls_validation,omitempty"`
	// The proxy URL. Format: scheme://host:port
	ProxyUrl NullableString `json:"proxy_url,omitempty"`
	PulpLabels *map[string]string `json:"pulp_labels,omitempty"`
	// Timestamp of the most recent update of the remote.
	PulpLastUpdated *time.Time `json:"pulp_last_updated,omitempty"`
	// Total number of simultaneous connections. If not set then the default value will be used.
	DownloadConcurrency NullableInt32 `json:"download_concurrency,omitempty"`
	// Maximum number of retry attempts after a download failure. If not set then the default value (3) will be used.
	MaxRetries NullableInt32 `json:"max_retries,omitempty"`
	Policy *PolicyDb6Enum `json:"policy,omitempty"`
	// aiohttp.ClientTimeout.total (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	TotalTimeout NullableFloat64 `json:"total_timeout,omitempty"`
	// aiohttp.ClientTimeout.connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	ConnectTimeout NullableFloat64 `json:"connect_timeout,omitempty"`
	// aiohttp.ClientTimeout.sock_connect (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	SockConnectTimeout NullableFloat64 `json:"sock_connect_timeout,omitempty"`
	// aiohttp.ClientTimeout.sock_read (q.v.) for download-connections. The default is null, which will cause the default from the aiohttp library to be used.
	SockReadTimeout NullableFloat64 `json:"sock_read_timeout,omitempty"`
	// Headers for aiohttp.Clientsession
	Headers []map[string]interface{} `json:"headers,omitempty"`
	// Limits requests per second for each concurrent downloader
	RateLimit NullableInt32 `json:"rate_limit,omitempty"`
	// List of hidden (write only) fields
	HiddenFields []RemoteResponseHiddenFieldsInner `json:"hidden_fields,omitempty"`
}

// NewAnsibleRoleRemoteResponse instantiates a new AnsibleRoleRemoteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnsibleRoleRemoteResponse(name string, url string) *AnsibleRoleRemoteResponse {
	this := AnsibleRoleRemoteResponse{}
	this.Name = name
	this.Url = url
	return &this
}

// NewAnsibleRoleRemoteResponseWithDefaults instantiates a new AnsibleRoleRemoteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnsibleRoleRemoteResponseWithDefaults() *AnsibleRoleRemoteResponse {
	this := AnsibleRoleRemoteResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *AnsibleRoleRemoteResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleRoleRemoteResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *AnsibleRoleRemoteResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *AnsibleRoleRemoteResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleRoleRemoteResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *AnsibleRoleRemoteResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetName returns the Name field value
func (o *AnsibleRoleRemoteResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnsibleRoleRemoteResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AnsibleRoleRemoteResponse) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *AnsibleRoleRemoteResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *AnsibleRoleRemoteResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *AnsibleRoleRemoteResponse) SetUrl(v string) {
	o.Url = v
}

// GetCaCert returns the CaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleRoleRemoteResponse) GetCaCert() string {
	if o == nil || IsNil(o.CaCert.Get()) {
		var ret string
		return ret
	}
	return *o.CaCert.Get()
}

// GetCaCertOk returns a tuple with the CaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleRoleRemoteResponse) GetCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaCert.Get(), o.CaCert.IsSet()
}

// HasCaCert returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasCaCert() bool {
	if o != nil && o.CaCert.IsSet() {
		return true
	}

	return false
}

// SetCaCert gets a reference to the given NullableString and assigns it to the CaCert field.
func (o *AnsibleRoleRemoteResponse) SetCaCert(v string) {
	o.CaCert.Set(&v)
}
// SetCaCertNil sets the value for CaCert to be an explicit nil
func (o *AnsibleRoleRemoteResponse) SetCaCertNil() {
	o.CaCert.Set(nil)
}

// UnsetCaCert ensures that no value is present for CaCert, not even an explicit nil
func (o *AnsibleRoleRemoteResponse) UnsetCaCert() {
	o.CaCert.Unset()
}

// GetClientCert returns the ClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleRoleRemoteResponse) GetClientCert() string {
	if o == nil || IsNil(o.ClientCert.Get()) {
		var ret string
		return ret
	}
	return *o.ClientCert.Get()
}

// GetClientCertOk returns a tuple with the ClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleRoleRemoteResponse) GetClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientCert.Get(), o.ClientCert.IsSet()
}

// HasClientCert returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasClientCert() bool {
	if o != nil && o.ClientCert.IsSet() {
		return true
	}

	return false
}

// SetClientCert gets a reference to the given NullableString and assigns it to the ClientCert field.
func (o *AnsibleRoleRemoteResponse) SetClientCert(v string) {
	o.ClientCert.Set(&v)
}
// SetClientCertNil sets the value for ClientCert to be an explicit nil
func (o *AnsibleRoleRemoteResponse) SetClientCertNil() {
	o.ClientCert.Set(nil)
}

// UnsetClientCert ensures that no value is present for ClientCert, not even an explicit nil
func (o *AnsibleRoleRemoteResponse) UnsetClientCert() {
	o.ClientCert.Unset()
}

// GetTlsValidation returns the TlsValidation field value if set, zero value otherwise.
func (o *AnsibleRoleRemoteResponse) GetTlsValidation() bool {
	if o == nil || IsNil(o.TlsValidation) {
		var ret bool
		return ret
	}
	return *o.TlsValidation
}

// GetTlsValidationOk returns a tuple with the TlsValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleRoleRemoteResponse) GetTlsValidationOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsValidation) {
		return nil, false
	}
	return o.TlsValidation, true
}

// HasTlsValidation returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasTlsValidation() bool {
	if o != nil && !IsNil(o.TlsValidation) {
		return true
	}

	return false
}

// SetTlsValidation gets a reference to the given bool and assigns it to the TlsValidation field.
func (o *AnsibleRoleRemoteResponse) SetTlsValidation(v bool) {
	o.TlsValidation = &v
}

// GetProxyUrl returns the ProxyUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleRoleRemoteResponse) GetProxyUrl() string {
	if o == nil || IsNil(o.ProxyUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ProxyUrl.Get()
}

// GetProxyUrlOk returns a tuple with the ProxyUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleRoleRemoteResponse) GetProxyUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProxyUrl.Get(), o.ProxyUrl.IsSet()
}

// HasProxyUrl returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasProxyUrl() bool {
	if o != nil && o.ProxyUrl.IsSet() {
		return true
	}

	return false
}

// SetProxyUrl gets a reference to the given NullableString and assigns it to the ProxyUrl field.
func (o *AnsibleRoleRemoteResponse) SetProxyUrl(v string) {
	o.ProxyUrl.Set(&v)
}
// SetProxyUrlNil sets the value for ProxyUrl to be an explicit nil
func (o *AnsibleRoleRemoteResponse) SetProxyUrlNil() {
	o.ProxyUrl.Set(nil)
}

// UnsetProxyUrl ensures that no value is present for ProxyUrl, not even an explicit nil
func (o *AnsibleRoleRemoteResponse) UnsetProxyUrl() {
	o.ProxyUrl.Unset()
}

// GetPulpLabels returns the PulpLabels field value if set, zero value otherwise.
func (o *AnsibleRoleRemoteResponse) GetPulpLabels() map[string]string {
	if o == nil || IsNil(o.PulpLabels) {
		var ret map[string]string
		return ret
	}
	return *o.PulpLabels
}

// GetPulpLabelsOk returns a tuple with the PulpLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleRoleRemoteResponse) GetPulpLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.PulpLabels) {
		return nil, false
	}
	return o.PulpLabels, true
}

// HasPulpLabels returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasPulpLabels() bool {
	if o != nil && !IsNil(o.PulpLabels) {
		return true
	}

	return false
}

// SetPulpLabels gets a reference to the given map[string]string and assigns it to the PulpLabels field.
func (o *AnsibleRoleRemoteResponse) SetPulpLabels(v map[string]string) {
	o.PulpLabels = &v
}

// GetPulpLastUpdated returns the PulpLastUpdated field value if set, zero value otherwise.
func (o *AnsibleRoleRemoteResponse) GetPulpLastUpdated() time.Time {
	if o == nil || IsNil(o.PulpLastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.PulpLastUpdated
}

// GetPulpLastUpdatedOk returns a tuple with the PulpLastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleRoleRemoteResponse) GetPulpLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpLastUpdated) {
		return nil, false
	}
	return o.PulpLastUpdated, true
}

// HasPulpLastUpdated returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasPulpLastUpdated() bool {
	if o != nil && !IsNil(o.PulpLastUpdated) {
		return true
	}

	return false
}

// SetPulpLastUpdated gets a reference to the given time.Time and assigns it to the PulpLastUpdated field.
func (o *AnsibleRoleRemoteResponse) SetPulpLastUpdated(v time.Time) {
	o.PulpLastUpdated = &v
}

// GetDownloadConcurrency returns the DownloadConcurrency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleRoleRemoteResponse) GetDownloadConcurrency() int32 {
	if o == nil || IsNil(o.DownloadConcurrency.Get()) {
		var ret int32
		return ret
	}
	return *o.DownloadConcurrency.Get()
}

// GetDownloadConcurrencyOk returns a tuple with the DownloadConcurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleRoleRemoteResponse) GetDownloadConcurrencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadConcurrency.Get(), o.DownloadConcurrency.IsSet()
}

// HasDownloadConcurrency returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasDownloadConcurrency() bool {
	if o != nil && o.DownloadConcurrency.IsSet() {
		return true
	}

	return false
}

// SetDownloadConcurrency gets a reference to the given NullableInt32 and assigns it to the DownloadConcurrency field.
func (o *AnsibleRoleRemoteResponse) SetDownloadConcurrency(v int32) {
	o.DownloadConcurrency.Set(&v)
}
// SetDownloadConcurrencyNil sets the value for DownloadConcurrency to be an explicit nil
func (o *AnsibleRoleRemoteResponse) SetDownloadConcurrencyNil() {
	o.DownloadConcurrency.Set(nil)
}

// UnsetDownloadConcurrency ensures that no value is present for DownloadConcurrency, not even an explicit nil
func (o *AnsibleRoleRemoteResponse) UnsetDownloadConcurrency() {
	o.DownloadConcurrency.Unset()
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleRoleRemoteResponse) GetMaxRetries() int32 {
	if o == nil || IsNil(o.MaxRetries.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxRetries.Get()
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleRoleRemoteResponse) GetMaxRetriesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRetries.Get(), o.MaxRetries.IsSet()
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasMaxRetries() bool {
	if o != nil && o.MaxRetries.IsSet() {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given NullableInt32 and assigns it to the MaxRetries field.
func (o *AnsibleRoleRemoteResponse) SetMaxRetries(v int32) {
	o.MaxRetries.Set(&v)
}
// SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil
func (o *AnsibleRoleRemoteResponse) SetMaxRetriesNil() {
	o.MaxRetries.Set(nil)
}

// UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
func (o *AnsibleRoleRemoteResponse) UnsetMaxRetries() {
	o.MaxRetries.Unset()
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *AnsibleRoleRemoteResponse) GetPolicy() PolicyDb6Enum {
	if o == nil || IsNil(o.Policy) {
		var ret PolicyDb6Enum
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleRoleRemoteResponse) GetPolicyOk() (*PolicyDb6Enum, bool) {
	if o == nil || IsNil(o.Policy) {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasPolicy() bool {
	if o != nil && !IsNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given PolicyDb6Enum and assigns it to the Policy field.
func (o *AnsibleRoleRemoteResponse) SetPolicy(v PolicyDb6Enum) {
	o.Policy = &v
}

// GetTotalTimeout returns the TotalTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleRoleRemoteResponse) GetTotalTimeout() float64 {
	if o == nil || IsNil(o.TotalTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.TotalTimeout.Get()
}

// GetTotalTimeoutOk returns a tuple with the TotalTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleRoleRemoteResponse) GetTotalTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalTimeout.Get(), o.TotalTimeout.IsSet()
}

// HasTotalTimeout returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasTotalTimeout() bool {
	if o != nil && o.TotalTimeout.IsSet() {
		return true
	}

	return false
}

// SetTotalTimeout gets a reference to the given NullableFloat64 and assigns it to the TotalTimeout field.
func (o *AnsibleRoleRemoteResponse) SetTotalTimeout(v float64) {
	o.TotalTimeout.Set(&v)
}
// SetTotalTimeoutNil sets the value for TotalTimeout to be an explicit nil
func (o *AnsibleRoleRemoteResponse) SetTotalTimeoutNil() {
	o.TotalTimeout.Set(nil)
}

// UnsetTotalTimeout ensures that no value is present for TotalTimeout, not even an explicit nil
func (o *AnsibleRoleRemoteResponse) UnsetTotalTimeout() {
	o.TotalTimeout.Unset()
}

// GetConnectTimeout returns the ConnectTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleRoleRemoteResponse) GetConnectTimeout() float64 {
	if o == nil || IsNil(o.ConnectTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.ConnectTimeout.Get()
}

// GetConnectTimeoutOk returns a tuple with the ConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleRoleRemoteResponse) GetConnectTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectTimeout.Get(), o.ConnectTimeout.IsSet()
}

// HasConnectTimeout returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasConnectTimeout() bool {
	if o != nil && o.ConnectTimeout.IsSet() {
		return true
	}

	return false
}

// SetConnectTimeout gets a reference to the given NullableFloat64 and assigns it to the ConnectTimeout field.
func (o *AnsibleRoleRemoteResponse) SetConnectTimeout(v float64) {
	o.ConnectTimeout.Set(&v)
}
// SetConnectTimeoutNil sets the value for ConnectTimeout to be an explicit nil
func (o *AnsibleRoleRemoteResponse) SetConnectTimeoutNil() {
	o.ConnectTimeout.Set(nil)
}

// UnsetConnectTimeout ensures that no value is present for ConnectTimeout, not even an explicit nil
func (o *AnsibleRoleRemoteResponse) UnsetConnectTimeout() {
	o.ConnectTimeout.Unset()
}

// GetSockConnectTimeout returns the SockConnectTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleRoleRemoteResponse) GetSockConnectTimeout() float64 {
	if o == nil || IsNil(o.SockConnectTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.SockConnectTimeout.Get()
}

// GetSockConnectTimeoutOk returns a tuple with the SockConnectTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleRoleRemoteResponse) GetSockConnectTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SockConnectTimeout.Get(), o.SockConnectTimeout.IsSet()
}

// HasSockConnectTimeout returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasSockConnectTimeout() bool {
	if o != nil && o.SockConnectTimeout.IsSet() {
		return true
	}

	return false
}

// SetSockConnectTimeout gets a reference to the given NullableFloat64 and assigns it to the SockConnectTimeout field.
func (o *AnsibleRoleRemoteResponse) SetSockConnectTimeout(v float64) {
	o.SockConnectTimeout.Set(&v)
}
// SetSockConnectTimeoutNil sets the value for SockConnectTimeout to be an explicit nil
func (o *AnsibleRoleRemoteResponse) SetSockConnectTimeoutNil() {
	o.SockConnectTimeout.Set(nil)
}

// UnsetSockConnectTimeout ensures that no value is present for SockConnectTimeout, not even an explicit nil
func (o *AnsibleRoleRemoteResponse) UnsetSockConnectTimeout() {
	o.SockConnectTimeout.Unset()
}

// GetSockReadTimeout returns the SockReadTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleRoleRemoteResponse) GetSockReadTimeout() float64 {
	if o == nil || IsNil(o.SockReadTimeout.Get()) {
		var ret float64
		return ret
	}
	return *o.SockReadTimeout.Get()
}

// GetSockReadTimeoutOk returns a tuple with the SockReadTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleRoleRemoteResponse) GetSockReadTimeoutOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SockReadTimeout.Get(), o.SockReadTimeout.IsSet()
}

// HasSockReadTimeout returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasSockReadTimeout() bool {
	if o != nil && o.SockReadTimeout.IsSet() {
		return true
	}

	return false
}

// SetSockReadTimeout gets a reference to the given NullableFloat64 and assigns it to the SockReadTimeout field.
func (o *AnsibleRoleRemoteResponse) SetSockReadTimeout(v float64) {
	o.SockReadTimeout.Set(&v)
}
// SetSockReadTimeoutNil sets the value for SockReadTimeout to be an explicit nil
func (o *AnsibleRoleRemoteResponse) SetSockReadTimeoutNil() {
	o.SockReadTimeout.Set(nil)
}

// UnsetSockReadTimeout ensures that no value is present for SockReadTimeout, not even an explicit nil
func (o *AnsibleRoleRemoteResponse) UnsetSockReadTimeout() {
	o.SockReadTimeout.Unset()
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *AnsibleRoleRemoteResponse) GetHeaders() []map[string]interface{} {
	if o == nil || IsNil(o.Headers) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleRoleRemoteResponse) GetHeadersOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []map[string]interface{} and assigns it to the Headers field.
func (o *AnsibleRoleRemoteResponse) SetHeaders(v []map[string]interface{}) {
	o.Headers = v
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnsibleRoleRemoteResponse) GetRateLimit() int32 {
	if o == nil || IsNil(o.RateLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.RateLimit.Get()
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AnsibleRoleRemoteResponse) GetRateLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RateLimit.Get(), o.RateLimit.IsSet()
}

// HasRateLimit returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasRateLimit() bool {
	if o != nil && o.RateLimit.IsSet() {
		return true
	}

	return false
}

// SetRateLimit gets a reference to the given NullableInt32 and assigns it to the RateLimit field.
func (o *AnsibleRoleRemoteResponse) SetRateLimit(v int32) {
	o.RateLimit.Set(&v)
}
// SetRateLimitNil sets the value for RateLimit to be an explicit nil
func (o *AnsibleRoleRemoteResponse) SetRateLimitNil() {
	o.RateLimit.Set(nil)
}

// UnsetRateLimit ensures that no value is present for RateLimit, not even an explicit nil
func (o *AnsibleRoleRemoteResponse) UnsetRateLimit() {
	o.RateLimit.Unset()
}

// GetHiddenFields returns the HiddenFields field value if set, zero value otherwise.
func (o *AnsibleRoleRemoteResponse) GetHiddenFields() []RemoteResponseHiddenFieldsInner {
	if o == nil || IsNil(o.HiddenFields) {
		var ret []RemoteResponseHiddenFieldsInner
		return ret
	}
	return o.HiddenFields
}

// GetHiddenFieldsOk returns a tuple with the HiddenFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnsibleRoleRemoteResponse) GetHiddenFieldsOk() ([]RemoteResponseHiddenFieldsInner, bool) {
	if o == nil || IsNil(o.HiddenFields) {
		return nil, false
	}
	return o.HiddenFields, true
}

// HasHiddenFields returns a boolean if a field has been set.
func (o *AnsibleRoleRemoteResponse) HasHiddenFields() bool {
	if o != nil && !IsNil(o.HiddenFields) {
		return true
	}

	return false
}

// SetHiddenFields gets a reference to the given []RemoteResponseHiddenFieldsInner and assigns it to the HiddenFields field.
func (o *AnsibleRoleRemoteResponse) SetHiddenFields(v []RemoteResponseHiddenFieldsInner) {
	o.HiddenFields = v
}

func (o AnsibleRoleRemoteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnsibleRoleRemoteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: pulp_href is readOnly
	// skip: pulp_created is readOnly
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	if o.CaCert.IsSet() {
		toSerialize["ca_cert"] = o.CaCert.Get()
	}
	if o.ClientCert.IsSet() {
		toSerialize["client_cert"] = o.ClientCert.Get()
	}
	if !IsNil(o.TlsValidation) {
		toSerialize["tls_validation"] = o.TlsValidation
	}
	if o.ProxyUrl.IsSet() {
		toSerialize["proxy_url"] = o.ProxyUrl.Get()
	}
	if !IsNil(o.PulpLabels) {
		toSerialize["pulp_labels"] = o.PulpLabels
	}
	// skip: pulp_last_updated is readOnly
	if o.DownloadConcurrency.IsSet() {
		toSerialize["download_concurrency"] = o.DownloadConcurrency.Get()
	}
	if o.MaxRetries.IsSet() {
		toSerialize["max_retries"] = o.MaxRetries.Get()
	}
	if !IsNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if o.TotalTimeout.IsSet() {
		toSerialize["total_timeout"] = o.TotalTimeout.Get()
	}
	if o.ConnectTimeout.IsSet() {
		toSerialize["connect_timeout"] = o.ConnectTimeout.Get()
	}
	if o.SockConnectTimeout.IsSet() {
		toSerialize["sock_connect_timeout"] = o.SockConnectTimeout.Get()
	}
	if o.SockReadTimeout.IsSet() {
		toSerialize["sock_read_timeout"] = o.SockReadTimeout.Get()
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if o.RateLimit.IsSet() {
		toSerialize["rate_limit"] = o.RateLimit.Get()
	}
	// skip: hidden_fields is readOnly
	return toSerialize, nil
}

type NullableAnsibleRoleRemoteResponse struct {
	value *AnsibleRoleRemoteResponse
	isSet bool
}

func (v NullableAnsibleRoleRemoteResponse) Get() *AnsibleRoleRemoteResponse {
	return v.value
}

func (v *NullableAnsibleRoleRemoteResponse) Set(val *AnsibleRoleRemoteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnsibleRoleRemoteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnsibleRoleRemoteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnsibleRoleRemoteResponse(val *AnsibleRoleRemoteResponse) *NullableAnsibleRoleRemoteResponse {
	return &NullableAnsibleRoleRemoteResponse{value: val, isSet: true}
}

func (v NullableAnsibleRoleRemoteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnsibleRoleRemoteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


