# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [backup.proto](#backup-proto)
    - [BackupCapabilitiesRequest](#cnpgi-backup-v1-BackupCapabilitiesRequest)
    - [BackupCapabilitiesResult](#cnpgi-backup-v1-BackupCapabilitiesResult)
    - [BackupCapability](#cnpgi-backup-v1-BackupCapability)
    - [BackupCapability.RPC](#cnpgi-backup-v1-BackupCapability-RPC)
    - [BackupRequest](#cnpgi-backup-v1-BackupRequest)
    - [BackupRequest.ParametersEntry](#cnpgi-backup-v1-BackupRequest-ParametersEntry)
    - [BackupResult](#cnpgi-backup-v1-BackupResult)
    - [BackupResult.MetadataEntry](#cnpgi-backup-v1-BackupResult-MetadataEntry)
  
    - [BackupCapability.RPC.Type](#cnpgi-backup-v1-BackupCapability-RPC-Type)
  
    - [Backup](#cnpgi-backup-v1-Backup)
  
- [identity.proto](#identity-proto)
    - [GetPluginCapabilitiesRequest](#cnpgi-identity-v1-GetPluginCapabilitiesRequest)
    - [GetPluginCapabilitiesResponse](#cnpgi-identity-v1-GetPluginCapabilitiesResponse)
    - [GetPluginMetadataRequest](#cnpgi-identity-v1-GetPluginMetadataRequest)
    - [GetPluginMetadataResponse](#cnpgi-identity-v1-GetPluginMetadataResponse)
    - [GetPluginMetadataResponse.ManifestEntry](#cnpgi-identity-v1-GetPluginMetadataResponse-ManifestEntry)
    - [PluginCapability](#cnpgi-identity-v1-PluginCapability)
    - [PluginCapability.Service](#cnpgi-identity-v1-PluginCapability-Service)
    - [ProbeRequest](#cnpgi-identity-v1-ProbeRequest)
    - [ProbeResponse](#cnpgi-identity-v1-ProbeResponse)
  
    - [PluginCapability.Service.Type](#cnpgi-identity-v1-PluginCapability-Service-Type)
  
    - [Identity](#cnpgi-identity-v1-Identity)
  
- [operator.proto](#operator-proto)
    - [OperatorCapabilitiesRequest](#cnpgi-operator-v1-OperatorCapabilitiesRequest)
    - [OperatorCapabilitiesResult](#cnpgi-operator-v1-OperatorCapabilitiesResult)
    - [OperatorCapability](#cnpgi-operator-v1-OperatorCapability)
    - [OperatorCapability.RPC](#cnpgi-operator-v1-OperatorCapability-RPC)
    - [OperatorMutateClusterRequest](#cnpgi-operator-v1-OperatorMutateClusterRequest)
    - [OperatorMutateClusterResult](#cnpgi-operator-v1-OperatorMutateClusterResult)
    - [OperatorValidateClusterChangeRequest](#cnpgi-operator-v1-OperatorValidateClusterChangeRequest)
    - [OperatorValidateClusterChangeResult](#cnpgi-operator-v1-OperatorValidateClusterChangeResult)
    - [OperatorValidateClusterCreateRequest](#cnpgi-operator-v1-OperatorValidateClusterCreateRequest)
    - [OperatorValidateClusterCreateResult](#cnpgi-operator-v1-OperatorValidateClusterCreateResult)
    - [ValidationError](#cnpgi-operator-v1-ValidationError)
  
    - [OperatorCapability.RPC.Type](#cnpgi-operator-v1-OperatorCapability-RPC-Type)
  
    - [Operator](#cnpgi-operator-v1-Operator)
  
- [operator_lifecycle.proto](#operator_lifecycle-proto)
    - [OperatorLifecycleCapabilities](#cnpgi-operator_lifecycle-v1-OperatorLifecycleCapabilities)
    - [OperatorLifecycleCapabilitiesRequest](#cnpgi-operator_lifecycle-v1-OperatorLifecycleCapabilitiesRequest)
    - [OperatorLifecycleCapabilitiesResponse](#cnpgi-operator_lifecycle-v1-OperatorLifecycleCapabilitiesResponse)
    - [OperatorLifecycleRequest](#cnpgi-operator_lifecycle-v1-OperatorLifecycleRequest)
    - [OperatorLifecycleResponse](#cnpgi-operator_lifecycle-v1-OperatorLifecycleResponse)
    - [OperatorOperationType](#cnpgi-operator_lifecycle-v1-OperatorOperationType)
  
    - [OperatorOperationType.Type](#cnpgi-operator_lifecycle-v1-OperatorOperationType-Type)
  
    - [OperatorLifecycle](#cnpgi-operator_lifecycle-v1-OperatorLifecycle)
  
- [reconciler.proto](#reconciler-proto)
    - [ReconcilerHooksCapabilitiesRequest](#cnpgi-reconciler-v1-ReconcilerHooksCapabilitiesRequest)
    - [ReconcilerHooksCapabilitiesResult](#cnpgi-reconciler-v1-ReconcilerHooksCapabilitiesResult)
    - [ReconcilerHooksCapability](#cnpgi-reconciler-v1-ReconcilerHooksCapability)
    - [ReconcilerHooksRequest](#cnpgi-reconciler-v1-ReconcilerHooksRequest)
    - [ReconcilerHooksResult](#cnpgi-reconciler-v1-ReconcilerHooksResult)
  
    - [ReconcilerHooksCapability.Kind](#cnpgi-reconciler-v1-ReconcilerHooksCapability-Kind)
    - [ReconcilerHooksResult.Behavior](#cnpgi-reconciler-v1-ReconcilerHooksResult-Behavior)
  
    - [ReconcilerHooks](#cnpgi-reconciler-v1-ReconcilerHooks)
  
- [wal.proto](#wal-proto)
    - [SetFirstRequiredRequest](#cnpgi-wal-v1-SetFirstRequiredRequest)
    - [SetFirstRequiredResult](#cnpgi-wal-v1-SetFirstRequiredResult)
    - [WALArchiveRequest](#cnpgi-wal-v1-WALArchiveRequest)
    - [WALArchiveRequest.ParametersEntry](#cnpgi-wal-v1-WALArchiveRequest-ParametersEntry)
    - [WALArchiveResult](#cnpgi-wal-v1-WALArchiveResult)
    - [WALCapabilitiesRequest](#cnpgi-wal-v1-WALCapabilitiesRequest)
    - [WALCapabilitiesResult](#cnpgi-wal-v1-WALCapabilitiesResult)
    - [WALCapability](#cnpgi-wal-v1-WALCapability)
    - [WALCapability.RPC](#cnpgi-wal-v1-WALCapability-RPC)
    - [WALRestoreRequest](#cnpgi-wal-v1-WALRestoreRequest)
    - [WALRestoreRequest.ParametersEntry](#cnpgi-wal-v1-WALRestoreRequest-ParametersEntry)
    - [WALRestoreResult](#cnpgi-wal-v1-WALRestoreResult)
    - [WALStatusRequest](#cnpgi-wal-v1-WALStatusRequest)
    - [WALStatusResult](#cnpgi-wal-v1-WALStatusResult)
    - [WALStatusResult.AdditionalInformationEntry](#cnpgi-wal-v1-WALStatusResult-AdditionalInformationEntry)
  
    - [WALCapability.RPC.Type](#cnpgi-wal-v1-WALCapability-RPC-Type)
  
    - [WAL](#cnpgi-wal-v1-WAL)
  
- [Scalar Value Types](#scalar-value-types)



<a name="backup-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## backup.proto



<a name="cnpgi-backup-v1-BackupCapabilitiesRequest"></a>

### BackupCapabilitiesRequest
Intentionally empty.






<a name="cnpgi-backup-v1-BackupCapabilitiesResult"></a>

### BackupCapabilitiesResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capabilities | [BackupCapability](#cnpgi-backup-v1-BackupCapability) | repeated | All the capabilities that the controller service supports. This field is OPTIONAL. |






<a name="cnpgi-backup-v1-BackupCapability"></a>

### BackupCapability



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rpc | [BackupCapability.RPC](#cnpgi-backup-v1-BackupCapability-RPC) |  |  |






<a name="cnpgi-backup-v1-BackupCapability-RPC"></a>

### BackupCapability.RPC



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [BackupCapability.RPC.Type](#cnpgi-backup-v1-BackupCapability-RPC-Type) |  |  |






<a name="cnpgi-backup-v1-BackupRequest"></a>

### BackupRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_definition | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the Cluster being backed up |
| backup_definition | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the Backup that is being taken |
| parameters | [BackupRequest.ParametersEntry](#cnpgi-backup-v1-BackupRequest-ParametersEntry) | repeated | This field is OPTIONAL. Value of this field is the configuration of this backup as set in the Backup or in the ScheduledBackup object |






<a name="cnpgi-backup-v1-BackupRequest-ParametersEntry"></a>

### BackupRequest.ParametersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="cnpgi-backup-v1-BackupResult"></a>

### BackupResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| backup_id | [string](#string) |  | This field is REQUIRED and contains a machine-readable ID of the backup that is being taken |
| backup_name | [string](#string) |  | This field is OPTIONAL and contains a human-readable name of the backup that is being taken |
| started_at | [int64](#int64) |  | This field is REQUIRED and contains the Unix timestamp of the start time of the backup |
| stopped_at | [int64](#int64) |  | This field is REQUIRED and contains the Unix timestamp of the end time of the backup |
| begin_wal | [string](#string) |  | This field is OPTIONAL and contains the current WAL when the backup was started |
| end_wal | [string](#string) |  | This field is OPTIONAL and contains the current WAL at the end of the backup |
| begin_lsn | [string](#string) |  | This field is OPTIONAL and contains the current LSN record when the backup was started |
| end_lsn | [string](#string) |  | This field is OPTIONAL and contains the current LSN record when the backup has finished |
| backup_label_file | [bytes](#bytes) |  | This field is OPTIONAL and contains the backup label of the backup that have been taken |
| tablespace_map_file | [bytes](#bytes) |  | This field is OPTIONAL and contains the tablespace map of the backup that have been taken |
| instance_id | [string](#string) |  | This field is OPTIONAL and contains the ID of the instance that have been backed up |
| online | [bool](#bool) |  | This field is REQUIRED and is set to true for online/hot backups and to false otherwise. |
| metadata | [BackupResult.MetadataEntry](#cnpgi-backup-v1-BackupResult-MetadataEntry) | repeated | This field is OPTIONAL and contains all the plugin specific information that needs to be stored |






<a name="cnpgi-backup-v1-BackupResult-MetadataEntry"></a>

### BackupResult.MetadataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 


<a name="cnpgi-backup-v1-BackupCapability-RPC-Type"></a>

### BackupCapability.RPC.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| TYPE_BACKUP | 1 | TYPE_BACKUP indicates that the Plugin is able to take physical backups. This feature is required for every plugin exposing the Backup service |


 

 


<a name="cnpgi-backup-v1-Backup"></a>

### Backup


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCapabilities | [BackupCapabilitiesRequest](#cnpgi-backup-v1-BackupCapabilitiesRequest) | [BackupCapabilitiesResult](#cnpgi-backup-v1-BackupCapabilitiesResult) | GetCapabilities gets the capabilities of the Backup service |
| Backup | [BackupRequest](#cnpgi-backup-v1-BackupRequest) | [BackupResult](#cnpgi-backup-v1-BackupResult) | Backup takes a physical backup of PostgreSQL. |

 



<a name="identity-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## identity.proto



<a name="cnpgi-identity-v1-GetPluginCapabilitiesRequest"></a>

### GetPluginCapabilitiesRequest
Intentionally empty.






<a name="cnpgi-identity-v1-GetPluginCapabilitiesResponse"></a>

### GetPluginCapabilitiesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capabilities | [PluginCapability](#cnpgi-identity-v1-PluginCapability) | repeated | All the capabilities that the controller service supports. This field is OPTIONAL. |






<a name="cnpgi-identity-v1-GetPluginMetadataRequest"></a>

### GetPluginMetadataRequest
Intentionally empty.






<a name="cnpgi-identity-v1-GetPluginMetadataResponse"></a>

### GetPluginMetadataResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | The name MUST follow domain name notation format (https://tools.ietf.org/html/rfc1035#section-2.3.1). It SHOULD include the plugin&#39;s host company name and the plugin name, to minimize the possibility of collisions. It MUST be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), dots (.), and alphanumerics between. This field is REQUIRED. |
| version | [string](#string) |  | This field is REQUIRED. Value of this field is opaque. |
| display_name | [string](#string) |  | A name to display for the plugin. This field is REQUIRED. |
| description | [string](#string) |  | A description for the plugin. This field is REQUIRED. |
| project_url | [string](#string) |  | URL of the home page of the plugin project. |
| repository_url | [string](#string) |  | URL of the source code repository for the plugin project. |
| license | [string](#string) |  | License of the plugin. This field is REQUIRED. |
| license_url | [string](#string) |  | URL of the license of the plugin. This field is REQUIRED. |
| maturity | [string](#string) |  | Maturity level (alpha, beta, stable) |
| vendor | [string](#string) |  | Provider/vendor of the plugin, e.g. an organization |
| manifest | [GetPluginMetadataResponse.ManifestEntry](#cnpgi-identity-v1-GetPluginMetadataResponse-ManifestEntry) | repeated | This field is OPTIONAL. Values are opaque. |






<a name="cnpgi-identity-v1-GetPluginMetadataResponse-ManifestEntry"></a>

### GetPluginMetadataResponse.ManifestEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="cnpgi-identity-v1-PluginCapability"></a>

### PluginCapability



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service | [PluginCapability.Service](#cnpgi-identity-v1-PluginCapability-Service) |  |  |






<a name="cnpgi-identity-v1-PluginCapability-Service"></a>

### PluginCapability.Service



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [PluginCapability.Service.Type](#cnpgi-identity-v1-PluginCapability-Service-Type) |  |  |






<a name="cnpgi-identity-v1-ProbeRequest"></a>

### ProbeRequest
Intentionally empty.






<a name="cnpgi-identity-v1-ProbeResponse"></a>

### ProbeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ready | [bool](#bool) |  | This field is OPTIONAL. If not present, the caller SHALL assume that the plugin is in a ready state and is accepting calls to its Controller and/or Node services (according to the plugin&#39;s reported capabilities). |





 


<a name="cnpgi-identity-v1-PluginCapability-Service-Type"></a>

### PluginCapability.Service.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| TYPE_OPERATOR_SERVICE | 1 | TYPE_OPERATOR_SERVICE indicated that the Plugin provider RPCs for the Operator service. The presence of this capability determines whether the CO will attempt to invoke the REQUIRED Operator RPCs, as well as specific RPCs as indicated by GetCapabilities. |
| TYPE_WAL_SERVICE | 2 | TYPE_WAL_SERVICE indicates that the Plugin provides RPCs for the WAL service. Plugins MAY provide this capability. The presence of this capability determines whether the CO will attempt to invoke the REQUIRED WALService RPCs, as well as specific RPCs as indicated by GetCapabilities. |
| TYPE_BACKUP_SERVICE | 3 | TYPE_BACKUP_SERVICE indicates that the Plugin provides RPCs for the Backup service. The presence of this capability determines whether the CO will attempt to invoke the REQUIRED Backup Service RPCs, as well as specific RPCs as indicated by GetCapabilities. |
| TYPE_LIFECYCLE_SERVICE | 4 | TYPE_LIFECYCLE_SERVICE indicates that the Plugin provides RPCs for the Lifecycle service. |
| TYPE_RECONCILER_HOOKS | 5 | TYPE_RECONCILER_HOOKS indicates that the Plugin provides RPCs to enhance the behavior of the reconcilers |


 

 


<a name="cnpgi-identity-v1-Identity"></a>

### Identity


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPluginMetadata | [GetPluginMetadataRequest](#cnpgi-identity-v1-GetPluginMetadataRequest) | [GetPluginMetadataResponse](#cnpgi-identity-v1-GetPluginMetadataResponse) | GetPluginMetadata gets the plugin metadata |
| GetPluginCapabilities | [GetPluginCapabilitiesRequest](#cnpgi-identity-v1-GetPluginCapabilitiesRequest) | [GetPluginCapabilitiesResponse](#cnpgi-identity-v1-GetPluginCapabilitiesResponse) | GetPluginCapabilities gets information about this plugin |
| Probe | [ProbeRequest](#cnpgi-identity-v1-ProbeRequest) | [ProbeResponse](#cnpgi-identity-v1-ProbeResponse) | Probe is used to tell if the plugin is ready to receive requests |

 



<a name="operator-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## operator.proto



<a name="cnpgi-operator-v1-OperatorCapabilitiesRequest"></a>

### OperatorCapabilitiesRequest
Intentionally empty.






<a name="cnpgi-operator-v1-OperatorCapabilitiesResult"></a>

### OperatorCapabilitiesResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capabilities | [OperatorCapability](#cnpgi-operator-v1-OperatorCapability) | repeated | All the capabilities that the controller service supports. This field is OPTIONAL. |






<a name="cnpgi-operator-v1-OperatorCapability"></a>

### OperatorCapability



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rpc | [OperatorCapability.RPC](#cnpgi-operator-v1-OperatorCapability-RPC) |  |  |






<a name="cnpgi-operator-v1-OperatorCapability-RPC"></a>

### OperatorCapability.RPC



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [OperatorCapability.RPC.Type](#cnpgi-operator-v1-OperatorCapability-RPC-Type) |  |  |






<a name="cnpgi-operator-v1-OperatorMutateClusterRequest"></a>

### OperatorMutateClusterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| definition | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the Cluster that should receive the default values |






<a name="cnpgi-operator-v1-OperatorMutateClusterResult"></a>

### OperatorMutateClusterResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| json_patch | [bytes](#bytes) |  | This field is OPTIONAL. Value of this field is a JSONPatch to be applied on the passed Cluster definition |






<a name="cnpgi-operator-v1-OperatorValidateClusterChangeRequest"></a>

### OperatorValidateClusterChangeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| old_cluster | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the current Cluster definition |
| new_cluster | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the updated Cluster definition |






<a name="cnpgi-operator-v1-OperatorValidateClusterChangeResult"></a>

### OperatorValidateClusterChangeResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| validation_errors | [ValidationError](#cnpgi-operator-v1-ValidationError) | repeated | This field is OPTIONAL. Value of this field is a set of validation errors |






<a name="cnpgi-operator-v1-OperatorValidateClusterCreateRequest"></a>

### OperatorValidateClusterCreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| definition | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the Cluster that is being created |






<a name="cnpgi-operator-v1-OperatorValidateClusterCreateResult"></a>

### OperatorValidateClusterCreateResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| validation_errors | [ValidationError](#cnpgi-operator-v1-ValidationError) | repeated | This field is OPTIONAL. Value of this field is a set of validation errors |






<a name="cnpgi-operator-v1-ValidationError"></a>

### ValidationError



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| path_components | [string](#string) | repeated | This field is REQUIRED. Value of this field is |
| value | [string](#string) |  | This field is REQUIRED. Value of this field is the value that caused a validation error |
| message | [string](#string) |  | This field is REQUIRED. Value of this field is a description of the validation error |





 


<a name="cnpgi-operator-v1-OperatorCapability-RPC-Type"></a>

### OperatorCapability.RPC.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| TYPE_VALIDATE_CLUSTER_CREATE | 1 | TYPE_VALIDATE_CLUSTER_CREATE indicates that the Plugin is able to reply to the ValidateClusterCreate RPC request |
| TYPE_VALIDATE_CLUSTER_CHANGE | 2 | TYPE_VALIDATE_CLUSTER_CHANGE indicates that the Plugin is able to reply to the ValidateClusterChange RPC request |
| TYPE_MUTATE_CLUSTER | 3 | TYPE_MUTATE_CLUSTER indicates that the Plugin is able to reply to the MutateCluster RPC request |
| TYPE_MUTATE_POD | 4 | TYPE_MUTATE_POD indicates that the Plugin is able to reply to the MutatePod RPC request |


 

 


<a name="cnpgi-operator-v1-Operator"></a>

### Operator


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCapabilities | [OperatorCapabilitiesRequest](#cnpgi-operator-v1-OperatorCapabilitiesRequest) | [OperatorCapabilitiesResult](#cnpgi-operator-v1-OperatorCapabilitiesResult) | GetCapabilities gets the capabilities of the WAL service |
| ValidateClusterCreate | [OperatorValidateClusterCreateRequest](#cnpgi-operator-v1-OperatorValidateClusterCreateRequest) | [OperatorValidateClusterCreateResult](#cnpgi-operator-v1-OperatorValidateClusterCreateResult) | ValidateCreate improves the behavior of the validating webhook that is called on creation of the Cluster resources |
| ValidateClusterChange | [OperatorValidateClusterChangeRequest](#cnpgi-operator-v1-OperatorValidateClusterChangeRequest) | [OperatorValidateClusterChangeResult](#cnpgi-operator-v1-OperatorValidateClusterChangeResult) | ValidateClusterChange improves the behavior of the validating webhook of is called on updates of the Cluster resources |
| MutateCluster | [OperatorMutateClusterRequest](#cnpgi-operator-v1-OperatorMutateClusterRequest) | [OperatorMutateClusterResult](#cnpgi-operator-v1-OperatorMutateClusterResult) | MutateCluster fills in the defaults inside a Cluster resource |

 



<a name="operator_lifecycle-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## operator_lifecycle.proto



<a name="cnpgi-operator_lifecycle-v1-OperatorLifecycleCapabilities"></a>

### OperatorLifecycleCapabilities



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [string](#string) |  | The Kubernetes resource group (such as &#34;apps&#34; or empty for core resources) |
| kind | [string](#string) |  | The Kubernetes Kind name (such as &#34;Cluster&#34; or &#34;Pod&#34;) |
| operation_types | [OperatorOperationType](#cnpgi-operator_lifecycle-v1-OperatorOperationType) | repeated | The operation type |






<a name="cnpgi-operator_lifecycle-v1-OperatorLifecycleCapabilitiesRequest"></a>

### OperatorLifecycleCapabilitiesRequest
This message is intentionally empty






<a name="cnpgi-operator_lifecycle-v1-OperatorLifecycleCapabilitiesResponse"></a>

### OperatorLifecycleCapabilitiesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| lifecycle_capabilities | [OperatorLifecycleCapabilities](#cnpgi-operator_lifecycle-v1-OperatorLifecycleCapabilities) | repeated | This message is OPTIONAL, containing the list of resources for which the lifecycle hook is called |






<a name="cnpgi-operator_lifecycle-v1-OperatorLifecycleRequest"></a>

### OperatorLifecycleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| operation_type | [OperatorOperationType](#cnpgi-operator_lifecycle-v1-OperatorOperationType) |  | This field is REQUIRED. |
| cluster_definition | [bytes](#bytes) |  | This field is REQUIRED |
| object_definition | [bytes](#bytes) |  | This field is REQUIRED. |






<a name="cnpgi-operator_lifecycle-v1-OperatorLifecycleResponse"></a>

### OperatorLifecycleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| json_patch | [bytes](#bytes) |  | This field is OPTIONAL. |






<a name="cnpgi-operator_lifecycle-v1-OperatorOperationType"></a>

### OperatorOperationType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [OperatorOperationType.Type](#cnpgi-operator_lifecycle-v1-OperatorOperationType-Type) |  |  |





 


<a name="cnpgi-operator_lifecycle-v1-OperatorOperationType-Type"></a>

### OperatorOperationType.Type
The operator type corresponds to the Kubernetes API method

| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| TYPE_PATCH | 1 |  |
| TYPE_UPDATE | 2 |  |
| TYPE_CREATE | 3 |  |
| TYPE_DELETE | 4 |  |


 

 


<a name="cnpgi-operator_lifecycle-v1-OperatorLifecycle"></a>

### OperatorLifecycle


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCapabilities | [OperatorLifecycleCapabilitiesRequest](#cnpgi-operator_lifecycle-v1-OperatorLifecycleCapabilitiesRequest) | [OperatorLifecycleCapabilitiesResponse](#cnpgi-operator_lifecycle-v1-OperatorLifecycleCapabilitiesResponse) | GetCapabilities gets the capabilities of the Lifecycle service |
| LifecycleHook | [OperatorLifecycleRequest](#cnpgi-operator_lifecycle-v1-OperatorLifecycleRequest) | [OperatorLifecycleResponse](#cnpgi-operator_lifecycle-v1-OperatorLifecycleResponse) | LifecycleHook allows the plugin to manipulate the Kubernetes resources before the CNPG operator applies them to the Kubernetes cluster. |

 



<a name="reconciler-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## reconciler.proto



<a name="cnpgi-reconciler-v1-ReconcilerHooksCapabilitiesRequest"></a>

### ReconcilerHooksCapabilitiesRequest







<a name="cnpgi-reconciler-v1-ReconcilerHooksCapabilitiesResult"></a>

### ReconcilerHooksCapabilitiesResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| reconciler_capabilities | [ReconcilerHooksCapability](#cnpgi-reconciler-v1-ReconcilerHooksCapability) | repeated | This message is OPTIONAL, containing the list of resources for which the lifecycle hook is called |






<a name="cnpgi-reconciler-v1-ReconcilerHooksCapability"></a>

### ReconcilerHooksCapability



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [ReconcilerHooksCapability.Kind](#cnpgi-reconciler-v1-ReconcilerHooksCapability-Kind) |  | kind is the controller Kind |






<a name="cnpgi-reconciler-v1-ReconcilerHooksRequest"></a>

### ReconcilerHooksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_definition | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the Cluster corresponding to the Pod being applied |
| resource_definition | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the resource being reconciled |






<a name="cnpgi-reconciler-v1-ReconcilerHooksResult"></a>

### ReconcilerHooksResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| behavior | [ReconcilerHooksResult.Behavior](#cnpgi-reconciler-v1-ReconcilerHooksResult-Behavior) |  | This field is REQUIRED, and indicates the behavior that should be used for the current reconciliation loop. |
| requeue_after | [int64](#int64) |  | This field is OPTIONAL. If true, the current reconciliation loop will be stopped and the operator will ensure that another one will be run in the requested number of seconds. IMPORTANT: the new reconciliation loop may start even before the number of specified seconds. |





 


<a name="cnpgi-reconciler-v1-ReconcilerHooksCapability-Kind"></a>

### ReconcilerHooksCapability.Kind


| Name | Number | Description |
| ---- | ------ | ----------- |
| KIND_UNSPECIFIED | 0 |  |
| KIND_CLUSTER | 1 | KIND_CLUSTER indicates that the plugin will plug the Cluster reconciler |
| KIND_BACKUP | 2 | KIND_BACKUP indicates that the plugin will plug the Backup reconciler |



<a name="cnpgi-reconciler-v1-ReconcilerHooksResult-Behavior"></a>

### ReconcilerHooksResult.Behavior


| Name | Number | Description |
| ---- | ------ | ----------- |
| BEHAVIOR_UNSPECIFIED | 0 |  |
| BEHAVIOR_CONTINUE | 1 | BEHAVIOR_CONTINUE indicates that this reconciliation loop will proceed running. |
| BEHAVIOR_REQUEUE | 2 | BEHAVIOR_REQUEUE indicates that this reconciliation loop will be stopped and a new one will be requested |
| BEHAVIOR_TERMINATE | 3 | BEHAVIOR_TERMINATE indicates that this reconciliation loop will be marked as succeeded and no other operations will be done. |


 

 


<a name="cnpgi-reconciler-v1-ReconcilerHooks"></a>

### ReconcilerHooks


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCapabilities | [ReconcilerHooksCapabilitiesRequest](#cnpgi-reconciler-v1-ReconcilerHooksCapabilitiesRequest) | [ReconcilerHooksCapabilitiesResult](#cnpgi-reconciler-v1-ReconcilerHooksCapabilitiesResult) | GetCapabilities gets the capabilities of the Backup service |
| Pre | [ReconcilerHooksRequest](#cnpgi-reconciler-v1-ReconcilerHooksRequest) | [ReconcilerHooksResult](#cnpgi-reconciler-v1-ReconcilerHooksResult) | Pre is executed before the operator executes the reconciliation loop |
| Post | [ReconcilerHooksRequest](#cnpgi-reconciler-v1-ReconcilerHooksRequest) | [ReconcilerHooksResult](#cnpgi-reconciler-v1-ReconcilerHooksResult) | Post is executed after the operator executes the reconciliation loop |

 



<a name="wal-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## wal.proto



<a name="cnpgi-wal-v1-SetFirstRequiredRequest"></a>

### SetFirstRequiredRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_definition | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the Cluster corresponding to the Pod being applied |
| first_required_wal | [string](#string) |  | This field is REQUIRED. Value of this field is the name of the first required WAL in the WAL archive for this cluster (normally based on the begin WAL of the first available base backup for the cluster) |






<a name="cnpgi-wal-v1-SetFirstRequiredResult"></a>

### SetFirstRequiredResult
Intentionally empty.






<a name="cnpgi-wal-v1-WALArchiveRequest"></a>

### WALArchiveRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_definition | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the Cluster corresponding to the Pod being applied |
| source_file_name | [string](#string) |  | This field is REQUIRED. Value of this field is the full path of the WAL file that should be archived |
| parameters | [WALArchiveRequest.ParametersEntry](#cnpgi-wal-v1-WALArchiveRequest-ParametersEntry) | repeated | This field is OPTIONAL. Values are opaque. |






<a name="cnpgi-wal-v1-WALArchiveRequest-ParametersEntry"></a>

### WALArchiveRequest.ParametersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="cnpgi-wal-v1-WALArchiveResult"></a>

### WALArchiveResult
Intentionally empty.






<a name="cnpgi-wal-v1-WALCapabilitiesRequest"></a>

### WALCapabilitiesRequest
Intentionally empty.






<a name="cnpgi-wal-v1-WALCapabilitiesResult"></a>

### WALCapabilitiesResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capabilities | [WALCapability](#cnpgi-wal-v1-WALCapability) | repeated | All the capabilities that the controller service supports. This field is OPTIONAL. |






<a name="cnpgi-wal-v1-WALCapability"></a>

### WALCapability



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rpc | [WALCapability.RPC](#cnpgi-wal-v1-WALCapability-RPC) |  |  |






<a name="cnpgi-wal-v1-WALCapability-RPC"></a>

### WALCapability.RPC



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [WALCapability.RPC.Type](#cnpgi-wal-v1-WALCapability-RPC-Type) |  |  |






<a name="cnpgi-wal-v1-WALRestoreRequest"></a>

### WALRestoreRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_definition | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the Cluster corresponding to the Pod being applied |
| source_wal_name | [string](#string) |  | This field is REQUIRED. Value of this field is the name of the WAL to be retrieved from the archive, such as: 000000010000000100000012 |
| destination_file_name | [string](#string) |  | This field is REQUIRED. Value of this field is the full path where the WAL file should be stored |
| parameters | [WALRestoreRequest.ParametersEntry](#cnpgi-wal-v1-WALRestoreRequest-ParametersEntry) | repeated | This field is OPTIONAL. Values are opaque. |






<a name="cnpgi-wal-v1-WALRestoreRequest-ParametersEntry"></a>

### WALRestoreRequest.ParametersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="cnpgi-wal-v1-WALRestoreResult"></a>

### WALRestoreResult
Intentionally empty.






<a name="cnpgi-wal-v1-WALStatusRequest"></a>

### WALStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_definition | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the Cluster corresponding to the Pod being applied |






<a name="cnpgi-wal-v1-WALStatusResult"></a>

### WALStatusResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| first_wal | [string](#string) |  | This field is REQUIRED. Value of this field is the base name of the oldest archived WAL, such as: 000000010000000100000012 |
| last_wal | [string](#string) |  | This field is REQUIRED. Value of this field is the base name of the newest archived WAL, such as: 000000010000000100000014 |
| additional_information | [WALStatusResult.AdditionalInformationEntry](#cnpgi-wal-v1-WALStatusResult-AdditionalInformationEntry) | repeated | This field is OPTIONAL. Value is opaque. |






<a name="cnpgi-wal-v1-WALStatusResult-AdditionalInformationEntry"></a>

### WALStatusResult.AdditionalInformationEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |





 


<a name="cnpgi-wal-v1-WALCapability-RPC-Type"></a>

### WALCapability.RPC.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| TYPE_ARCHIVE_WAL | 1 | TYPE_ARCHIVE_WAL indicates that the Plugin is able to reply to the Archive RPC request |
| TYPE_RESTORE_WAL | 2 | TYPE_RESTORE_WAL indicates that the Plugin is able to reply to the Restore RPC request |
| TYPE_STATUS | 3 | TYPE_STATUS indicates that the Plugin is able to reply to the Status RPC request |
| TYPE_SET_FIRST_REQUIRED | 4 | TYPE_SET_FIRST_REQUIRED indicates that the Plugin is able to reply to the SetFirstRequired RPC request |


 

 


<a name="cnpgi-wal-v1-WAL"></a>

### WAL


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCapabilities | [WALCapabilitiesRequest](#cnpgi-wal-v1-WALCapabilitiesRequest) | [WALCapabilitiesResult](#cnpgi-wal-v1-WALCapabilitiesResult) | GetCapabilities gets the capabilities of the WAL service |
| Archive | [WALArchiveRequest](#cnpgi-wal-v1-WALArchiveRequest) | [WALArchiveResult](#cnpgi-wal-v1-WALArchiveResult) | Archive copies one WAL file into the archive |
| Restore | [WALRestoreRequest](#cnpgi-wal-v1-WALRestoreRequest) | [WALRestoreResult](#cnpgi-wal-v1-WALRestoreResult) | Restores copies WAL file from the archive to the data directory |
| Status | [WALStatusRequest](#cnpgi-wal-v1-WALStatusRequest) | [WALStatusResult](#cnpgi-wal-v1-WALStatusResult) | Status gets the statistics of the WAL file archive |
| SetFirstRequired | [SetFirstRequiredRequest](#cnpgi-wal-v1-SetFirstRequiredRequest) | [SetFirstRequiredResult](#cnpgi-wal-v1-SetFirstRequiredResult) | SetFirstRequired sets the first required WAL for the cluster |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

