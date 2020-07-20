kubeadm - is a toolkit by kubernates to create a cluster 
			it works on linux os.

for laptop or single or dev env or single node env you can use minikube , docker-for-desktop, or  docker-desktop

kops - for aws to create a cluster.

vagrant or VirtualBox (for windows to run liunx )

```bash
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
``` 

How to push it to docker hub:

```bash
docker login
docker tag imageid adisai123/myfirstgo

docker push adisai123/myfirstgo
```

or imidiately tag an image during the build process.:
```bash
docker build -t adisai123/myfirstgo
docker push adisai123/myfirstgo
```

All the data in the container is not preserved, when a container stops, all the changes within a container are lost.
	You can preserve data, using volumnes.


create and start pod :kubectl create -f myfirstgo.yaml

```bash
kubectl get pod = Get information about all running pods
kubectl describe pod <podname> = describe one pod
kubectl expose pod <podname> --port=444 --name frontend     = expose port and create new service

kubectl port-forward <pod> 8080 =Port forward the exposed pod port to your local machine.
kubectl attach <podname> -i  =Attach to the pod
kubectl exec <pod> --command = Execute a command on the pod
```

scaling : ReplactionController
If your application is stateless you can horizontally scale it 
	1. Most web applications can be made staeless: 
	2. session management needs to be done outside the container
	3. on continer any files that need to be saved cant be saved locally on the container. 

scale :
```bash
kubectl scale --replicas=4 -f replicationC.yaml (do it only in case your app is stateless)

kubectl get rc

admalpan kubernates]$kubectl get rc
NAME           DESIRED   CURRENT   READY     AGE
myfirstgoapp   2         2         2         5d
```

To delete replca controller 
```bash
[admalpan kubernates]$k delete rc/myfirstgoapp
replicationcontroller "myfirstgoapp" deleted
```
Replica set is the next-generation Replication Controller. (used by deployment object)

# Deployment: Allow you to do app deployments and updates .

When using the deployment object, you define the state of your application.
kubernates then make sure that clusters matches your desire state.

With deployment you can Create , update new verison and do rolling updates.

rollback to prevision of your app.

```bash
kubectl get deployments           (show deployments)
kubectl get rs                    (replica set)
kubectl get pods --show-labels     (get pods, and also show labels attached to those pods)
kubectl rollout status deployment/myfirstgoapp      (get deployment status)
kubectl set image deployment/myfirstgoapp myapp=adisai123/myfirstgo:2   (change version)
kubectl edit deployment/myfirstgoapp         (edit deployment object)
kubectl rollout status deployment/myfirstgoapp (get status of changed object status)
kubectl rollout history deployment/myfirstgoapp (get rollout object history)
kubectl rollout undo deployment/myfitstgoapp (rollback to previous version)
kubectl rollout undo deployment/myfirstgoapp --to-revision=n (Rollback to any version)

k expose deployment myfirstgoapp --type=NodePort 
```

why pods should never be accessed directly:
	when using a replication controller , pods are terminated and created during scaling operation.
	When using deployments, when updating  the images version, pods are terminated and new pods get created.

	thats why create services:
Service is a direct bridge between the pod and other services or end-users.
Creating new service will create an endpoint for your pods, so that it could be accessible externally.
	A clusterip: virtual ip address only reachable from within the cluster.
	NodePort : to reach externally.

# Labels
Labels: key value pair attached to object.
Lables are not unique and multiple labels can be added to one oject.

you can label to tag nodes

```bash
addLabel$ kubectl label node docker-for-desktop environment=dev
removeLabel$ kubectl label node docker-for-desktop environment- 
```

once nodes are tagged, you can use label selectors to let pods only run on specific nodes.

There are 2 steps required to run a pod on a specific set of nodes:
	first you tag the node
	then you add a nodeselector to your configuration.


Two types for health check:
	1. Running a command in the container periodically
	2. Periodic check on a URL.

That way it insure availiblity and resiliency testing

# secrets
Secret provides a way in kubernates to distribute cerdentials, keys and passwords or "secret" data
to the pods

Secrets can be used in following ways:
	1. As a enviromnent variables
	2. Use Secrets as a file in pod
		1. This setup uses volumes to be mounted in a container.
		2. in this volume you have files
		3. Can be used for instance for dotenv files or your app just read this file.
	3. Use an exernal image to pull Secrets (from a private image registery)

to generate secrets using files:
```bash
echo -n "root" > ./username.txt
echo -n "password" > .password.txt

kubectl create secret generic db-user-pass --from-file=./username.txt  --from-file=./password.txt
```
db-user-pass is secret name

or it could be secret ssh key or an ssl cert

```bash
kubectl create secret generic ssl-certificate --from-file=ssh-privatekey=~/.ssh/id_rsa --ssl-cert=mysslcert.crt
```
password : echo -n "root" | base64


readynessProbes vs livenessProbes:

livenessProbes : whether a container is running

readynessProbes: whether a container is ready to serve request.


