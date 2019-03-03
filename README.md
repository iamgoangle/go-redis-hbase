# GO HBase and Redis

## Run

```sh
docker-compose up --build
```

## Hbase Dockerize

`docker` https://github.com/PierreZ/hbase-docker

`default port` https://blog.cloudera.com/blog/2013/07/guide-to-using-apache-hbase-ports

## Cloudera

`docker` https://github.com/emirkorkmaz/cloudera-quickstart-docker-compose

## Hue

`docker` https://github.com/cloudera/hue/tree/master/tools/docker

## Port

```sh
- "8020:8020"   # HDFS 
- "8022:22"     # SSH
- "7180:7180"   # Cloudera Manager
- "8888:8888"   # Hue
- "11000:11000" # Oozie
- "50070:50070" # HDFS Rest Namenode
- "50075:50075" # HDFS Rest Datanode
- "2181:2181"   # Zookeeper
- "8088:8088"   # YARN Resource Manager
- "19888:19888" # MapReduce Job History
- "50030:50030" # MapReduce Job Tracker
- "8983:8983"   # Solr
- "16000:16000" # Sqoop Metastore
- "8042:8042"   # YARN Node Manager
- "60010:60010" # HBase Master
- "60030:60030" # HBase Region
- "9090:9090"   # HBase Thrift
- "8080:8080"   # HBase Rest
- "7077:7077"   # Spark Master
```

