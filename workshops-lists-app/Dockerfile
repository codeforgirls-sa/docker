FROM node:current-slim

WORKDIR /app

COPY . .

RUN npm install

EXPOSE 8080

ENTRYPOINT ["npm", "start"]