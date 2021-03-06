# remotedebug in container

## IDE
This project requires [vscode](https://code.visualstudio.com/) to run.

## Requirements
| Plugin | Reference |
| ------ | --------- |
| Remote Development | [Ref](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) |
| Docker | [Ref](https://www.docker.com/) |
| Docker-compose | [Ref](https://docs.docker.com/compose/) |
| Go Test Explorer | [Ref](https://marketplace.visualstudio.com/items?itemName=ethan-reesor.vscode-go-test-adapter) |
| PlantUML | [Ref](https://plantuml.com/en/) |
| goreleaser | [Ref](https://goreleaser.com/customization/build/) |

## Useage - remote debug
Please modify the ```"host": "127.0.0.1"``` to which IP you want to deploy to at ```/.vscode/launch.json``` .
And toggle the Debugging to choose ```Attach remote debug mode``` to start it.

## Useage -- release

According to the default setting for goreleaser as below :
GoReleaser will use the latest Git tag of your repository.

```sh
git tag -a $VERSION -m "Release message"
git push origin $VERSION
goreleaser release
```

## Useage -- minikube for Loadbalancer type of helm chart

Please type the below commands to toggle minukube connect to LoadBalancer services

```sh
sh ./script/installhelm.sh &> /dev/null 2>&1 &
```

And then waitting termial feedback done

```sh
sh ./script/showipforminikube.sh
```

If you want to uninstall this helm chart, type below commands on terminal localhost

```sh
sh ./script/uninstallhelm.sh
```

