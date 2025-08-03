# Kubernetes Service: ClusterIP

## What it is

A `ClusterIP` is the default and most common type of Kubernetes Service. It exposes a set of Pods on an internal, cluster-only IP address. This means the Service is only reachable from within the cluster, such as from other Pods. You cannot access a `ClusterIP` Service from outside the cluster without using a proxy or port-forwarding. Its primary purpose is to enable communication between different internal components of your application.

## Key Characteristics

- **Internal Exposure:** Provides a stable IP address and DNS name that is only accessible from within the cluster.
- **Default Service Type:** If you create a Service manifest without specifying a `type`, Kubernetes will create a `ClusterIP` Service.
- **Load Balancing:** Automatically load-balances network traffic across all the Pods that match its label selector.
- **Stable Endpoint:** The `ClusterIP` address is stable. Even if the Pods behind the service are created or destroyed, the IP address of the service remains the same, allowing for reliable service discovery.
- **Service Discovery:** Integrates with the cluster's internal DNS service (like CoreDNS), so other Pods can access it using a consistent DNS name (e.g., `my-service.my-namespace`).

## Common Use Cases

- **Backend Services:** Exposing a database, cache, or other backend microservice to a frontend application running in the same cluster.
- **Internal APIs:** Providing an internal API endpoint for other microservices to consume.
- **Service-to-Service Communication:** Facilitating communication between different layers or components of a multi-tiered application.

## Common `kubectl` Commands

- **Create a Service from a manifest file:**

  ```bash
  kubectl apply -f service-manifest.yaml
  ```

  _Explanation: Creates or updates a Service using a YAML manifest file._

- **Get a list of Services:**

  ```bash
  kubectl get services -n <namespace>
  ```

  _Explanation: Lists all Services in the specified namespace, showing their type and ClusterIP._

- **Describe a Service to see its details:**

  ```bash
  kubectl describe service <service-name> -n <namespace>
  ```

  _Explanation: Shows detailed information about a Service, including its labels, selector, IP address, and endpoints (the IP addresses of the Pods it's routing traffic to)._

- **Temporarily access a ClusterIP Service from your local machine:**

  ```bash
  kubectl port-forward service/<service-name> <local-port>:<service-port> -n <namespace>
  ```

  _Explanation: Forwards a local port on your machine to a port on the Service, allowing you to debug or interact with it directly._

- **Delete a Service:**
  ```bash
  kubectl delete service <service-name> -n <namespace>
  ```
  _Explanation: Deletes the specified Service. This does not delete the Pods that the Service was routing to._

## Example Manifest

Here is an example of a `ClusterIP` Service that exposes a backend application running in a set of Pods with the label `app: my-backend`.

```yaml
# backend-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: my-backend-service
spec:
  # The type is ClusterIP by default, but it can be set explicitly.
  type: ClusterIP
  selector:
    # This selector must match the labels of the Pods you want to expose.
    app: my-backend
  ports:
    - protocol: TCP
      # The port that the Service will be exposed on within the cluster.
      port: 80
      # The port on the Pods that the traffic will be forwarded to.
      targetPort: 8080
```
