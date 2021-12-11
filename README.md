# Phalanx

Phalanx is a cloud-native full-text search and indexing server written in [Go](https://golang.org) built on top of [Bluge](https://github.com/blugelabs/bluge) that provides endpoints through [gRPC](https://grpc.io/) and traditional RESTful API.  
Phalanx implements a cluster formation by [hashicorp/memberlist](https://github.com/hashicorp/memberlist) and managing index metadata on [etcd](https://etcd.io/), so it is easy to bring up a fault-tolerant cluster.  
Metrics for system operation can also be output in Prometheus exposition format, so that monitoring can be done immediately using [Prometheus](https://prometheus.io/).  
Phalanx is using object storage for the storage layer, it is only responsible for the computation layer, such as indexing and retrieval processes. Therefore, scaling is easy, and you can simply add new nodes to the cluster.  
Currently, it is an alpha version and only supports [MinIO](https://min.io/) as the storage layer, but in the future it will support [Amazon S3](https://aws.amazon.com/s3/), [Google Cloud Storage](https://cloud.google.com/storage), and [Azure Blob Storage](https://azure.microsoft.com/en-us/services/storage/blobs/).  


## Build

Building Phalanx as following:

```bash
% git clone https://github.com/mosuka/phalanx.git
% cd phalanx
% make build
```


## Binary

You can see the binary file when build successful like so:

```bash
$ ls ./bin
phalanx
```


## Start Phalanx on local machine with local file system

Phalanx can be started on a local machine as if it were using local storage instead of object storage.  
A configuration file is available for starting on the local machine. You can start Phalanx by using it or by specifying the contents described in it with flags in the CLI. The following command starts with a configuration file:

```
% phalanx --config-file=./examples/phalanx_local.yml
```


## Start Phalanx on local machine with MinIO and etcd

To experience the features of Phalanx, let's try to start Phalanx using MinIO and etcd.  
First of all, let's start MinIO and etcd that Phalanx depends on. You can start them yourself, but there is a docker-compose.yml in this repository. You can use this to easily run MinIO and etcd on Docker.

```
% docker-compose up
```

Once the container has been started, you can check the MinIO and etcd data in your browser at the following URL.

- MinIO  
http://localhost:9001/dashboard

- ETCD Keeper  
http://localhost:8080/etcdkeeper/

Then, let's start Phalanx. A configuration file is available for starting on the local machine with MinIO and etcd. You can start Phalanx by using it or by specifying the contents described in it with flags in the CLI. The following command starts with a configuration file:

```
% phalanx --config-file=./examples/phalanx.yml
% phalanx --config-file=./examples/phalanx.yml --bind-port=2001 --grpc-port=5001 --http-port=8001 --seed-addresses=0.0.0.0:2000
% phalanx --config-file=./examples/phalanx.yml --bind-port=2002 --grpc-port=5002 --http-port=8002 --seed-addresses=0.0.0.0:2000
```

The above command overwrites the values in the configuration file with CLI flags in order to start multiple Phalanx nodes. Above example shows each Phalanx node running on the same host, so each node must listen on different ports. This would not be necessary if each node ran on a different host.  
This instructs each new node to join an existing node, each node recognizes the joining clusters when started. So you have a 3-node cluster.


## Health check

### Liveness check

```
% curl -XGET http://localhost:8000/livez | jq .
```

```json
{
  "state":"alive"
}
```

### Readiness check

```
% curl -XGET http://localhost:8000/readyz | jq .
```

```json
{
  "state":"ready"
}
```

There are endpoints, but it is not yet fully implemented.


## Metrics exposition

```
% curl -XGET http://localhost:8000/metrics
```

```text
# HELP phalanx_grpc_server_handled_total Total number of RPCs completed on the server, regardless of success or failure.
# TYPE phalanx_grpc_server_handled_total counter
phalanx_grpc_server_handled_total{grpc_code="Aborted",grpc_method="AddDocuments",grpc_service="index.Index",grpc_type="unary"} 0
phalanx_grpc_server_handled_total{grpc_code="Aborted",grpc_method="Cluster",grpc_service="index.Index",grpc_type="unary"} 0
...
```


## Cluster status

```
% curl -XGET http://localhost:8000/cluster | jq .
```

```json
{
  "indexer_assignment": {},
  "indexes": {},
  "nodes": {
    "node-KWkaw0jJ": {
      "addr": "0.0.0.0",
      "meta": {
        "grpc_port": 5000,
        "http_port": 8000,
        "roles": [
          "indexer",
          "searcher"
        ]
      },
      "port": 2000,
      "state": "alive"
    }
  },
  "searcher_assignment": {}
}
```


## Create index

### Create index on local file system

If you have started Phalanx to use the local file system, you can use this command to create an index.

```
% curl -XPUT -H 'Content-type: application/json' http://localhost:8000/v1/indexes/example_en --data-binary @./examples/create_index_example_en_local.json
```

### Create index with MinIO and etcd

If you have started Phalanx to use MinIO and etcd, use this command to create the index.

```
% curl -XPUT -H 'Content-type: application/json' http://localhost:8000/v1/indexes/example_en --data-binary @./examples/create_index_example_en.json
```

The difference between the above commands is the difference between `index_uri` and `lock_uri` in the configuration file. This parameter specifies where the index and its lock file will be created.


## Delete index

```
% curl -XDELETE http://localhost:8000/v1/indexes/example_en
```


## Add / Update documents

```
% curl -XPUT -H 'Content-type: application/x-ndjson' http://localhost:8000/v1/indexes/example_en/documents --data-binary @./examples/add_documents.ndjson
```


## Delete documents

```
% curl -XDELETE -H 'Content-type: text/plain' http://localhost:8000/v1/indexes/example_en/documents --data-binary @./examples/delete_ids.txt
```


## Search

```
% curl -XPOST -H 'Content-type: text/plain' http://localhost:8000/v1/indexes/example_en/_search --data-binary @./examples/search.json | jq .
```

```json
{
  "documents": [
    {
      "fields": {
        "_timestamp": "2021-12-10T13:03:18Z",
        "category": "/language/rust",
        "description": "Sonic is a fast, lightweight and schema-less search backend.",
        "name": "Sonic",
        "popularity": 7895,
        "publish_date": "2019-12-10T14:13:00Z",
        "url": "https://github.com/valeriansaliou/sonic"
      },
      "id": "7",
      "score": 0.37863163826497015
    },
    {
      "fields": {
        "_timestamp": "2021-12-10T13:03:18Z",
        "category": "/language/python",
        "description": "Whoosh is a fast, pure Python search engine library.",
        "name": "Whoosh",
        "popularity": 0,
        "publish_date": "2019-10-07T20:30:26Z",
        "url": "https://bitbucket.org/mchaput/whoosh/wiki/Home"
      },
      "id": "11",
      "score": 0.3731338946601548
    },
    {
      "fields": {
        "_timestamp": "2021-12-10T13:03:18Z",
        "category": "/language/java",
        "description": "Apache Lucene is a high-performance, full-featured text search engine library written entirely in Java.",
        "name": "Lucene",
        "popularity": 3135,
        "publish_date": "2019-12-19T05:08:00Z",
        "url": "https://lucene.apache.org/"
      },
      "id": "9",
      "score": 0.3710793549141038
    },
    {
      "fields": {
        "_timestamp": "2021-12-10T13:03:18Z",
        "category": "/language/go",
        "description": "Riot is Go Open Source, Distributed, Simple and efficient full text search engine.",
        "name": "Riot",
        "popularity": 4948,
        "publish_date": "2019-12-15T22:12:00Z",
        "url": "https://github.com/go-ego/riot"
      },
      "id": "5",
      "score": 0.3611255085637879
    },
    {
      "fields": {
        "_timestamp": "2021-12-10T13:03:18Z",
        "category": "/language/rust",
        "description": "Tantivy is a full-text search engine library inspired by Apache Lucene and written in Rust.",
        "name": "Tantivy",
        "popularity": 3142,
        "publish_date": "2019-12-19T01:07:00Z",
        "url": "https://github.com/quickwit-inc/tantivy"
      },
      "id": "8",
      "score": 0.34530979286026436
    },
    {
      "fields": {
        "_timestamp": "2021-12-10T13:03:18Z",
        "category": "/language/java",
        "description": "Elasticsearch is a distributed, open source search and analytics engine for all types of data, including textual, numerical, geospatial, structured, and unstructured.",
        "name": "Elasticsearch",
        "popularity": 46054,
        "publish_date": "2019-12-18T23:19:00Z",
        "url": "https://www.elastic.co/products/elasticsearch"
      },
      "id": "3",
      "score": 0.13076457838717315
    },
    {
      "fields": {
        "_timestamp": "2021-12-10T13:03:18Z",
        "category": "/language/go",
        "description": "Phalanx is a cloud-native full-text search and indexing server written in Go built on top of Bluge that provides endpoints through gRPC and traditional RESTful API.",
        "name": "Phalanx",
        "popularity": 0,
        "publish_date": "2021-12-10T12:00:00Z",
        "url": "https://github.com/mosuka/phalanx"
      },
      "id": "1",
      "score": 0.13076457838717315
    },
    {
      "fields": {
        "_timestamp": "2021-12-10T13:03:18Z",
        "category": "/language/rust",
        "description": "Toshi is meant to be a full-text search engine similar to Elasticsearch. Toshi strives to be to Elasticsearch what Tantivy is to Lucene.",
        "name": "Toshi",
        "popularity": 2448,
        "publish_date": "2019-12-01T19:00:00Z",
        "url": "https://github.com/toshi-search/Toshi"
      },
      "id": "6",
      "score": 0.13076457838717315
    },
    {
      "fields": {
        "_timestamp": "2021-12-10T13:03:18Z",
        "category": "/language/go",
        "description": "Blast is a full text search and indexing server, written in Go, built on top of Bleve.",
        "name": "Blast",
        "popularity": 654,
        "publish_date": "2019-10-18T10:50:00Z",
        "url": "https://github.com/mosuka/blast"
      },
      "id": "4",
      "score": 0.08523749485612774
    },
    {
      "fields": {
        "_timestamp": "2021-12-10T13:03:18Z",
        "category": "/language/rust",
        "description": "Quickwit is a distributed search engine built from the ground up to offer cost-efficiency and high reliability.",
        "name": "quickwit",
        "popularity": 0,
        "publish_date": "2021-07-13T15:07:00Z",
        "url": "https://github.com/quickwit-inc/quickwit"
      },
      "id": "13",
      "score": 0.08063697039612684
    }
  ],
  "hits": 10,
  "index_name": "example_en"
}
```


## Docker container

### Build Docker container image

You can build the Docker container image like so:

```
% make docker-build
```

### Pull Docker container image from docker.io

You can also use the Docker container image already registered in docker.io like so:

```
% docker pull mosuka/phalanx:latest
```

See https://hub.docker.com/r/mosuka/phalanx/tags/

### Start on Docker

Running a Blast data node on Docker. Start Blast node like so:

```bash
$ docker run --rm --name phalanx-node1 \
    -p 2000:2000 \
    -p 5000:5000 \
    -p 8000:8000 \
    mosuka/phalanx:latest start \
      --host=0.0.0.0 \
      --bind-port=2000 \
      --grpc-port=5000 \
      --http-port=8000 \
      --roles=indexer,searcher \
      --index-metastore-uri=file:///tmp/phalanx/metadata
```
