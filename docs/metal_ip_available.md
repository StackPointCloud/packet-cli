## metal ip available

Retrieves a list of IP resevations for a single project.

### Synopsis

Example:

metal ip available --reservation-id [reservation_id] --cidr [size_of_subnet]

  

```
metal ip available [flags]
```

### Options

```
  -c, --cidr int                Size of subnet
  -h, --help                    help for available
  -r, --reservation-id string   UUID of the reservation
```

### Options inherited from parent commands

```
      --config string     Path to JSON or YAML configuration file
      --exclude strings   Comma seperated Href references to collapse in results, may be dotted three levels deep
      --include strings   Comma seperated Href references to expand in results, may be dotted three levels deep
  -j, --json              JSON output
      --search string     Search keyword for use in 'get' actions. Search is not supported by all resources.
      --token string      Metal API Token (METAL_AUTH_TOKEN)
  -y, --yaml              YAML output
```

### SEE ALSO

* [metal ip](metal_ip.md)	 - IP operations

