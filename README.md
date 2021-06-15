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
