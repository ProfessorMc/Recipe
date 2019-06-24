# Environment Setup
The following are the prerequisites for executing the go examples provided.
# GIT
### Install Link
- To install Git for Windows use the link [Git](https://git-scm.com/download/win)
- To verify installation, run from a command line:
  ```
  git version
  ```
### Configure Proxy

- If operating behind a corporate proxy, set the proxy settings using the following command:
  
```
git config --global http.proxy http://<username>:<password>@<proxy-server-url>:<port>
```
- Note: Proxy settings require the use of URL safe encoding, so special characters may need to be encoded
  - Example: 12@34 would be encoded as 12%4034
  - You can encode it using powershell by opening a powershell script and typing:
```
Add-Type -AssemblyName System.Web
[System.Web.HttpUtility]::UrlEncode("<password>")
```
# Go
### Install Line
- To install Git for Windows use the link [Go](https://golang.org/doc/install)
- To verify installation, run from a command line:
  ```
  go version
  ```
- Note: You may have to refresh/restart the command line before it will work

### Environment Variables
- To Run Go Correcty, verify that go environment variables are set correctly
``` Powershell
Get-ChildItem Env:GOPATH
Get-ChildItem Env:GOROOT
```
- You can also use the UI to open environment variables window
  - Type: `enivronment variables` in the search bar
  - Click the Button to view the environment variables.
  - Add the following to your PATH Variable (User or system)
```
%GOPATH%\bin
%GOROOT%\bin
```
### Configure Proxy
Go uses the proxy settings set in environment variables.  It requires, http_proxy and https_proxy.  It is also recommended that you include the noproxy variable be set as these are commonly used and will break applications such as postmand if not set properly.

```
http_proxy = http://<username>:<password>@<proxy-server-url>:<port>
https_proxy = http://<username>:<password>@<proxy-server-url>:<port>
noproxy = "127.0.*, localhost, *.corp.myinternal.com"
```

# IDE Options
The IDE is optional and there are several choices, but for this class, we will use VSCode.  It is free and open source.

### IDEs
- [GoLand](https://www.jetbrains.com/go/) - IMO the best IDE out there now for go projects, 30 day trial, but paid beyond that.
- [Visual Studio Code](https://code.visualstudio.com/) - Free, open source and has go plugins supported by microsoft.
- [VIM Go](https://github.com/fatih/vim-go) - If you use this, we can't be friends.
- [Atom](https://ide.atom.io/) - Gree and open source has community go plugin support.

