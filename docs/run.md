# Run

> Back-up is uploaded by streaming to the Storj network.

The following flags can be used with the `store` command:

* `accesskey` - Connects to the Storj network using a serialized access key instead of an API key, satellite url and encryption passphrase.
* `share` - Generates a restricted shareable serialized access with the restrictions specified in the Storj configuration file.

Once you have built the project you can run the following:

## Get help

```
$ ./connector-nextcloud --help
```

## Check version

```
$ ./connector-nextcloud --version
```

## Upload back-up data to Storj

```
$ ./connector-nextcloud store --local <path_to_nextcloud_config_file> --storj <path_to_storj_config_file>
```

## Upload back-up data to Storj bucket using Access Key

```
$ ./connector-nextcloud store --accesskey
```

## Upload back-up data to Storj and generate a Shareable Access Key based on restrictions in `storj_config.json`

```
$ ./connector-nextcloud store --share
```
