# dxid-tool
This repo contains cli tool a collection of useful cli command under linux deployment
```
.
├── cmd
│   └── dxid
│       ├── install.go
│       └── main.go
├── dist
├── go.mod
├── pkg
│   └── util
│       ├── os_check.go
│       └── docker_install.go
└── README.md
```

# Installation 
Prequisite 
- golang installed

Build from source: 
- clone the repo
```

```
- build
```
go build -o dxid ./cmd/dxid
```
- post built
```
sudo mv dxid /usr/local/bin/
```

# How to use it

verify the version
```
dxid --version
```

support tools
```
dxid install --target docker
```