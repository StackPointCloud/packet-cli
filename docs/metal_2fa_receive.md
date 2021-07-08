## metal 2fa receive

Receive two factor authentication token

### Synopsis

Example:
Issue the token via SMS:
metal 2fa receive -s 

Issue the token via app:
metal 2fa receive -a



```
metal 2fa receive [flags]
```

### Options

```
  -a, --app    Issues otp uri for auth application
  -h, --help   help for receive
  -s, --sms    Issues SMS otp token to user's phone
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

* [metal 2fa](metal_2fa.md)	 - Two Factor Authentication operations

