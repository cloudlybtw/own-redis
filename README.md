
# own-redis

Custom key-value databases are crucial for efficiently storing and retrieving data using unique keys. This application will operate in RAM for fast access, utilizing the UDP protocol to facilitate communication between the storage and the client. Key operations will include SET, GET, and PING.

## Installation
### Prerequisites
- Go 1.22.6 (Ensure Go is properly installed on your machine).
### Steps
Follow these steps to install the project locally on your machine:

1. Clone the repository:
   ```bash
   $ git clone git@git.platform.alem.school:dakabirov/own-redis.git
2. Navigate into the project directory:
    ```bash
    $ cd own-redis
3. Build the programme.
    ``` bash
    $ go build -o own-redis .
4. Run the programme. You can include flags if required.
    ``` bash
    $ ./own-redis [--help] [--port]
5. You can use operations in a new terminal window.
    ``` bash
    $ nc 0.0.0.0 8080
## Flags
1. port
Receives port number, on which UDP protocol will be held.

2. help
Returns help message to console.

## Operations
### PING
Check if the server is alive.

- Request:
    ```bash
    PING
- Responce:
    ```bash
    PONG
### SET
Insert a key-value pair into the store. You can also set an expiration time in milliseconds using the PX argument.
- Request:
    ```bash
    SET foo bar
- Request with expiration, the pair will expire in 10 seconds:
    ```bash
    SET foo bar px 10000
- Responce:
    ```bash
    OK
### GET
Retrieve the value for a given key.
- Request:
    ```bash
    GET foo
- Responce:
    ```bash
    bar
If the key doesn't exist:
- Responce:
    ```bash
    GET RandomKey
    (nil)
## Author
Daniyar Kabirov

cloudlypower@gmail.com

Made for Alem school Foundation.


