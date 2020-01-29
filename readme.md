# Bkpbot

My personal backup solution. Archives a list of files into a single file. Supported formats are `zip` and `tar`.

## How to use?

Usually we backup into `~/.backup`. Therefore we put all the necessary files into this folder. There should be the following:

1. bkpbot binary
1. config.yaml
1. run.sh (look into `scripts`)

Backups may be launched by crontab. Examples below.

## Examples

### Crontab

```crontab
5 5 * * * /home/olaf/.backup/run.sh daily 9
3 3 * * 1 /home/olaf/.backup/run.sh weekly 11
1 1 1 * * /home/olaf/.backup/run.sh monthly 13
```

### Bkpbot Config

```yaml
-
    to_backup:
        - ${HOME}/documents
    detect: false
    output_folder: <CONFIGDIR>
    format: zip
-
    to_backup:
        - ${HOME}/rolling/aip
    detect: true
    exclusions:
        - .*\/github$
    output_folder: <CONFIGDIR>
    format: zip
```
