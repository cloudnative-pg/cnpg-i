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
  
- [metrics.proto](#metrics-proto)
    - [CollectMetric](#cnpgi-metrics-v1-CollectMetric)
    - [CollectMetricsRequest](#cnpgi-metrics-v1-CollectMetricsRequest)
    - [CollectMetricsRequest.ParametersEntry](#cnpgi-metrics-v1-CollectMetricsRequest-ParametersEntry)
    - [CollectMetricsResult](#cnpgi-metrics-v1-CollectMetricsResult)
    - [DefineMetricsRequest](#cnpgi-metrics-v1-DefineMetricsRequest)
    - [DefineMetricsResult](#cnpgi-metrics-v1-DefineMetricsResult)
    - [Metric](#cnpgi-metrics-v1-Metric)
    - [Metric.ConstLabelsEntry](#cnpgi-metrics-v1-Metric-ConstLabelsEntry)
    - [MetricType](#cnpgi-metrics-v1-MetricType)
    - [MetricsCapabilitiesRequest](#cnpgi-metrics-v1-MetricsCapabilitiesRequest)
    - [MetricsCapabilitiesResult](#cnpgi-metrics-v1-MetricsCapabilitiesResult)
    - [MetricsCapability](#cnpgi-metrics-v1-MetricsCapability)
    - [MetricsCapability.RPC](#cnpgi-metrics-v1-MetricsCapability-RPC)
  
    - [MetricType.Type](#cnpgi-metrics-v1-MetricType-Type)
    - [MetricsCapability.RPC.Type](#cnpgi-metrics-v1-MetricsCapability-RPC-Type)
  
    - [Metrics](#cnpgi-metrics-v1-Metrics)
  
- [operator.proto](#operator-proto)
    - [DeregisterRequest](#cnpgi-operator-v1-DeregisterRequest)
    - [DeregisterResponse](#cnpgi-operator-v1-DeregisterResponse)
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
    - [SetStatusInClusterRequest](#cnpgi-operator-v1-SetStatusInClusterRequest)
    - [SetStatusInClusterResponse](#cnpgi-operator-v1-SetStatusInClusterResponse)
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
  
- [postgres.proto](#postgres-proto)
    - [EnrichConfigurationRequest](#cnpgi-identity-v1-EnrichConfigurationRequest)
    - [EnrichConfigurationRequest.ConfigsEntry](#cnpgi-identity-v1-EnrichConfigurationRequest-ConfigsEntry)
    - [EnrichConfigurationResult](#cnpgi-identity-v1-EnrichConfigurationResult)
    - [EnrichConfigurationResult.ConfigsEntry](#cnpgi-identity-v1-EnrichConfigurationResult-ConfigsEntry)
    - [OperationType](#cnpgi-identity-v1-OperationType)
    - [PostgresCapabilitiesRequest](#cnpgi-identity-v1-PostgresCapabilitiesRequest)
    - [PostgresCapabilitiesResult](#cnpgi-identity-v1-PostgresCapabilitiesResult)
    - [PostgresCapability](#cnpgi-identity-v1-PostgresCapability)
    - [PostgresCapability.RPC](#cnpgi-identity-v1-PostgresCapability-RPC)
  
    - [OperationType.Type](#cnpgi-identity-v1-OperationType-Type)
    - [PostgresCapability.RPC.Type](#cnpgi-identity-v1-PostgresCapability-RPC-Type)
  
    - [Postgres](#cnpgi-identity-v1-Postgres)
  
- [reconciler.proto](#reconciler-proto)
    - [ReconcilerHooksCapabilitiesRequest](#cnpgi-reconciler-v1-ReconcilerHooksCapabilitiesRequest)
    - [ReconcilerHooksCapabilitiesResult](#cnpgi-reconciler-v1-ReconcilerHooksCapabilitiesResult)
    - [ReconcilerHooksCapability](#cnpgi-reconciler-v1-ReconcilerHooksCapability)
    - [ReconcilerHooksRequest](#cnpgi-reconciler-v1-ReconcilerHooksRequest)
    - [ReconcilerHooksResult](#cnpgi-reconciler-v1-ReconcilerHooksResult)
  
    - [ReconcilerHooksCapability.Kind](#cnpgi-reconciler-v1-ReconcilerHooksCapability-Kind)
    - [ReconcilerHooksResult.Behavior](#cnpgi-reconciler-v1-ReconcilerHooksResult-Behavior)
  
    - [ReconcilerHooks](#cnpgi-reconciler-v1-ReconcilerHooks)
  
- [restore_job.proto](#restore_job-proto)
    - [RestoreJobHooksCapabilitiesRequest](#cnpgi-wal-v1-RestoreJobHooksCapabilitiesRequest)
    - [RestoreJobHooksCapabilitiesResult](#cnpgi-wal-v1-RestoreJobHooksCapabilitiesResult)
    - [RestoreJobHooksCapability](#cnpgi-wal-v1-RestoreJobHooksCapability)
    - [RestoreRequest](#cnpgi-wal-v1-RestoreRequest)
    - [RestoreResponse](#cnpgi-wal-v1-RestoreResponse)
  
    - [RestoreJobHooksCapability.Kind](#cnpgi-wal-v1-RestoreJobHooksCapability-Kind)
  
    - [RestoreJobHooks](#cnpgi-wal-v1-RestoreJobHooks)
  
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
| TYPE_RESTORE_JOB | 6 | TYPE_RESTORE_JOB_HOOKS indicates that the Plugin provides RPCs to enhance the behavior of the restore jobs |
| TYPE_POSTGRES | 7 | TYPE_POSTGRES indicates that the Plugin provides RPCs to enhance the behavior of PostgreSQL |
| TYPE_INSTANCE_SIDECAR_INJECTION | 8 | TYPE_INSTANCE_SIDECAR_INJECTION indicates that the Plugin provides a instance sidecar container |
| TYPE_INSTANCE_JOB_SIDECAR_INJECTION | 9 | TYPE_INSTANCE_JOB_SIDECAR_INJECTION indicates that the Plugin provides a job sidecar container |
| TYPE_METRICS | 10 | TYPE_METRICS indicates that the Plugin provides metrics to the instance container |


 

 


<a name="cnpgi-identity-v1-Identity"></a>

### Identity


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetPluginMetadata | [GetPluginMetadataRequest](#cnpgi-identity-v1-GetPluginMetadataRequest) | [GetPluginMetadataResponse](#cnpgi-identity-v1-GetPluginMetadataResponse) | GetPluginMetadata gets the plugin metadata |
| GetPluginCapabilities | [GetPluginCapabilitiesRequest](#cnpgi-identity-v1-GetPluginCapabilitiesRequest) | [GetPluginCapabilitiesResponse](#cnpgi-identity-v1-GetPluginCapabilitiesResponse) | GetPluginCapabilities gets information about this plugin |
| Probe | [ProbeRequest](#cnpgi-identity-v1-ProbeRequest) | [ProbeResponse](#cnpgi-identity-v1-ProbeResponse) | Probe is used to tell if the plugin is ready to receive requests |

 



<a name="metrics-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## metrics.proto



<a name="cnpgi-metrics-v1-CollectMetric"></a>

### CollectMetric



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  | value is the metric value |
| fq_name | [string](#string) |  | fq_name is the fully qualified name of the metric |
| variable_labels | [string](#string) | repeated | variable_labels are the values for the variable labels of this metric |






<a name="cnpgi-metrics-v1-CollectMetricsRequest"></a>

### CollectMetricsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_definition | [bytes](#bytes) |  | This field is REQUIRED and contains the JSON serialization of the Cluster being monitored |
| metrics_definition | [bytes](#bytes) |  | This field is REQUIRED and contains the JSON serialization of the Metrics that are being collected |
| parameters | [CollectMetricsRequest.ParametersEntry](#cnpgi-metrics-v1-CollectMetricsRequest-ParametersEntry) | repeated | This field is OPTIONAL and contains the configuration of this collection |






<a name="cnpgi-metrics-v1-CollectMetricsRequest-ParametersEntry"></a>

### CollectMetricsRequest.ParametersEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="cnpgi-metrics-v1-CollectMetricsResult"></a>

### CollectMetricsResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metrics | [CollectMetric](#cnpgi-metrics-v1-CollectMetric) | repeated | This field is REQUIRED and contains the JSON serialization of the collected metrics |






<a name="cnpgi-metrics-v1-DefineMetricsRequest"></a>

### DefineMetricsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_definition | [bytes](#bytes) |  | This field is REQUIRED and contains the JSON serialization of the Cluster being monitored |






<a name="cnpgi-metrics-v1-DefineMetricsResult"></a>

### DefineMetricsResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| metrics | [Metric](#cnpgi-metrics-v1-Metric) | repeated | This field is REQUIRED and contains the JSON serialization of the defined metrics |






<a name="cnpgi-metrics-v1-Metric"></a>

### Metric



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fq_name | [string](#string) |  | fqName is the fully qualified name of the metric |
| help | [string](#string) |  | help provides a human-readable description of the metric |
| variable_labels | [string](#string) | repeated | variable_labels are the label names that can vary for this metric |
| const_labels | [Metric.ConstLabelsEntry](#cnpgi-metrics-v1-Metric-ConstLabelsEntry) | repeated | const_labels are the constant labels applied to this metric |
| value_type | [MetricType](#cnpgi-metrics-v1-MetricType) |  | value_type indicates the Prometheus value type for this metric |






<a name="cnpgi-metrics-v1-Metric-ConstLabelsEntry"></a>

### Metric.ConstLabelsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="cnpgi-metrics-v1-MetricType"></a>

### MetricType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [MetricType.Type](#cnpgi-metrics-v1-MetricType-Type) |  |  |






<a name="cnpgi-metrics-v1-MetricsCapabilitiesRequest"></a>

### MetricsCapabilitiesRequest
Intentionally empty.






<a name="cnpgi-metrics-v1-MetricsCapabilitiesResult"></a>

### MetricsCapabilitiesResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capabilities | [MetricsCapability](#cnpgi-metrics-v1-MetricsCapability) | repeated | All the capabilities that the controller service supports. This field is OPTIONAL. |






<a name="cnpgi-metrics-v1-MetricsCapability"></a>

### MetricsCapability



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rpc | [MetricsCapability.RPC](#cnpgi-metrics-v1-MetricsCapability-RPC) |  |  |






<a name="cnpgi-metrics-v1-MetricsCapability-RPC"></a>

### MetricsCapability.RPC



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [MetricsCapability.RPC.Type](#cnpgi-metrics-v1-MetricsCapability-RPC-Type) |  |  |





 


<a name="cnpgi-metrics-v1-MetricType-Type"></a>

### MetricType.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| TYPE_COUNTER | 1 | COUNTER represents a monotonically increasing value |
| TYPE_GAUGE | 2 | GAUGE represents a value that can go up and down |
| TYPE_UNTYPED | 3 | UNTYPED represents an untyped metric |



<a name="cnpgi-metrics-v1-MetricsCapability-RPC-Type"></a>

### MetricsCapability.RPC.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| TYPE_METRICS | 1 | TYPE_METRICS indicates that the Plugin is able to collect metrics. This feature is required for every plugin exposing the Metrics service |


 

 


<a name="cnpgi-metrics-v1-Metrics"></a>

### Metrics


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCapabilities | [MetricsCapabilitiesRequest](#cnpgi-metrics-v1-MetricsCapabilitiesRequest) | [MetricsCapabilitiesResult](#cnpgi-metrics-v1-MetricsCapabilitiesResult) | GetCapabilities gets the capabilities of the Metrics service |
| Define | [DefineMetricsRequest](#cnpgi-metrics-v1-DefineMetricsRequest) | [DefineMetricsResult](#cnpgi-metrics-v1-DefineMetricsResult) | Define the metrics that are collected |
| Collect | [CollectMetricsRequest](#cnpgi-metrics-v1-CollectMetricsRequest) | [CollectMetricsResult](#cnpgi-metrics-v1-CollectMetricsResult) | Collect the defined metrics |

 



<a name="operator-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## operator.proto



<a name="cnpgi-operator-v1-DeregisterRequest"></a>

### DeregisterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| definition | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the Cluster that should receive the default values |






<a name="cnpgi-operator-v1-DeregisterResponse"></a>

### DeregisterResponse







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






<a name="cnpgi-operator-v1-SetStatusInClusterRequest"></a>

### SetStatusInClusterRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the current Cluster definition |






<a name="cnpgi-operator-v1-SetStatusInClusterResponse"></a>

### SetStatusInClusterResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| json_status | [bytes](#bytes) |  | This field is OPTIONAL. No value means no-op, otherwise the passed json is set inside the .status.plugins[pluginName] key |






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
| TYPE_SET_STATUS_IN_CLUSTER | 5 | TYPE_SET_STATUS_IN_CLUSTER indicates that the Plugin is able to reply to the SetStatusInCluster RPC request |
| TYPE_DEREGISTER | 6 | TYPE_DEREGISTER indicates that the Plugin is able to execute the cleanup logic once it is removed from the cluster definition |


 

 


<a name="cnpgi-operator-v1-Operator"></a>

### Operator


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCapabilities | [OperatorCapabilitiesRequest](#cnpgi-operator-v1-OperatorCapabilitiesRequest) | [OperatorCapabilitiesResult](#cnpgi-operator-v1-OperatorCapabilitiesResult) | GetCapabilities gets the capabilities of the WAL service |
| ValidateClusterCreate | [OperatorValidateClusterCreateRequest](#cnpgi-operator-v1-OperatorValidateClusterCreateRequest) | [OperatorValidateClusterCreateResult](#cnpgi-operator-v1-OperatorValidateClusterCreateResult) | ValidateCreate improves the behavior of the validating webhook that is called on creation of the Cluster resources |
| ValidateClusterChange | [OperatorValidateClusterChangeRequest](#cnpgi-operator-v1-OperatorValidateClusterChangeRequest) | [OperatorValidateClusterChangeResult](#cnpgi-operator-v1-OperatorValidateClusterChangeResult) | ValidateClusterChange improves the behavior of the validating webhook of is called on updates of the Cluster resources |
| MutateCluster | [OperatorMutateClusterRequest](#cnpgi-operator-v1-OperatorMutateClusterRequest) | [OperatorMutateClusterResult](#cnpgi-operator-v1-OperatorMutateClusterResult) | MutateCluster fills in the defaults inside a Cluster resource |
| SetStatusInCluster | [SetStatusInClusterRequest](#cnpgi-operator-v1-SetStatusInClusterRequest) | [SetStatusInClusterResponse](#cnpgi-operator-v1-SetStatusInClusterResponse) | SetStatusInCluster is invoked at the end of the reconciliation loop and it is used to set the plugin status inside the .status.plugins[pluginName] map key |
| Deregister | [DeregisterRequest](#cnpgi-operator-v1-DeregisterRequest) | [DeregisterResponse](#cnpgi-operator-v1-DeregisterResponse) | Deregister is invoked when the plugin is removed from the cluster definition. It is expected that the plugin executes its cleanup logic when this method is invoked. |

 



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
| TYPE_DEREGISTER | 5 |  |
| TYPE_EVALUATE | 6 |  |


 

 


<a name="cnpgi-operator_lifecycle-v1-OperatorLifecycle"></a>

### OperatorLifecycle


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCapabilities | [OperatorLifecycleCapabilitiesRequest](#cnpgi-operator_lifecycle-v1-OperatorLifecycleCapabilitiesRequest) | [OperatorLifecycleCapabilitiesResponse](#cnpgi-operator_lifecycle-v1-OperatorLifecycleCapabilitiesResponse) | GetCapabilities gets the capabilities of the Lifecycle service |
| LifecycleHook | [OperatorLifecycleRequest](#cnpgi-operator_lifecycle-v1-OperatorLifecycleRequest) | [OperatorLifecycleResponse](#cnpgi-operator_lifecycle-v1-OperatorLifecycleResponse) | LifecycleHook allows the plugin to manipulate the Kubernetes resources before the CNPG operator applies them to the Kubernetes cluster. |

 



<a name="postgres-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## postgres.proto



<a name="cnpgi-identity-v1-EnrichConfigurationRequest"></a>

### EnrichConfigurationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configs | [EnrichConfigurationRequest.ConfigsEntry](#cnpgi-identity-v1-EnrichConfigurationRequest-ConfigsEntry) | repeated | This field is REQUIRED and represent the PostgreSQL configuration parameters as generated by the instance manager |
| cluster_definition | [bytes](#bytes) |  | This field is REQUIRED |
| operation_type | [OperationType](#cnpgi-identity-v1-OperationType) |  | This field is REQUIRED. |






<a name="cnpgi-identity-v1-EnrichConfigurationRequest-ConfigsEntry"></a>

### EnrichConfigurationRequest.ConfigsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="cnpgi-identity-v1-EnrichConfigurationResult"></a>

### EnrichConfigurationResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configs | [EnrichConfigurationResult.ConfigsEntry](#cnpgi-identity-v1-EnrichConfigurationResult-ConfigsEntry) | repeated | This field is OPTIONAL. Returns the new configuration. |






<a name="cnpgi-identity-v1-EnrichConfigurationResult-ConfigsEntry"></a>

### EnrichConfigurationResult.ConfigsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="cnpgi-identity-v1-OperationType"></a>

### OperationType



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [OperationType.Type](#cnpgi-identity-v1-OperationType-Type) |  |  |






<a name="cnpgi-identity-v1-PostgresCapabilitiesRequest"></a>

### PostgresCapabilitiesRequest
Intentionally empty






<a name="cnpgi-identity-v1-PostgresCapabilitiesResult"></a>

### PostgresCapabilitiesResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capabilities | [PostgresCapability](#cnpgi-identity-v1-PostgresCapability) | repeated | All the capabilities that the controller service supports. This field is OPTIONAL. |






<a name="cnpgi-identity-v1-PostgresCapability"></a>

### PostgresCapability



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rpc | [PostgresCapability.RPC](#cnpgi-identity-v1-PostgresCapability-RPC) |  |  |






<a name="cnpgi-identity-v1-PostgresCapability-RPC"></a>

### PostgresCapability.RPC



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [PostgresCapability.RPC.Type](#cnpgi-identity-v1-PostgresCapability-RPC-Type) |  |  |





 


<a name="cnpgi-identity-v1-OperationType-Type"></a>

### OperationType.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| TYPE_INIT | 1 |  |
| TYPE_RESTORE | 2 |  |
| TYPE_RECONCILE | 3 |  |
| TYPE_UPGRADE | 4 |  |



<a name="cnpgi-identity-v1-PostgresCapability-RPC-Type"></a>

### PostgresCapability.RPC.Type


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPE_UNSPECIFIED | 0 |  |
| TYPE_ENRICH_CONFIGURATION | 1 | TYPE_BACKUP indicates that the Plugin is able to enrich the PostgreSQL configuration via the EnrichConfiguration endpoint |


 

 


<a name="cnpgi-identity-v1-Postgres"></a>

### Postgres


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCapabilities | [PostgresCapabilitiesRequest](#cnpgi-identity-v1-PostgresCapabilitiesRequest) | [PostgresCapabilitiesResult](#cnpgi-identity-v1-PostgresCapabilitiesResult) | GetCapabilities gets the capabilities of the Backup service |
| EnrichConfiguration | [EnrichConfigurationRequest](#cnpgi-identity-v1-EnrichConfigurationRequest) | [EnrichConfigurationResult](#cnpgi-identity-v1-EnrichConfigurationResult) | EnrichConfiguration is called before applying the configuration to PostgreSQL |

 



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
| cluster_definition | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the Cluster tied to the `Kind` being reconciled |
| resource_definition | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the resource being reconciled. Please note that in case of Cluster Reconciliation, the `resource_definition` will match the `cluster_definition` |






<a name="cnpgi-reconciler-v1-ReconcilerHooksResult"></a>

### ReconcilerHooksResult
ReconcilerHooksResult response is used to instruct then the CNPG controller on which action
take after the plugin execution.


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
| BEHAVIOR_CONTINUE | 1 | BEHAVIOR_CONTINUE indicates that this reconciliation loop will proceed running. BEHAVIOR_CONTINUE is useful when the plugin executes changes on internal status or resources not directly managed by the main reconciliation loop |
| BEHAVIOR_REQUEUE | 2 | BEHAVIOR_REQUEUE indicates that this reconciliation loop will be stopped and a new one will be requested. BEHAVIOR_REQUEUE should always be set when the plugin applies changes on resources that are directly managed by the main reconciliation loop |
| BEHAVIOR_TERMINATE | 3 | BEHAVIOR_TERMINATE indicates that the main reconciliation loop needs to be marked as succeeded and no further operations needs to be taken. This should be used when the invoked Reconcile hook act as a substitute for the main CNPG reconciliation loop. An example would be a plugin that substitutes the `Backup` logic of CNPG. |


 

 


<a name="cnpgi-reconciler-v1-ReconcilerHooks"></a>

### ReconcilerHooks
ReconcilerHooks offers a way for the plugins to directly execute changes on the resources
through the kube-api server.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCapabilities | [ReconcilerHooksCapabilitiesRequest](#cnpgi-reconciler-v1-ReconcilerHooksCapabilitiesRequest) | [ReconcilerHooksCapabilitiesResult](#cnpgi-reconciler-v1-ReconcilerHooksCapabilitiesResult) | GetCapabilities gets the capabilities of the Backup service |
| Pre | [ReconcilerHooksRequest](#cnpgi-reconciler-v1-ReconcilerHooksRequest) | [ReconcilerHooksResult](#cnpgi-reconciler-v1-ReconcilerHooksResult) | Pre is executed before the operator executes the reconciliation loop It is a way for the plugins to directly execute changes on the resources through the kube-api server. |
| Post | [ReconcilerHooksRequest](#cnpgi-reconciler-v1-ReconcilerHooksRequest) | [ReconcilerHooksResult](#cnpgi-reconciler-v1-ReconcilerHooksResult) | Post is executed after the operator executes the reconciliation loop It is a way for the plugins to directly execute changes on the resources through the kube-api server. |

 



<a name="restore_job-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## restore_job.proto



<a name="cnpgi-wal-v1-RestoreJobHooksCapabilitiesRequest"></a>

### RestoreJobHooksCapabilitiesRequest







<a name="cnpgi-wal-v1-RestoreJobHooksCapabilitiesResult"></a>

### RestoreJobHooksCapabilitiesResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| capabilities | [RestoreJobHooksCapability](#cnpgi-wal-v1-RestoreJobHooksCapability) | repeated | This field is REQUIRED and contains the describe capability |






<a name="cnpgi-wal-v1-RestoreJobHooksCapability"></a>

### RestoreJobHooksCapability



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kind | [RestoreJobHooksCapability.Kind](#cnpgi-wal-v1-RestoreJobHooksCapability-Kind) |  |  |






<a name="cnpgi-wal-v1-RestoreRequest"></a>

### RestoreRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cluster_definition | [bytes](#bytes) |  | This field is REQUIRED. Value of this field is the JSON serialization of the Cluster. |






<a name="cnpgi-wal-v1-RestoreResponse"></a>

### RestoreResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| restore_config | [string](#string) |  | This field is REQUIRED. Value of this field is the string representation of PostgreSQL configuration parameters to use for the restore. |
| envs | [string](#string) | repeated | This field if REQUIRED. Environment variables to be set in the restore job, expressed as NAME=VALUE. |





 


<a name="cnpgi-wal-v1-RestoreJobHooksCapability-Kind"></a>

### RestoreJobHooksCapability.Kind


| Name | Number | Description |
| ---- | ------ | ----------- |
| KIND_UNSPECIFIED | 0 |  |
| KIND_RESTORE | 1 | KIND_RESTORE means that the plugin is able to handle Restore requests |


 

 


<a name="cnpgi-wal-v1-RestoreJobHooks"></a>

### RestoreJobHooks
RestoreJobHooks offers a way for the plugins to enhance the cluster
recovery process

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCapabilities | [RestoreJobHooksCapabilitiesRequest](#cnpgi-wal-v1-RestoreJobHooksCapabilitiesRequest) | [RestoreJobHooksCapabilitiesResult](#cnpgi-wal-v1-RestoreJobHooksCapabilitiesResult) | GetCapabilities gets the capabilities of the Backup service |
| Restore | [RestoreRequest](#cnpgi-wal-v1-RestoreRequest) | [RestoreResponse](#cnpgi-wal-v1-RestoreResponse) | Restore is called to restore a PGDATA |

 



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

