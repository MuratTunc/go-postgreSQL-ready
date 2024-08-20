# BASIC POSTGRESQL APPLICATION VIA GOLANG

In this study, I will try to introduce the simplest postgreSQL applicaiton using the Golang programming language. Intermediate level knowledge of Golang language will be sufficient.

> ## Working environment
>
> - Debian Operating System
> - Golang sw language
> - PostgreSQL
> - DBeaver
> - VS Code 

## General hierarchy of the project
![hierarchy.PNG](pictures/hierarchy.PNG)


## Dbeaver database table creation

### Step 1
![dbeaver1.png](pictures/dbeaver1.png)

### Step 2
![dbeaver2.png](pictures/dbeaver2.png)

### Step 3
![dbeaver3.png](pictures/dbeaver3.png)

### Step 4
![dbeaver4.png](pictures/dbeaver4.png)

### Step 5
![dbeaver5.png](pictures/dbeaver5.png)

### Step 6
![dbeaver6.png](pictures/dbeaver6.png)

### Step 7
![dbeaver7.png](pictures/dbeaver7.png)

### Step 8
![dbeaver8.png](pictures/dbeaver8.png)

### Step 9
![dbeaver9.png](pictures/dbeaver9.png)

### Step 10
![dbeaver10.png](pictures/dbeaver10.png)

### Step 11
![dbeaver11.png](pictures/dbeaver11.png)

### Step 12
![dbeaver12.png](pictures/dbeaver12.png)

### Step 13
![dbeaver13.png](pictures/dbeaver13.png)


Now our Database table is ready !
Lets insert some values to it via Golang.



## VS code
```
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
	host     = "localhost"
	port     = 5432
)

func main() {
	// Connection to database
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// Open the database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected!")
	// Close the database
	defer db.Close()
	insert := `insert into "icecream"("Name", "Price") values($1, $2)`
	// insert new raw to database table
	_, err = db.Exec(insert, "Vanilla", 4)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(insert, "Chocolate", 5)
	if err != nil {
		panic(err)
	}
	fmt.Println("Writing to database completed!")
}
```

## Run go file
![vs_code_run.png](pictures/vs_code_run.png)

## Database table result
![dbeaver14.png](pictures/dbeaver14.png)



## Github source code

[github_source_code](https://github.com/MuratTunc/go-postgreSQL-ready.git)

__**I wish you good coding, thinking that it is educational and fun.**__
