# CreateHostIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | A Red Hat Account number that owns the host. | 
**AnsibleHost** | **string** | The ansible host name for remediations | [optional] 
**BiosUuid** | **string** | A UUID of the host machine BIOS.  This field is considered to be a canonical fact. | [optional] 
**DisplayName** | **string** | A host’s human-readable display name, e.g. in a form of a domain name. | [optional] 
**ExternalId** | **string** | Host’s reference in the external source e.g. AWS EC2, Azure, OpenStack, etc. This field is considered to be a canonical fact. | [optional] 
**Facts** | [**[]FactSet**](FactSet.md) | A set of facts belonging to the host. | [optional] 
**Fqdn** | **string** | A host’s Fully Qualified Domain Name.  This field is considered to be a canonical fact. | [optional] 
**InsightsId** | **string** | An ID defined in /etc/insights-client/machine-id. This field is considered a canonical fact. | [optional] 
**IpAddresses** | **[]string** | Host’s network IP addresses.  This field is considered to be a canonical fact. | [optional] 
**MacAddresses** | **[]string** | Host’s network interfaces MAC addresses.  This field is considered to be a canonical fact. | [optional] 
**Reporter** | **string** | Reporting source of the host. Used when updating the stale_timestamp. | [optional] 
**RhelMachineId** | **string** | A Machine ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 
**SatelliteId** | **string** | A Red Hat Satellite ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 
**StaleTimestamp** | **string** | Timestamp from which the host is considered stale. | [optional] 
**SubscriptionManagerId** | **string** | A Red Hat Subcription Manager ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 
**SystemProfile** | [**SystemProfileIn**](SystemProfileIn.md) |  | [optional] 
**Tags** | [**[]StructuredTag**](StructuredTag.md) | The tags on a host | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


