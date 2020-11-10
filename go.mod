module github.com/triole/bkpbot

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/alecthomas/template v0.0.0-20160405071501-a0175ee3bccc // indirect
	github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf // indirect
	github.com/dsnet/compress v0.0.1 // indirect
	github.com/frankban/quicktest v1.11.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/mholt/archiver v3.1.1+incompatible
	github.com/nwaples/rardecode v1.1.0 // indirect
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/triole/bkpbot/env v0.0.0-00010101000000-000000000000
	github.com/triole/bkpbot/logging v0.0.0-00010101000000-000000000000
	github.com/triole/bkpbot/rx v0.0.0-00010101000000-000000000000
	github.com/triole/bkpbot/syslib v0.0.0-00010101000000-000000000000
	github.com/ulikunitz/xz v0.5.7 // indirect
	github.com/xi2/xz v0.0.0-20171230120015-48954b6210f8 // indirect
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

replace github.com/triole/bkpbot/env => ./src/env

replace github.com/triole/bkpbot/logging => ./src/logging

replace github.com/triole/bkpbot/rx => ./src/rx

replace github.com/triole/bkpbot/syslib => ./src/syslib
