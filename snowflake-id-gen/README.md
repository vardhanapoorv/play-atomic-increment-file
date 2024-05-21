## Snowflake

It is used to generate unique IDs for objects within a distributed system. It is a 64-bit integer which consists of:

- Time
- Sequence number
- Machine ID

Exact details of the number of bits used vary by implementation. Twitter was the first one to introduce this.

- Twitter - https://blog.x.com/engineering/en_us/a/2010/announcing-snowflake
- Sonyflake - OSS library in Go https://github.com/sony/sonyflake
- Discord - https://discord.com/developers/docs/reference#snowflakes

## Interesting facts

- Sonyflake uses 39 bits for time vs 42 bits used by Discord. But still has a longer lifetime than Discord.
- This happens because Sonyflake each unit of time is 10ms vs 1ms in Discord. Hence it can represent a larger time range.
- But in order to get higher lifetime it compromises on the number of IDs it can generate per millisecond as compared to Discord is lower.
- Twitter's implementation used directly epoch time in milliseconds hence the lifetime is lower(69 years) than both Sonyflake and Discord. Both of which use time elpased since a custom epoch.
- None of the implementation use the most significant bit (sign bit). Hence can convert safely from `uint` to `int` without losing any information.
