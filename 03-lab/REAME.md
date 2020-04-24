# Multistage Dockerfile

```bash
# 1. Copy Dockerfile
cp 03-lab/Dockerfile .

# 2. Build docker image again with a new tag
docker build -t my-awasom-app:v3 .

# 3. List images
docker image ls

# 4. Do you see the difference in the image size? 👸🏻

```

> Did you notice how we also reordered layers  to utilize layer caching.
