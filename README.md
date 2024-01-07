# Go Contact Server

## Introduction

This is a lightweight HTTP(S) server for handling contact form submissions. With just a few enviornment variables set up you can run the server and start seeing emails routed appropriately.

## Setup
Since the server is written in Go, you will need to have a Go compiler installed on your machine to generate the executable. That can be found at the [official website](https://go.dev/doc/install).

Once Go is installed, you should be able to compile the source code by running the following command at the root.
```
go build main.go
```
This should generate an executable at your current working directory, which can then be run to start the server. You may see some errors if you try to run the program without configuring the environment properly. The next section covers all the necessary environment variables to get the server ready for end-to-end service.

## List of Environment Variables
| Name | Description |
| - | - |
| `CONTACTSERVER_SMTP_SERVER` | URL for the SMTP server. If you are routing through a popular email service, you can probably find this server URL through some quick search engine queries. |
| `CONTACTSERVER_SMTP_PORT` | The port the server listens for SMTP traffic. `587` is the typical. |
| `CONTACTSERVER_SMTP_SENDER` | The email address you want to send the automated emails from. |
| `CONTACTSERVER_SMTP_RECEIVER` | The email address you want the automated emails to be sent to. |
| `CONTACTSERVER_SMTP_PASSWORD` | Password to connect to the SMTP server. Some services will require an authentication token separate from your usual email password (like Yahoo app passwords). Make sure to check if your SMTP server has these restrictions. |
| `CONTACTSERVER_USE_HTTPS` | Indicates if traffic to the server should use HTTP or HTTPS. `1`=HTTPS, `0`=HTTP |
| `CONTACTSERVER_TLS_CERT` | Path to the TLS certificate on your server. |
| `CONTACTSERVER_TLS_KEY` | Path to the TLS private key on your server. |

## Sending a Request
If the server is now running, it is time to try sending a request! The URL you should send your request to should print out to standard output when you run the executable. The endpoint accepts JSON and the body should appear in the following format

```json
{
    "name": "Some Name",
    "contact": "Some Contact Info (can be email, phone, etc)",
    "message": "Some Message"
}
```
If the server received the request it will output some information to standard output and eventually a response should be sent to the client.
```json
{
    "success": true
}
```
If the response shows that it did not succeed, try checking the console on the server for troubleshooting.
## Contact
If you would like to contact the author, please visit [my website](https://david-allan-jones.github.io/personal-website/contact) and leave a message via the contact form there.
