# Groupie Tracker


## Table of Contents
1. [Description](#description)
2. [Authors](#authors)
3. [Usage:](#usage)
4. [Implementation details: algorithm](#implementation-details-algorithm)
5. [Testing](#testing)

### Description:
***
Hi *Talent*!
Groupie Trackers consists in creating and running a server which, based on the usage of API given and display informations after manipulating JSON files.
This project's server (back-end) is written using GO and the web documents (front-end) using HTML & CSS

#### Go,Html,Css:
***
- **GO**, also called Golang or Go language, is an open source programming language that Google developed.
* **HTML**,The HyperText Markup Language is the standard markup language for documents designed to be displayed in at web browser.
+ **CSS**,Cascading Style Sheets, form a computer language that describes the presentation of HTML documents.

### Authors:
***
- Vincent Félix Ndour (Writer and test-file) -README/TEST_FILE- https://learn.zone01dakar.sn/git/vindour
* Masseck Thiaw (the genius and our captain) -BACKEND- https://learn.zone01dakar.sn/git/mthiaw
+ Seynabou Niang (The artist Michel-Ange) -FRONTEND- https://learn.zone01dakar.sn/git/sniang

### Usage:
***
A little intro about the installation.
```
$ git clone https://learn.zone01dakar.sn/git/mthiaw/groupie-tracker
$ cd groupie-tracker
$ go run main.go
```

[def]: #usage-how-to-run

## Implementation-details-algorithm:
***
```
Our program define structs that helps to store datas from the API after retrieving them by using the get_link function and according to user's request transmits this data to the concerned function (eg: Api_artists, Artists, Api_dates or Api_locations) then send it to the appropriate handler.
```
- **func Api_artists** Get all informations of the API's artist part from "https://groupietrackers.herokuapp.com/api/artists"

- **func api_locations** Gets datas from API's locations part --> "https://groupietrackers.herokuapp.com/api/locations".

+ **func api_dates** Gets datas from API's dates part --> "https://groupietrackers.herokuapp.com/api/dates" .

* **func api_relation** Gets datas from API's dates part --> "https://groupietrackers.herokuapp.com/api/dates" returns relations.

## Testing :
***
```
*** For the unit test file two file are used one that helps to test the API with using the Rest API extension and the other for our programme.

- NB: For unit testing of the API the Rest Client extension must be installed on your PC.