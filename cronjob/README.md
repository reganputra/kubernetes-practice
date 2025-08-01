# Kubernetes CronJob

## What it is

A CronJob is a Kubernetes resource that creates Jobs on a repeating schedule. It is used for running periodic and recurring tasks, such as backups, report generation, or automated tests. A CronJob is like a classic cron (time-based job scheduler) on a Linux or UNIX system, but it runs containers instead of scripts.

## Key Characteristics

- **Scheduled Execution:** Runs Jobs based on a defined cron schedule.
- **Job Management:** Automatically creates and manages the lifecycle of Jobs.
- **Concurrency Policy:** Defines how to handle concurrent executions of a Job if the previous one is still running. Options include `Allow`, `Forbid`, and `Replace`.
- **History Limits:** You can configure how many successful and failed Job histories to retain, which helps in managing resource usage.
- **Deadline:** A starting deadline (`startingDeadlineSeconds`) can be set to prevent a Job from starting if it has missed its schedule for too long.

## Common `kubectl` Commands

- **Create a CronJob from a manifest file:**

  ```bash
  kubectl apply -f manifest.yaml
  ```

  _Explanation: Creates or updates a CronJob using a YAML or JSON manifest file._

- **Get a list of CronJobs:**

  ```bash
  kubectl get cronjobs -n <namespace>
  ```

  _Explanation: Lists all CronJobs in the specified namespace._

- **Describe a CronJob to see its details:**

  ```bash
  kubectl describe cronjob <cronjob-name> -n <namespace>
  ```

  _Explanation: Shows detailed information about a specific CronJob, including its schedule, last run, and created jobs._

- **Get the Jobs created by a CronJob:**

  ```bash
  kubectl get jobs --watch -l "job-name in (<cronjob-name>)"
  ```

  _Explanation: Lists the Jobs created by the specified CronJob. You can often find the right label with `kubectl describe cronjob`._

- **Manually trigger a Job from a CronJob:**

  ```bash
  kubectl create job --from=cronjob/<cronjob-name> <manual-job-name>
  ```

  _Explanation: Creates a new Job from the CronJob's template, independent of its schedule._

- **Suspend/Resume a CronJob:**

  ```bash
  # Suspend
  kubectl patch cronjob <cronjob-name> -p '{"spec" : {"suspend" : true}}'

  # Resume
  kubectl patch cronjob <cronjob-name> -p '{"spec" : {"suspend" : false}}'
  ```

  _Explanation: Toggles the `suspend` field to stop or resume the scheduling of new jobs._

- **Delete a CronJob:**
  ```bash
  kubectl delete cronjob <cronjob-name> -n <namespace>
  ```
  _Explanation: Deletes the specified CronJob. By default, this also deletes the jobs it created._

## Example Manifest

Here is a basic example of a CronJob manifest that runs a simple command every minute:

```yaml
# cronjob-example.yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: hello-cronjob
spec:
  # Cron schedule format: minute hour day-of-month month day-of-week
  # This schedule runs every minute.
  schedule: "*/1 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
            - name: hello
              image: busybox
              args:
                - /bin/sh
                - -c
                - date; echo "Hello from the Kubernetes CronJob"
          restartPolicy: OnFailure
  # Optional: Keep the last 3 successful and 1 failed job history
  successfulJobsHistoryLimit: 3
  failedJobsHistoryLimit: 1
```
