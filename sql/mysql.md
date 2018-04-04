* 展示数据表字符集
```sql
show create table t_name;
```

* 数据库字符集
```sql
show create database d_name;
```

+ 创建数据库
```sql
CREATE DATABASE IF NOT EXISTS arvinDB DEFAULT CHARSET utf8 COLLATE utf8_general_ci
```