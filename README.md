
# Cross-compile

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