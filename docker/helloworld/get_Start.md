## Sample application

```dockerfile
# syntax=docker/dockerfile:1
FROM node:12-alpine
RUN apk add --no-cache python g++ make
WORKDIR /app
COPY .. .
RUN yarn install --production
CMD ["node", "src/index.js"]
```
### build
eg:
```shell
docker build -t getting-started .
```
the -t flag tags our image. Think of this simply as a human-readable name for the final image. 
Since we named the image getting-started, we can refer to that image when we run a container.

The `CMD` directive specifies the default command to run when starting a container from this image.

`.` at the end of the docker build command tells that Docker should look for the Dockerfile in the current directory.
### run
```shell
 docker run -dp 3000:3000 getting-started
```
Remember the -d and -p flags? 
We’re running the new container in “detached”(-d) mode (in the background) and creating a mapping between 
the **host’s port 3000** to the **container’s port 3000**. (-p)
Without the port mapping, we wouldn’t be able to access the application.

## Update the application
先 remove 原来的 docker 再 

remove 的方式
1. `docker stop <the-container-id>` 然后 `docker rm <the-container-id>`
2. `docker rm -f <the-container-id>`

> You can stop and remove a container in a single command by adding the “force” flag to the docker rm command.

然后再运行更新的 docker ， `docker run -dp 3000:3000 getting-started`

## Share the application
#### Create a repo

#### Push the image
1. In the command line, try running the push command you see on Docker Hub. Note that your command will be using your namespace, not “docker”.
2. Login to the Docker Hub using the command `docker login -u YOUR-USER-NAME`.
3. Use the docker tag command to give the getting-started image a new name. Be sure to swap out YOUR-USER-NAME with your Docker ID.
```shell
$ docker tag getting-started YOUR-USER-NAME/getting-started
```
4. Now try your push command again. If you’re copying the value from Docker Hub, you can drop the tagname portion, 
   as we didn’t add a tag to the image name. If you don’t specify a tag, Docker will use a tag called `latest`.
```shell
$ docker push YOUR-USER-NAME/getting-started
```

#### Run the image on a new instance
login

```shell
$ docker run -dp 3000:3000 YOUR-USER-NAME/getting-started
```
You should see the image get pulled down and eventually start up!

## Persist the DB


