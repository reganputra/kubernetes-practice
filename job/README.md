# Kubernetes Job

## What it is

A Job is a Kubernetes resource that creates one or more Pods and ensures that a specified number of them successfully terminate. Unlike a Deployment or ReplicaSet, which are designed for continuous, long-running applications, a Job is meant for short-lived, task-oriented workloads that run to completion. Once the desired number of successful completions is reached, the Job is complete.

## Key Characteristics

- **Run to Completion:** Designed for tasks that have a clear start and end.
- **Completion Tracking:** A Job tracks the successful completions of its Pods.
- **Parallelism:** Can be configured to run multiple Pods in parallel (`parallelism`) to speed up a task.
- **Retry Mechanism:** If a Pod fails, the Job can be configured to retry the task by creating a new Pod (`backoffLimit`).
- **TTL for Finished Jobs:** Jobs can be automatically cleaned up after they finish using a TTL mechanism (`ttlSecondsAfterFinished`).

## Common `kubectl` Commands

- **Create a Job from a manifest file:**

  ```bash
  kubectl apply -f manifest.yaml
  ```

  _Explanation: Creates or updates a Job using a YAML or JSON manifest file._

- **Get a list of Jobs:**

  ```bash
  kubectl get jobs -n <namespace>
  ```

  _Explanation: Lists all Jobs in the specified namespace._

- **Describe a Job to see its status:**

  ```bash
  kubectl describe job <job-name> -n <namespace>
  ```

  _Explanation: Shows detailed information about a Job, including its parallelism, completions, and events._

- **Get the Pods created by a Job:**

  ```bash
  kubectl get pods --selector=job-name=<job-name> -n <namespace>
  ```

  _Explanation: Lists the Pods that were created by the specified Job._

- **View logs from a Job's Pod:**

  ```bash
  # First, get the pod name
  POD_NAME=$(kubectl get pods --selector=job-name=<job-name> -n <namespace> -o jsonpath='{.items[0].metadata.name}')
  # Then, view its logs
  kubectl logs $POD_NAME -n <namespace>
  ```

  _Explanation: Fetches the logs from one of the Pods managed by the Job._

- **Delete a Job:**
  ```bash
  kubectl delete job <job-name> -n <namespace>
  ```
  _Explanation: Deletes the specified Job. By default, this also deletes the pods it created._

## Example Manifest

Here is a basic example of a Job that calculates pi to 2000 places and prints it:

```yaml
# job-example.yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: pi-calculation-job
spec:
  template:
    spec:
      containers:
        - name: pi
          image: perl:5.34.0
          command: ["perl", "-Mbignum=bpi", "-wle", "print bpi(2000)"]
      # The restart policy for a Job must be OnFailure or Never.
      # It cannot be Always.
      restartPolicy: OnFailure
  # Optional: Number of times to retry the Job if a pod fails
  backoffLimit: 4
  # Optional: Automatically clean up the Job 100 seconds after it finishes
  ttlSecondsAfterFinished: 100
```
