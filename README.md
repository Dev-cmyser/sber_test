# Calculate Ipoteca

## test project for sber

## Installation

3. Build and run the application:

    ### Development
    To run the application in development mode using Docker:
    ```bash
    make dev
    ```

    ### Production
    To build and run the application in production mode:
    ```bash
    make prod
    ```

## API Documentation (Swagger)

Once the application is running, you can access the Swagger documentation at:

`http://localhost:8080/docs/index.html#/`


## Testing
```bash
make test
```

## Linter
```bash
make lint
```


## Coverage 
<img width="1427" alt="image" src="https://github.com/user-attachments/assets/8f4e77b4-ea52-4d88-86d0-d758a16d5de5">

## Container Size
```bash
make prod 
```
```bash
docker ps --size
```
![image](https://github.com/user-attachments/assets/e3f80cd5-2668-4b9b-9dee-712a6f81a111)
