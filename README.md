# pratica_kubernetes
<img src="https://matthewpalmer.net/kubernetes-app-developer/articles/networking-overview.png"/>

### O que é um Node?
Basicamente um node é uma maquina dentro do cluster que executa os pods e fornece os recursos necessarios para as cargas de trabalho. 

### O que é um Pod?
É uma unidade simples execução e implantação de containers, agrupando um ou mais containers em um unico nó. Podem ser gerenciados, escalados e removido de forma. Basicamente, pods são grupos de containers relacionados.

## brincando com kubernetes

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
