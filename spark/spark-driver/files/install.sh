#!/usr/bin/env sh

mkdir -p /opt/spark
tar xf /work-dir/spark.tgz -C /opt/spark --strip-components=1 && rm /work-dir/spark.tgz
cp /conf/* /opt/spark/conf/
export SPARK_HOME=/opt/spark
