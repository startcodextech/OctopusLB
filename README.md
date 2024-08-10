# Octopus LB
```shell
brew install buf
brew install protobuf
brew install protoc-gen-js
brew install protoc-gen-grpc-web
go get github.com/grpc-ecosystem/grpc-gateway/v2@v2.20.0
```

## Caddy
```shell
go install github.com/caddyserver/xcaddy/cmd/xcaddy@latest
xcaddy build --with github.com/mholt/caddy-l4
```


## Cross-compile

For cross-compiling, you need to install the following packages.

### Compile for Linux
#### APT
```shell
sudo apt install gcc-multilib g++-multilib
sudo apt install gcc-aarch64-linux-gnu g++-aarch64-linux-gnu
```
#### DNF
```shell
sudo dnf install gcc gcc-c++ gcc-aarch64-linux-gnu g++-aarch64-linux-gnu
```

### Compile for MacOS
#### Brew
```shell
brew install FiloSottile/musl-cross/musl-cross
```

## Compile binaries
```shell
sudo make build
```

sudo systemctl status firewalld
sudo systemctl start firewalld
sudo systemctl enable firewalld
sudo firewall-cmd --zone=public --add-port=80/tcp --permanent
sudo firewall-cmd --zone=public --add-port=443/tcp --permanent
sudo firewall-cmd --reload
sudo firewall-cmd --list-ports
nmap -p 80,443 <tu-dominio-o-ip>
