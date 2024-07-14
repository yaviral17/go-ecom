# About Project
**This is a simple project to demonstrate implementation of auth using JWT on GoLang and having mysql as database.**

# How to run
- Clone the project
- Set the environment variables in **`.env`** file
    > - **`PUBLIC_HOST`** - The host on which the server will run
    > - **`PORT`** - The port on which the server will run
    > - **`DB_USER`** - The username of the database
    > - **`DB_PASS`** - The password of the database
    > - **`DB_NAME`** - The name of the database
    > - **`DB_ADDR`** - The host of the database
- Setup database by running **`dbsetup.sql`** commands in the database
- Run **`make run`** to start the server
- Run **`make test`** to run the tests 
- Run **`make build`** to build the project