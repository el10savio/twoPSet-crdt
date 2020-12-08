# 2PSet-crdt

2PSet CRDT Cluster implemented in Go & Docker

## Introduction

CRDTs (Commutative Replicated Data Types) are a certain form of data types that when replicated across several nodes over a network achieve eventual consistency without the need for a consensus round. 2PSets, unlike GSets sets are CRDT sets modified to add & remove data into it and becomes consistent across nodes in a cluster having replicated the set.

## Example

After building a cluster of 2PSet nodes we can now write/remove values to either one or many nodes in the cluster.

```
$ curl -i -X POST localhost:8080/twopset/add/user1
$ curl -i -X POST localhost:8081/twopset/add/user2
$ curl -i -X POST localhost:8080/twopset/remove/user1
$ curl -i -X POST localhost:8081/twopset/remove/user2
```

When reading the list of values in the set they then sync up with each other and thus return consistent values every time from any node in the cluster

```
$ curl -i -X GET localhost:8081/twopset/list
{
    set: [
        user1,
        user2
    ]
}
```

The values remain consistent for nodes in the cluster that have never added any values into it

```
$ curl -i -X GET localhost:8082/twopset/list
{
    set: [
        user1,
        user2
    ]
}
```

We can also lookup if certain values are present in the set

```
$ curl -i -X GET localhost:8082/twopset/lookup/user1
{
    present: true
}

$ curl -i -X GET localhost:8082/twopset/lookup/user3
{
    present: false
}
```

## Steps

After cloning the repo. To provision the cluster:

```
$ make provision
```

This creates a 3 node 2PSet cluster established in their own docker network.

To view the status of the cluster

```
$ make info
```

Now we can send requests to append, list, and lookup values to any peer node using its port allocated.

```
$ curl -i -X POST localhost:<peer-port>/twopset/add/<value>
$ curl -i -X POST localhost:<peer-port>/twopset/remove/<value>
$ curl -i -X GET localhost:<peer-port>/twopset/lookup/<value>
$ curl -i -X GET localhost:<peer-port>/twopset/list
```

In the logs for each peer docker container, we can see the logs of the peer nodes getting in sync during read operations.

To tear down the cluster and remove the built docker images:

```
$ make clean
```

This is not certain to clean up all the locally created docker images at times. You can do a docker rmi to delete them.

## References

- [A comprehensive study of Convergent and Commutative Replicated Data Types](https://hal.inria.fr/inria-00555588/document) [Marc Shapiro et al]
- [Strong Eventual Consistency and Conflict-free Replicated Data Types](https://www.youtube.com/watch?v=oyUHd894w18&t=3902s) [Marc Shapiro]
