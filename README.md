The Gator Blog Aggregator will require that Postgress and Go are installed on your machine.

Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local Postgres database. You can then install gator with:

go install ...


Create a .gatorconfig.json file in your home directory with the following structure:

{
  "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}


Replace the values with your database connection string.


Create a new user:


gator register <name>


Add a feed:


gator addfeed <url>


Start the aggregator:


gator agg 30s

There are a few other commands you'll need as well:

- `gator login <name>` - Log in as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds


