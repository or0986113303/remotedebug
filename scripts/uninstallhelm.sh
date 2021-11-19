#!/bin/bash
# helm uninstall remotedebug
kill $(ps -ef | grep "/usr/local/bin/minikube tunnel" | awk '{print $2}')
helm uninstall remotedebug
# echo $TMP