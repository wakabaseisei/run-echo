CREATE TABLE IF NOT EXISTS
    go_database.posts(
        id serial,
        title VARCHAR(255) NOT NULL,
        content VARCHAR(255) NOT NULL,
        publish_date int NOT NULL,
        thumbnail_url VARCHAR(255),
        user_id bigint UNSIGNED NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY(id),
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
    ) ENGINE=INNODB DEFAULT CHARSET=utf8;