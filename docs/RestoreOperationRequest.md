# RestoreOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | type of operation, one has to set it to &#39;restore&#39; in case he wants this request to be considered as a backup one  | 
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

### NewRestoreOperationRequest

`func NewRestoreOperationRequest(type_ string, storageLocation string, snapshotTag string, ) *RestoreOperationRequest`

NewRestoreOperationRequest instantiates a new RestoreOperationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreOperationRequestWithDefaults

`func NewRestoreOperationRequestWithDefaults() *RestoreOperationRequest`

NewRestoreOperationRequestWithDefaults instantiates a new RestoreOperationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RestoreOperationRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreOperationRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreOperationRequest) SetType(v string)`

SetType sets Type field to given value.


### GetStorageLocation

`func (o *RestoreOperationRequest) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *RestoreOperationRequest) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *RestoreOperationRequest) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.


### GetConcurrentConnections

`func (o *RestoreOperationRequest) GetConcurrentConnections() int32`

GetConcurrentConnections returns the ConcurrentConnections field if non-nil, zero value otherwise.

### GetConcurrentConnectionsOk

`func (o *RestoreOperationRequest) GetConcurrentConnectionsOk() (*int32, bool)`

GetConcurrentConnectionsOk returns a tuple with the ConcurrentConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentConnections

`func (o *RestoreOperationRequest) SetConcurrentConnections(v int32)`

SetConcurrentConnections sets ConcurrentConnections field to given value.

### HasConcurrentConnections

`func (o *RestoreOperationRequest) HasConcurrentConnections() bool`

HasConcurrentConnections returns a boolean if a field has been set.

### GetLockFile

`func (o *RestoreOperationRequest) GetLockFile() string`

GetLockFile returns the LockFile field if non-nil, zero value otherwise.

### GetLockFileOk

`func (o *RestoreOperationRequest) GetLockFileOk() (*string, bool)`

GetLockFileOk returns a tuple with the LockFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockFile

`func (o *RestoreOperationRequest) SetLockFile(v string)`

SetLockFile sets LockFile field to given value.

### HasLockFile

`func (o *RestoreOperationRequest) HasLockFile() bool`

HasLockFile returns a boolean if a field has been set.

### GetCassandraDirectory

`func (o *RestoreOperationRequest) GetCassandraDirectory() string`

GetCassandraDirectory returns the CassandraDirectory field if non-nil, zero value otherwise.

### GetCassandraDirectoryOk

`func (o *RestoreOperationRequest) GetCassandraDirectoryOk() (*string, bool)`

GetCassandraDirectoryOk returns a tuple with the CassandraDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraDirectory

`func (o *RestoreOperationRequest) SetCassandraDirectory(v string)`

SetCassandraDirectory sets CassandraDirectory field to given value.

### HasCassandraDirectory

`func (o *RestoreOperationRequest) HasCassandraDirectory() bool`

HasCassandraDirectory returns a boolean if a field has been set.

### GetCassandraConfigDirectory

`func (o *RestoreOperationRequest) GetCassandraConfigDirectory() string`

GetCassandraConfigDirectory returns the CassandraConfigDirectory field if non-nil, zero value otherwise.

### GetCassandraConfigDirectoryOk

`func (o *RestoreOperationRequest) GetCassandraConfigDirectoryOk() (*string, bool)`

GetCassandraConfigDirectoryOk returns a tuple with the CassandraConfigDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCassandraConfigDirectory

`func (o *RestoreOperationRequest) SetCassandraConfigDirectory(v string)`

SetCassandraConfigDirectory sets CassandraConfigDirectory field to given value.

### HasCassandraConfigDirectory

`func (o *RestoreOperationRequest) HasCassandraConfigDirectory() bool`

HasCassandraConfigDirectory returns a boolean if a field has been set.

### GetRestoreSystemKeyspace

`func (o *RestoreOperationRequest) GetRestoreSystemKeyspace() bool`

GetRestoreSystemKeyspace returns the RestoreSystemKeyspace field if non-nil, zero value otherwise.

