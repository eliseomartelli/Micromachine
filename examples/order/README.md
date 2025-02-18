# Order Processing Example

The example models a simplified order processing system.  The states are
`OrderCreated`, `OrderPaid`, `OrderShipped`, `OrderDelivered`, and
`OrderCancelled`.

## HTTP Endpoints:

`/order`: Returns the current order information as JSON.
`/transition`: Transitions the order to a new state based on the request body (e.g., `{"to": "paid"}`).

## Look at it in action

Get Order:

```Bash
curl http://localhost:8080/order
```

Pay for Order:

```Bash
curl -X POST -H "Content-Type: application/json" -d '{"to": "paid"}' \
http://localhost:8080/transition
```

Ship Order:

```Bash
curl -X POST -H "Content-Type: application/json" -d '{"to": "shipped"}' \
http://localhost:8080/transition
```

Deliver Order:

```Bash
curl -X POST -H "Content-Type: application/json" -d '{"to": "delivered"}' \
http://localhost:8080/transition
```

Cancel Order (from Created):

```Bash
curl -X POST -H "Content-Type: application/json" -d '{"to": "cancelled"}' \
http://localhost:8080/transition
```

Cancel Order (from Paid):

```Bash
curl -X POST -H "Content-Type: application/json" -d '{"to": "cancelled"}' \
http://localhost:8080/transition
```

 Each transition will print a message to the server's console, and the `/order`
 endpoint will reflect the updated status.
