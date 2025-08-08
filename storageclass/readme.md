# Kubernetes StorageClass

## What it is

A `StorageClass` is a Kubernetes API object that allows administrators to describe the "classes" of storage they offer. Different classes might map to quality-of-service levels, backup policies, or arbitrary policies determined by the cluster administrators. Each `StorageClass` contains the fields `provisioner`, `parameters`, and `reclaimPolicy`, which are used when a `PersistentVolume` belonging to the class needs to be dynamically provisioned. In essence, `StorageClass` provides a way for users to request a certain type of storage without needing to know the underlying details of the storage provider.

## Key Concepts

- **Dynamic Provisioning:** This is the primary purpose of a `StorageClass`. It allows storage volumes to be created on-demand when a user creates a `PersistentVolumeClaim` (PVC), rather than requiring a cluster administrator to manually create `PersistentVolume` (PV)s in advance.

- **Provisioner:** Determines what volume plugin is used for provisioning PVs. Kubernetes provides several internal provisioners (e.g., `kubernetes.io/gce-pd`, `kubernetes.io/aws-ebs`), but you can also use external provisioners for storage systems not supported natively.

- **`reclaimPolicy`:** Specifies what happens to the underlying storage volume after the `PersistentVolume` it is bound to is released (i.e., the PVC is deleted).

  - `Delete`: The volume is deleted. This is the default for most dynamic provisioners.
  - `Retain`: The volume is kept. The administrator must manually clean up the resource.

- **`volumeBindingMode`:** Controls when volume binding and dynamic provisioning should occur.

  - `Immediate`: (Default) Provisions and binds the volume as soon as the PVC is created.
  - `WaitForFirstConsumer`: Delays the binding and provisioning of a `PersistentVolume` until a Pod using the `PersistentVolumeClaim` is created. This is useful for topology-aware storage backends.

- **`allowVolumeExpansion`:** If set to `true`, it allows the `PersistentVolume` to be resized by editing the corresponding PVC.

## Common `kubectl` Commands

- **List all StorageClasses:**

  ```bash
  kubectl get storageclass
  ```

  _Explanation: Shows all available StorageClasses in the cluster. You can also see if one is marked as the default._

- **Describe a StorageClass:**

  ```bash
  kubectl describe storageclass <storageclass-name>
  ```

  _Explanation: Provides detailed information about a specific StorageClass, including its provisioner, parameters, and reclaim policy._

- **Create a StorageClass from a manifest file:**

  ```bash
  kubectl apply -f storageclass-manifest.yaml
  ```

  _Explanation: Creates or updates a StorageClass using a YAML file._

- **Delete a StorageClass:**
  ```bash
  kubectl delete storageclass <storageclass-name>
  ```
  _Explanation: Deletes the StorageClass object. This does not affect existing PVs created with this class._

## Example Manifest

Here is an example of a `StorageClass` for General Purpose SSD (`gp2`) volumes on AWS. A `PersistentVolumeClaim` requesting this class would dynamically provision an AWS EBS volume.

```yaml
# aws-gp2-storageclass.yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: standard
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp2
  # Optional: specify filesystem type
  fsType: ext4
# reclaimPolicy determines what happens to the volume when the PVC is deleted.
reclaimPolicy: Retain
# allowVolumeExpansion allows users to resize the volume by editing their PVC.
allowVolumeExpansion: true
# volumeBindingMode delays provisioning until a pod uses the PVC.
volumeBindingMode: WaitForFirstConsumer
```
