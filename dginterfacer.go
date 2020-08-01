package main

import (

	"context"
	"fmt"
	"encoding/json"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"

)

func DGPopulate(query string) (string, error) {

	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	
	defer d.Close()
	
	if err != nil {
		
		return "", err

	}

	dgClient := dgo.NewDgraphClient(api.NewDgraphClient(d),)

	err = dgClient.Alter(context.Background(), &api.Operation{

		Schema: `

			age: int .
			name: string .
			bid: string .
			trans: [uid] @reverse .
			tid: string .
			ip: string @index(hash) .
			dev: string .
			pids: [string] .
			prods: [uid] .
	
		`,

	})

	if err != nil {

		return "", err

	}

	txn := dgClient.NewTxn()

	defer txn.Discard(context.Background())

	_, err = txn.Mutate(context.Background(), &api.Mutation{SetJson: []byte(query)})

	if err != nil {

		return "", err

	}

	err = txn.Commit(context.Background())

	if err != nil {

		return "", err

	}

	return "succesful", nil

}

func DGQueryBuyers() ([]byte, error) {

	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	
	defer d.Close()
	
	if err != nil {
		
		return nil, err

	}

	dgClient := dgo.NewDgraphClient(api.NewDgraphClient(d),)

	err = dgClient.Alter(context.Background(), &api.Operation{

		Schema: `

			age: int .
			name: string .
			bid: string .
			trans: [uid] @reverse .
			tid: string .
			ip: string @index(hash) .
			dev: string .
			pids: [string] .
			prods: [uid] .
	
		`,

	})

	if err != nil {

		return nil, err

	}

	txn := dgClient.NewTxn()

	defer txn.Discard(context.Background())

	const q = `

		{	
			getBuyers(func: has(age)) {
				name
				age
				uid
			}
		}

	`

	resp, err := txn.Query(context.Background(), q)

	if err != nil {

		return nil, err

	}

	err = txn.Commit(context.Background())

	if err != nil {

		return nil, err

	}

	return resp.GetJson(), nil

}

func DGQueryBuyer(uid string) ([]byte, error) {

	d, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
	
	defer d.Close()
	
	if err != nil {
		
		return nil, err

	}

	dgClient := dgo.NewDgraphClient(api.NewDgraphClient(d),)

	err = dgClient.Alter(context.Background(), &api.Operation{

		Schema: `

			age: int .
			name: string .
			bid: string .
			trans: [uid] @reverse .
			tid: string .
			ip: string @index(hash) .
			dev: string .
			pids: [string] .
			prods: [uid] .
		
		`,

	})

	if err != nil {

		return nil, err

	}

	txn := dgClient.NewTxn()

	defer txn.Discard(context.Background())

	fq := fmt.Sprintf(`
	
		{
			getBuyer(func: uid(%s)) {
				name
				age
				trans {
					ip
					prods {
						name
						price
					}
				}
			}
		}

	`, uid)

	resp, err := txn.Query(context.Background(), fq)

	if err != nil {

		return nil, err

	}

	var buyer map[string]interface{}

	json.Unmarshal(resp.GetJson(), &buyer)

	trans := buyer["getBuyer"].([]interface{})[0].(map[string]interface{})["trans"].([]interface{})

	var ips []string

	for _, tran := range trans {

		ips = append(ips, tran.(map[string]interface{})["ip"].(string))

	}

	ips = norep(ips)

	var uidlist []string

	for _, ip := range ips {

		var getuids map[string]interface{}

		ipq, err := txn.Query(context.Background(), fmt.Sprintf(`
		
			{
				getuids(func: eq(ip, %s)) {

					~trans {

						uid

					}

				}
			}

		`, ip))

		if err != nil {

			return nil, err

		}

		json.Unmarshal(ipq.GetJson(), &getuids)

		for _, trans := range getuids["getuids"].([]interface{}) {

			uidlist = append(uidlist, trans.(map[string]interface{})["~trans"].([]interface{})[0].(map[string]interface{})["uid"].(string))

		}

	}

	uidlist = norep(uidlist)

	var uidlistclean []string

	for _, i := range uidlist {

		if i != uid {

			uidlistclean = append(uidlistclean, i)

		} 

	}

	var otherBuyersSameIp []interface{}

	for _, uidq := range uidlistclean {

		var getBuyerSameIp map[string]interface{}

		buyersameip, err := txn.Query(context.Background(), fmt.Sprintf(`
		
			{
				buyersameip(func: uid(%s)) {
					name
					age
				}
			}

		`, uidq))

		if err != nil {

			return nil, err

		}

		json.Unmarshal(buyersameip.GetJson(), &getBuyerSameIp)

		otherBuyersSameIp = append(otherBuyersSameIp, getBuyerSameIp["buyersameip"].([]interface{})[0])

	}

	buyer["otherBuyersSameIp"] = otherBuyersSameIp

	var recomProds []interface{}

	for _, uidprod := range uidlistclean {

		var recomProdsPerUid map[string]interface{}

		recomProdsResp, err := txn.Query(context.Background(), fmt.Sprintf(`
		
			{
				getrecomprod(func: uid(%s)) {

					trans {

						prods {

							name
							price

						}

					}

				}
			}

		`, uidprod))

		if err != nil {

			return nil, err

		}

		json.Unmarshal(recomProdsResp.GetJson(), &recomProdsPerUid)

		for _, recomProdset := range recomProdsPerUid["getrecomprod"].([]interface{})[0].(map[string]interface{})["trans"].([]interface{}) {

			recomProds = append(recomProds, recomProdset.(map[string]interface{})["prods"].([]interface{})...)

		}

	}

	buyer["recommendedProducts"] = fiveBest(recomProds)
	
	err = txn.Commit(context.Background())

	if err != nil {
	
		return nil, err
	
	}

	final, err := json.Marshal(buyer)

	if err != nil {

		return nil, err

	}

	return final, nil

}

func norep(arr []string) []string {

	check := make(map[string]int)

	noreparr := make([]string, 0)

	for _, val := range arr {
		check[val] = 1
	}

	for letter, _ := range check {
		noreparr = append(noreparr,letter)
	}

	return noreparr

}

func fiveBest(arr []interface{}) []interface{} {

	stringPriceMap := make(map[string]string, 0)

	for _, i := range arr {

		stringPriceMap[i.(map[string]interface{})["name"].(string)] = i.(map[string]interface{})["price"].(string)

	}

	var stringArray []string

	for _, i := range arr {

		stringArray = append(stringArray, i.(map[string]interface{})["name"].(string))

	}

	stringFreqMap := make(map[string]int, 0) 

	for _, i := range stringArray {

		stringFreqMap[i] += 1

	} 

	var fiveMostFreq []string

	for i := 1; i <= 5 ; i +=1 {

		mostPopulousKey := ""

		mostPopulousValue := 0

		for key, value := range stringFreqMap {

			if value > mostPopulousValue {

				mostPopulousValue = value
				mostPopulousKey = key

			}

		}

		fiveMostFreq = append(fiveMostFreq, mostPopulousKey)

		delete(stringFreqMap, mostPopulousKey)

	} 	

	var reconstruction []interface{}

	for _, i := range fiveMostFreq {

		m := make(map[string]string, 0)

		m["name"] = i
		m["price"] = stringPriceMap[i]

		reconstruction = append(reconstruction, m)

	}

	return reconstruction

}
