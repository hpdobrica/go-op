# 1Password Package for Golang

The this project is a wrapper around 1password CLI (https://support.1password.com/command-line/) with the goal of making an easy-to-use programatic interface towards 1Password for Go applications.

Goals of the project:
- Mimic 1Password CLI API as closely as possible
- Provide fully typed responses from the 1Password CLI 
- Make quality of life improvements (e.g. refreshing sessions)


Expected environment variables:

```
OP_DEVICE
OP_SIGNIN_ADDRESS
OP_EMAIL
OP_SECRET_KEY
OP_MASTER_PASSWORD
```