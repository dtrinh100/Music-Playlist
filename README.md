# MP

The Music Project is a project that allows users to listen to a variety of music (in this case, just sample pieces of music that are royalty free) and upload their own music.  This project was created by David Trinh and Hector Lovo to help us improve our development skills, while learning new technologies. It utilizes the latest cutting edge technologies. Codes in this project may not have top quality,
since we were limited on time, but we did strive to do our best with a small timeframe.

## Technology Stack

We use the following technology:

**Front-end:**

 - HTML5
 - CSS3/SASS
 - Javascript/Typescript with the Angular 2 framework

**Back-end:**

 - Golang with Gorilla
 - MongoDB

**Misc:**

 - Docker - to have a development environment
 - Caddy - super simple web server used to serve static content and reverse proxy our API

## Running Our App

To run our app, you only need 2 things installed on your system:

 - Docker
 - Docker Compose

Once you have installed docker and docker compose, run the following command:

    docker-compose build && docker-compose up
Note: You may need admin privileges to run these commands

Once you have run the command, simply open up your web browser and go to [0.0.0.0:2015](0.0.0.0:2015)
