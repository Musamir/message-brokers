[![Test statuses: ](https://github.com/Musamir/message-brokers/workflows/Test%20statuses/badge.svg??branch=master)](https://github.com/Musamir/message-brokers/actions)

## Contents
- [Description of the example](#Description-of-the-example)
- [Quick start](#Quick-start)

### Description of the example
    
In this example we have one producer and one consumer. 
Our producer sends a few messages (you can find them on service/service.go testData slice) 
and our consumer receives these messages. Messages are sent as slice of bytes ([]bytes), 
so the producer marshals messages before sending them via the protobuf protocol, and the consumer unmarshal them and prints.
I chose the protobuf protocol because it is one of the fastest protocols 
(you can see  an example of using this protocol and also a comparison of this protocol with others in my [repository](https://github.com/Musamir/performance/tree/main/encoding-decoding)).
    
### Quick start
If you wish to run the example, you need Docker (if you don’t have, you can download it from the official website [Docker](https://www.docker.com/get-started)):

1. Use the following command to run
    ```sh
    $ docker-compose up --build app-go
    ```
2. You might face with such error ![err conn](https://user-images.githubusercontent.com/43841786/145704473-88347618-040d-4410-b7e5-c5d37950ba21.png)
   This usually happens because RabbitMQ needs some time get ready before we can connect to it, so you just have to wait a few seconds and restart the program with the previous command without flag --build in order not to build again our go application
    ```sh
    $ docker-compose up app-go
    ```
   You can also run the following command to make sure that RabbitMQ is already running before starting go application to connect to it again.
    ```sh
    $ docker ps
    ```
   ![docker ps](https://user-images.githubusercontent.com/43841786/145704920-a1a99ba4-1c3b-4298-9958-faf7be4e9efd.png)
3. RabbitMQ also has management plugin provides an HTTP-based API for management and monitoring of RabbitMQ nodes and clusters, along with a browser-based UI and a command line tool, rabbitmqadmin.
   The management UI in our case can be accessed by following the link http://localhost:15673
    
    <p align="center" width="100%">
       <img src="https://user-images.githubusercontent.com/43841786/145705321-6903b0be-2cfd-44e8-940e-860ac48b2a40.png" width="50%">t
    </p>
    By default username - user, password - pass (you can change them in docker-compose.yaml).
    After authentication, the rabbitmq UI will be available to you.
   
   ![ui](https://user-images.githubusercontent.com/43841786/145705739-6cbcd047-dcc8-4a0e-84c8-5680edec7338.png)


I hope this example helped you, if you have any questions, feel free to ask them,
I would love to answer them also If you find any bugs, I would appreciate it if you let me know.

My contacts for communication: (mail - mirovmusamir@gmail.com), (telegram - @MusamirMirov)