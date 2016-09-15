# Version 0.1
FROM ipython/scipystack
MAINTAINER Anirudh Ramanathan "ramanathana@google.com"

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update
RUN apt-get dist-upgrade -y

# Setup
RUN sudo apt-key adv --keyserver keyserver.ubuntu.com --recv E56151BF
RUN apt-get -y update
RUN apt-get install -y python libnss3 curl

RUN curl http://d3kbcqa49mib13.cloudfront.net/spark-1.6.0-bin-hadoop2.6.tgz \
    | tar -xzC /opt && \
    mv /opt/spark* /opt/spark

RUN apt-get clean

# Fix pypspark six error.
RUN pip2 install -U six
RUN pip2 install msgpack-python
RUN pip2 install avro

COPY spark-conf/* /opt/spark/conf/

ENV SPARK_HOME /opt/spark