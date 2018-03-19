# create_repository

## Usage
```
$ cr -n
flag needs an argument: -n
Usage of C:\app\bin\cr.exe:
  -github_token string
        Github token
  -n string
        Repository Name
```
1. create [Personal access tokens](https://github.com/settings/tokens) (Require Repository access)  
2. set Environment **GITHUB_TOKEN** for Personal access tokens
3. set Environment **GITHUB_API_URL** (Default set https://api.github.com/)
4. and hit this command....

`cr -n <New Repository Name>`

## Requirements
Nop

## Installation
Go to the releases page, find the version you want, and download the zip file. Unpack the zip file, and put the binary to somewhere you want (on UNIX-y systems, /usr/local/bin or the like). Make sure it has execution bits turned on. Yes, it is a single binary! You can put it anywhere you want :)

## Author
R.Maejima(dorastone@gmail.com)
