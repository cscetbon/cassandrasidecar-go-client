# BackupOperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | type of operation, one has to set it to &#39;backup&#39; in case he wants this request to be considered as a backup one  | 
**Id** | Pointer to **string** | unique identifier of an operation, a random id is assigned to each operation after a request is submitted, from caller&#39;s perspective, an id is sent back as a response to his request so he can further query state of that operation, referencing id, by operations/{id} endpoint  | 
**State** | Pointer to **string** | state of an operation, operation might be in various states, PENDING - this operation is pending for being submitted. RUNNING - this operation is actively doing its job, COMPLETED - this operation has finished successfully, CANCELLED - this operation was interrupted while being run, FAILED - this operation has finished errorneously  | 
**Progress** | Pointer to **float32** | float from 0.0 to 1.0, 1.0 telling that operation is completed, either successfully or with errors.  | 
**CreationTime** | Pointer to [**time.Time**](time.Time.md) | timestamp telling when this operation was created on Sidecar&#39;s side  | 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | timestamp telling when this operation was started by Sidecar, if an operation is created, it does not necessarily mean that it will be started right away, in most cases it is the case but if e.g. ExecutorService is full on its working thread, an execution of an operation is postponed and start time is updated only after that  | [optional] 
**CompletionTime** | Pointer to [**time.Time**](time.Time.md) | timestamp telling when an operation has finished, irrelevant of its result, an operation can be failed and it would still have this field populated.  | [optional] 
**FailureCause** | Pointer to [**map[string]interface{}**](.md) | This field contains serialized java.lang.Throwable in case this operation has failed  | [optional] 
**StorageLocation** | Pointer to **string** | location where SSTables will be uploaded. A value of the storageLocation property has to have exact format which is &#39;protocol://bucket/clusterName/dcName/nodeName&#39;. protocol is either &#39;gcp&#39;, &#39;s3&#39;, &#39;azure&#39; or &#39;file:/&#39;. For global requests, dcName and nodeName are changed automatically as these values are read from Cassandra and storageLocation is updated automatically for every node a specific backup request will be submitted to so the value of dcName and nodeName is irrelevant for global requests as they will be modified every time, a bucket does not need to exist, it will be created automatically if it does not, clusterName has to be specified. There might be automatic resolution of clusterName in the future however for now, one has to supply this property on his own.  | 
**CassandraDirectory** | Pointer to **string** | directory of Cassandra, by default it is /var/lib/cassandra, in this path, one expects there is &#39;data&#39; directory  | [optional] 
**Duration** | Pointer to **string** | Based on this field, there will be throughtput per second computed based on what size data we want to upload we have. The formula is \&quot;size / duration\&quot;. The lower the duration is, the higher throughput per second we will need and vice versa. This will influence e.g. responsiveness of a node to its business requests so one can control how much bandwidth is used for backup purposes in case a cluster is fully operational. The format of this field is \&quot;amount unit\&quot;. &#39;unit&#39; is just a (case-insensitive) java.util.concurrent.TimeUnit enum value. If not used, there will not be any restrictions as how fast an upload can be.  | [optional] 
**Bandwidth** | Pointer to [**DataRate**](DataRate.md) |  | [optional] 
**ConcurrentConnections** | Pointer to **int32** | number of threads used for upload, there might be at most so many uploading threads at any given time, when not set, it defaults to 10  | [optional] 
**LockFile** | Pointer to **string** | path to file which will be used for locking the critical logic dealing with backups, this locking is done by locking a file on a file system so other execution will not proceed until the former one has finished. By default, this path is System.getProperty(\&quot;java.io.tmpdir\&quot;) + \&quot;/global-transfer-lock\&quot;.  | [optional] 
**SnapshotTag** | Pointer to **string** | name of snapshot to make so this snapshot will be uploaded to storage location. If not specified, the name of snapshot will be automatically generated and it will have name &#39;autosnap-milliseconds-since-epoch&#39;  | [optional] 
**Dc** | Pointer to **string** | name of datacenter to backup, nodes in the other datacenter(s) will not be involved  | [optional] 
**Entities** | Pointer to **string** | database entities to backup, it might be either only keyspaces or only tables (from different keyspaces if needed), e.g. &#39;k1,k2&#39; if one wants to backup whole keyspaces and &#39;ks1.t1,ks2,t2&#39; if one wants to backup tables. These formats can not be used together so &#39;k1,k2.t2&#39; is invalid. If this field is empty, all keyspaces are backed up. | [optional] 
**K8sNamespace** | Pointer to **string** | name of Kubernetes namespace to fetch Kubernetes secret for backups from, when not specified, it defaults to &#39;default&#39;  | [optional] 
**K8sBackupSecretName** | Pointer to **string** | name of Kubernetes secret from which credentials used for the communication to cloud storage providers are read, if not specified, secret name to be read will be automatically derived in form &#39;cassandra-backup-restore-secret-cluster-{name-of-cluster}&#39;. These secrets are used only in case protocol in storageLocation is gcp, azure or s3.  | [optional] 
**GlobalRequest** | Pointer to **bool** | flag saying if this request is meant to be global or not, once a global backup request is submitted to Sidecar, it will coordinate backup for all other nodes in a cluster (including itself) so from a point of view of a caller, one might just backup whole cluster by one request and repeatedly query its status based on returned operation id.  | [optional] 

