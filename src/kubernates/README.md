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
k run -it --rm --image=adisai123/daemonset:1 sh -- another way to run 
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

 
 # Affinity and anti-affinity:

The affinity and anti-affinity feature allows you to do complex scheduling than the nodeSelector and also works on Pods.

You can create rules that are not hard  requirements, but rather a preferred rule, meaning that the scheduler will still be able to schedule your pod, even if the rules cannot be met.

You can create rules that take other pod labels into account.
Affinity exist not only for node but pod as well.

Node affinity is similar to the nodeSelector.

Pod affinity/anti-affinity allows you to create rules how pods should be scheduled taking into account other running pods.

Affinity/anti-affinity mechanism is only relevant during scheduling, once a pod is running, it'll need to be  recreated to apply the rules again.

node affility  2 types:
1. requiredDuringSchedulingIgonredDuingExecution
2. preferredDuringSchedulingIgnoredDuringExecution

first one sets hard requirement like nodeSelector.
Secont type will try to enforece rule, but it will not guarentee it
	Even if the rule is not met, the pod can still be scheduled, its soft reuirement, a preference.
weighing is specified in preferredDuringSchedlingIgnoredDuringExection 
The heigher this weighting, the more weight is given to that rule.

If you have different rules with weight 1 and 5  so node will have total 6 
And if other node has only one weight 
The node that has the heigher total score, that's where the pod will be scheduled on.



# Operator:
An operator is a method of packaging, deploying, and managing a kubernates application.

It provides great way to deploy Stateful services on kubernates.(due to that lot of complexities will be hidden from the end-user)

There are operators for Prometheus, Vault, Rook (storage), Mysql, PostgreSQL, and so on.

If you use postgreSql operator, it'will allow you to also create replicas, initiate a failover, create backups, scale

ex: https://github.com/CrunchyData/postgres-operator



Toleration and Taints:
Toleration mean - opposite to affinity (pod can not be run on these nodes)

Taints mark a node, tolerations are applied to pods to influence the scheduling of the pods.

One use case for taints is to make sure that when you create a new pod, they're not scheduled on master.
	The master has a taint: (node-role.kubernetes.io/master:NoSchedule)

To add a new taint to a node, you can use kubectl taint:

```
kubectl taint nodes node1 key=value:NoSchedule
```

This will make sure that no pods will be scheduled on node1, as long as they don't have a matching toleration.

The following toleration would allow a new pod to be scheduled on the tainted node1:

```yaml
tolerations:
- key: "key"
  operator: "Equal"
  value: "value"
  effect: "NoSchedule"
```

Just like affinity, taints can also be a preference (or "soft") rather than a requirement:

1. NoSchedule: a hard requirement that a pod will not be scheduled unless there is a matching tolaration

2. PreferNoSchedule: Kubernates will try and avoid placing a pod that doesn't have a matching tolaration, but it's not a hard requirement.

If the taint is applied while there are already running	pods, these will not be evicted, unless the following taint type is used:
	NoExecute: evict pods with non matching tolerations

When specify NoExecute, you can specify within your toleration how long the pod can run on a tainted node before being evicted:
```yaml
tolerations:
- key: "key"
  operator: "Equal"
  value: "value"
  effect: "NoExecute"
  tolerationSeconds: 3600
```

If you dont specify the tolerationSeconds, the toleration will match and the pod will keep running on the node.

Taint can be applied on below keys:
```yaml
tolerations: 
-	key:	"node.alpha.kubernetes.io/unreachable"
	operator:	"Exists"
	effect:	"NoExecute"
	tolerationSeconds: 3600
```
similary:
node.kubernetes.io/not-ready
node.kubernetes.io/unreachable
node.kubernetes.io/out-of-disk
node.kubernetes.io/memory-pressure
node.kubernetes.io/disk-pressure
node.kubernetes.io/network-pressure



to remove taint use - sign after key:

```
kubectl taints node mynodename testkey-

```


# helm :

helm is the best way to find, share and use software built for kubernetes:

Helm is a packagemanager for kubernetes
Helm uses packaging format called charts

A charts a collection of files that describe a set of kubernetes resources.

A single chart can deploy an app, a piece of software, or a database

It can have dependencies.

Charts use templates that are typically developed by a package maintainer

They will generate yaml files that kubernetes understands

you can think of templates as dynamic yaml files, which can contain logic and variables

``` 
helm init ---- Install tiller on the cluster
helm reset --- Remove tiller from cluster
helm install --- install chart
helm search
helm list
helm upgrade
helm rollback
helm install . --generate-name
helm list
helm delete chart-1595698716
```


# kubeless:
Way to provide serverless in kubernetes 
Rather than using contaniers to start applications on kubernetes, you can use functions.
With this you can easily deploy function on kubernetes.
# Serverless:
	You dont need to manage distribution (windows/Linux) 
	You dont need to build container
	You only will have to pay for the time your function running.
	Ex: Azure Funtions , AWS Lambda, Google Cloud function.
With these products, you dont need to manage the underling infrastructure.

alternative to kubeless ; openFaas , Fission, OpenWhisk

Monoliths that dont have anything to do with each other.


# isto

# canary deployment ( 10% latest patch %90 old verison with the help of weight	)

# virtulService

# retry feature of isto

# skaffold - for builing cicd , code changes monitoring tool

# raft algo -> to select leader
	suppose 5 out of 1 need to select
	all 5 will wait for random time
	one which finished its time will send msg for to be get elected to be as a leader as it is the first one to finish wait and sending heart bit and in that we it will come to know that leader is still active.


	

[admalpan bundles]$time for ((i=0;i<10;i++)); do sleep 1; echo "$i";done
0
1
2
3
4
5
6
7
8
9

real	0m10.059s
user	0m0.008s
sys	0m0.019s
[admalpan bundles]$

