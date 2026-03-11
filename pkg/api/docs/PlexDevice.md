# PlexDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Product** | **string** |  | 
**ProductVersion** | **string** |  | 
**Platform** | **string** |  | 
**PlatformVersion** | Pointer to **string** |  | [optional] 
**Device** | **string** |  | 
**ClientIdentifier** | **string** |  | 
**CreatedAt** | **string** |  | 
**LastSeenAt** | **string** |  | 
**Provides** | **[]string** |  | 
**Owned** | **bool** |  | 
**OwnerID** | Pointer to **string** |  | [optional] 
**Home** | Pointer to **bool** |  | [optional] 
**SourceTitle** | Pointer to **string** |  | [optional] 
**AccessToken** | Pointer to **string** |  | [optional] 
**PublicAddress** | Pointer to **string** |  | [optional] 
**HttpsRequired** | Pointer to **bool** |  | [optional] 
**Synced** | Pointer to **bool** |  | [optional] 
**Relay** | Pointer to **bool** |  | [optional] 
**DnsRebindingProtection** | Pointer to **bool** |  | [optional] 
**NatLoopbackSupported** | Pointer to **bool** |  | [optional] 
**PublicAddressMatches** | Pointer to **bool** |  | [optional] 
**Presence** | Pointer to **bool** |  | [optional] 
**Connection** | [**[]PlexConnection**](PlexConnection.md) |  | 

## Methods

### NewPlexDevice

`func NewPlexDevice(name string, product string, productVersion string, platform string, device string, clientIdentifier string, createdAt string, lastSeenAt string, provides []string, owned bool, connection []PlexConnection, ) *PlexDevice`

NewPlexDevice instantiates a new PlexDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlexDeviceWithDefaults

`func NewPlexDeviceWithDefaults() *PlexDevice`

NewPlexDeviceWithDefaults instantiates a new PlexDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlexDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlexDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlexDevice) SetName(v string)`

SetName sets Name field to given value.


### GetProduct

`func (o *PlexDevice) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PlexDevice) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PlexDevice) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetProductVersion

`func (o *PlexDevice) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *PlexDevice) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *PlexDevice) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.


### GetPlatform

`func (o *PlexDevice) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PlexDevice) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PlexDevice) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPlatformVersion

`func (o *PlexDevice) GetPlatformVersion() string`

GetPlatformVersion returns the PlatformVersion field if non-nil, zero value otherwise.

### GetPlatformVersionOk

`func (o *PlexDevice) GetPlatformVersionOk() (*string, bool)`

GetPlatformVersionOk returns a tuple with the PlatformVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformVersion

`func (o *PlexDevice) SetPlatformVersion(v string)`

SetPlatformVersion sets PlatformVersion field to given value.

### HasPlatformVersion

`func (o *PlexDevice) HasPlatformVersion() bool`

HasPlatformVersion returns a boolean if a field has been set.

### GetDevice

`func (o *PlexDevice) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PlexDevice) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PlexDevice) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetClientIdentifier

`func (o *PlexDevice) GetClientIdentifier() string`

GetClientIdentifier returns the ClientIdentifier field if non-nil, zero value otherwise.

### GetClientIdentifierOk

`func (o *PlexDevice) GetClientIdentifierOk() (*string, bool)`

GetClientIdentifierOk returns a tuple with the ClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifier

`func (o *PlexDevice) SetClientIdentifier(v string)`

SetClientIdentifier sets ClientIdentifier field to given value.


### GetCreatedAt

`func (o *PlexDevice) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PlexDevice) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PlexDevice) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastSeenAt

`func (o *PlexDevice) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *PlexDevice) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *PlexDevice) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.


### GetProvides

`func (o *PlexDevice) GetProvides() []string`

GetProvides returns the Provides field if non-nil, zero value otherwise.

### GetProvidesOk

`func (o *PlexDevice) GetProvidesOk() (*[]string, bool)`

GetProvidesOk returns a tuple with the Provides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvides

`func (o *PlexDevice) SetProvides(v []string)`

SetProvides sets Provides field to given value.


### GetOwned

`func (o *PlexDevice) GetOwned() bool`

GetOwned returns the Owned field if non-nil, zero value otherwise.

### GetOwnedOk

`func (o *PlexDevice) GetOwnedOk() (*bool, bool)`

GetOwnedOk returns a tuple with the Owned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwned

`func (o *PlexDevice) SetOwned(v bool)`

SetOwned sets Owned field to given value.


### GetOwnerID

`func (o *PlexDevice) GetOwnerID() string`

GetOwnerID returns the OwnerID field if non-nil, zero value otherwise.

### GetOwnerIDOk

`func (o *PlexDevice) GetOwnerIDOk() (*string, bool)`

GetOwnerIDOk returns a tuple with the OwnerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerID

`func (o *PlexDevice) SetOwnerID(v string)`

SetOwnerID sets OwnerID field to given value.

### HasOwnerID

`func (o *PlexDevice) HasOwnerID() bool`

HasOwnerID returns a boolean if a field has been set.

### GetHome

`func (o *PlexDevice) GetHome() bool`

GetHome returns the Home field if non-nil, zero value otherwise.

### GetHomeOk

`func (o *PlexDevice) GetHomeOk() (*bool, bool)`

GetHomeOk returns a tuple with the Home field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHome

`func (o *PlexDevice) SetHome(v bool)`

SetHome sets Home field to given value.

### HasHome

`func (o *PlexDevice) HasHome() bool`

HasHome returns a boolean if a field has been set.

### GetSourceTitle

`func (o *PlexDevice) GetSourceTitle() string`

GetSourceTitle returns the SourceTitle field if non-nil, zero value otherwise.

### GetSourceTitleOk

`func (o *PlexDevice) GetSourceTitleOk() (*string, bool)`

GetSourceTitleOk returns a tuple with the SourceTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTitle

`func (o *PlexDevice) SetSourceTitle(v string)`

SetSourceTitle sets SourceTitle field to given value.

### HasSourceTitle

`func (o *PlexDevice) HasSourceTitle() bool`

HasSourceTitle returns a boolean if a field has been set.

### GetAccessToken

`func (o *PlexDevice) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *PlexDevice) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *PlexDevice) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *PlexDevice) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetPublicAddress

