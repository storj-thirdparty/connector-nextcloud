# connector-nextcloud (uplink v1.0.5)

[![Go Report Card](https://goreportcard.com/badge/github.com/storj-thirdparty/connector-nextcloud)](https://goreportcard.com/report/github.com/storj-thirdparty/connector-nextcloud)

## Overview

The nextcloud Connector connects to nextcloud and transfers all the files with the file structure to Storj network.


```bash
Usage:
  connector-nextcloud [command] <flags>

Available Commands:
  help        Help about any command
  store       Command to upload data to a Storj V3 network.
  version     Prints the version of the tool

```

`store` - Connect to the specified database (default: `nextcloud_property.json`). Back-up of the database is generated using tooling provided by Nextcloud and then uploaded to the Storj network. Connect to a Storj v3 network using the access specified in the Storj configuration file (default: `storj_config.json`).

Back-up data is iterated through and upload in 32KB chunks to the Storj network.

The following flags  can be used with the `store` command:

* `accesskey` - Connects to the Storj network using a serialized access key instead of an API key, satellite url and encryption passphrase.
* `shared` - Generates a restricted shareable serialized access with the restrictions specified in the Storj configuration file.
* `debug` - Download the uploaded backup files to local disk inside project_folder/debug folder.



Sample configuration files are provided in the `./config` folder.


## Requirements and Install

To build from scratch, [install the latest Go](https://golang.org/doc/install#install).

> Note: Ensure go modules are enabled (GO111MODULE=on)

#### Option #1: clone this repo (most common)

To clone the repo

```
git clone https://github.com/storj-thirdparty/connector-nextcloud.git
```

Then, build the project using the following:

```
cd connector-nextcloud
go build
```

#### Option #2:  ``go get`` into your gopath

 To download the project inside your GOPATH use the following command:

```
go get github.com/storj-thirdparty/connector-nextcloud
```

**NOTE**: The above command may show the following warnings and can be ingnored:
```
github.com/storj-thirdparty/connector-nextcloud/nextcloud
../../go/src/github.com/storj-thirdparty/connector-nextcloud/nextcloud/nextcloud.go:78:2: cannot use nextcloudClient (type gonextcloud.Client) as type *gonextcloud.Client in return argument:
 *gonextcloud.Client is pointer to interface, not interface
 
../../go/src/github.com/storj-thirdparty/connector-nextcloud/nextcloud/nextcloud.go:83:33: nextcloudClient.WebDav undefined (type *gonextcloud.Client is pointer to interface, not interface)

../../go/src/github.com/storj-thirdparty/connector-nextcloud/nextcloud/nextcloud.go:109:34: nextcloudClient.WebDav undefined (type *gonextcloud.Client is pointer to interface, not interface)

../../go/src/github.com/storj-thirdparty/connector-nextcloud/nextcloud/nextcloud.go:130:41: nextcloudClient.WebDav undefined (type *gonextcloud.Client is pointer to interface, not interface)
```


## Run (short version)

Once you have built the project run the following commands as per your requirement:

##### Get help

```
$ ./connector-nextcloud --help
```

##### Check version

```
$ ./connector-nextcloud version
```

##### Transfer files from Nextcloud to Storj

```
$ ./connector-nextcloud store
```


## Documentation

For more information on runtime flags, configuration, testing, and diagrams, check out the [Detail](//github.com/storj-thirdparty/connector-nextcloud/wiki/) or jump to:

* [Config Files](//github.com/storj-thirdparty/connector-nextcloud/wiki/#config-files)
* [Run (long version)](//github.com/storj-thirdparty/connector-nextcloud/wiki/#run)
* [Testing](//github.com/storj-thirdparty/connector-nextcloud/wiki/#testing)
* [Flow Diagram](//github.com/storj-thirdparty/connector-nextcloud/wiki/#flow-diagram)
