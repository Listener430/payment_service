# Payment Notification Service


## Run 
```
make docker-up
```
## Testing 

1. **Check if service is running**
```
curl http://localhost:8080
```
You should get an empty response or not found.

2. **Send Apple Notification**
```
curl -X POST -H "Content-Type: application/json" \
-d {"receiptData":"someData","userId":"bob"} \
http://localhost:8080/apple
```
If all is well, you get "unlocked".

3. **Check DB**