## Methods

### NewBackupOperationResponse

`func NewBackupOperationResponse(type_ string, id string, state string, progress float32, creationTime time.Time, storageLocation string, ) *BackupOperationResponse`

NewBackupOperationResponse instantiates a new BackupOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupOperationResponseWithDefaults

`func NewBackupOperationResponseWithDefaults() *BackupOperationResponse`

NewBackupOperationResponseWithDefaults instantiates a new BackupOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BackupOperationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupOperationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupOperationResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BackupOperationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupOperationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupOperationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *BackupOperationResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackupOperationResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BackupOperationResponse) SetState(v string)`

SetState sets State field to given value.


### GetProgress

`func (o *BackupOperationResponse) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *BackupOperationResponse) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *BackupOperationResponse) SetProgress(v float32)`

SetProgress sets Progress field to given value.


### GetCreationTime

`func (o *BackupOperationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *BackupOperationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *BackupOperationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetStartTime

`func (o *BackupOperationResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupOperationResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupOperationResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BackupOperationResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetCompletionTime

`func (o *BackupOperationResponse) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *BackupOperationResponse) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *BackupOperationResponse) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *BackupOperationResponse) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetFailureCause

`func (o *BackupOperationResponse) GetFailureCause() map[string]interface{}`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *BackupOperationResponse) GetFailureCauseOk() (*map[string]interface{}, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *BackupOperationResponse) SetFailureCause(v map[string]interface{})`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *BackupOperationResponse) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.

### GetStorageLocation

`func (o *BackupOperationResponse) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *BackupOperationResponse) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *BackupOperationResponse) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.


### GetCassandraDirectory

`func (o *BackupOperationResponse) GetCassandraDirectory() string`

GetCassandraDirectory returns the CassandraDirectory field if non-nil, zero value otherwise.

### GetCassandraDirectoryOk

`func (o *BackupOperationResponse) GetCassandraDirectoryOk() (*string, bool)`

GetCassandraDirectoryOk returns a tuple with the CassandraDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraDirectory

`func (o *BackupOperationResponse) SetCassandraDirectory(v string)`

SetCassandraDirectory sets CassandraDirectory field to given value.

### HasCassandraDirectory

`func (o *BackupOperationResponse) HasCassandraDirectory() bool`

HasCassandraDirectory returns a boolean if a field has been set.

### GetDuration

