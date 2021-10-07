# [covid-tracking-app](https://covid-tracking-apps.herokuapp.com)

This repository contains the code of a covid tracking app that targets the US states and show the data of covid-19 on Google Map. The data will contain positive and negative cases, hospitalized and recovered patients, and the number of deaths.

The data is taken from the [postman.com](https://www.postman.com/) which contans different collections of covid data. These collections contain the data from different countries and their states. Although the data was not just limited to positive, negative cases, etc. It contains a bunch of data that I ignored and just took some of them. 

This app uses the Json Web Token(JWT) to authenticate a user. It generates, stores and deletes a token in a cookie. This app also uses Redis and MongoDB database which stores the data. Redis stores the data in a cache and MongoDB stores the data in different servers like AWS, Azure, Google Cloud.

This software is written in Golang, HTML, and CSS. The directory structure is based on **MVC model**. The software contains seven directories in which Go code, Pandemic data, Database, HTML and CSS templates are written separately. These directories are.

**1. controllers**: This directory contains two files. `controllers.go` and `handlers.go`. These files handle the main functionality of different pages, like calling the templates, generating the JWT tokens, refreshing them, convert passwords into hash values, etc.
    
**2. database**: This directory contains two files. `db.go` and `redis.go`. The first file contains the code of a MongoDB database that will connect, insert and find the data. The second file contains the code of redis(It stores the values in memory) that will set, get and delete a value.
    
**3. middleware**: This directory has one file `middleware.go`. It contains handles the errors in most of the functions in this app.
   
**4. models**: This directory has one file `models.go`. It contains a structure which has the user credentials.

**5. pandemic**: This directory has two files. `pandemic.go` and `pandemic.csv`. The first file stores the covid data in a CSV file. The second file contains all the data.

**6. routes**: This directory has one file `routes.go`. It handles the paths that should be visited and uses middleware to check for errors.

**7. views**: This directory contains two subdirectories **tempalates** and **static**. The first subdirectory contains all the HTML files. The second subdirectory contains all the CSS files and images also. 
