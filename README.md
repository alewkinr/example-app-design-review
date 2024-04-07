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
curl -X POST --location "http://localhost:8080/orders" \
    -H "Content-Type: application/json" \
    -d '{
          "hotel_id": "reddison",
          "room_id": "lux1",
          "email": "guest@mail.ru",
          "from": "2024-01-05T00:00:00Z",
          "to": "2024-01-06T00:00:00Z"
        }'
```

## Solution
### Design issues

1. The code is poorly structured, all written in one `main.go`, which leads to additional costs in maintaining and developing the service. As a solution, I used [go-standard-layout](https://github.com/golang-standards/project-layout) to organize the codebase. This layout is widely used in the Go community and provides a good starting point for a new project.
2. The core business logic of the service is implemented in the handler. This is a bad practice because the handler should only be responsible for handling transportation logic with requests and responses. I moved the business logic to the `internal/*` packages and split it in different domains, such as `booking`, `order`.
3. No dependency injection. I added a dependency injection container to the service in `internal/app.go`. This container is responsible for creating and storing all dependencies of the service. The logger, config, database connection, and other dependencies are created in the container and passed to the necessary components.
4. No graceful shutdown. I added a graceful shutdown to the service in `pkg/graceful`
The code is not covered by tests. I added unit tests for the main business logic of the service.
5. No validation of incoming requests. I decided not to add it, because it is costly to implement without external libraries, but still I point it out.
6. Configuration is hardcoded. I added a configuration loader in `internal/config`. The configuration is loaded from environment variables described in `.env-example`.
7. Concurrent access to the in-memory store was unsafe. I added a mutex to the in-memory store and wrapped store methods with interfaces so that the store can be easily replaced with another one. In-memory store is implemented in `pkg/store/inmemory` package.
8. No logging. I added logging to the service using the `log/slog` library. The logger is created in the container and passed to the necessary components. It is thread-safe and performance can be improved when it is necessary upon switching to another logging handler (i.e. [uber-go/zap](https://github.com/uber-go/zap/blob/master/exp/zapslog/doc.go)).