FROM node:alpine
LABEL maintainer = "Hector Lovo <lovohh@gmail.com>"

# =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=
# =-=-=-=-=-=-=-=-= ANGULAR2 CONFIG =-=-=-=-=-=-=-=-=-=-=-=
# =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=

ENV USER app
ENV HOME /home/${USER}
ENV APP mp_app

RUN npm install -g @angular/cli

COPY package.json ${HOME}/${APP}/
# TODO: it may be necessary to use/copy => "npm-shrinkwrap.json" in production

WORKDIR ${HOME}/${APP}

RUN npm install

COPY . ${HOME}/${APP}

EXPOSE 4200

CMD ["npm", "start"]
