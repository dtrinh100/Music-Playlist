FROM golang:alpine
LABEL maintainer "Hector Lovo <lovohh@gmail.com>"

# Needed to download Go-packages
RUN apk add --no-cache --virtual git

# Directory is based upon current Go-file's references
ENV app_dir /go/src/github.com/dtrinh100/Music-Playlist/src/api/
RUN mkdir -p ${app_dir}
WORKDIR ${app_dir}

# Adding Go files
COPY . ${app_dir}

# Download & install Go-dependencies
RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
