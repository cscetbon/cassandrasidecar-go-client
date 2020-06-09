# BackupOperationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | type of operation, one has to set it to &#39;backup&#39; in case he wants this request to be considered as a backup one  | 
**StorageLocation** | **string** | location where SSTables will be uploaded. A value of the storageLocation property has to have exact format which is &#39;protocol://bucket/clusterName/dcName/nodeName&#39;. protocol is either &#39;gcp&#39;, &#39;s3&#39;, &#39;azure&#39; or &#39;file:/&#39;. For global requests, dcName and nodeName are changed automatically as these values are read from Cassandra and storageLocation is updated automatically for every node a specific backup request will be submitted to so the value of dcName and nodeName is irrelevant for global requests as they will be modified every time, a bucket does not need to exist, it will be created automatically if it does not, clusterName has to be specified. There might be automatic resolution of clusterName in the future however for now, one has to supply this property on his own.  | 
**CassandraDirectory** | **string** | directory of Cassandra, by default it is /var/lib/cassandra, in this path, one expects there is &#39;data&#39; directory  | [optional] 
**Duration** | **string** | Based on this field, there will be throughtput per second computed based on what size data we want to upload we have. The formula is \&quot;size / duration\&quot;. The lower the duration is, the higher throughput per second we will need and vice versa. This will influence e.g. responsiveness of a node to its business requests so one can control how much bandwidth is used for backup purposes in case a cluster is fully operational. The format of this field is \&quot;amount unit\&quot;. &#39;unit&#39; is just a (case-insensitive) java.util.concurrent.TimeUnit enum value. If not used, there will not be any restrictions as how fast an upload can be.  | [optional] 
**Bandwidth** | [**DataRate**](DataRate.md) |  | [optional] 
**ConcurrentConnections** | **int32** | number of threads used for upload, there might be at most so many uploading threads at any given time, when not set, it defaults to 10  | [optional] 
**LockFile** | **string** | path to file which will be used for locking the critical logic dealing with backups, this locking is done by locking a file on a file system so other execution will not proceed until the former one has finished. By default, this path is System.getProperty(\&quot;java.io.tmpdir\&quot;) + \&quot;/global-transfer-lock\&quot;.  | [optional] 
**SnapshotTag** | **string** | name of snapshot to make so this snapshot will be uploaded to storage location. If not specified, the name of snapshot will be automatically generated and it will have name &#39;autosnap-milliseconds-since-epoch&#39;  | [optional] 
**Dc** | **string** | name of datacenter to backup, nodes in the other datacenter(s) will not be involved  | [optional] 
**Entities** | **string** | database entities to backup, it might be either only keyspaces or only tables (from different keyspaces if needed), e.g. &#39;k1,k2&#39; if one wants to backup whole keyspaces and &#39;ks1.t1,ks2,t2&#39; if one wants to backup tables. These formats can not be used together so &#39;k1,k2.t2&#39; is invalid. If this field is empty, all keyspaces are backed up. | [optional] 
**K8sNamespace** | **string** | name of Kubernetes namespace to fetch Kubernetes secret for backups from, when not specified, it defaults to &#39;default&#39;  | [optional] 
**K8sBackupSecretName** | **string** | name of Kubernetes secret from which credentials used for the communication to cloud storage providers are read, if not specified, secret name to be read will be automatically derived in form &#39;cassandra-backup-restore-secret-cluster-{name-of-cluster}&#39;. These secrets are used only in case protocol in storageLocation is gcp, azure or s3.  | [optional] 
**GlobalRequest** | **bool** | flag saying if this request is meant to be global or not, once a global backup request is submitted to Sidecar, it will coordinate backup for all other nodes in a cluster (including itself) so from a point of view of a caller, one might just backup whole cluster by one request and repeatedly query its status based on returned operation id.  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


