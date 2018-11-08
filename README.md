## Example request
```
curl -X POST \
-H "Content-Type: application/json" \
-d '{"jsonrpc": "1.0", "method": "HelloService.SayHi", "params": [{"Who":"Joe"}], "id": 1}' \
"http://localhost:8888/rpc"

{"result":{"Message":"Hi, Joe!"},"error":null,"id":1}
```