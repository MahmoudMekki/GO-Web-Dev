FROM docker/whalesay:latest
RUN apt-get -y update && apt-get install -y fortunes
ENTRYPOINT /usr/games/fortune -a |cowsay