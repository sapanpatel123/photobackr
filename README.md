# photobackr
Easy backup CLI for your Synology NAS.

## Setup

1. Create a config named `~/.photobackr.yaml` in the following format:

```
synology:
  username: <synology-username>
  privateKeyPath: <private-key-path>
  host: <synology-hostname>
```

## Dev

To run photobackr, use the following command:

```
go run main.go backup --source <source-dir> --destination <destination-dir>
```


Note: This is still a work in progress (WIP)!