# Pod status:
	1.  Running
	2.	Pending : network slow (image is still downloading) or 
				pod cannot be scheduled becuase of resource constrains.

	3. succeeded: All the containers within this pod has been terminated. and will not be restarted.

	4. Failed: All containers within this pod have been Terminated, and at least one container returned a failure code.

# Pod lifecycle:
				initdelayin		 -------------------
                second          |  readinessProbes	|														
								|  livenessProbes	|
                                 -------------------
				post start hook 						per stop hook
---------------   ....................................................
init container |  |    				main container                   |
--------------    ....................................................

----------------------------------------------------------------------->
										time


pre and post can be defined in yaml file in pod specification to execute certain commands

service discovery:

inside the pod sh
$nslookup <service name>

```bash
[admalpan app-infra]$k get svc
NAME                TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)           AGE
kubernetes          ClusterIP   10.96.0.1       <none>        443/TCP           131d
myfirstgo           NodePort    10.107.209.78   <none>        8080:32193/TCP    19h
myfirstgo-service   NodePort    10.98.81.198    <none>        31001:31001/TCP   1d
[admalpan app-infra]$k exec -it myfirstgoapp-5989d5d746-grpl6 sh
/app # nslookup myfirstgo-service
Server:		10.96.0.10
Address:	10.96.0.10:53

Name:	myfirstgo-service.default.svc.cluster.local
Address: 10.98.81.198
```


# ConfigMap:
	Configuration parameters that are not secret, can be put in a ConfigMap.

To create configmap using files:
```bash
[admalpan configmap]$cat <<aditya>> myprop.properties 
> aditya=adiya
> sai=sai
> aditya=aditya
> aditya
[admalpan configmap]k create configmap app-config --from-file=myprop.properties
[admalpan configmap] curl http://localhost:30001/ -vvvv   (-vvvv - gives who is serving that request.)
```

# Ingress:
	allows inbound connections to the cluster.
It is an alternative to the exxternal Loadbalancer and NodePort.
You can write your own ingress controller (for creating Loadbalancer rule)


# Volume:
	volumnes in kubernetes allow you to store data outside the continer.

Persistent volumes in kubernates allow you attach a volume to a continer that will exists even when
the container stops.

# stateful Sets:
	on shutdown it wont shutdown , pod name will be fixed i.e. no random string 

# Daemon sets:
It ensures that every single node in the cluster runs the same resource
This is useful when you want to ensure that a certain pod is running on every single kubernates node.

When a node is added to cluster, a new pod will be started automatically.

Same when a node is removed, the pod will be not be rescheduled on another node.

Typically use cases:
	Logging aggregators , Monitoring, Load Balancers / Reverse Proxies/ API Gateways

// influx db : heapster grafana

//Autoscaling 
```bash
[admalpan podautoscle]$k get hpa
NAME          REFERENCE                 TARGETS        MINPODS   MAXPODS   REPLICAS   AGE
newautscale   Deployment/myfirstgoapp   <unknown>/2%   1         10        1          17s
[admalpan podautoscle]$
```

# Master component:
	APIs contains : 
		Scheduling actuator:
		Rest:
	Scheduler
	Controller manager (node controller)

	Every time you send create request to kubectl it using Rest API part  these objects are saved into etcd.
	Rest interface communicate with etcd.
	rest communicate with kubelet available in each node (slave node).

# Helm :
Helm is the package manager in kubernetes

#Resource quota:
You can manage resource with ResourceQuota and ObjectQuota objects.
	Each container can specify request capacity and capacity limit.
	Rquest capacity is an explcit request for resources.
	Scheduler can use the request capacity to make decision on where to put the pod on.

you can see it as a minimum amount of resources	the pod need.

Resource limit is a limit imposed to the container.

If a capacity quota (eg. mem/cpu) has been specified by the administrator, then each pod needs to specify capacity quota during creation.

The administrator can specify default request values  for pods that dont specify any values for capacity.

same is valid for limit below limits within a namespace:

requests.cpu  	the sum of cpu requests of all pods can not exceed this value.
requests.mem 	the sum of mem  requests of all pods can not exceed this value.
requests.storage the sum of storage requests of all Persistent volumn claims cannot exceed  this value.
limit.cpu
limit.memory

configmaps total number of configmaps that can exist in a namespace
persistentvolumnclaims
pods 		total number of pods that can exist in a namespace
replicationcontrollers total number of ReplactionController 
Resourcequotas
services    
services.loadbalancer
service.nodeports
secrets


# Namespaces :
 allow to create logical clusters within  the same physical cluster
 Logically seperates your clusters
 standard namespace is called default namespace where all resources are launched in by default.

 modify default namespace:
 export CONTEXT=$(k config view | awk '/current-context/ {print $2}')
 kubectl config set-context $CONTEXT --namespace=aditya


 admalpan ResourcelimitObject]$k get resourcequota
 NAME                 AGE
 computer-resources   4m
 object-counts        12s
 [admalpan ResourcelimitObject]$

 