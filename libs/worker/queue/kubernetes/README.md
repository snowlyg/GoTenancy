# Kubernetes Backend for QOR Worker

This package provides [QOR worker](http://GoTenancy/libs/worker)'s backend based on [Kubernetes Job](https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/)

## Basic Usage

In the basic usage mode, we will collect currently running pod's information like containers, volumes to create a new job pod to complete your requests.

```go
kubernetesBackend, err := kubernetes.New(&kubernetes.Config{})

Worker := worker.New(&worker.Config{
  Queue: kubernetesBackend,
})

Worker.RegisterJob(&worker.Job{
  Name: "Send Newsletter",
  Handler: func(argument interface{}, qorJob worker.QorJobInterface) error {
})
```

## Advanced Usage

For advanced requirements, like manage job's priority, cpu/memory's requests/limits, you could customize Job's template based on your requirements.

For example:

```go
func main() {
  kubernetesBackend, err := kubernetes.New(&kubernetes.Config{
    JobTemplateMaker: func(qorJob worker.QorJobInterface) string {
      // Job `process order` has higher priority than other jobs
      if qorJob.GetJobName() == "process_order" {
        return `
apiVersion: batch/v1
kind: Job
metadata:
  name: jobname
spec:
  template:
    spec:
      containers:
      - name: app
        image: my_image
      priorityClassName: high-priority
`
      }

      return `
apiVersion: batch/v1
kind: Job
metadata:
  name: jobname
spec:
  template:
    spec:
      containers:
      - name: app
        image: my_image
        resources:
          limits:
            cpu: "750m"
      priorityClassName: low-priority
`
    },
  })

  Worker := worker.New(&worker.Config{
    Queue: kubernetesBackend,
  })
}
```
