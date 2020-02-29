# arvan
final commit

```
localhost:8080 wallet service
localhost:8081 discount service
```

```
localhost:8080/get-balance get balance of user with given phone number(user id):

{
	"phone_number":"9126116"
}
```
with result:
```
{
    "balance": 2110
}
```
```
localhost:8080/update-balance updates balance of user that submitted voucher code(if valid).
```
input:

```
{
	"phone_number":"9126116",
	"amount":10
}
```

result :
```
{
    "result": "Ok"
}
```
mean all done and it's fine.
Ps. Actually it should've been grpc internal communication from voucher(discount) service to wallet. Due to building microservice separate from monolith and learning lots of new stuff, I had to skip implementing it.

```
localhost:8081/enable-voucher-code
```

I assumed that we have limited voucher codes and each one has specific amount of gift credit. This  api gets the code and the credit we want and enables it in database.

input:

```
{
	"voucher_code":"98CB-7558-JF9U",
	"amount": 1000000
}

output:
{
    "result": "OK"
}
```

```
localhost:8081/submit-voucher-code submits specific voucher code for specific user. 
```
input:
```
{
	"phone_number":"9126116",
	"voucher_code":"98CB-7558-JF9U"
}
output:
{
    "result": "OK"
}
```

This API should call update-balance internally.

```
localhost:8081/voucher-status/{voucher-code}
```
returns all voucher entities that are active and the users that used it.
I tried to  handle concurrency isolations as much as I could.






