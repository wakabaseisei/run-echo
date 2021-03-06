CREATE TABLE IF NOT EXISTS
    go_database.users(
        id serial,
        sex int NOT NULL,
        introduction VARCHAR(255) NOT NULL,
        is_forum_need_paid boolean,
        billing_price int NOT NULL DEFAULT 0,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY(id)
    ) ENGINE=INNODB DEFAULT CHARSET=utf8;