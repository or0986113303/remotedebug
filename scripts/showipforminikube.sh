export SERVICE_IP=$(kubectl get svc --namespace default remotedebug-basicapp --template "{{ range (index .status.loadBalancer.ingress 0) }}{{.}}{{ end }}")
echo http://$SERVICE_IP:5000
echo "test to make http request"
curl http://$SERVICE_IP:5000/api/v1/test/message