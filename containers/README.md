# containers
## Run presentation: (Linux)


```
  # git clone https://github.com/ggerman/pandora.git .
  # cd containers
  # sudo docker build  -t name/presentation .
  # docker run -p 3999:3999 name/presentation
```
Browser: http://172.17.42.1:3999/presentation.slide  

## Interesting

## Ignore entrypont in docker

  ```docker run -ti --entrypoint=/bin/bash rubylit/presentation -s```

## Remove all images running

  ```# docker rm $(docker ps -a -q)```
 
# Docker

![first slide](https://raw.githubusercontent.com/ggerman/pandora/master/containers/images/Screenshot%20-%2008282015%20-%2011%3A49%3A55%20AM.png)

* ¿Que son los containers?

* Tecnologías
 - Docker
 - Docker machine
 - Compose

  MANUEL: Vagrant
  MANUEL: puppet
  MANUEL: chef

* Licencias

* Boot2docker (multiplataformas)

* Mi primer container (Packman)
 - Dockerfile
 - build
 - images
 - rmi / rmi -f
 - run (X server, ports, permisos)

* Aislamiento (Isolation)
  - Creemos un proyecto
  - Database container
  - Application container
    - commit
    - diff

* Docker hub
  - push 
  - pull
