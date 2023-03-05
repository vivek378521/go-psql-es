# go-psql-es

## Setting up the repo

1. Clone the repo
2. `go get .` download all packages
3. Setup the .env file
4. Set up the postgres db (https://medium.com/coding-blocks/creating-user-database-and-adding-access-on-postgresql-8bfcd2f4a91e)
5. Set up ES, currently I have used a cloud hosted ES instance but you will have to replace some code to make the repo adjust to local ES instance (es installation: https://www.elastic.co/guide/en/elasticsearch/reference/current/install-elasticsearch.html)
6. After you have setup everything run `go run main.go`
7. Export the thunder collection json (added in repo) and call the APIs

