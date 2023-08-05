                                                         
██╗░░░██╗███╗░░██╗░██████╗███████╗███████╗███╗░░██╗      
██║░░░██║████╗░██║██╔════╝██╔════╝██╔════╝████╗░██║      
██║░░░██║██╔██╗██║╚█████╗░█████╗░░█████╗░░██╔██╗██║      
██║░░░██║██║╚████║░╚═══██╗██╔══╝░░██╔══╝░░██║╚████║      
╚██████╔╝██║░╚███║██████╔╝███████╗███████╗██║░╚███║      
░╚═════╝░╚═╝░░╚══╝╚═════╝░╚══════╝╚══════╝╚═╝░░╚══╝      
                                                         
[![Go Reference](https://pkg.go.dev/badge/github.com/snowkluster/unseen.svg)](https://pkg.go.dev/github.com/snowkluster/unseen)

# unseen
A modern go cli tool for hashing and analysing the type of hash 

# Installation 

## On windows

- run `go install github.com/snowkluster/unseen@latest` 

## On linux

- make sure you have your `GOBIN` setup in your path 
- firstly find out your `$GOPATH` by running `go env GOPATH`
- then paste the output of the above command {$GOPATH} into the command below 
- finally add <code>export PATH=${PATH}:{$GOPATH}/bin</code> to your `.bashrc` or `.profile` 
- this should allow you to use `unseen` on linux 
