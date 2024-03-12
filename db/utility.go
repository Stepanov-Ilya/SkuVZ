package db

import (
	"purple_hack_tree/service"
	"sort"
	"sync"
)

var wg sync.WaitGroup
var STORAGE CurrentStorage

func GetPrice(request service.Request) service.Response {
	discountIds := GetSegmentsByUserID(request.UserId)
	sort.Slice(discountIds, func(i, j int) bool { return discountIds[i] > discountIds[j] })
	var response
	if discountIds != nil {
		for _, discountId := range discountIds {
			// Todo search in storage of discount
			// FindInDiscount(&response, discountId, request.LocationId, request.MicroCategoryId)
			// if response != nil {
			//	return response
			//}

		}
	}

	// Todo search in storage of baseline
	//FindInBaseline(&response, request.LocationId, request.MicroCategoryId)

	return response
}

func UpdateStorage(storage service.Storage) {

	wg.Add(1)
	wg.Add(1)
	go AddBaseline(storage.Baseline)
	go AddDiscounts(storage.Discounts)

	wg.Wait()
}

func UpdateBaseline(lines []service.Matrix) {
	//Todo create baseline
	//Todo add data
	//STORAGE.Baseline = CreateBaseline()
	//AddBaselineData(lines)
	defer wg.Done()
}

func UpdateDiscounts(discounts map[uint64][]service.Matrix) {
	//STORAGE.Discounts = make(map[uint64]uint64, 1)
	//wg.add(len(discounts))
	//AddDiscountData()
	for discountId, lines := range discounts {
		//Todo create discount
		//STORAGE.Discounts[discountId] = CreateDiscount()
		//go AddDiscountData(lines)
	}

	defer wg.Done()
}

