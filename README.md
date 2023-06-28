# bluebell
imitate reddit

Controller -> Logic -> DAO

## Environment
Go 1.20.4 windows/amd64
Redis 7.0.11  
MySQL 8.0

## todo
In voting, if the insertion into MySQL is successful but the insertion 
into Redis fails, the MySQL insertion should be rolled back.

use refresh-token to get access-token