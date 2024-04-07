# Application Design Review Task

## Task definition
The code presents a prototype of a hotel room booking service, which includes the ability to reserve a vacant room in a hotel.

The task is to refactor the structure and code of the application, fix existing problems in the logic, define and split codebase into different layers. It is also required to show the ability to use best practices and design patterns in the code.

Eventually, think about that the application will evolve soon, for example:

* There is a plan to implement feature to send confirmation email after booking.
* Discounts, promo codes, loyalty programs will be introduced.
* The ability to book multiple rooms is considered to be added.

> **✨Definition of Done:**
> 
> As a result of completing the task, structured code of the service is expected, with correctly functioning logic for hotel room booking scenarios.

## Limitations:
* The main store of application is expected to be in memory but is easily replaceable with any external storage.
* Use only the standard Go library packages and router (i.e., [chi](https://github.com/go-chi/chi)).


## Run examples

The application is expected to start with command presented below from the main package directory or with `Makefile` command if it wis provided.:
```sh
go run main.go
```
The HTTP API of the application is expected to correctly handle this cURL request. The request body can be improved, but no changes are expected with required fields. You can use the following cURL request to test the application:

```sh
curl --location --request POST 'localhost:8080/orders' \
--header 'Content-Type: application/json' \
--data-raw '{
    "hotel_id": "reddison",
    "room_id": "lux",
    "email": "guest@mail.ru",
    "from": "2024-01-02T00:00:00Z",
    "to": "2024-01-04T00:00:00Z"
}'
```