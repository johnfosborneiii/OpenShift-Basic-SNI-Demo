# hello-sni

Demonstrates a TLS tcp client connecting through the OpenShift router to a TLS tcp server.

1.  Build the docker image with `build.sh`
1.  Start OpenShift
1.  Install the router
1.  Create a new app from the template
1.  Run the client and pass in route as a parameter(ie, ***go run client.go sni.service.com***)

If you're not using the single machine vagrant environment you may need to adjust the ip address in the client.  The client
points to 10.0.2.15 which is the default ip for the vagrant machine and is also the entry point for the router since the
router binds to host ports.

```
[vagrant@openshiftdev paul_temp]$ oc new-app -f snidemo-template.json 
[vagrant@openshiftdev paul_temp]$ oc get pods && oc get services && oc get routes
POD                 IP                  CONTAINER(S)                   IMAGE(S)                          HOST                           LABELS              STATUS
router              172.17.0.2          origin-haproxy-router-router   openshift/origin-haproxy-router   openshiftdev.local/127.0.0.1   <none>              Running
tls-server          172.17.0.4          tls-server                     pweil/tls-server                  openshiftdev.local/127.0.0.1   name=tls-server     Running
NAME                LABELS                                    SELECTOR            IP                  PORT
kubernetes          component=apiserver,provider=kubernetes   <none>              172.30.17.2         443
kubernetes-ro       component=apiserver,provider=kubernetes   <none>              172.30.17.1         80
tls-service         <none>                                    name=tls-server     172.30.17.20        2999
NAME                HOST/PORT           PATH                SERVICE             LABELS
route-passthrough   my-tls-server                           tls-service   

[vagrant@openshiftdev paul_temp]$ go run client.go tls-service.route.ocpmaster.com
Hello TLS
[vagrant@openshiftdev paul_temp]$
```
