CREATE TABLE IF NOT EXISTS
    go_database.follows(
        id serial,
        user_id BIGINT UNSIGNED,
        viewer_id BIGINT UNSIGNED,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY(id),
        FOREIGN KEY(user_id) REFERENCES users(id),
        FOREIGN KEY(viewer_id) REFERENCES viewers(id)
    ) ENGINE=INNODB DEFAULT CHARSET=utf8;