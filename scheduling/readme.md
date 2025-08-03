# Kubernetes Scheduling

## What it is

Kubernetes scheduling is the process of assigning Pods to Nodes so that the Kubelet can run them. The component responsible for this is the `kube-scheduler`, which is part of the control plane. For every new Pod, the scheduler makes a decision based on a variety of factors, including resource requirements, policy constraints, affinity and anti-affinity specifications, data locality, and inter-workload interference. The scheduler's goal is to find the most suitable Node for each Pod.

## Key Scheduling Concepts

- **Node Selector (`nodeSelector`):** The simplest way to constrain Pods to run on nodes with specific labels. You add a `nodeSelector` field to your Pod specification with key-value pairs that must match the labels on a target node.

- **Affinity and Anti-Affinity:** A more expressive and flexible set of rules for constraining which nodes your Pods can be scheduled on.

  - **Node Affinity:** Similar to `nodeSelector` but allows for more complex rules. It has "hard" (`requiredDuringSchedulingIgnoredDuringExecution`) and "soft" (`preferredDuringSchedulingIgnoredDuringExecution`) requirements.
  - **Pod Affinity/Anti-Affinity:** Schedules Pods based on the labels of other Pods already running on a node. This is useful for co-locating services (affinity) or spreading them out to avoid single points of failure (anti-affinity).

- **Taints and Tolerations:** A mechanism to repel Pods from certain nodes.

  - **Taint:** Applied to a Node to mark that it should not accept any Pods that do not tolerate the taint. Effects include `NoSchedule` (won't schedule new pods), `PreferNoSchedule` (soft version), and `NoExecute` (evicts running pods that don't tolerate it).
  - **Toleration:** Applied to a Pod to allow it to be scheduled on a node with a matching taint.

- **Topology Spread Constraints (`topologySpreadConstraints`):** Provides fine-grained control over how Pods are spread across failure-domains like regions, zones, and nodes. This helps with high availability and resource utilization by preventing Pods from concentrating in a single domain.

## Common `kubectl` Commands

- **View Node labels:**

  ```bash
  kubectl get nodes --show-labels
  ```

  _Explanation: Displays all nodes in the cluster along with their labels, which can be used for scheduling._

- **Add a label to a Node:**

  ```bash
  kubectl label node <node-name> <label-key>=<label-value>
  ```

  _Explanation: Attaches a new label to a node, making it targetable by selectors and affinity rules._

- **Add a taint to a Node:**

  ```bash
  kubectl taint nodes <node-name> <key>=<value>:<effect>
  ```

  _Explanation: Applies a taint to a node, which will repel pods that do not have a matching toleration._

- **Describe a Node to see its taints:**

  ```bash
  kubectl describe node <node-name> | grep Taints
  ```

  _Explanation: Shows the taints currently applied to a specific node._

- **Describe a Pod to see scheduling events:**
  ```bash
  kubectl describe pod <pod-name> -n <namespace>
  ```
  _Explanation: Provides details about a Pod, including events that show why it was scheduled on a particular node or why it might be pending._

## Example Manifests

### Pod with `nodeAffinity`

```yaml
# pod-with-node-affinity.yaml
apiVersion: v1
kind: Pod
metadata:
  name: with-node-affinity
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
          - matchExpressions:
              - key: disktype
                operator: In
                values:
                  - ssd
  containers:
    - name: nginx
      image: nginx
```

### Pod with `podAntiAffinity` and `tolerations`

```yaml
# pod-with-anti-affinity-and-toleration.yaml
apiVersion: v1
kind: Pod
metadata:
  name: with-pod-anti-affinity
spec:
  affinity:
    podAntiAffinity:
      preferredDuringSchedulingIgnoredDuringExecution:
        - weight: 100
          podAffinityTerm:
            labelSelector:
              matchExpressions:
                - key: app
                  operator: In
                  values:
                    - web-store
            topologyKey: kubernetes.io/hostname
  tolerations:
    - key: "app"
      operator: "Equal"
      value: "critical"
      effect: "NoSchedule"
  containers:
    - name: with-pod-affinity
      image: k8s.gcr.io/pause
```
