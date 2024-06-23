

https://www.spinkube.dev/docs/spin-operator/quickstart/

```sh
k3d registry create myregistry.localhost --port 12345
k3d cluster create wasm-cluster \
  --image ghcr.io/spinkube/containerd-shim-spin/k3d:v0.14.1 \
  --port "8081:80@loadbalancer" \
  --agents 2 \
  --registry-use k3d-myregistry.localhost:12345

```


```sh
spin registry push -k localhost:12345/hello-spin:latest
spin kube deploy --from k3d-myregistry.localhost:12345/hello-spin:latest
```