`func (o *PlexDevice) GetPublicAddress() string`

GetPublicAddress returns the PublicAddress field if non-nil, zero value otherwise.

### GetPublicAddressOk

`func (o *PlexDevice) GetPublicAddressOk() (*string, bool)`

GetPublicAddressOk returns a tuple with the PublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAddress

`func (o *PlexDevice) SetPublicAddress(v string)`

SetPublicAddress sets PublicAddress field to given value.

### HasPublicAddress

`func (o *PlexDevice) HasPublicAddress() bool`

HasPublicAddress returns a boolean if a field has been set.

### GetHttpsRequired

`func (o *PlexDevice) GetHttpsRequired() bool`

GetHttpsRequired returns the HttpsRequired field if non-nil, zero value otherwise.

### GetHttpsRequiredOk

`func (o *PlexDevice) GetHttpsRequiredOk() (*bool, bool)`

GetHttpsRequiredOk returns a tuple with the HttpsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsRequired

`func (o *PlexDevice) SetHttpsRequired(v bool)`

SetHttpsRequired sets HttpsRequired field to given value.

### HasHttpsRequired

`func (o *PlexDevice) HasHttpsRequired() bool`

HasHttpsRequired returns a boolean if a field has been set.

### GetSynced

`func (o *PlexDevice) GetSynced() bool`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *PlexDevice) GetSyncedOk() (*bool, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *PlexDevice) SetSynced(v bool)`

SetSynced sets Synced field to given value.

### HasSynced

`func (o *PlexDevice) HasSynced() bool`

HasSynced returns a boolean if a field has been set.

### GetRelay

`func (o *PlexDevice) GetRelay() bool`

GetRelay returns the Relay field if non-nil, zero value otherwise.

### GetRelayOk

`func (o *PlexDevice) GetRelayOk() (*bool, bool)`

GetRelayOk returns a tuple with the Relay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelay

`func (o *PlexDevice) SetRelay(v bool)`

SetRelay sets Relay field to given value.

### HasRelay

`func (o *PlexDevice) HasRelay() bool`

HasRelay returns a boolean if a field has been set.

### GetDnsRebindingProtection

`func (o *PlexDevice) GetDnsRebindingProtection() bool`

GetDnsRebindingProtection returns the DnsRebindingProtection field if non-nil, zero value otherwise.

### GetDnsRebindingProtectionOk

`func (o *PlexDevice) GetDnsRebindingProtectionOk() (*bool, bool)`

GetDnsRebindingProtectionOk returns a tuple with the DnsRebindingProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRebindingProtection

`func (o *PlexDevice) SetDnsRebindingProtection(v bool)`

SetDnsRebindingProtection sets DnsRebindingProtection field to given value.

### HasDnsRebindingProtection

`func (o *PlexDevice) HasDnsRebindingProtection() bool`

HasDnsRebindingProtection returns a boolean if a field has been set.

### GetNatLoopbackSupported

`func (o *PlexDevice) GetNatLoopbackSupported() bool`

GetNatLoopbackSupported returns the NatLoopbackSupported field if non-nil, zero value otherwise.

### GetNatLoopbackSupportedOk

`func (o *PlexDevice) GetNatLoopbackSupportedOk() (*bool, bool)`

GetNatLoopbackSupportedOk returns a tuple with the NatLoopbackSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatLoopbackSupported

`func (o *PlexDevice) SetNatLoopbackSupported(v bool)`

SetNatLoopbackSupported sets NatLoopbackSupported field to given value.

### HasNatLoopbackSupported

`func (o *PlexDevice) HasNatLoopbackSupported() bool`

HasNatLoopbackSupported returns a boolean if a field has been set.

### GetPublicAddressMatches

`func (o *PlexDevice) GetPublicAddressMatches() bool`

GetPublicAddressMatches returns the PublicAddressMatches field if non-nil, zero value otherwise.

### GetPublicAddressMatchesOk

`func (o *PlexDevice) GetPublicAddressMatchesOk() (*bool, bool)`

GetPublicAddressMatchesOk returns a tuple with the PublicAddressMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAddressMatches

`func (o *PlexDevice) SetPublicAddressMatches(v bool)`

SetPublicAddressMatches sets PublicAddressMatches field to given value.

### HasPublicAddressMatches

`func (o *PlexDevice) HasPublicAddressMatches() bool`

HasPublicAddressMatches returns a boolean if a field has been set.

### GetPresence

`func (o *PlexDevice) GetPresence() bool`

GetPresence returns the Presence field if non-nil, zero value otherwise.

### GetPresenceOk

`func (o *PlexDevice) GetPresenceOk() (*bool, bool)`

GetPresenceOk returns a tuple with the Presence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresence

`func (o *PlexDevice) SetPresence(v bool)`

SetPresence sets Presence field to given value.

### HasPresence

`func (o *PlexDevice) HasPresence() bool`

HasPresence returns a boolean if a field has been set.

### GetConnection

`func (o *PlexDevice) GetConnection() []PlexConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *PlexDevice) GetConnectionOk() (*[]PlexConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *PlexDevice) SetConnection(v []PlexConnection)`

SetConnection sets Connection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


