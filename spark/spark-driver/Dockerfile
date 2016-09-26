# Version 0.4
FROM debian:jessie
MAINTAINER Anirudh Ramanathan "ramanathana@google.com"

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update
RUN apt-get dist-upgrade -y
RUN apt-get install -y wget curl default-jre
RUN export JAVA_HOME=/usr/lib/jvm/java-7-openjdk-amd64
RUN wget https://1f5b3e8dda16d464bd3e2da586fd08bff4e30ed2.googledrive.com/host/0BwTtSeDsboEFTFpqNHNQQTdwcWs/original-spark-examples_2.11-2.1.0-SNAPSHOT.jar


# spark snapshot copy
RUN wget -P /opt/ https://storage.googleapis.com/foxish-spark-distro/spark.tgz
RUN mkdir /opt/spark
RUN tar xf /opt/spark.tgz -C /opt/spark --strip-components=1 && rm /opt/spark.tgz
COPY spark-conf/* /opt/spark/conf/
ENV SPARK_HOME /opt/spark