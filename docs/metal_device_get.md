## metal device get

Retrieves device list or device details

### Synopsis

Example:
	
packet device get --id [device_UUID]

	

```
metal device get [flags]
```

### Options

```
  -h, --help                help for get
  -i, --id string           UUID of the device
  -p, --project-id string   UUID of the project
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
  -y, --yaml              YAML output
```

### SEE ALSO

* [metal device](metal_device.md)	 - Device operations
