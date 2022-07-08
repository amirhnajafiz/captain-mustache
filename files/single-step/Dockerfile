# syntax=docker/dockerfile:experimental
# from our base image
FROM golang:${VERSION}-${IMAGE_NAME}

# in work directoru app
WORKDIR /app

# copy all files
COPY * ./

# building our golang app
RUN go build -o /docker ${MAIN_PATH}

# exposing on port
EXPOSE ${PORT}

# execuitng the file
CMD [ "/docker" ]
