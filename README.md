# cookiecutter-grpcgo

Powered by [Cookiecutter](https://github.com/audreyr/cookiecutter), Cookiecutter Golang is a framework for jumpstarting production-ready go projects quickly.

## Features

- [x] Makefile
- [x] Dockerfile
- [ ] Docker-compose (prometheus, jaeger)
- [x] build time and git hash at build time
- [x] .gitignore
- [x] README.md
- [ ] Prototool config
- [x] Cobra cli
- [x] Viper config
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

You will be asked about your basic info (name, project name, app name, etc.). This info will be used to customize your new project.

Warning: After this point, change all default options to your own information.

Answer the prompts with your own desired [options]().

Enter the project and take a look around:
```console
$ cd echoserver/
$ ls
```

Run `make help` to see the available management commands, or just run `make build` to build your project.
```console
$ make help
$ make build
$ ./bin/echoserver
```
