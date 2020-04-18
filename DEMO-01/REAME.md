# Instructions

```bash
# 1. Copy docker file
cp DEMO-01/Dockerfile Dockerfile

# 2. Build docker image
docker build -t my-awasom-app .

# 3. Run instance container from the built image and name it my-awesome-app-container
docker run my-awasom-app --name my-awesome-app-container

# 4. Navigate to http://localhost:8080, nothing right? That's because we haven't mapped the container port to our local machine port. Let's stop the container to re-run it"
docker container stop my-awesome-app-container
# or simply `Ctrl+C`

# 5. Run the container the option -p to map local host port 8081 to the container port 8080
docker run my-awasom-app --name my-awesome-app-container -p 8081:8080

```

> _*Pro Tip*_: Add `--rm` to `docker run` to automatically delete the container after stopping with `Ctrl+C` or the command `docker container stop <CONTAINER>`