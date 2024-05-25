## mysqldump

- By default, functions/procedures are not dumped by `mysqldump`. Use `--routines` to include them. [docs](https://dev.mysql.com/doc/refman/8.0/en/mysqldump-stored-programs.html)
- Interestingly the trigger didn't fail even though the function was missing. It just created the trigger without the function.
   - No error, successful dump in the new instance.
   - Failure when the trigger was fired, since the function was missing.
   - Seems like the trigger body is not validated during creation, since I was also set the value to a string for id column.

### Commands 

```bash
# Dump the database
docker exec -it test-mysql-master sh -c 'exec mysqldump --routines -uroot insta1' > insta1_dump_with_functions.sql

# New instance 
docker run -d --name test-mysql-dump-func -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -p 3311:3306 mysql

# Copy the dump file to the new instance
docker cp insta1_dump_with_functions.sql test-mysql-dump-func:/tmp/insta1.sql

# Connect to new instance to create database
mysql --host=127.0.0.1 --port=3311

# Load the dump

docker exec -it test-mysql-dump-func sh -c 'exec mysql -uroot insta1 < /tmp/insta1.sql'
```

- Tried inserting a new record that worked using the trigger with the function to create snowflake id.
- The dump file is available [here](./insta1_dump_with_functions.sql)
