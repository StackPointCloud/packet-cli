## packet ssh-key delete

Deletes an SSH key

### Synopsis

Example:

packet ssh-key delete --id [ssh-key_UUID]



```
packet ssh-key delete [flags]
```

### Options

```
  -f, --force       Force removal of the SSH key
  -h, --help        help for delete
  -i, --id string   UUID of the SSH key
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

* [packet ssh-key](packet_ssh-key.md)	 - SSH key operations

