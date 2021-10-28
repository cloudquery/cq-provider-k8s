
# Table: k8s_apps_deployment_template_volumes
Volume represents a named volume in a pod that may be accessed by any container in the pod.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|deployment_cq_id|uuid|Unique CloudQuery ID of k8s_apps_deployments table (FK)|
|name|text|Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names|
|host_path|text|Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath|
|host_path_type|text|Type for HostPath Volume Defaults to "" More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath +optional|
|empty_dir_medium|text|What type of storage medium should back this directory. The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir +optional|
|empty_dir_size_limit_format|text||
|gce_persistent_disk|jsonb|GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk +optional|
|aws_elastic_block_store|jsonb|AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore +optional|
|git_repo|jsonb|GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated|
|secret|jsonb|Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret +optional|
|nfs|jsonb|NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs +optional|
|iscsi|jsonb|ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md +optional|
|glusterfs|jsonb|Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md +optional|
|persistent_volume_claim|jsonb|PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims +optional|
|rbd|jsonb|RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md +optional|
|flex_volume|jsonb|FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin. +optional|
|cinder|jsonb|Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md +optional|
|ceph_fs|jsonb|CephFS represents a Ceph FS mount on the host that shares a pod's lifetime +optional|
|flocker|jsonb|Flocker represents a Flocker volume attached to a kubelet's host machine|
|downward_api|jsonb|DownwardAPI represents downward API about the pod that should populate this volume +optional|
|fc|jsonb|FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod. +optional|
|azure_file|jsonb|AzureFile represents an Azure File Service mount on the host and bind mount to the pod. +optional|
|config_map|jsonb|ConfigMap represents a configMap that should populate this volume +optional|
|vsphere_volume|jsonb|VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine +optional|
|quobyte|jsonb|Quobyte represents a Quobyte mount on the host that shares a pod's lifetime +optional|
|azure_disk|jsonb|AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod. +optional|
|photon_persistent_disk|jsonb|PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine|
|projected|jsonb|Items for all in one resources secrets, configmaps, and downward API|
|portworx_volume|jsonb|PortworxVolume represents a portworx volume attached and mounted on kubelets host machine +optional|
|scale_io|jsonb|ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes. +optional|
|storage_os|jsonb|StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes. +optional|
|csi|jsonb|CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature). +optional|
|ephemeral|jsonb|Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more    information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.  This is a beta feature and only available when the GenericEphemeralVolume feature gate is enabled.  +optional|
