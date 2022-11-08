CREATE TABLE users(

    id SERIAL PRIMARY KEY NOT NULL,

    user_name TEXT NOT NULL,

    password TEXT NOT NULL

);



CREATE TABLE books(

    id SERIAL PRIMARY KEY NOT NULL,

    title TEXT NOT NULL,

    author TEXT NOT NULL,

    description TEXT NOT NULL

);