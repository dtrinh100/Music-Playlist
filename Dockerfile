FROM abiosoft/caddy
LABEL maintainer "Hector Lovo <lovohh@gmail.com>"

# Copies the "production" version of the app into /srv.
# Note: This assumes that the directory ./dist exists locally.
COPY ./dist/ /srv/

COPY ./Caddyfile /etc/Caddyfile
