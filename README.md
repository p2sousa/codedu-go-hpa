### Modulo DevOps - Kubernets (k8s) - HPA e CI\CD

### Cluster do servidor web GO
Endpoint do servidor web em GO com processo de loop antes de renderizar:

- **[servidor web GO](http://34.68.239.24/)**

### Imagem GoLang Servidor Web (LOOP)

 - **[p2sousa/go-hpa](https://hub.docker.com/r/p2sousa/go-hpa)**

 Para executar o processo entre na pasta k8s e rode os seguintes comandos:

```
  $ kubectl apply -f deployment.yml
  $ kubectl apply -f service.yml
  $ kubectl apply -f hpa.yml
```
Para acessar o servidor em loop:

```
  $ kubectl run -it loader --image=busybox /bin/sh
  $ while true; do wget -q -O- http://go-hpa.default.svc.cluster.local; done;
```

Para verificar o processo de Autosclaler rode:

```
  $ watch kubectl get hpa
```

Caso nao tenha o watch instalado, no macOS rode:

```
  $ brew install watch
```



