# RestoreOperationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | type of operation, one has to set it to &#39;restore&#39; in case he wants this request to be considered as a backup one  | 
**Id** | Pointer to **string** | unique identifier of an operation, a random id is assigned to each operation after a request is submitted, from caller&#39;s perspective, an id is sent back as a response to his request so he can further query state of that operation, referencing id, by operations/{id} endpoint  | 
**State** | Pointer to **string** | state of an operation, operation might be in various states, PENDING - this operation is pending for being submitted. RUNNING - this operation is actively doing its job, COMPLETED - this operation has finished successfully, CANCELLED - this operation was interrupted while being run, FAILED - this operation has finished errorneously  | 
**Progress** | Pointer to **float32** | float from 0.0 to 1.0, 1.0 telling that operation is completed, either successfully or with errors.  | 
**CreationTime** | Pointer to [**time.Time**](time.Time.md) | timestamp telling when this operation was created on Sidecar&#39;s side  | 
**StartTime** | Pointer to [**time.Time**](time.Time.md) | timestamp telling when this operation was started by Sidecar, if an operation is created, it does not necessarily mean that it will be started right away, in most cases it is the case but if e.g. ExecutorService is full on its working thread, an execution of an operation is postponed and start time is updated only after that  | [optional] 
**CompletionTime** | Pointer to [**time.Time**](time.Time.md) | timestamp telling when an operation has finished, irrelevant of its result, an operation can be failed and it would still have this field populated.  | [optional] 
**FailureCause** | Pointer to [**map[string]interface{}**](.md) | This field contains serialized java.lang.Throwable in case this operation has failed  | [optional] 
**StorageLocation** | Pointer to **string** | similar to field in backup request but used for telling from where files should be downloaded, not uploaded, in case globalRequest field is set to true, it does not matter what dc and node id is used, these components in storageLocation path will be automatically changed.  | 
**ConcurrentConnections** | Pointer to **int32** | similar to field in backup request but used for downloading files, not uploading them  | [optional] 
**LockFile** | Pointer to **string** | similar to field in backup request  | [optional] 
**CassandraDirectory** | Pointer to **string** | similar to field in backup request  | [optional] 
**CassandraConfigDirectory** | Pointer to **string** | directory where one expects to find &#39;conf/cassandra.yaml&#39; file in case we need to update it with initial tokens in case restoration strategy is IN_PLACE.  | [optional] 
**RestoreSystemKeyspace** | Pointer to **bool** | a flag saying if we should restore system keyspaces as well, relevant only for IN_PLACE restoration  | [optional] 
**SnapshotTag** | Pointer to **string** | name of snapshot to restore  | 
**Entities** | Pointer to **string** | similar to field in backup request, when empty, all entities in given snapshot will be restored  | [optional] 
**UpdateCassandraYaml** | Pointer to **bool** | flag telling if cassandra.yaml should be updated with initial_tokens, relevant only in case of IN_PLACE strategy  | [optional] 
**RestorationStrategyType** | Pointer to **string** | strategy telling how we should go about restoration, please refer to details in backup and sidecar documentation  | [optional] 
**RestorationPhase** | Pointer to **string** | phase telling what should we do, this field has to be set just once as DOWNLOAD if globalRequest if true and coordinator of that request will take care of all other phases automatically on its own  | [optional] 
**NoDeleteTruncates** | Pointer to **bool** | flag saying if we should not delete truncated SSTables after they are imported, as part of CLEANUP phase, defaults to false  | [optional] 
**NoDeleteDownloads** | Pointer to **bool** | flag saying if we should not delete downloaded SSTables from remote location, as part of CLEANUP phase, defaults to false  | [optional] 
**NoDownloadData** | Pointer to **bool** | flag saying if we should not download data from remote location as we expect them to be there already, defaults to false, setting this to true has sense only in case noDeleteDownloads was set to true in previous restoration requests  | [optional] 
**SchemaVersion** | Pointer to **string** | version of schema we want to restore from, upon backup, a schema version is automatically appended to snapshot name and its manifest is uploaded under that name. In case we have two snapshots having same name, we might distinguish between them by this schema version. If schema version is not specified, we expect that there will be one and only one backup taken with respective snapshot name. This schema version has to match the version of a Cassandra node we are doing restore for (hence, by proxy, when global request mode is used, all nodes have to be on exact same schema version).  | [optional] 
**ExactSchemaVersion** | Pointer to **bool** | flag saying if we indeed want a schema version of a running node match with schema version a snapshot is taken on. there might be cases when we want to restore a table for which its CQL schema has not changed but it has changed for other table / keyspace but a schema for that node has changed by doing that.  | [optional] 
**Import** | Pointer to [**ImportOperationRequest**](ImportOperationRequest.md) |  | [optional] 
**GlobalRequest** | Pointer to **bool** | flag saying that this request is a global one, meaning that a Sidecar this request is sent to will act as a restoration coordinator sending all other requests to each node in a cluster, for each phase.  | [optional] 

