# fib

1. install golang and psql on Linux env

2. change the psql password to "abc123" for user "postgres"

3. create database "testdb", creat table "newfib", this table has two fields "mykey" and "value"    mykey is the Fib input number, value is the Fibonacci value.

4. run "make fib" to start Fibnacci server

5. run "make test" to run test cases, it will do test twice, we can see the time difference.

Notice: I create a docker file. Since bring both psql server and fib server to docker env will ccause some extra work, we will discuss this in the interview process.

