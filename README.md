# go-memhack
[![](https://goreportcard.com/badge/github.com/andygeiss/go-memhack)](https://goreportcard.com/report/github.com/andygeiss/go-memhack)

An example of using ptrace to hack process memory and implement security measures.

## Build binaries

    make

    ls build
    
    hackme  memhack  memsearch

## Run target binary

Start the compiled with:

    ./build/hackme

The executable will run in an infinite loop printing the address of the Player and the Players Score (address) every second:

    Hackme [Player: 0xc82002c020, Score: 12345678 (0xc82000e450)]
    Hackme [Player: 0xc82002c020, Score: 12345678 (0xc82000e450)]
    Hackme [Player: 0xc82002c020, Score: 12345678 (0xc82000e450)]
    Hackme [Player: 0xc82002c020, Score: 12345678 (0xc82000e450)]
    Hackme [Player: 0xc82002c020, Score: 12345678 (0xc82000e450)]

## Run memsearch to find the value

Open a new terminal window and start memsearch:

    ./build/memsearch -addr 0xc82000e450 -len 8 -value 12345678 -pid 14653
    
    Process [14653] attached.
    Searching for value [dec: 12345678][hex: 0xbc614e] ...
    Value found at [0xc82000e450].
    Process [14653] detached.
    Value found [1] times.

Please ensure to replace the PID 14653 by your current hackme using the -pid flag.

## Hack the players score by using memhack

Set Players score from 12345678 to 87654321:

    ./build/memhack -addr 0xc82000e450 -value 87654321 -pid 14653

Now look at the looping hackme:

    Hackme [Player: 0xc82002c020, Score: 12345678 (0xc82000e450)]
    Hackme [Player: 0xc82002c020, Score: 12345678 (0xc82000e450)]
    Hackme [Player: 0xc82002c020, Score: 87654321 (0xc82000e450)] <---- SUCCESS
    Hackme [Player: 0xc82002c020, Score: 87654321 (0xc82000e450)]

## Protecting the player

Modify the source by replacing the following lines at platform/hackme/main.go:

    s = score.NewDefaultService()
	//s = score.NewSecurityService() <---- Use this!

This service uses a simple but powerful way to protect the score value.
The function NewSecurityService creates a new service with a random key.
Each access to the players score will be encrypted/decrypted by XORing the key with the real score value. The key will never be known by the user or persisted to the filesystem.
  