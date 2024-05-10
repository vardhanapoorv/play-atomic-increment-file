## UUID Benchmark

- Created 3 tables (`ID`, `age`)-
  - UUID VARCHAR(36)
  - UUID BINARY(16)
  - Integer (auto increment)

- Inserted 1 million rows into each table, got similar insertion times. UUIDBINARY(16) was slightly slower by 40seconds.

![Insertion Time](./uuid-insertion.png)


- Created index on `age` column for all 3 tables.

- Time taken for table with UUID VARCHAR(36)

![Index creation time on uuid varchar table](./index-creation-on-uuid-varchar.png)


- Time taken for table with UUID BINARY(16)

![Index creation time on uuid binary table](./index-on-uuid-binary.png)


- Time taken for table with Integer (auto increment)

![Index creation time on integer table](./index-on-int-auto-increment.png)


- Wanted to check `index_length` for each table.
- Ran `SHOW TABLE STATUS` on each table.
- `index_length` was showing zero for all tables then ran `ANALYZE TABLE` on each table.
- After that all tables got correct `index_length` value.

Approx `index_length` for each table per record (rough calculation):

- UUID VARCHAR(36) ~ 55 bytes => 36 + 4 + bytes for internal_stuff(eg: pointer)
- UUID BINARY(16) ~ 31 bytes => 16 + 4 + bytes for internal_stuff
- Integer (auto increment) ~ 17 bytes => 4 + 4 + bytes for internal_stuff

![Total index length UUID VARCHAR](./total-index-length-uuid-varchar.png)


![Total index length UUID BINARY](./total-index-length-uuid-binary.png)


![Total index length Integer](./total-index-length-int-auto-increment.png)
