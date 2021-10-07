# [covid-tracking-app](https://covid-tracking-apps.herokuapp.com)

![Screenshot 2021-10-07 at 20-53-46 Main](https://user-images.githubusercontent.com/64713734/136428321-5ec128d6-39b2-4fb4-afe5-e30863f637d6.png)

## Intro

This repository contains the code of a covid tracking app that targets the US states and show the data of covid-19 on Google Map. The data will contain positive and negative cases, hospitalized and recovered patients, and the number of deaths.

The data is taken from the [postman.com](https://www.postman.com/) which contans different collections of covid data. These collections contain the data from different countries and their states. Although the data was not just limited to positive, negative cases, etc. It contains a bunch of other data that I ignored and just took some of them. 

This app uses the JSON Web Token(JWT) to authenticate a user. It generates, stores and deletes a token in a cookie. This app also uses Redis and MongoDB database which stores the data. Redis stores the data in a cache and MongoDB stores the data in different servers like AWS, Azure, and Google Cloud.

---

## App Structure

This software is written in Golang, HTML, and CSS. The directory structure is based on **MVC model**. The software contains seven directories in which Go code, Pandemic data, Database, HTML and CSS templates are written separately. These directories are.

- **controllers**: This directory contains two files. `controllers.go` and `handlers.go`. These files handle the main functionality of different pages, like calling the templates, generating the JWT tokens, refreshing them, converting passwords into hash values, etc.
    
- **database**: This directory contains two files. `db.go` and `redis.go`. The first file contains the code of a MongoDB database that will connect, insert and find the data. The second file contains the code of redis(It stores the values in memory) that will set, get and delete a value.
    
- **middleware**: This directory has one file `middleware.go`. It handles the errors in most of the functions.
   
- **models**: This directory has one file `models.go`. It contains a structure which has the user credentials.

- **pandemic**: This directory has two files. `pandemic.go` and `pandemic.csv`. The first file stores the covid data in a CSV file. The second file contains all the data.

- **routes**: This directory has one file `routes.go`. It handles the paths that should be visited and uses middleware to check for errors.

- **views**: This directory contains two subdirectories **templates** and **static**. The first subdirectory contains all the HTML files. The second subdirectory contains all the CSS files and images also. 

## Author Info

- YouTube - [ibilalkayy](https://www.youtube.com/channel/UCBLTfRg0Rgm4FtXkvql7DRQ)
- LinkedIn - [ibilalkayy](https://www.linkedin.com/in/ibilalkayy/)
- Twitter - [ibilalkayy](https://twitter.com/ibilalkayy)

[Back to Top](#covid-tracking-app)
