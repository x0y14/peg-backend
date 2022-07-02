### chatroomの種類。
- 1-1はDirectChat
- N-N?(討論会)はGroupChat

### idの種類
- User(string)
  - firebase UID(20~128?)  = 128(?)

- Operation(BIGINT,int64)
  - snowflake


- Message(string)
  - "ms|{xid}"
  - 2(ms) + 1(|) + 20(xid) = 23 -> max 23

- DirectChat(string)
  - "di|{ソートしたuser_idをドットで繋ぐ}
  - 2(di) + 1(|) + 128(uid) + 1(.) + 128(uid) = 260
- Group(string)
  - "gr|{xid}"
  - 2(gr) + 1(|) + 20(xid) = 23

### databaseについて
- operation.destination, friends, group_memberは1-1のテーブルにする。