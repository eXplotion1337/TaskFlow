CREATE TABLE tasks (
    "id" INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
    "uid" TEXT DEFAULT '' NOT NULL,
    "name" TEXT DEFAULT '' NOT NULL,
    "description" TEXT DEFAULT '' NOT NULL,
    "user" TEXT DEFAULT '' NOT NULL,
    "time_start" TEXT DEFAULT '' NOT NULL,
    "time_end" TEXT DEFAULT '' NOT NULL,
    "creator" "time_start" TEXT DEFAULT '' NOT NULL,
    "project" TEXT DEFAULT '' NOT NULL
);

CREATE TABLE users (
   "id" INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
   "login" TEXT DEFAULT '' NOT NULL,
   "password" TEXT DEFAULT '' NOT NULL,
   "token" TEXT DEFAULT '' NOT NULL
);

CREATE TABLE projects (
   "id" INTEGER PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
   "name" TEXT DEFAULT '' NOT NULL,
   "description" TEXT DEFAULT '' NOT NULL,
   "collaborators" TEXT DEFAULT '' NOT NULL,
   "token" TEXT DEFAULT '' NOT NULL
);
