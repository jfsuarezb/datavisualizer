package main 


func Concatenate(buyersData *[]map[string]interface{}, productsData []map[string]interface{}, transactionsData *[]map[string]interface{}) {

	for i, trans := range *transactionsData {

		prodArray := make([]map[string]interface{}, 0)
		prodBlacklist := make(map[string]bool, 0)

		if trans["pids"] != nil {

			for _, transProdsId := range trans["pids"].([]string) {

				for _, prod := range productsData {
	
					if transProdsId == prod["pid"] && !prodBlacklist[prod["pid"].(string)] {
	
						prodArray = append(prodArray, prod)
						prodBlacklist[prod["pid"].(string)] = true
	
					}
	
				}
	
			}

		}

		(*transactionsData)[i]["prods"] = prodArray

	}

	for i, buyer := range *buyersData {

		transArray := make([]map[string]interface{}, 0)

		for _, trans := range *transactionsData {

			if trans["bid"] == buyer["bid"] {

				transArray = append(transArray, trans)

			}

		}

		(*buyersData)[i]["trans"] = transArray

	}

}