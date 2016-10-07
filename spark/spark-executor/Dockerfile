# Version 0.4
FROM openjdk:alpine
MAINTAINER Anirudh Ramanathan "ramanathana@google.com"

# spark snapshot copy
RUN apk update && apk add tar bash
RUN mkdir /conf #redo2
COPY files/conf/* /conf/
COPY /files/install.sh /
RUN chmod +x /install.sh