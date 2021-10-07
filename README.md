# [covid-tracking-app](https://covid-tracking-apps.herokuapp.com)

This repository contains the code of a covid tracking app that targets the US states and show the data of covid-19 on Google Map. The data will contain positive and negative cases, hospitalized and recovered patients, and the number of deaths.

The data is taken from the [postman.com](https://www.postman.com/) which contans different collections of covid data. These collections contain the data from different countries and their states. Although the data was not just limited to positive, negative cases, etc. It contains a bunch of data that I ignored and just took some of them. 

This app uses the Json Web Token(JWT) to authenticate a user. It generates, stores and deletes a token in a cookie. This app also uses Redis and MongoDB database which stores the data. Redis stores the data in a cache and MongoDB stores the data in different servers like AWS, Azure, Google Cloud.

This software is written in Golang, HTML, and CSS. It contains four directories in which Go code, Database, HTML and CSS templates are written separately. Those directories are.

    database. It handles the database stuff and I used MongoDB database in this code.
    handler. It contains many functions like home, signup, login, about, etc for performing various tasks.
    templates. It contains all the HTML files.
    static. It contains all CSS files to design the pages.
