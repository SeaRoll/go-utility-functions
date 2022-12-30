# Golang utility functions

The folders include some methods / applications that I would like to use
if I ever would implemenet my own backend in go. These are helper functions
such as SQL Migrations, JWT auth etc.

### /migrator

A Migration library that takes a `*sql.DB` and creates migrations
for the specific database. the features are almost the same as flyway migrations
by sorting by versions given. Example of the sql folder can be found in `sql` folder.

Note: the sql folder needs to be in the same directory as the binary file. There is
also a Dockerfile if you would ever want to try to make it work.
