# MyFitnessTomodachi-api
The backend API for MyFitnessTomodachi, a calorie calculator app inspired by MyFitnessPal (obviously)

## Table of Contents
- [MyFitnessTomodachi-api](#myfitnesstomodachi-api)
  - [Table of Contents](#table-of-contents)
  - [Technologies Used](#technologies-used)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Credit](#credit)

## Technologies Used
1. MySQL
2. Golang
3. Railway

## Installation
0. Ensure you have [Docker and Docker Compose](https://docs.docker.com/engine/install/ubuntu/) installed.

1. Clone the repo with:
```
$ git clone git@github.com:wrewsama/MyFitnessTomodachi-api.git
```

2. Navigate to the repo with
```
$ cd MyFitnessTomodachi-api
```

3. Create a file named `.env` and paste in the following credentials:
```
SQLUSER=root
PASSWORD=pwd
HOSTNAME=127.0.0.1
DBPORT=3306
DBNAME=mft
```

4. Start the database and the api with:
```
$ docker compose up
```

5. Run the following command to set up the schema:
```
$ go run migrate/migrate.go
```

6. Then stop the containers with:
```
$ docker compose down
```
_The containers need to be restarted for the api to start properly_

## Usage
Run
```
$ docker compose up
```

## Credit
Andrew Lo Zhi Sheng 

[![Github Badge](https://img.shields.io/badge/GitHub-100000?style=for-the-badge&logo=github&logoColor=white)](https://github.com/wrewsama)
[![LinkedIn Badge](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/andrewlozhisheng/)