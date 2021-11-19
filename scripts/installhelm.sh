#!/bin/bash
# helm install remotedebug ./deployments/helm/basicapp &> /dev/null 2>&1 &
# /usr/local/bin/minikube tunnel --cleanup &> /dev/null 2>&1 &
helm install remotedebug ./deployments/helm/basicapp
sleep 2
/usr/local/bin/minikube tunnel --cleanup
sleep 10