`func (o *BackupOperationResponse) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *BackupOperationResponse) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *BackupOperationResponse) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *BackupOperationResponse) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetBandwidth

`func (o *BackupOperationResponse) GetBandwidth() DataRate`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *BackupOperationResponse) GetBandwidthOk() (*DataRate, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *BackupOperationResponse) SetBandwidth(v DataRate)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *BackupOperationResponse) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetConcurrentConnections

`func (o *BackupOperationResponse) GetConcurrentConnections() int32`

GetConcurrentConnections returns the ConcurrentConnections field if non-nil, zero value otherwise.

### GetConcurrentConnectionsOk

`func (o *BackupOperationResponse) GetConcurrentConnectionsOk() (*int32, bool)`

GetConcurrentConnectionsOk returns a tuple with the ConcurrentConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentConnections

`func (o *BackupOperationResponse) SetConcurrentConnections(v int32)`

SetConcurrentConnections sets ConcurrentConnections field to given value.

### HasConcurrentConnections

`func (o *BackupOperationResponse) HasConcurrentConnections() bool`

HasConcurrentConnections returns a boolean if a field has been set.

### GetLockFile

`func (o *BackupOperationResponse) GetLockFile() string`

GetLockFile returns the LockFile field if non-nil, zero value otherwise.

### GetLockFileOk

`func (o *BackupOperationResponse) GetLockFileOk() (*string, bool)`

GetLockFileOk returns a tuple with the LockFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockFile

`func (o *BackupOperationResponse) SetLockFile(v string)`

SetLockFile sets LockFile field to given value.

### HasLockFile

`func (o *BackupOperationResponse) HasLockFile() bool`

HasLockFile returns a boolean if a field has been set.

### GetSnapshotTag

`func (o *BackupOperationResponse) GetSnapshotTag() string`

GetSnapshotTag returns the SnapshotTag field if non-nil, zero value otherwise.

### GetSnapshotTagOk

`func (o *BackupOperationResponse) GetSnapshotTagOk() (*string, bool)`

GetSnapshotTagOk returns a tuple with the SnapshotTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTag

`func (o *BackupOperationResponse) SetSnapshotTag(v string)`

SetSnapshotTag sets SnapshotTag field to given value.

### HasSnapshotTag

`func (o *BackupOperationResponse) HasSnapshotTag() bool`

HasSnapshotTag returns a boolean if a field has been set.

### GetDc

`func (o *BackupOperationResponse) GetDc() string`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *BackupOperationResponse) GetDcOk() (*string, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *BackupOperationResponse) SetDc(v string)`

SetDc sets Dc field to given value.

### HasDc

`func (o *BackupOperationResponse) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetEntities

`func (o *BackupOperationResponse) GetEntities() string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *BackupOperationResponse) GetEntitiesOk() (*string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *BackupOperationResponse) SetEntities(v string)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *BackupOperationResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetK8sNamespace

`func (o *BackupOperationResponse) GetK8sNamespace() string`

GetK8sNamespace returns the K8sNamespace field if non-nil, zero value otherwise.

### GetK8sNamespaceOk

`func (o *BackupOperationResponse) GetK8sNamespaceOk() (*string, bool)`

GetK8sNamespaceOk returns a tuple with the K8sNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNamespace

`func (o *BackupOperationResponse) SetK8sNamespace(v string)`

SetK8sNamespace sets K8sNamespace field to given value.

### HasK8sNamespace

`func (o *BackupOperationResponse) HasK8sNamespace() bool`

HasK8sNamespace returns a boolean if a field has been set.

### GetK8sBackupSecretName

`func (o *BackupOperationResponse) GetK8sBackupSecretName() string`

GetK8sBackupSecretName returns the K8sBackupSecretName field if non-nil, zero value otherwise.

### GetK8sBackupSecretNameOk

`func (o *BackupOperationResponse) GetK8sBackupSecretNameOk() (*string, bool)`

GetK8sBackupSecretNameOk returns a tuple with the K8sBackupSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sBackupSecretName

`func (o *BackupOperationResponse) SetK8sBackupSecretName(v string)`

SetK8sBackupSecretName sets K8sBackupSecretName field to given value.

### HasK8sBackupSecretName

`func (o *BackupOperationResponse) HasK8sBackupSecretName() bool`

HasK8sBackupSecretName returns a boolean if a field has been set.

### GetGlobalRequest

`func (o *BackupOperationResponse) GetGlobalRequest() bool`

GetGlobalRequest returns the GlobalRequest field if non-nil, zero value otherwise.

### GetGlobalRequestOk

`func (o *BackupOperationResponse) GetGlobalRequestOk() (*bool, bool)`

GetGlobalRequestOk returns a tuple with the GlobalRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRequest

`func (o *BackupOperationResponse) SetGlobalRequest(v bool)`

SetGlobalRequest sets GlobalRequest field to given value.

### HasGlobalRequest

`func (o *BackupOperationResponse) HasGlobalRequest() bool`

HasGlobalRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


