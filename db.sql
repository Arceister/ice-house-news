CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY ,
    email VARCHAR(60) NOT NULL UNIQUE ,
    password VARCHAR(255) NOT NULL ,
    name VARCHAR(60) NOT NULL ,
    bio VARCHAR(200) ,
    web VARCHAR(50) ,
    picture VARCHAR (255)
);

CREATE TABLE IF NOT EXISTS categories (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY ,
    name VARCHAR(30) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS news (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY ,
    users_id uuid REFERENCES users(id) ,
    category_id uuid REFERENCES categories(id) ,
    title VARCHAR(80) NOT NULL UNIQUE ,
    isi TEXT NOT NULL ,
    slug_url VARCHAR(200) NOT NULL UNIQUE ,
    cover_image VARCHAR(255) NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    nsfw BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS news_counter (
    news_id uuid REFERENCES news(id) ON DELETE CASCADE ,
    upvote INTEGER NOT NULL DEFAULT 0 ,
    downvote INTEGER NOT NULL DEFAULT 0 ,
    comment INTEGER NOT NULL DEFAULT 0 ,
    view INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS news_additional_images (
    news_id uuid REFERENCES news(id) ON DELETE CASCADE ,
    image VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS news_comment (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY ,
    news_id uuid REFERENCES news(id) ON DELETE CASCADE ,
    users_id uuid REFERENCES users(id) ON DELETE CASCADE ,
    description VARCHAR(200) NOT NULL ,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);