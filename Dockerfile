FROM node:alpine
LABEL maintainer "Hector Lovo <lovohh@gmail.com>"

# =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
# =-=-=-=-=-=-=-=-= ANGULAR2 CONFIG =-=-=-=-=-=-=-=-=-=-=-=
# =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

# Builds "production" version to serve
ENV app_dir /home/app/mp_app
RUN mkdir -p ${app_dir}
WORKDIR ${app_dir}

COPY . ${app_dir}

RUN npm install -g @angular/cli \
 && npm install \
 && npm cache verify \
 && npm run build \
 && mv dist/* /srv \
 && mv Caddyfile /etc/Caddyfile \
 && cd /srv \
 && rm -fr ${app_dir}


# =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
# =-=-=-=-=-=-=-=-=-= CADDY CONFIG -=-=-=-=-=-=-=-=-=-=-=-=
# =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

# Installs & starts Caddy
ARG plugins=http.git

RUN apk add --no-cache openssh-client git tar curl
RUN curl --silent --show-error --fail --location \
      --header "Accept: application/tar+gzip, application/x-gzip, application/octet-stream" -o - \
      "https://caddyserver.com/download/linux/amd64?plugins=${plugins}" \
    | tar --no-same-owner -C /usr/bin/ -xz caddy \
 && chmod 0755 /usr/bin/caddy \
 && /usr/bin/caddy -version

EXPOSE 80 443 2015

WORKDIR /srv

ENTRYPOINT ["/usr/bin/caddy"]
CMD ["--conf", "/etc/Caddyfile", "--log", "stdout"]
