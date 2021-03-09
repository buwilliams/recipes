# Recipes

Experimenting with [Go](https://golang.org) and [Dolt](https://www.dolthub.com/), a version controlled database.

## Why do I find dolt interesting?

- Managing SQL databases is always a pain. Dolt provides a better UX.
- Therefore, prototyping ideas becomes easier, especially since Dolt supports CSV import. This allows me to gradually add features to the database.
- For smaller projects, performance isn't the most important thing.
- Having access to public databases through DoltHub is handy

## Getting Started

- Install dolt `sudo curl -L https://github.com/dolthub/dolt/releases/latest/download/install.sh | sudo bash`
- Clone dolt repo `dolt clone buwilliams/twitter-db` (will become recipes later)
- Start mysql server `dolt sql-server`
- Clone recipes repo `https://github.com/buwilliams/twitter-clone.git`
- Install deps `go get`
- Start Gin server `go run .`

## API

- GET [/user](http://localhost:3001/user)
- GET [/tweet](http://localhost:3001/tweet)

## Links

- [original hacker news post](https://news.ycombinator.com/item?id=22731928)
- [twitter-db dolt repo](https://www.dolthub.com/repositories/buwilliams/twitter-db)