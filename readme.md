# Bookrest MongoDB Based REST-Api

## Configuration
Configuration is stored on the ```config.toml``` file,setup your MongoDB connection string, ports, database name, and books collection name

```
Port = 39822
Database = "bookrest"
DatabaseConnectionString = "mongodb://localhost:27017/admin"
BooksCollection = "book"
```

## Starting up the server
```
go run main.go
```

## Endpoints
### AllBooksEndpoint (GET)
#### Path
```
/books
```
#### Command
```
curl http://bookrest/books
```
#### Response
```
[
   {
      "ISBN":"1477246444",
      "Title":"What Is Sql ?: Fundamentals Of Sql,t-sql,pl/sql And Datawarehousing.",
      "Author":"Victor Ebai",
      "Status":"available",
      "Publisher":"Authorhouseuk",
      "Year":2019,
      "IssueCount":60,
      "ViewCount":424,
      "CurrentIssuer":""
   },
   {
      "ISBN":"147724644433",
      "Title":"Golang and The Future",
      "Author":"Victor Ebai",
      "Status":"available",
      "Publisher":"Authorhouseuk",
      "Year":2019,
      "IssueCount":422,
      "ViewCount":3000,
      "CurrentIssuer":""
   },
   {
      "ISBN":"303",
      "Title":"Yet Another Book",
      "Author":"Victor Ebai",
      "Status":"available",
      "Publisher":"Authorhouseuk",
      "Year":2019,
      "IssueCount":1,
      "ViewCount":1,
      "CurrentIssuer":""
   }
]
```
This will return all existing books in the database

### CreateBookEndpoint (POST)
#### Path
```
/books
```
#### Payload
```
{
    "title": "Yet Another Book 2",
    "author": "Victor Ebai",
    "isbn": "3023",
    "status": "available",
    "publisher": "Authorhouseuk",
    "year": 2019
}
```
#### Command
```
curl --location --request POST "bookrest/books" --header "Content-Type: application/json" --data "{\"title\":\"Yet Another Book 2\",\"author\":\"Victor Ebai\",\"isbn\":\"3023\",\"status\":\"\",\"publisher\":\"Authorhouseuk\",\"year\":2019}"
```
#### Response
```
200 OK
```
This will create a new book in the database

### TopBookEndpoint (GET)
#### Path
```
/books/top
```
#### Command
```
curl http://bookrest/books/top
```
#### Response
```
[ 
   { 
      "ISBN":"147724644433",
      "Title":"Golang and The Future",
      "Author":"Victor Ebai",
      "Status":"available",
      "Publisher":"Authorhouseuk",
      "Year":2019,
      "IssueCount":422,
      "ViewCount":3000,
      "CurrentIssuer":""
   }
]
```
This will return the book with the most views

### MostIssuedBookEndpoint (GET)
#### Path
```
/books/mostissued
```
#### Command
```
curl http://bookrest/books/mostissued
```
#### Response
```
[ 
   { 
      "ISBN":"147724644433",
      "Title":"Golang and The Future",
      "Author":"Victor Ebai",
      "Status":"available",
      "Publisher":"Authorhouseuk",
      "Year":2019,
      "IssueCount":422,
      "ViewCount":3000,
      "CurrentIssuer":""
   }
]
```
This will return the most issued book


### FindBookEndpoint (GET)
#### Path
```
/books/{isbn}
```
#### Command
```
curl http://bookrest/books/147724644433
```
#### Response
```
[ 
   { 
      "ISBN":"147724644433",
      "Title":"Golang and The Future",
      "Author":"Victor Ebai",
      "Status":"available",
      "Publisher":"Authorhouseuk",
      "Year":2019,
      "IssueCount":422,
      "ViewCount":3000,
      "CurrentIssuer":""
   }
]
```

This will find the book based on the ISBN

### IssueEndPoint (PUT)
#### Path
```
/issue/{isbn}
```
#### Payload
```
{"issuer" : "reaganiwadha"}
```
#### Command
```
curl --location --request PUT "localhost:39822/issue/147724644433" --header "Content-Type: application/json" --data "{\"issuer\": \"reaganiwadha\"}"
```
#### Response (If book is available)
```
{"ISBN":"147724644433","Title":"Golang and The Future","Author":"Victor Ebai","Status":"issued","Publisher":"Authorhouseuk","Year":2019,"IssueCount":423,"ViewCount":3000,"CurrentIssuer":"reaganiwadha"}
```
#### Response (If book is already issued)
```
{"error":"Book already issued"}
```

This will set the current issuer of the book and will increment the issuer count

### DeleteIssuerEndpoint
#### Path
```
/issue/{isbn}
```
#### Command
```
curl -X DELETE bookrest/issue/147724644433
```
#### Response
```
200 OK
```
This will empty the current issuer of the book and make the book available

## Sorry for writing this garbage, i've just learned go recently