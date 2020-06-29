# goodluck

![Size](https://img.shields.io/github/repo-size/2kabhishek/goodluck?style=plastic&color=0f0&label=Size)
![Updated](https://img.shields.io/github/last-commit/2kabhishek/goodluck?style=plastic&color=f00&label=Updated)
![Stars](https://img.shields.io/github/stars/2kabhishek/goodluck?style=plastic&color=ffc801&label=Stars)
![Forks](https://img.shields.io/github/forks/2kabhishek/goodluck?style=plastic&color=003cff&label=Forks)
![Watchers](https://img.shields.io/github/watchers/2kabhishek/goodluck?style=plastic&color=ff5500&label=Watchers)
![Contributors](https://img.shields.io/github/contributors/2kabhishek/goodluck?style=plastic&color=f0f&label=Contributors)
![License](https://img.shields.io/github/license/2kabhishek/goodluck?style=plastic&color=555&label=License)

Are you lucky today? Run goodluck to find out.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed the latest version of `go` & `fortune`.

## Installing goodluck

To install goodluck, follow these steps:

```bash
git clone https://github.com/2kabhishek/goodluck
cd goodluck
go build
go install
```

## Using goodluck

```bash
USAGE:
    goodluck
```

You may also have to export your fortune files path as an environment variable.
It can be done like this:

```bash
fortune -f # Copy the path in output
export FORTUNE_PATH="/copied/path/"
```
