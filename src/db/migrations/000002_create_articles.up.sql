CREATE TABLE IF NOT EXISTS
    go_database.articles(
        id serial,
        title VARCHAR(255) NOT NULL,
        content VARCHAR(255) NOT NULL,
        is_paid boolean,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        user_id BIGINT UNSIGNED,
        PRIMARY KEY(id),
        CONSTRAINT fk_user_id
        FOREIGN KEY (user_id) 
        REFERENCES users (id)
        ON DELETE CASCADE 
        ON UPDATE CASCADE
    ) ENGINE=INNODB DEFAULT CHARSET=utf8;