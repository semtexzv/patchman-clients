# HostOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | A Red Hat Account number that owns the host. | 
**AnsibleHost** | Pointer to **string** | The ansible host name for remediations | [optional] 
**BiosUuid** | Pointer to **string** | A UUID of the host machine BIOS.  This field is considered to be a canonical fact. | [optional] 
**Created** | **string** | A timestamp when the entry was created. | [optional] 
**CulledTimestamp** | Pointer to **string** | Timestamp from which the host is considered deleted. | [optional] 
**DisplayName** | Pointer to **string** | A host’s human-readable display name, e.g. in a form of a domain name. | [optional] 
**ExternalId** | Pointer to **string** | Host’s reference in the external source e.g. AWS EC2, Azure, OpenStack, etc. This field is considered to be a canonical fact. | [optional] 
**Facts** | [**[]FactSet**](FactSet.md) | A set of facts belonging to the host. | [optional] 
**Fqdn** | Pointer to **string** | A host’s Fully Qualified Domain Name.  This field is considered to be a canonical fact. | [optional] 
**Id** | **string** | A durable and reliable platform-wide host identifier. Applications should use this identifier to reference hosts. | [optional] 
**InsightsId** | Pointer to **string** | An ID defined in /etc/insights-client/machine-id. This field is considered a canonical fact. | [optional] 
**IpAddresses** | Pointer to **[]string** | Host’s network IP addresses.  This field is considered to be a canonical fact. | [optional] 
**MacAddresses** | Pointer to **[]string** | Host’s network interfaces MAC addresses.  This field is considered to be a canonical fact. | [optional] 
**Reporter** | Pointer to **string** | Reporting source of the host. Used when updating the stale_timestamp. | [optional] 
**RhelMachineId** | Pointer to **string** | A Machine ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 
**SatelliteId** | Pointer to **string** | A Red Hat Satellite ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 
**StaleTimestamp** | Pointer to **string** | Timestamp from which the host is considered stale. | [optional] 
**StaleWarningTimestamp** | Pointer to **string** | Timestamp from which the host is considered too stale to be listed without an explicit toggle. | [optional] 
**SubscriptionManagerId** | Pointer to **string** | A Red Hat Subcription Manager ID of a RHEL host.  This field is considered to be a canonical fact. | [optional] 
**Tags** | [**[]StructuredTag**](StructuredTag.md) | An array of the tags on the host | [optional] 
**Updated** | **string** | A timestamp when the entry was last updated. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


