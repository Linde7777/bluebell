# bluebell
The backend part of a online forum, imitating Reddit

Controller -> Logic -> DAO

## Environment
Go 1.20.4 windows/amd64
Redis 7.0.11  
MySQL 8.0

## Todo
1. In voting part, if the insertion into MySQL is successful but the insertion 
into Redis fails, the MySQL insertion should be rolled back.

2. Use refresh-token to get access-token