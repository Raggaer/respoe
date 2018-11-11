# Respoe

Path of Exile website client. Useful for automating tasks on the Path of Exile website or forum (bots or just gathering data).

It makes requests and then parses the response to fetch all the needed data. The library is separated in three packages:

- Client: Related to all account features.
- Forum: Related to all forum data.
- Trade: Related to trading websites
- Util: Related to util methods for developing tools.

## Usage

First get the package:

```go
go get github.com/raggaer/respoe/...
```

You can then use this as a library for developing tools related to the Path of Exile website and APIs, forum or the trade site.
You can automate the reading of forums, creating or replying threads or even automating private messaging (just like poe.trade bot does),
gather currency prices or gather ladder information.

## Features

Below is a list of the features that are already completed:

### Client

[![GoDoc](https://godoc.org/github.com/Raggaer/respoe/client?status.svg)](http://godoc.org/github.com/Raggaer/respoe/client)

- [x] Login
- [x] Logout
- [x] Change password
- [x] Retrieve inbox by page
- [x] Send private messages
- [x] Retrieve special offer list
- [x] Retrieve account profile
- [x] Retrieve account character list
- [x] Retrieve account character items
- [x] Retrieve active leagues

### Forum

[![GoDoc](https://godoc.org/github.com/Raggaer/respoe/forum?status.svg)](http://godoc.org/github.com/Raggaer/respoe/forum)

- [x] Forum list
- [x] Retrieve specific forum pagination information (first, current and last page)
- [x] Retrieve all threads from a forum (with status, views, author, creation date and replies)
- [x] Retrieve all posts from a thread (content, author, creation date, items)
- [x] Work with forum items (sockets)
- [x] Reply to a thread

### Trade

[![GoDoc](https://godoc.org/github.com/Raggaer/respoe/trade?status.svg)](http://godoc.org/github.com/Raggaer/respoe/trade)

- [x] Retrieve currency exchange offers

### Ladder

[![GoDoc](https://godoc.org/github.com/Raggaer/respoe/ladder?status.svg)](http://godoc.org/github.com/Raggaer/respoe/ladder)

- [x] Retrieve entries from a league ladder

### Util

[![GoDoc](https://godoc.org/github.com/Raggaer/respoe/util?status.svg)](http://godoc.org/github.com/Raggaer/respoe/util)

- [x] Retrieve pagination from any valid page
- [x] Retrieve hash value from any valid form
- [x] Retrieve hash value from any reply thread form
- [x] Retrieve errors from a submitted form

## Testing

All methods are testable. To run the tests you need to set some environment variables:

- `RESPOE_EMAIL`: Your Path of Exile account email (used for login).
- `RESPOE_PASSWORD`: Your Path of Exile account password (used for login).
- `RESPOE_NEW_PASSWORD`: New password for your account (used for change password).
- `RESPOE_RECIPIENT`: Recipient for sending private messages (used for private message sending).

**This variables are only needed if you want to run the tests. These are not needed for regular usage**

## Example of usage

Currently I developed this library mainly to parse [Path Of Exile forums](https://www.pathofexile.com/forum) to be able to create
a [mobile-friendly forum](https://respoe.xyz). With the trading package I created [Currency status](https://respoe.xyz/currency/chaos) to view currency prices compared to chaos

## License

**Respoe** is licensed under the **MIT** license.
