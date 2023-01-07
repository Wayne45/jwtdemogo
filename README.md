Quick start
-----------

![setup](https://github.com/Wayne45/jwtdemogo/blob/master/jwtdemogo.png)

### Run:
1. `docker-compose build`
2. `docker-compose up -d`

### Test:
1. Run test case on test/request/test1.http
2. Run test case on test/request/test2.http

### API:
1. Generate jwt token:
   ```
   curl --request POST 'localhost:8081/v1/jwt-tokens'
   ```
2. Verify jwt token by go service:
   ```
   curl --request GET 'localhost:8081/v1/jwt-tokens' \
   --header 'Authorization: Bearer {token}'
   ```
3. Verify jwt token by envoy:
   ```
   curl --request GET 'localhost:8081/v1/self' \
   --header 'Authorization: Bearer {token}'
   ```
