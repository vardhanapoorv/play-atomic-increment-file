## Database function for Snowflake generation

Few things to note: 

- `READS SQL DATA` - This was needed because my MYSQL instance (pun intended) was running in strict mode, binary logging enabled.
- Delimiter change because default is `;` and we are using `;` in the function body.

```sql

delimiter //

//
CREATE FUNCTION insta1.next_id()
RETURNS BIGINT
READS SQL DATA
BEGIN
    DECLARE epoch BIGINT DEFAULT 1314220021721;
    DECLARE seq_id BIGINT;
    DECLARE now_ms BIGINT;
    DECLARE shard_id INT DEFAULT 1;
    DECLARE result BIGINT;

    -- Get the next sequence value
    SET seq_id = insta1.nextval() % 1024;

    -- Get the current timestamp in milliseconds
    SET now_ms = ROUND(UNIX_TIMESTAMP(NOW(3)) * 1000);

    -- Calculate the result
    SET result = (now_ms - epoch) << 23;
    SET result = result | (shard_id << 10);
    SET result = result | (seq_id);

    RETURN result;
END //
```

- Needed to create another function for sequence since mysql doesn't support sequences out of the box.
- For that we need to create a table to store the sequence values.
- `LAST_INSERT_ID()` - Last inserted value for a auto increment column [docs](https://dev.mysql.com/doc/refman/8.0/en/information-functions.html#function_last-insert-id).

```sql

CREATE TABLE insta1.table_id_seq (
    id BIGINT NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (id)
);

INSERT INTO insta1.table_id_seq VALUES (NULL);

delimiter //

CREATE FUNCTION insta1.nextval()
RETURNS BIGINT
READS SQL DATA
BEGIN
    INSERT INTO insta1.table_id_seq VALUES (NULL);
    RETURN LAST_INSERT_ID();
END //
```

- MYSQL doesn't support of default values coming from a function. Hence we need to use a trigger.
- MYSQL 8 added support for expressions and constants(built-in) for default values but not custom functions. [docs](https://dev.mysql.com/doc/refman/8.0/en/data-type-defaults.html) , [link](https://stackoverflow.com/questions/270309/can-i-use-a-function-for-a-default-value-in-mysql)

```sql
delimiter //

CREATE TRIGGER insta1.before_insert_posts
BEFORE INSERT ON insta1.posts
FOR EACH ROW
BEGIN
    IF NEW.post_id IS NULL THEN
        SET NEW.post_id = insta1.next_id();
    END IF;
END //
```

- Insert a row into the table and see the post_id being generated.
