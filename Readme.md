## steps for auto generate gql code
    go run github.com/99designs/gqlgen init

    go get github.com/99designs/gqlgen@v0.17.39

    go run github.com/99designs/gqlgen generate

## Mongo operation

[mongo_db_url](https://www.geeksforgeeks.org/mongodb-insertone-method-db-collection-insertone/?ref=lbp)
    
    mongo "mongourl"
    show dbs;
    use db_name
    show collections
    