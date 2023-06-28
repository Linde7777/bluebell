# bluebell
The backend of a online forum, imitating Reddit

The code are constructed in such layer: Controller -> Logic -> DAO

## Environment
Go 1.20.4  
Redis 7.0.11  
MySQL 8.0

## Deployment
1. Go to settings/config.yaml, modify the configuration as you need
2. `redis-server` (or `redis-server.exe` in Windows)
3. `go build -o main`
4. `./main`

## Todo
1. In voting part, if the insertion into MySQL is successful but the insertion 
into Redis fails, the MySQL insertion should be rolled back.

2. Use refresh-token to get access-token