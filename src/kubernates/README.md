kubeadm - is a toolkit by kubernates to create a cluster 
			it works on linux os.

for laptop or single or dev env or single node env you can use minikube , docker-for-desktop, or  docker-desktop

kops - for aws to create a cluster.

vagrant or VirtualBox (for windows to run liunx )
>cat Dockerfile
From golang:alpine 
RUN mkdir /app
COPY . /app
WORKDIR /app
go run build -o main .
CMD ["/app/main"]

>Docker build .
you will get an image id 

> Docker run -p 8080:8080 -it f8401af5a13d
 

How to push it to docker hub:

docker login
docker tag imageid adisai123/myfirstgo

docker push adisai123/myfirstgo

or imidiately tag an image during the build process.:

docker build -t adisai123/myfirstgo
docker push adisai123/myfirstgo

All the data in the container is not preserved, when a container stops, all the changes within a container are lost.
	You can preserve data, using volumnes.


create and start pod :k create -f myfirstgo.yaml

kubectl get pod = Get information about all running pods
kubectl describe pod <podname> = describe one pod
kubectl expose pod <podname> --port=444 --name frontend     = expose port and create new service

kubectl port-forward <pod> 8080 =Port forward the exposed pod port to your local machine.
kubectl attach <podname> -i  =Attach to the pod
kubectl exec <pod> --command = Execute a command on the pod



