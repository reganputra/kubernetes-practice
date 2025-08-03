# Kubernetes Ingress

## What it is

An Ingress is a Kubernetes API object that manages external access to the services in a cluster, typically HTTP and HTTPS. Ingress can provide load balancing, SSL termination, and name-based virtual hosting. It acts as a smart router or an entry point to your cluster, allowing you to define rules that control how external traffic is routed to internal services.

## How it Works: The Ingress Controller

An Ingress resource on its own doesn't do anything. You need an **Ingress controller** running in your cluster to fulfill the Ingress rules. The Ingress controller is a pod that watches the Kubernetes API for Ingress resources and configures a load balancer (like NGINX, HAProxy, or a cloud provider's load balancer) accordingly.

Different Ingress controllers (e.g., NGINX, Traefik, HAProxy, GKE Ingress) have different capabilities and configuration options. You must have an Ingress controller running in your cluster for your Ingress resources to work.

## Key Characteristics

- **Layer 7 Routing:** Operates at the application layer (HTTP/HTTPS) of the network stack.
- **Host-Based Routing:** Can route traffic to different services based on the requested hostname (e.g., `foo.example.com` goes to the `foo-service`, `bar.example.com` goes to the `bar-service`).
- **Path-Based Routing:** Can route traffic to different services based on the request path (e.g., `example.com/api` goes to the `api-service`, `example.com/ui` goes to the `ui-service`).
- **SSL/TLS Termination:** Can terminate SSL/TLS connections at the Ingress point, offloading the encryption/decryption work from your application pods.
- **Single Entry Point:** Provides a single, stable external IP address for multiple services, which is more efficient and cost-effective than creating a `LoadBalancer` service for every service you want to expose.

## Common `kubectl` Commands

- **Create an Ingress from a manifest file:**

  ```bash
  kubectl apply -f ingress-manifest.yaml
  ```

  _Explanation: Creates or updates an Ingress resource using a YAML manifest file._

- **Get a list of Ingresses:**

  ```bash
  kubectl get ingress -n <namespace>
  ```

  _Explanation: Lists all Ingress resources in the specified namespace, showing their hosts, address, and ports._

- **Describe an Ingress to see its details:**

  ```bash
  kubectl describe ingress <ingress-name> -n <namespace>
  ```

  _Explanation: Shows detailed information about an Ingress, including its rules, default backend, and associated events._

- **Delete an Ingress:**
  ```bash
  kubectl delete ingress <ingress-name> -n <namespace>
  ```
  _Explanation: Deletes the specified Ingress resource. The Ingress controller will de-provision the routing rules._

## Example Manifest

Here is an example of an Ingress that routes traffic based on the hostname and path. It assumes you have an Ingress controller running and two services named `service-one` and `service-two`.

```yaml
# example-ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: example-ingress
  annotations:
    # Annotations are often used to configure the Ingress controller
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  # IngressClass specifies which Ingress controller should handle this Ingress
  ingressClassName: nginx-example
  rules:
    - host: "foo.example.com"
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: service-one
                port:
                  number: 80
    - host: "bar.example.com"
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: service-two
                port:
                  number: 80
  # TLS configuration for HTTPS
  tls:
    - hosts:
        - foo.example.com
        - bar.example.com
      secretName: my-tls-secret # A secret containing the TLS certificate and key
```
