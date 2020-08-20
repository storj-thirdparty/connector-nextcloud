## Flow Diagram

![](https://github.com/storj-thirdparty/connector-nextcloud/blob/master/README.assets/arch.drawio.png)

## Config Files

There are two config files that contain Storj network and Nextcloud connection information. The tool is designed so you can specify a config file as part of your tooling/workflow.



##### `nextcloud_property.json`

Inside the `./config` directory there is a `nextcloud_property.json` file, with following information about your nextcloud account instance:

* `url`- Login url of the user nextcloud account of the corresponding service provider
* `username` - Username of nextcloud account
* `password` - Password of nextcloud account


##### `storj_config.json`

Inside the `./config` directory a `storj_config.json` file, with Storj network configuration information in JSON format:

* `apikey` - API Key created in Storj Satellite GUI (mandatory)
* `satellite` - Storj Satellite URL (mandatory)
* `encryptionpassphrase` - Storj Encryption Passphrase (mandatory)
* `bucket` - Name of the bucket to upload data into (mandatory)
* `uploadPath` - Path on Storj Bucket to store data (optional) or "" or "/". (mandatory)
* `serializedAccess` - Serialized access shared while uploading data used to access bucket without API Key (mandatory)
* `allowDownload` - Set *true* to create serialized access with restricted download (mandatory while using *share* flag)
* `allowUpload` - Set *true* to create serialized access with restricted upload (mandatory while using *share* flag)
* `allowList` - Set *true* to create serialized access with restricted list access
* `allowDelete` - Set *true* to create serialized access with restricted delete
* `notBefore` - Set time that is always before *notAfter*
* `notAfter` - Set time that is always after *notBefore*



## Run

Once you have built the project run the following commands as per your requirement:

##### Get help

```
$ ./connector-nextcloud help
```

##### Check version

```
$ ./connector-nextcloud version
```

##### Connect to a desired Nextcloud account and take complete back-up to desired Storj network bucket using API key from `storj_config.json`
```
$ ./connector-nextcloud store --nextcloud <path-to-nextcloud-config> --storj <path-to-storj-config>
```

##### Connect to to a desired Nextcloud account and take complete back-up to desired Storj network bucket Serialized Access Key from `storj_config.json`
```
$ ./connector-nextcloud store --accesskey
```

##### Connect to a desired Nextcloud account and take complete back-up to desired Storj network bucket using API key and generates a Shareable Serialized Access Key with restrictions in `storj-config.json`.
```
$ ./connector-nextcloud store --share
```



## Testing
* The project has been tested on the following operating systems:
```
* Windows
	* Version: 10 Pro
	* Processor: Intel(R) Core(TM) i3-5005U CPU @ 2.00GHz 2.00GHz

* macOS Catalina
	* Version: 10.15.4
	* Processor: 2.5 GHz Dual-Core Intel Core i5

* ubuntu
	* Version: 16.04 LTS
	* Processor: AMD A6-7310 APU with AMD Radeon R4 Graphics × 4
```