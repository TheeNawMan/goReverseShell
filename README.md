# goReverseShell


To use this shell you need to install and update some current pacakages to be effective.

# Debian
```bash
sudo apt install git golang upx
```

# Clone the repo

```bash
# install golang and upx
git clone https://github.com/TheeNawMan/goReverseShell
cd goReverseShell
```

# Edit main.go to put in the information we want to use.
```bash
nano main.go
```

# Edit lines 64-65
```go
	Addr := "127.0.0.1"
	Port := "9001"
```

# Lets build :)

```bash
chmod +x build.sh
./build.sh
```