## Methods

### NewRestoreOperationResponse

`func NewRestoreOperationResponse(type_ string, id string, state string, progress float32, creationTime time.Time, storageLocation string, snapshotTag string, ) *RestoreOperationResponse`

NewRestoreOperationResponse instantiates a new RestoreOperationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreOperationResponseWithDefaults

`func NewRestoreOperationResponseWithDefaults() *RestoreOperationResponse`

NewRestoreOperationResponseWithDefaults instantiates a new RestoreOperationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RestoreOperationResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreOperationResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreOperationResponse) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *RestoreOperationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreOperationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreOperationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *RestoreOperationResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RestoreOperationResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RestoreOperationResponse) SetState(v string)`

SetState sets State field to given value.


### GetProgress

`func (o *RestoreOperationResponse) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *RestoreOperationResponse) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *RestoreOperationResponse) SetProgress(v float32)`

SetProgress sets Progress field to given value.


### GetCreationTime

`func (o *RestoreOperationResponse) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RestoreOperationResponse) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RestoreOperationResponse) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetStartTime

`func (o *RestoreOperationResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RestoreOperationResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RestoreOperationResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *RestoreOperationResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetCompletionTime

`func (o *RestoreOperationResponse) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *RestoreOperationResponse) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *RestoreOperationResponse) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *RestoreOperationResponse) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetFailureCause

`func (o *RestoreOperationResponse) GetFailureCause() map[string]interface{}`

GetFailureCause returns the FailureCause field if non-nil, zero value otherwise.

### GetFailureCauseOk

`func (o *RestoreOperationResponse) GetFailureCauseOk() (*map[string]interface{}, bool)`

GetFailureCauseOk returns a tuple with the FailureCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCause

`func (o *RestoreOperationResponse) SetFailureCause(v map[string]interface{})`

SetFailureCause sets FailureCause field to given value.

### HasFailureCause

`func (o *RestoreOperationResponse) HasFailureCause() bool`

HasFailureCause returns a boolean if a field has been set.

### GetStorageLocation

