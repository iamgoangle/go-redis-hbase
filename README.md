# GO HBase and Redis

## Run

```sh
docker-compose up --build
```

## Hbase Dockerize

`docker` https://github.com/PierreZ/hbase-docker

`default port` https://blog.cloudera.com/blog/2013/07/guide-to-using-apache-hbase-ports

| **Component** | **Configuration parameter**           | **Default value** | **Used places** |
| ------------- | ------------------------------------- | ----------------- | --------------- |
| ZooKeeper     | `hbase.zookeeper.property.clientPort` | 2181              | 1,5,6           |
| Master        | `hbase.master.port`                   | 60000             | 3,7             |
| Master        | `hbase.master.info.port`              | 60010             | 11              |
| Region server | `hbase.regionserver.port`             | 60020             | 2,4,8           |
| Region server | `hbase.regionserver.info.port`        | 60030             | 12              |
| REST server   | `hbase.rest.port`**                   | 8080              | 9               |
| REST server   | `hbase.rest.info.port`*               | 8085              | 13              |
| Thrift server | `hbase.regionserver.thrift.port`**    | 9090              | 10              |
| Thrift server | `hbase.thrift.info.port`*             | 9095              | 14              |

