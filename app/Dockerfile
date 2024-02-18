FROM node:19-slim
LABEL authors="Federico Guerinoni<guerinoni.federico@gmail.com>"

RUN apt-get update && apt-get install -y \
    git \
    && rm -rf /var/lib/apt/lists/*

RUN npm install -g next
