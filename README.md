Go-Hadoop
========

Go-Hadoop is native go clients for Apache Hadoop YARN.

It includes an early version of Hadoop IPC client and requisite YARN client libraries to implement YARN applications completely in go (both YARN application-client and application-master).

### Why go-hadoop?
Firstly, thanks for `hortonworks/gohadoop`, this package make it easier for us to access YARN components.   
But, it seems no longer being maintained, and the protobuf files are too old, so what difference in this repo:

* Protobuf files have been upgraded to YARN-2.8.2 .
* New RPC service: ResourceManagerAdministrationProtocolService.
* We will keep it up-to-date.

### Notes:
* Set HADOOP_CONF_DIR environment variable, and ensure the conf directory contains `yarn-site.xml`.
* hadoop_yarn/examples/dist_shell is an example go YARN application: client.go is the submission client and applicationmaster.go is the application-master.

To run:

### Pull Repo
```
mkdir -p $GOPATH/src/github.com/stormgbs
cd $GOPATH/src/github.com/stormgbs
git clone https://github.com/stormgbs/go-hadoop.git
```

### Run regular dist-shell with AM in-cluster
```
HADOOP_CONF_DIR=conf go run hadoop_yarn/examples/dist_shell/client.go
```

### Run dist-shell with unmanaged AM
```
HADOOP_CONF_DIR=conf go run hadoop_yarn/examples/dist_shell/unmanaged.go
```
