# Q&A-management

Using golang Gin Framework to build project

## Command

- ### Application
    - Install node modules
      ```sh
      $ go get .
      ```
    - Build app
      ```sh
      $ go build -o main
      ```
    - Start app in development
      ```sh
      $ go run main.go
      ```

- ### Docker
    - Build container
    ```sh
    $ docker-compose build
    ```
    - Run docker with flags
    ```sh
     $ sudo docker-compose up -d --<flags name>
    ```
    - Run container build with flags
    ```sh
    $ docker-compose up -d --build --<flags name>
    ```
    - Run container
    ```sh
    $ docker-compose up -d --build
    ```
    - Stop container
    ```sh
    $ docker-compose down
    ```