`func (o *RestoreOperationResponse) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *RestoreOperationResponse) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *RestoreOperationResponse) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.


### GetConcurrentConnections

`func (o *RestoreOperationResponse) GetConcurrentConnections() int32`

GetConcurrentConnections returns the ConcurrentConnections field if non-nil, zero value otherwise.

### GetConcurrentConnectionsOk

`func (o *RestoreOperationResponse) GetConcurrentConnectionsOk() (*int32, bool)`

GetConcurrentConnectionsOk returns a tuple with the ConcurrentConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentConnections

`func (o *RestoreOperationResponse) SetConcurrentConnections(v int32)`

SetConcurrentConnections sets ConcurrentConnections field to given value.

### HasConcurrentConnections

`func (o *RestoreOperationResponse) HasConcurrentConnections() bool`

HasConcurrentConnections returns a boolean if a field has been set.

### GetLockFile

`func (o *RestoreOperationResponse) GetLockFile() string`

GetLockFile returns the LockFile field if non-nil, zero value otherwise.

### GetLockFileOk

`func (o *RestoreOperationResponse) GetLockFileOk() (*string, bool)`

GetLockFileOk returns a tuple with the LockFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockFile

`func (o *RestoreOperationResponse) SetLockFile(v string)`

SetLockFile sets LockFile field to given value.

### HasLockFile

`func (o *RestoreOperationResponse) HasLockFile() bool`

HasLockFile returns a boolean if a field has been set.

### GetCassandraDirectory

`func (o *RestoreOperationResponse) GetCassandraDirectory() string`

GetCassandraDirectory returns the CassandraDirectory field if non-nil, zero value otherwise.

### GetCassandraDirectoryOk

`func (o *RestoreOperationResponse) GetCassandraDirectoryOk() (*string, bool)`

GetCassandraDirectoryOk returns a tuple with the CassandraDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraDirectory

`func (o *RestoreOperationResponse) SetCassandraDirectory(v string)`

SetCassandraDirectory sets CassandraDirectory field to given value.

### HasCassandraDirectory

`func (o *RestoreOperationResponse) HasCassandraDirectory() bool`

HasCassandraDirectory returns a boolean if a field has been set.

### GetCassandraConfigDirectory

`func (o *RestoreOperationResponse) GetCassandraConfigDirectory() string`

GetCassandraConfigDirectory returns the CassandraConfigDirectory field if non-nil, zero value otherwise.

### GetCassandraConfigDirectoryOk

`func (o *RestoreOperationResponse) GetCassandraConfigDirectoryOk() (*string, bool)`

GetCassandraConfigDirectoryOk returns a tuple with the CassandraConfigDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraConfigDirectory

`func (o *RestoreOperationResponse) SetCassandraConfigDirectory(v string)`

SetCassandraConfigDirectory sets CassandraConfigDirectory field to given value.

### HasCassandraConfigDirectory

`func (o *RestoreOperationResponse) HasCassandraConfigDirectory() bool`

HasCassandraConfigDirectory returns a boolean if a field has been set.

### GetRestoreSystemKeyspace

`func (o *RestoreOperationResponse) GetRestoreSystemKeyspace() bool`

GetRestoreSystemKeyspace returns the RestoreSystemKeyspace field if non-nil, zero value otherwise.

### GetRestoreSystemKeyspaceOk

`func (o *RestoreOperationResponse) GetRestoreSystemKeyspaceOk() (*bool, bool)`

GetRestoreSystemKeyspaceOk returns a tuple with the RestoreSystemKeyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreSystemKeyspace

`func (o *RestoreOperationResponse) SetRestoreSystemKeyspace(v bool)`

SetRestoreSystemKeyspace sets RestoreSystemKeyspace field to given value.

### HasRestoreSystemKeyspace

`func (o *RestoreOperationResponse) HasRestoreSystemKeyspace() bool`

HasRestoreSystemKeyspace returns a boolean if a field has been set.

### GetSnapshotTag

`func (o *RestoreOperationResponse) GetSnapshotTag() string`

GetSnapshotTag returns the SnapshotTag field if non-nil, zero value otherwise.

### GetSnapshotTagOk

`func (o *RestoreOperationResponse) GetSnapshotTagOk() (*string, bool)`

GetSnapshotTagOk returns a tuple with the SnapshotTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTag

`func (o *RestoreOperationResponse) SetSnapshotTag(v string)`

SetSnapshotTag sets SnapshotTag field to given value.


### GetEntities

`func (o *RestoreOperationResponse) GetEntities() string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *RestoreOperationResponse) GetEntitiesOk() (*string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *RestoreOperationResponse) SetEntities(v string)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *RestoreOperationResponse) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetUpdateCassandraYaml

`func (o *RestoreOperationResponse) GetUpdateCassandraYaml() bool`

GetUpdateCassandraYaml returns the UpdateCassandraYaml field if non-nil, zero value otherwise.

### GetUpdateCassandraYamlOk

`func (o *RestoreOperationResponse) GetUpdateCassandraYamlOk() (*bool, bool)`

GetUpdateCassandraYamlOk returns a tuple with the UpdateCassandraYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateCassandraYaml

`func (o *RestoreOperationResponse) SetUpdateCassandraYaml(v bool)`

SetUpdateCassandraYaml sets UpdateCassandraYaml field to given value.

### HasUpdateCassandraYaml

`func (o *RestoreOperationResponse) HasUpdateCassandraYaml() bool`

HasUpdateCassandraYaml returns a boolean if a field has been set.

### GetRestorationStrategyType

`func (o *RestoreOperationResponse) GetRestorationStrategyType() string`

GetRestorationStrategyType returns the RestorationStrategyType field if non-nil, zero value otherwise.

### GetRestorationStrategyTypeOk

`func (o *RestoreOperationResponse) GetRestorationStrategyTypeOk() (*string, bool)`

GetRestorationStrategyTypeOk returns a tuple with the RestorationStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorationStrategyType

`func (o *RestoreOperationResponse) SetRestorationStrategyType(v string)`

SetRestorationStrategyType sets RestorationStrategyType field to given value.

### HasRestorationStrategyType

`func (o *RestoreOperationResponse) HasRestorationStrategyType() bool`

HasRestorationStrategyType returns a boolean if a field has been set.

### GetRestorationPhase

`func (o *RestoreOperationResponse) GetRestorationPhase() string`

GetRestorationPhase returns the RestorationPhase field if non-nil, zero value otherwise.

### GetRestorationPhaseOk

`func (o *RestoreOperationResponse) GetRestorationPhaseOk() (*string, bool)`

GetRestorationPhaseOk returns a tuple with the RestorationPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorationPhase

