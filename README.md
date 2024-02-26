# React-project
Keyspace is my_keyspace
CREATE TABLE user (user_id INT PRIMARY KEY,username TEXT,email TEXT);
user_id | email             | username
---------+-------------------+----------
       1 | user1@example.com |    user1
       2 | user2@example.com |    user2
       3 | user3@example.com |    user3

CREATE TABLE product (product_id INT PRIMARY KEY,product_name TEXT);
 product_id | product_name
------------+--------------
        102 |      course2
        101 |      course1
        103 |      course3
CREATE TABLE user_stats (id INT PRIMARY KEY,user_id INT,product_id INT,time_taken BIGINT);

 id   | product_id | time_taken | user_id
------+------------+------------+---------
 1001 |        101 |        500 |       1
 1003 |        103 |        766 |       3
 1002 |        102 |        605 |       2

create table marks)id int primary key,marks int,name text);
 id | marks | name
----+-------+------
  5 |    90 |  eee
  1 |    50 |  aaa
  2 |    70 |  bbb
  4 |    40 |  ddd
  7 |    79 |  ggg
  6 |    50 |  fff
  3 |    90 |  ccc
