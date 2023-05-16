# pratica_kubernetes
brincando com kubernetes

Link da [imagem](https://hub.docker.com/r/gabrielronei/hello-go) 

Pra lembrar no futuro:

- Criar o cluster com kind:
`kind create cluster --config=k8s/kind.yaml --name=hellogo`

- Listar nodes
`kubectl get nodes`

- Para rodar
`kubectl cluster-info --context kind-hellogo`

- Criar um pod
`kubectl apply -f k8s/pod.yaml`

- Listar os pod
`kubectl get po`

- Apontando para porta do pod
`kubectl port-forward pod/goserver 8080:8082`
 
- Aplicando criação da replica
`kubectl apply -f k8s/replicaset.yaml`

`docker run --rm  -p8082:8082 gabrielronei/hello-go`
