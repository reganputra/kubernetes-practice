# Kubernetes Pod Connectivity

## What it is

Kubernetes has a distinct networking model that enables clean and straightforward communication between Pods. In this model, every Pod gets its own unique IP address, and all containers within a Pod share that same IP address and network namespace. This means containers within a Pod can communicate with each other using `localhost`. More importantly, it creates a flat, cluster-wide network where every Pod can communicate with every other Pod directly, without needing Network Address Translation (NAT).

## Key Concepts

- **Pod IP Address:** Every Pod is assigned a unique IP address from the cluster's Pod CIDR range. This IP is routable within the cluster network.

- **Container Network Interface (CNI):** Kubernetes itself does not implement the networking layer. Instead, it relies on third-party CNI plugins to configure the network. Popular CNI plugins include Calico, Flannel, Cilium, and Weave Net. These plugins are responsible for assigning IP addresses to Pods and enabling traffic flow between them.

- **Communication Types:**

  - **Intra-Pod:** Containers within the same Pod share a network namespace and can communicate over `localhost`.
  - **Inter-Pod (Same Node):** Pods on the same node can communicate directly via a virtual ethernet bridge on the node.
  - **Inter-Pod (Different Nodes):** Pods on different nodes communicate via the cluster's network fabric, which is managed by the CNI plugin. This typically involves an overlay network or direct routing.

- **Service Discovery: The Role of Services**
  While Pods can communicate directly using their IP addresses, these IPs are ephemeral and change when a Pod is recreated. For stable communication, Kubernetes uses **Services**.

  - **Service:** A Kubernetes Service provides a stable, virtual IP address (called the `ClusterIP`) and a DNS name for a set of Pods. When traffic is sent to the Service, it load-balances the requests among the healthy Pods that match its label selector.
  - **DNS:** Kubernetes provides a built-in DNS service (usually CoreDNS) that creates DNS records for Services. A Pod can resolve a Service by its name (e.g., `my-service.my-namespace.svc.cluster.local`), and the DNS server will return the Service's `ClusterIP`. This is the primary way applications discover and connect to each other within the cluster.

- **NetworkPolicy:** A NetworkPolicy object allows you to define firewall rules for Pods. You can specify which Pods are allowed to connect to other Pods, effectively segmenting the network and securing your applications.

## Common `kubectl` Commands for Debugging

- **Get Pod IP addresses:**

  ```bash
  kubectl get pods -o wide -n <namespace>
  ```

  _Explanation: Shows the IP address of each Pod and the node it's running on._

- **Run a temporary debug container:**

  ```bash
  kubectl run -it --rm --image=busybox:1.28 dns-test -- /bin/sh
  ```

  _Explanation: Starts a temporary container with shell access, which you can use to test network connectivity from within the cluster._

- **Test DNS resolution from a Pod:**

  ```bash
  # Inside a debug container or an existing pod
  nslookup <service-name>
  ```

  _Explanation: Checks if the DNS name for a Service resolves to its ClusterIP._

- **Test connectivity to a Service:**
  ```bash
  # Inside a debug container or an existing pod
  wget -O- <service-name>:<port>
  ```
  _Explanation: Attempts to connect to a Service on a specific port._

## Example: Client-Server Communication

This example shows a backend `Deployment` exposed by a `Service`, and a `Pod` acting as a client that connects to it.

### 1. Backend Deployment and Service

```yaml
# backend-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend-app
spec:
  replicas: 2
  selector:
    matchLabels:
      app: backend
  template:
    metadata:
      labels:
        app: backend
    spec:
      containers:
        - name: echo-server
          image: ealen/echo-server
          ports:
            - containerPort: 80
---
# backend-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: backend-service
spec:
  selector:
    app: backend # This must match the labels on the backend pods
  ports:
    - protocol: TCP
      port: 80 # The port the service will be available on
      targetPort: 80 # The port the container is listening on
```

### 2. Client Pod

This Pod will connect to the `backend-service` by its DNS name.

```yaml
# client-pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: client-pod
spec:
  containers:
    - name: client
      image: busybox:1.28
      # Keep the pod running
      command: ["sh", "-c", "while true; do sleep 3600; done"]
```

You can then `exec` into the `client-pod` and run `wget -O- http://backend-service` to see the response from the `echo-server`.
