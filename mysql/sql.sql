1)SHOW CHARACTER SET;

2)show collection;

3)CREATE DATABASE db_name CHARACTER SET latin1 COLLATE latin1_swedish_ci;

4)CREATE TABLE tbl1;

5)CREATE TABLE tbl1 (....) DEFAULT CHARSET=utf8 DEFAULT COLLATE=utf8_bin;

6)CREATE TABLE tbl1;

7)CREATE TABLE tbl1 (col1 VARCHAR(5) CHARACTER SET latin1 COLLATE latin1_german1_ci);

8)
-- 查看字符集
show [global] variables like 'character%';
show [global] variables like 'collation%';

-- 修改字符集
set global character_set_server=utf8; -- 全局
alter table xxx convert to character set xxx; -- 表

9)set names utf8;

10)create table `stu1`(
`name` varchar(50) not null default '',
`course` varchar(50) default null,
`score` int(11) default null,
primary key(`name`)
) engine=innodb default charset=latin1;

11)insert into stu1 values('pw','你好',44);(报错)

12)create table `stu2`(
`name` varchar(50) not null default '',
`course` varchar(50) default null,
`score` int(11) default null,
primary key(`name`)
) engine=innodb default charset=utf8;

13)回话连接设为utf8 & set names utf8;

14)insert into stu2 values('pw','你好',33);(正确)

15)select * from stu2;

16)set names gbk;

17)insert into stu2 values('pw2','你好吗',33);(报错)

18)insert into stu2 values('pw2','你好',33);(成功执行，存储错误)

19)set names utf8;

20)select * from stu2;

21)回话连接设为gbk & set names gbk;

22)insert into stu2 values('pw3','你好',33);(正确)

23)select * from stu2;

24)select length(course) from stu2;

25)create table `tb1`(

`name` varchar(50) not null,

) engine=innodb default charset=gbk;

26)
show global varibales like '%char%';
set character_set_database=gbk;
load data infile 'XXXX'（文件路径）;(文件编码是utf8)
select * from tb1;

27)
delete from tb1;
show varibales like '%char%';
set character_set_database=utf8;
load data infile 'XXXX'（文件路径）;(文件编码是utf8)
select * from tb1;