# pratica_kubernetes
<img src="https://matthewpalmer.net/kubernetes-app-developer/articles/networking-overview.png"/>

### O que é um cluster?
Um cluster no Kubernetes é um grupo de maquinas que trabalham em conjunto para executar e gerenciar aplicativos em containers. O cluster inclui um nó mestre que controla o cluster e varios nós de trabalho que executam os aplicativos
 
### O que é um Nó(Node)?
Basicamente um node é uma maquina dentro do cluster que executa os pods e fornece os recursos necessarios para as cargas de trabalho. 

### O que é um Pod?
É uma unidade simples execução e implantação de containers, agrupando um ou mais containers em um unico nó. Podem ser gerenciados, escalados e removido de forma. Basicamente, pods são grupos de containers relacionados.

## Como funciona o deployment no kubernetes?
Basicamente é Deployment > Replicaset > Pods, o deployment cria um replicaset e ele cria os pods, se você alterar o deployment ele vai repetir esse processo e atualizar a versão automaticamente.

## brincando com kubernetes

Link da [imagem no dockerhub](https://hub.docker.com/r/gabrielronei/hello-go) 

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

- Deletar pods
`kubectl delete pod goserver`
 
- Aplicando criação da replica
`kubectl apply -f k8s/replicaset.yaml`

- Listar as replicas
`kubectl get replicasets`

- Aplicando criação do deployment
`kubectl apply -f k8s/deployment.yaml`

- Ver o que está acontecendo no deployment
`kubectl describe deployment goserver`

- Ver tudo que mudou
`kubectl rollout history deployment goserver`

- Voltar para versão anterior
`kubectl rollout undo deployment goserver`

- Voltar para versão especifica 
`kubectl rollout undo deployment goserver --to-revision=versao_especifica`

`docker run --rm  -p8082:8082 gabrielronei/hello-go`
