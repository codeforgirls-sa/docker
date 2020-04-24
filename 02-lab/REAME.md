# .dockerignore first step into optimizing docker image size

.dockerignore is like .gitignore

```bash
# 1. Build docker image
docker build -t my-awasom-app:v1 .

# 2. Copy .dockerignore
cp 02-lab/.dockerignore . 

# 3. Build docker image again with a new tag
docker build -t my-awasom-app:v2 .

# 3. List images
docker image ls

# 4. Do you see the difference? 😎

```

> _*Pro Tip*_: Make sure to always tag your images otherwise it will be tagged as latest which means it will be overwritten each time you build the image.

> NODE: .dockerignore is not only for optimization it's for security two you don't want sensitive files/directories to be copied over into your image, like .git directory, specially if you are going to publish your image publicly.