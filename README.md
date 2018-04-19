# Courier

## Configuration options

`COURIER_PORT`
> COURIER PORT tells courier which port to run on

`SMTP_HOST`
> SMTP Host is the domain name associated with your SMTP server
> for gmail, it is smtp.gmail.com

`SMTP_PASSWORD`
> This is your authentication method to the SMTP server, for gmail, it is your google password

`SMTPUser`
> This is your authentication method to the SMTP server, for gmail, it is your google username/emailaddress

All configuration options above are provided in the docker-compose.yml and may be overidden at your discretion


## Getting Started

1. overide the environment variables in the docker-compose file,
2. `docker-compose up`
  * Courier is now running on your system!
* you now need a method to communicate to the courier server, you may use the example client located in `cmd/client/main/go` or you may write your own

## Creating your own client

1. Courier exposes a gRPC interface, so you are able to use the generated code in the pb package.
2. Create a connection to the Courier instance
    ```go
    // Set up a connection to the server.
	  conn, err := grpc.Dial(address, grpc.WithInsecure())
	  if err != nil {
		  log.Fatalf("did not connect: %v", err)
	  }
	  defer conn.Close()
	  c := pb.NewMailClient(conn)
    ```
3. Send a message to whoever you wish by, sending a message like:
  ```go
    // Contact the server and print out its response.
	  r, err := c.Send(context.Background(), &pb.SendBody{
		  Recipients: []string{"email@email.com"},
		  Sender:     "email@email.com",
		  Subject:    "Hello, to my friends",
		  BodyText:   "Just saying hi",
		  BodyType:   "text/html",
	  })
  ```
4. That's all it takes