# Kubernetes DaemonSet

## What it is

A DaemonSet ensures that all (or some) Nodes run a copy of a Pod. As nodes are added to the cluster, Pods are added to them. As nodes are removed from the cluster, those Pods are garbage collected. Deleting a DaemonSet will clean up the Pods it created. This is useful for deploying system-level daemons such as log collectors, monitoring agents, or cluster storage daemons that must run on every node.

## Key Characteristics

- **Node Coverage:** Ensures one Pod per node, or on nodes matching a `nodeSelector`.
- **Automatic Scaling:** Automatically deploys Pods to new nodes as they join the cluster.
- **System-Level Tasks:** Ideal for background services that need to run on all or most nodes.
- **Update Strategy:** Supports rolling updates (`RollingUpdate`) to update Pods with zero downtime, or `OnDelete` to update only when the old Pod is manually deleted.
- **No Scheduling by Default:** The Kubernetes scheduler is not involved in placing these Pods; the DaemonSet controller handles it directly.

## Common `kubectl` Commands

- **Create a DaemonSet from a manifest file:**

  ```bash
  kubectl apply -f manifest.yaml
  ```

  _Explanation: Creates or updates a DaemonSet using a YAML or JSON manifest file._

- **Get a list of DaemonSets:**

  ```bash
  kubectl get daemonsets -n <namespace>
  ```

  _Explanation: Lists all DaemonSets in the specified namespace._

- **Describe a DaemonSet to see its status:**

  ```bash
  kubectl describe daemonset <daemonset-name> -n <namespace>
  ```

  _Explanation: Shows detailed information about a DaemonSet, including the number of pods scheduled, ready, and available._

- **Check the rollout status of a DaemonSet:**

  ```bash
  kubectl rollout status daemonset/<daemonset-name> -n <namespace>
  ```

  _Explanation: Watches the status of a DaemonSet's rolling update until it's complete._

- **Edit a DaemonSet:**

  ```bash
  kubectl edit daemonset <daemonset-name> -n <namespace>
  ```

  _Explanation: Opens the DaemonSet's manifest in the default editor for live edits._

- **Delete a DaemonSet:**
  ```bash
  kubectl delete daemonset <daemonset-name> -n <namespace>
  ```
  _Explanation: Deletes the specified DaemonSet and the Pods it manages._

## Example Manifest

Here is a basic example of a DaemonSet that runs a simple logging agent on every node:

```yaml
# daemonset-example.yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: fluentd-logging-agent
  namespace: kube-system
  labels:
    k8s-app: fluentd-logging
spec:
  selector:
    matchLabels:
      name: fluentd-logging-agent
  template:
    metadata:
      labels:
        name: fluentd-logging-agent
    spec:
      tolerations:
        # This toleration is to have the daemonset runnable on control plane nodes
        # Remove it if you don't want to run pods on the control plane nodes.
        - key: node-role.kubernetes.io/control-plane
          operator: Exists
          effect: NoSchedule
      containers:
        - name: fluentd
          image: fluent/fluentd:v1.14.5-1.0
          resources:
            limits:
              memory: 200Mi
            requests:
              cpu: 100m
              memory: 200Mi
          volumeMounts:
            - name: varlog
              mountPath: /var/log
            - name: varlibdockercontainers
              mountPath: /var/lib/docker/containers
              readOnly: true
      terminationGracePeriodSeconds: 30
      volumes:
        - name: varlog
          hostPath:
            path: /var/log
        - name: varlibdockercontainers
          hostPath:
            path: /var/lib/docker/containers
```
