# containers

![Flisol 2016 - Questions] (./images/CgnKOanWgAEbvtg.jpg)

Volt and Mongodb [Presentation example]

```
# docker pull ggerman/voltframework-docker
https://hub.docker.com/r/ggerman/voltframework-docker/~/dockerfile/
```

## 02 Sep 2015 update Dockerfile for new docker version:

```
  Client:
   Version:      1.9.0-dev
   API version:  1.21
   Go version:   go1.4.2
   Git commit:   3111aa7
   Built:        Tue Sep  1 20:35:25 UTC 2015
   OS/Arch:      linux/amd64

  Server:
   Version:      1.9.0-dev
   API version:  1.21
   Go version:   go1.4.2
   Git commit:   3111aa7
   Built:        Tue Sep  1 20:35:25 UTC 2015
   OS/Arch:      linux/amd64
```
Patch for skipping a failing test than rendered our docker image usel
https://gist.githubusercontent.com/xiam/f50f6dd6085f9a07ccfd/raw/5e0f472221f9c1556fe34788ff01724b63980337/docker_golang

## Run presentation: (Linux)


```
  # git clone https://github.com/ggerman/pandora.git .
  # cd containers
  # sudo docker build  -t name/presentation .
  # docker run -p 3999:3999 name/presentation
```
Browser: http://172.17.42.1:3999/presentation.slide  

Detached (-d) [RUNNING in background]

```
  # docker run -p 3999:3999 -d  rubylit/presentation:01
```

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

* Docker
  - ¿Que es docker?
  - Isolation
  - Usando

* Creando un container  
  - Instalar
  - search
  - run / build
    - ps -a -q / rm / images / rmi
  - hagamos correr clojure (lein)

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


![Agradecimentos](https://github.com/ggerman/pandora/blob/master/containers/thanks.png)
