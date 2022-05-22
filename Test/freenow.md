# Platform Developer Test

The team SRE Platform has created a Kubernetes cluster and the first few services have been deployed to it. Now it is
time to expose some information on the current state of each service.

## Tasks

For any of the following tasks please use git to track the status and progress of your development and include the git repository (.git/) in the upload.

You can choose whatever programming language you feel most comfortable with. Please include instructions on how to build and run the
application.

You will need the following:

- [`kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- A Kubernetes Cluster, created e.g. via [kind](https://kind.sigs.k8s.io/), [minikube](https://kubernetes.io/docs/tasks/tools/#minikube) or [Docker for Mac](https://docs.docker.com/docker-for-mac/kubernetes/).

Please apply the services to your local Kubernetes cluster by executing `kubectl apply -f ./services.yaml`.

### 1. Expose information on all pods in the cluster

Add an endpoint to the service that exposes the number of pods running in the cluster in namespace `default` per service
and per application group:

```
GET `/services`
[
  {
    "name": "<service>",
    "applicationGroup": "alpha",
    "runningPodsCount": 2
  },
  {
    "name": "<service>",
    "applicationGroup": "beta",
    "runningPodsCount": 1
  },
  ...
]
```

### 2. Expose information on a group of applications in the cluster

Create an endpoint in your service that exposes the pods in the cluster in namespace `default` that are part of the same `applicationGroup`:

```
GET `/services/{applicationGroup}`
[
  {
    "name": "<service>",
    "applicationGroup": "<applicationGroup>",
    "runningPodsCount": 1
  },
  ...
]
```
