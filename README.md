# HelloWorld Chain

It is an Application specific blockchain, built using the `Cosmos SDK` on top of `Tindermint`, that handles `greeting` messages, written in `Go`.

### Features

- Add your personal greeting message with your name on the ledger. Nobody can have the same greeting as you :D
- List the available greetings.
- Get the sender of a specific greeting.

### How to use

## Build

- Clone the repository:
```ssh
    $ git clone https://github.com/SweeXordious/HelloWorldChain && cd HelloWorldChain
```
- Build:
```ssh
    $ make
```

## Run
To run this blockchain, we have two binaries:
- `hellod` : which is the daemon starting `tindermint`.
- `hellocli` :  the `CLI` used to interact with the blockchain.

#### Setup
To start the chain, some setup is necessary. the `init.sh` script is there for you:
```ssh
    $ ./init.sh
```
#### Start
- keep the terminal where you run the `init` script and open a new one
- in the new terminal run:
```ssh
    $ hellod start
```

## Interact
Now, from the terminal where you run `init`, you can start interacting with it.

##### Send a transaction
To send a transaction:
```ssh
    $ hellocli tx helloworld  setHello My_Favourite_Greeting me --from me --keyring-backend test
```
The `me` account is already created in the `init.sh` script along with a `you` account. 

You will receive the transaction hash in the response. To get more details:
```ssh
    $ hellocli q tx <txhash>
```

Also, the logs of the other terminal should state that a valid transaction was validated in a block etc.
##### List existing greetings
To send a transaction:
```ssh
    $ hellocli query helloworld list
```
The `me` account is already created in the `init.sh` script along with a `you` account. 
