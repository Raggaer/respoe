# Respoe

Path of Exile forum client. Useful for automating tasks on the Path of Exile website or forum (bots or just gathering data).

It makes requests and then parses the response to fetch all the needed data. The library is separated in three packages:

- Client: Related to all account features (login, logout, private messages, ...).
- Forum: Related to all forum data (threads, posts, replying, ...).
- Util: Related to util methods for developing tools (get XSRF data, get form errors, ...).

## Usage

First get the package:

`go get github.com/raggaer/respoe/...`

You can then use this as a library for developing tools related to the Path of Exile website / forum.

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

### Forum

[![GoDoc](https://godoc.org/github.com/Raggaer/respoe/forum?status.svg)](http://godoc.org/github.com/Raggaer/respoe/forum)

- [x] List with stats
- [x] Retrieve specific forum pagination information (first, current and last page)
- [x] Retrieve all threads from a forum (with status, views, author, creation date and replies)
- [x] Retrieve all posts from a thread (content, author and creation date)
- [x] Reply to a thread

### Util

[![GoDoc](https://godoc.org/github.com/Raggaer/respoe/util?status.svg)](http://godoc.org/github.com/Raggaer/respoe/util)

- [x] Retrieve pagination from any valid page
- [x] Retrieve hash value from any valid form
- [x] Retrieve hash value from any reply thread form
- [x] Retrieve errors from a submitted form

## Testing

To run the tests you need to set some environment variables:

- `RESPOE_EMAIL`: Your Path of Exile account email (used for login).
- `RESPOE_PASSWORD`: Your Path of Exile account password (used for login).
- `RESPOE_NEW_PASSWORD`: New password for your account (used for change password).
- `RESPOE_RECIPIENT`: Recipient for sending private messages.