# Bkpbot

My personal backup solution. Archives a list of files into a single file. Supported formats are `zip` and `tar`.

## Config Example

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
