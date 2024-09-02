## migration

本文件夹文件名示例：
```
|── migration
|   |── 1_init.sql
|   |── 2_update.sql
|   └── 3_xxx.sql
```

会按照数字顺序依次检查是否已经执行

理论上，表结构修改时，只能通过增加`.sql`文件达到目的。不能删除`.sql`文件


