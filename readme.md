# Core Kubernetes Concepts

## Kubernetes Objects

In Kubernetes, an "object" is a persistent entity in the Kubernetes system. Kubernetes uses these entities to represent the state of your cluster. Specifically, they can describe:

- What containerized applications are running (and on which nodes)
- The resources available to those applications
- The policies around how those applications behave, such as restart policies, upgrades, and fault-tolerance

Once you create an object, the Kubernetes system will constantly work to ensure that object exists. By creating an object, you're effectively telling the Kubernetes system what you want your cluster's workload to look like; this is your cluster's **desired state**.

The most common Kubernetes objects include:

- **Pod:** The smallest and simplest unit in the Kubernetes object model that you create or deploy.
- **Service:** An abstract way to expose an application running on a set of Pods as a network service.
- **Volume:** A directory containing data, accessible to the containers in a Pod.
- **Namespace:** A way to divide cluster resources between multiple users.

## Kubernetes Workloads

A "workload" is an application running on Kubernetes. Whether your workload is a single component or several that work together, on Kubernetes you run it inside a set of **Pods**. A workload resource is a type of Kubernetes object that you use to manage a set of Pods.

Workload resources configure controllers that make sure the right number of the right kind of Pod are running to match the state you've specified. Kubernetes provides several built-in workload resources:

- **Deployment** and **ReplicaSet** (replacing the legacy ReplicationController). Deployment is a good fit for managing a stateless application workload on your cluster.
- **StatefulSet** lets you run one or more related Pods that do track state.
- **DaemonSet** defines Pods that provide node-local facilities.
- **Job** and **CronJob** define tasks that run to completion and then stop.
