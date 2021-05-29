CREATE DATABASE IF NOT EXISTS UserEvents;

USE UserEvents;

# CREATE TABLE IF NOT EXISTS User
# (
#     id        bigint unsigned primary key not null,
#     createdAt DATETIME                    not null,
#     updatedAt DATETIME                    not null
# );

CREATE TABLE IF NOT EXISTS Schedule
(
    userID    bigint unsigned             not null,
    drug      varchar(255)                not null,
    date      DATETIME                    not null,
    createdAt DATETIME                    not null,
    updatedAt DATETIME                    not null
);

# CREATE TABLE IF NOT EXISTS Drugs
# (
#     id        bigint unsigned primary key not null,
#     userID    bigint unsigned             not null,
#     createdAt DATETIME                    not null,
#     updatedAt DATETIME                    not null
# );
