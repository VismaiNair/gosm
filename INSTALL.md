# wasmate Installation Guide

## Installing via `go install`
### Prerequisites (Required for Installation)
1. Have Go installed. We recommend version 1.24 and later.

### Installation Process
1. In your terminal, run the following command: 
```
go install github.com/vismainair/wasmate@latest
```

## Uninstalling
### Prerequisites (Required for Uninstallation)
1. Have wasmate installed.
2. Have Go installed.

### Uninstallation Process
#### Windows
Run the following command in *PowerShell or Command Prompt*:
```
del %GOPATH%\bin\wasmate.exe
```

If you run into your problems relating to your GOPATH not being set, we recommend trying
```
del %USERPROFILE%\go\bin\wasmate.exe
```
We recommend trying the first method (%GOPATH%) first.
### macOS or Linux
Run the following command:
```
rm $GOPATH/bin/wasmate
```

If you run into problems relating to your GOPATH not being set, try
```
rm ~/go/bin/wasmate
```
We recommend trying the first method ($GOPATH) first.