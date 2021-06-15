# Nebula

**ALERT: This is not finished so don't expect it to acctually install things!**

This used to be a repository for an alternative [Homebrew](https://brew.sh/) client, but I decided to make a different _thing_. I do love `brew` however and use it daily, so go check them out!

Before you run this application, be sure to set the `NEBULA_PROCESSOR_ARCH` environment variable to either `arm64` for an M1 mac or `amd64` for an Intel one! This is **very** important, since it let's nebula know if you're running on arm64 or amd64! To set the processor arch run the following command in the terminal:

_Tip: the `nebula` command will warn you when installing a valid package without having the processor arch set._

M1 macs:
```bash
export NEBULA_PROCESSOR_ARCH=arm64
```

or

Intel macs:
```bash
export NEBULA_PROCESSOR_ARCH=amd64
```

## Installation

There are 2 direct requirements other than having GoLang installed. 

A quick way to install dependancies, build, and install nebula is by using Mage. Make sure that your GOPATH is set to `/Users/$USER/go` in most cases ($USER should be replaced with your user name). 

Here are the quick instructions for installing Mage:

1. Clone Mage: 

```bash
git clone https://github.com/magefile/mage
```

2. `cd` into the `mage` directory

3. Run the bootstrap script:

```bash
go run bootstrap.go
```

_You might have to run it with `sudo` if it doesn't work the first time; that fixes access denied errors._

4. Check if $GOPATH/bin is in your $PATH by running:

```bash
echo $PATH
```

5. And checking if `/Users/$USER/go/bin` is there; however it might not be there, if so run:

```bash
export PATH=/Users/$USER/go/bin:$PATH
```

6. Then go into the directory you would like to have the source code for nebula, then run:

```bash
git clone https://github.com/Reboot-Codes/nebula.git
```

7. `cd` into the `nebula` directory

8. Then run the following to install nebula:

```bash
mage install
```

_An error may happen where an access denied error happens, if so, run with `sudo`._

9. You should be finished! Try running `nebula`!