`func (o *RestoreOperationResponse) SetRestorationPhase(v string)`

SetRestorationPhase sets RestorationPhase field to given value.

### HasRestorationPhase

`func (o *RestoreOperationResponse) HasRestorationPhase() bool`

HasRestorationPhase returns a boolean if a field has been set.

### GetNoDeleteTruncates

`func (o *RestoreOperationResponse) GetNoDeleteTruncates() bool`

GetNoDeleteTruncates returns the NoDeleteTruncates field if non-nil, zero value otherwise.

### GetNoDeleteTruncatesOk

`func (o *RestoreOperationResponse) GetNoDeleteTruncatesOk() (*bool, bool)`

GetNoDeleteTruncatesOk returns a tuple with the NoDeleteTruncates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDeleteTruncates

`func (o *RestoreOperationResponse) SetNoDeleteTruncates(v bool)`

SetNoDeleteTruncates sets NoDeleteTruncates field to given value.

### HasNoDeleteTruncates

`func (o *RestoreOperationResponse) HasNoDeleteTruncates() bool`

HasNoDeleteTruncates returns a boolean if a field has been set.

### GetNoDeleteDownloads

`func (o *RestoreOperationResponse) GetNoDeleteDownloads() bool`

GetNoDeleteDownloads returns the NoDeleteDownloads field if non-nil, zero value otherwise.

### GetNoDeleteDownloadsOk

`func (o *RestoreOperationResponse) GetNoDeleteDownloadsOk() (*bool, bool)`

GetNoDeleteDownloadsOk returns a tuple with the NoDeleteDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDeleteDownloads

`func (o *RestoreOperationResponse) SetNoDeleteDownloads(v bool)`

SetNoDeleteDownloads sets NoDeleteDownloads field to given value.

### HasNoDeleteDownloads

`func (o *RestoreOperationResponse) HasNoDeleteDownloads() bool`

HasNoDeleteDownloads returns a boolean if a field has been set.

### GetNoDownloadData

`func (o *RestoreOperationResponse) GetNoDownloadData() bool`

GetNoDownloadData returns the NoDownloadData field if non-nil, zero value otherwise.

### GetNoDownloadDataOk

`func (o *RestoreOperationResponse) GetNoDownloadDataOk() (*bool, bool)`

GetNoDownloadDataOk returns a tuple with the NoDownloadData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDownloadData

`func (o *RestoreOperationResponse) SetNoDownloadData(v bool)`

SetNoDownloadData sets NoDownloadData field to given value.

### HasNoDownloadData

`func (o *RestoreOperationResponse) HasNoDownloadData() bool`

HasNoDownloadData returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *RestoreOperationResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RestoreOperationResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RestoreOperationResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *RestoreOperationResponse) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetExactSchemaVersion

`func (o *RestoreOperationResponse) GetExactSchemaVersion() bool`

GetExactSchemaVersion returns the ExactSchemaVersion field if non-nil, zero value otherwise.

### GetExactSchemaVersionOk

`func (o *RestoreOperationResponse) GetExactSchemaVersionOk() (*bool, bool)`

GetExactSchemaVersionOk returns a tuple with the ExactSchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactSchemaVersion

`func (o *RestoreOperationResponse) SetExactSchemaVersion(v bool)`

SetExactSchemaVersion sets ExactSchemaVersion field to given value.

### HasExactSchemaVersion

`func (o *RestoreOperationResponse) HasExactSchemaVersion() bool`

HasExactSchemaVersion returns a boolean if a field has been set.

### GetImport

`func (o *RestoreOperationResponse) GetImport() ImportOperationRequest`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *RestoreOperationResponse) GetImportOk() (*ImportOperationRequest, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *RestoreOperationResponse) SetImport(v ImportOperationRequest)`

SetImport sets Import field to given value.

### HasImport

`func (o *RestoreOperationResponse) HasImport() bool`

HasImport returns a boolean if a field has been set.

### GetGlobalRequest

`func (o *RestoreOperationResponse) GetGlobalRequest() bool`

GetGlobalRequest returns the GlobalRequest field if non-nil, zero value otherwise.

### GetGlobalRequestOk

`func (o *RestoreOperationResponse) GetGlobalRequestOk() (*bool, bool)`

GetGlobalRequestOk returns a tuple with the GlobalRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRequest

`func (o *RestoreOperationResponse) SetGlobalRequest(v bool)`

SetGlobalRequest sets GlobalRequest field to given value.

### HasGlobalRequest

`func (o *RestoreOperationResponse) HasGlobalRequest() bool`

HasGlobalRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


