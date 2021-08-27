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
#### The container’s filesystem

When a container runs, it uses the various layers from an image for its filesystem. 
Each container also gets its own “scratch space” to create/update/remove files. 
Any changes won’t be seen in another container, even if they are using the same image.


#### Container volumes
With the previous experiment, we saw that each container starts from the image definition each time it starts. While containers can create, update, and delete files, those changes are lost when the container is removed and all changes are isolated to that container. With volumes, we can change all of this.

Volumes provide the ability to connect specific filesystem paths of the container back to the host machine. If a directory in the container is mounted, changes in that directory are also seen on the host machine. If we mount that same directory across container restarts, we’d see the same files.

There are two main types of volumes. We will eventually use both, but we will start with named volumes.

#### Persist the todo data

With the database being a single file, if we can persist that file on the host and make it available to the next container, it should be able to pick up where the last one left off. By creating a volume and attaching (often called “mounting”) it to the directory the data is stored in, we can persist the data. As our container writes to the todo.db file, it will be persisted to the host in the volume.

As mentioned, we are going to use a named volume. Think of a named volume as simply a bucket of data. Docker maintains the physical location on the disk and you only need to remember the name of the volume. Every time you use the volume, Docker will make sure the correct data is provided.

#### Dive into the volume
A lot of people frequently ask “Where is Docker actually storing my data when I use a named volume?” If you want to know, you can use the docker volume inspect command.
```shell
$ docker volume inspect todo-db
```

## Use bind mounts


