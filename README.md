# message-brokers

Hi, my name is Musamir, and I am sharing here examples that interest me, and I hope you find them interesting. In this repository I give examples of using message brokers. If you find any bugs, I would appreciate it if you let me know. Also, if you have any questions, I would love to answer them.

## Contents
- [Message broker](#Message-broker)
    - [RabbitMQ](https://github.com/Musamir/message-brokers/tree/main/rabbitMQ/hello%20world)
    - [NATS](#NATS)

### Message broker
Message broker is a service that accepts messages and forwards them. The illustration is given below

  <p align="center" width="100%">
       <img src="https://user-images.githubusercontent.com/43841786/145713076-554de953-bbf4-4224-a15d-24f84f6b863e.png" width="50%">t
  </p>

There are two roles in this process:

1. Producer - sends messages
2. Consumer - accepts messages and processes them.

An example of using message brokers is when, for example, 
you need to send emails. Since emails are not sent quickly, 
in order to not wait for the completion of sending emails, 
we pass(Publish) them to the message broker and the services 
that send messages will take(Consume) themselves messages from the 
broker when they are free.

My contacts for communication: (mail - mirovmusamir@gmail.com), (telegram - @MusamirMirov)
