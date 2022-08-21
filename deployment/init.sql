CREATE EXTENSION pgcrypto;

CREATE TABLE IF NOT EXISTS users (
id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY NOT NULL,
username TEXT NOT NULL UNIQUE,
role TEXT NOT NULL,
status TEXT NOT NULL,
chat_id TEXT NOT NULL,
created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE IF NOT EXISTS sessions (
id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY NOT NULL,
user_id UUID REFERENCES users(id),
created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE IF NOT EXISTS codes (
id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY NOT NULL,
code TEXT,
attempts INT,
user_id UUID REFERENCES users(id),
created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE IF NOT EXISTS categories (
id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY NOT NULL,
name TEXT NOT NULL UNIQUE,
created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP);

CREATE TABLE IF NOT EXISTS items (
id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY NOT NULL,
title TEXT NOT NULL,
description TEXT NOT NULL,
status TEXT NOT NULL,
category TEXT REFERENCES categories(name),
image_link TEXT UNIQUE,
raw_price integer,
start_price integer,
owner_id UUID REFERENCES users(id),
current_price integer,
created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP);
