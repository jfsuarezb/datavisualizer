package main

import (

	"fmt"
	"context"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"

)

func DGQuery(query string) (string, error) {

	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	
	defer d.Close()
	
	if err != nil {
		
		fmt.Println("e1")
		return "", err

	}

	dgClient := dgo.NewDgraphClient(api.NewDgraphClient(d),)

	err = dgClient.Alter(context.Background(), &api.Operation{

		Schema: `

			age: int .
			name: string .
			bid: string .
			trans: [uid] .
		
		`,

	})

	if err != nil {

		fmt.Println("e2")
		return "", err

	}

	txn := dgClient.NewTxn()

	defer txn.Discard(context.Background())

	_, err = txn.Mutate(context.Background(), &api.Mutation{SetJson: []byte(query)})

	if err != nil {

		fmt.Println("e3")
		return "", err

	}

	err = txn.Commit(context.Background())

	if err != nil {

		fmt.Println("e4")
		return "", err

	}

	return "succesful", nil

} 
