# Product & Order REST API (Go v1.23)

A simple REST API to manage products and orders, built with Go v1.23 (Recommend to use the same version to minimize error). This API provides endpoints to create and list products and orders.

## Installation

1. Ensure that you have Go v1.23 installed on your machine.
2. Clone the repository:

    ```bash
    git clone https://github.com/fachriyanh/simple-ecommerce.git
    cd simple-ecommerce
    ```

3. Install the required Go modules:

    ```bash
    go mod tidy
    ```

## Usage

To run the application, simply execute the following command:

```bash
make run
```

The server will start and listen on localhost:8080.


## API Endpoints
| Method | Endpoint                | Description           | Request Body                                                                                 | Example Response                                                                                                           |
|--------|-------------------------|-----------------------|-----------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------|
| POST   | `/v1/api/products`       | Create a new product  | `{ "name": "Product Name", "stock": 100, "price": 29.99, "sku": "PROD001" }`                   | `{ "message": "Product created successfully" }`                                                                            |
| GET    | `/v1/api/products`       | Get all products      | N/A                                                                                           | `[ { "name": "Product Name", "stock": 100, "price": 29.99, "sku": "PROD001" }, { "name": "Another Product", "stock": 50, "price": 19.99, "sku": "PROD002" } ]`    |
| POST   | `/v1/api/orders`         | Create a new order    | `{ "name": "Order Name", "stock": 10, "items": 5 }`                                           | `{ "message": "Order created successfully" }`                                                                              |
| GET    | `/v1/api/orders`         | Get all orders        | N/A                                                                                           | `[ { "name": "Order Name", "stock": 10, "items": 5 }, { "name": "Another Order", "stock": 20, "items": 15 } ]`             |
