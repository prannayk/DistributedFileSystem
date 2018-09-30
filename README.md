# Distributed file system

[![build status](https://travis-ci.org/prannayk/DistributedFileSystem.svg?branch=master)](https://travis-ci.org/prannayk/DistributedFileSystem)

A work in progress implementation of a Chord based Distributed File System

Completed : 
1. Consistent Ring based Hashing
2. [WIP] Implementation of Distributed Hash Table 

## Implementation details
1. Consistent hashing using the fast `xxhash` in order to get well distributed but fast hashing
2. Dependency of Protobuf for creating Client-Server API interface
3. Implemented circular finger update instead of Random to get a faster implementation
4. Field-wise Mutex allow two non intersecting R/W operations to happen simultaneously 

Based on implementations of consistent hashing by [xring](https://github.com/arriqaaq/xring), [xxhash](https://github.com/cespare/xxhash), [chord](https://github.com/r-medina/gmaj/blob/82b35fdc07a559af639411b261d1756314f95926/finger.go), [rbtree](https://github.com/arriqaaq/rbt)
