# Respoe

Path of Exile forum client. Useful for automating tasks on the forum such as bots or gathering data.

## Client

- [x] Login
- [x] Logout
- [x] Change password

## Forum

- [x] List with stats
- [x] Retrieve specific forum pagination information (first, current and last page)
- [x] Retrieve all threads from a forum (with status, views, author, creation date and replies)
- [x] Retrieve all posts from a thread (content, author and creation date)

### Testing

To run the tests you need to set some environment variables:

- `RESPOE_EMAIL`: Your Path of Exile account email (used for login).
- `RESPOE_PASSWORD`: Your Path of Exile account password (used for login).
- `RESPOE_NEW_PASSWORD`: New password for your account (used for change password).