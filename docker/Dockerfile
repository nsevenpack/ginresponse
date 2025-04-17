# changer la version de golang si besoin 
FROM golang:1.24-alpine 
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN apk add --no-cache bash git curl
CMD [ "sh" ]