### GetRestoreSystemKeyspaceOk

`func (o *RestoreOperationRequest) GetRestoreSystemKeyspaceOk() (*bool, bool)`

GetRestoreSystemKeyspaceOk returns a tuple with the RestoreSystemKeyspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreSystemKeyspace

`func (o *RestoreOperationRequest) SetRestoreSystemKeyspace(v bool)`

SetRestoreSystemKeyspace sets RestoreSystemKeyspace field to given value.

### HasRestoreSystemKeyspace

`func (o *RestoreOperationRequest) HasRestoreSystemKeyspace() bool`

HasRestoreSystemKeyspace returns a boolean if a field has been set.

### GetSnapshotTag

`func (o *RestoreOperationRequest) GetSnapshotTag() string`

GetSnapshotTag returns the SnapshotTag field if non-nil, zero value otherwise.

### GetSnapshotTagOk

`func (o *RestoreOperationRequest) GetSnapshotTagOk() (*string, bool)`

GetSnapshotTagOk returns a tuple with the SnapshotTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTag

`func (o *RestoreOperationRequest) SetSnapshotTag(v string)`

SetSnapshotTag sets SnapshotTag field to given value.


### GetEntities

`func (o *RestoreOperationRequest) GetEntities() string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *RestoreOperationRequest) GetEntitiesOk() (*string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *RestoreOperationRequest) SetEntities(v string)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *RestoreOperationRequest) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetUpdateCassandraYaml

`func (o *RestoreOperationRequest) GetUpdateCassandraYaml() bool`

GetUpdateCassandraYaml returns the UpdateCassandraYaml field if non-nil, zero value otherwise.

### GetUpdateCassandraYamlOk

`func (o *RestoreOperationRequest) GetUpdateCassandraYamlOk() (*bool, bool)`

GetUpdateCassandraYamlOk returns a tuple with the UpdateCassandraYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateCassandraYaml

`func (o *RestoreOperationRequest) SetUpdateCassandraYaml(v bool)`

SetUpdateCassandraYaml sets UpdateCassandraYaml field to given value.

### HasUpdateCassandraYaml

`func (o *RestoreOperationRequest) HasUpdateCassandraYaml() bool`

HasUpdateCassandraYaml returns a boolean if a field has been set.

### GetRestorationStrategyType

`func (o *RestoreOperationRequest) GetRestorationStrategyType() string`

GetRestorationStrategyType returns the RestorationStrategyType field if non-nil, zero value otherwise.

### GetRestorationStrategyTypeOk

`func (o *RestoreOperationRequest) GetRestorationStrategyTypeOk() (*string, bool)`

GetRestorationStrategyTypeOk returns a tuple with the RestorationStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorationStrategyType

`func (o *RestoreOperationRequest) SetRestorationStrategyType(v string)`

SetRestorationStrategyType sets RestorationStrategyType field to given value.

### HasRestorationStrategyType

`func (o *RestoreOperationRequest) HasRestorationStrategyType() bool`

HasRestorationStrategyType returns a boolean if a field has been set.

### GetRestorationPhase

`func (o *RestoreOperationRequest) GetRestorationPhase() string`

GetRestorationPhase returns the RestorationPhase field if non-nil, zero value otherwise.

### GetRestorationPhaseOk

`func (o *RestoreOperationRequest) GetRestorationPhaseOk() (*string, bool)`

GetRestorationPhaseOk returns a tuple with the RestorationPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorationPhase

`func (o *RestoreOperationRequest) SetRestorationPhase(v string)`

SetRestorationPhase sets RestorationPhase field to given value.

### HasRestorationPhase

`func (o *RestoreOperationRequest) HasRestorationPhase() bool`

HasRestorationPhase returns a boolean if a field has been set.

### GetNoDeleteTruncates

`func (o *RestoreOperationRequest) GetNoDeleteTruncates() bool`

GetNoDeleteTruncates returns the NoDeleteTruncates field if non-nil, zero value otherwise.

### GetNoDeleteTruncatesOk

`func (o *RestoreOperationRequest) GetNoDeleteTruncatesOk() (*bool, bool)`

GetNoDeleteTruncatesOk returns a tuple with the NoDeleteTruncates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDeleteTruncates

`func (o *RestoreOperationRequest) SetNoDeleteTruncates(v bool)`

SetNoDeleteTruncates sets NoDeleteTruncates field to given value.

### HasNoDeleteTruncates

`func (o *RestoreOperationRequest) HasNoDeleteTruncates() bool`

HasNoDeleteTruncates returns a boolean if a field has been set.

### GetNoDeleteDownloads

`func (o *RestoreOperationRequest) GetNoDeleteDownloads() bool`

GetNoDeleteDownloads returns the NoDeleteDownloads field if non-nil, zero value otherwise.

### GetNoDeleteDownloadsOk

`func (o *RestoreOperationRequest) GetNoDeleteDownloadsOk() (*bool, bool)`

GetNoDeleteDownloadsOk returns a tuple with the NoDeleteDownloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDeleteDownloads

`func (o *RestoreOperationRequest) SetNoDeleteDownloads(v bool)`

SetNoDeleteDownloads sets NoDeleteDownloads field to given value.

### HasNoDeleteDownloads

`func (o *RestoreOperationRequest) HasNoDeleteDownloads() bool`

HasNoDeleteDownloads returns a boolean if a field has been set.

### GetNoDownloadData

`func (o *RestoreOperationRequest) GetNoDownloadData() bool`

GetNoDownloadData returns the NoDownloadData field if non-nil, zero value otherwise.

### GetNoDownloadDataOk

`func (o *RestoreOperationRequest) GetNoDownloadDataOk() (*bool, bool)`

GetNoDownloadDataOk returns a tuple with the NoDownloadData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDownloadData

`func (o *RestoreOperationRequest) SetNoDownloadData(v bool)`

SetNoDownloadData sets NoDownloadData field to given value.

### HasNoDownloadData

`func (o *RestoreOperationRequest) HasNoDownloadData() bool`

HasNoDownloadData returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *RestoreOperationRequest) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *RestoreOperationRequest) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *RestoreOperationRequest) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *RestoreOperationRequest) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetExactSchemaVersion

`func (o *RestoreOperationRequest) GetExactSchemaVersion() bool`

GetExactSchemaVersion returns the ExactSchemaVersion field if non-nil, zero value otherwise.

### GetExactSchemaVersionOk

`func (o *RestoreOperationRequest) GetExactSchemaVersionOk() (*bool, bool)`

GetExactSchemaVersionOk returns a tuple with the ExactSchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactSchemaVersion

`func (o *RestoreOperationRequest) SetExactSchemaVersion(v bool)`

SetExactSchemaVersion sets ExactSchemaVersion field to given value.

### HasExactSchemaVersion

`func (o *RestoreOperationRequest) HasExactSchemaVersion() bool`

HasExactSchemaVersion returns a boolean if a field has been set.

### GetImport

`func (o *RestoreOperationRequest) GetImport() ImportOperationRequest`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *RestoreOperationRequest) GetImportOk() (*ImportOperationRequest, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *RestoreOperationRequest) SetImport(v ImportOperationRequest)`

SetImport sets Import field to given value.

### HasImport

`func (o *RestoreOperationRequest) HasImport() bool`

HasImport returns a boolean if a field has been set.

### GetGlobalRequest

`func (o *RestoreOperationRequest) GetGlobalRequest() bool`

GetGlobalRequest returns the GlobalRequest field if non-nil, zero value otherwise.

### GetGlobalRequestOk

`func (o *RestoreOperationRequest) GetGlobalRequestOk() (*bool, bool)`

GetGlobalRequestOk returns a tuple with the GlobalRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalRequest

`func (o *RestoreOperationRequest) SetGlobalRequest(v bool)`

SetGlobalRequest sets GlobalRequest field to given value.

### HasGlobalRequest

`func (o *RestoreOperationRequest) HasGlobalRequest() bool`

HasGlobalRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


