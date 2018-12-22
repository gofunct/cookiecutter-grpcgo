# cookiecutter-grpcgo

Powered by [Cookiecutter](https://github.com/audreyr/cookiecutter), Cookiecutter Golang is a framework for jumpstarting production-ready go projects quickly.

## Features

- [x] Makefile
- [x] Dockerfile
- [ ] Docker-compose (prometheus, jaeger)
- [x] build time and git hash at build time
- [x] README.md
- [ ] Prototool config
- [x] Cobra cli
- [ ] Viper config
- [ ] Starter proto file
- [ ] Gokit http server
- [ ] Gokit grpc server
- [ ] Grpc gateway
- [ ] UI
- [ ] Kubernetes Deployment

## Usage

First, get Cookiecutter. Trust me, it's awesome:
```console
$ pip install cookiecutter
```

Alternatively, you can install `cookiecutter` with homebrew:
```console
$ brew install cookiecutter
```

Finally, to run it based on this template, type:
```console
$ cookiecutter https://github.com/gofunct/cookiecutter-grpcgo
```

## options
You will be asked about your basic info (name, project name, app name, etc.). This info will be used to customize your new project.

Warning: After this point, change all default options to your own information.

Answer the prompts with your own desired options.

```console
full_name [Coleman Word]: 
github_username [gofunct]: 
app_name [example]: 
project_short_description [example description]: 
docker_hub_username [gofunct]: 
docker_image [alpine]: 
docker_build_image [gofunct/cookiecutter-grpcgo]: 
Select docker_build_image_version:
1 - 1.11
2 - 1.10.3
3 - 1.9.7
Choose from 1, 2, 3 [1]: 
viper_config_name [config]: 
grpc_port [9090]: 
http_port [8080]: 
log_level [debug]: 
json_logs [y]: 
use_docker [y]: 
use_tls [y]: 
use_go_mod [y]: 
Select use_ci:
1 - travis
2 - circle
3 - none
Choose from 1, 2, 3 [1]: 2

```
