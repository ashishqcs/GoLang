package service

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"internet.provider/client"
	"internet.provider/postgres"
)

func NewInternetProviderService(internetClient client.Client,
	repository postgres.Repository) *InternetProviderService {
	return &InternetProviderService{
		client:         internetClient,
		billRepository: repository,
	}
}

type InternetProviderService struct {
	client         client.Client
	billRepository postgres.Repository
}

func (ips *InternetProviderService) GetInternetBills(startId int, endId int) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	numOfThreads := runtime.GOMAXPROCS(runtime.NumCPU())

	bills := []postgres.BillEntity{}

	partitions := createPartitions(startId, endId, numOfThreads)

	for _, res := range partitions {
		wg.Add(1)
		go func(start int, end int) {
			defer wg.Done()
			be := getBillsForIds(ips, start, end)
			mu.Lock()
			bills = append(bills, be...)
			mu.Unlock()
		}(res[0], res[1])
	}

	wg.Wait()

	ips.billRepository.SaveBills(bills)

	fmt.Print(ips.GetInternetBillsForName("Ram"))
}

func (ips *InternetProviderService) GetInternetBillsForName(name string) []postgres.BillEntity {
	be, _ := ips.billRepository.GetBillsByName(name)
	return be
}

func getBillsForIds(ips *InternetProviderService, start int, end int) []postgres.BillEntity {
	bills := []postgres.BillEntity{}

	for i := start; i < end; i++ {
		body, err := ips.client.GetBillById(i)
		if err == nil {
			res := string(body)
			splitRes := strings.Split(res, "\n")
			bill := mapToBillEntity(splitRes[1])
			fmt.Print(bill)
			fmt.Println()
			bills = append(bills, bill)
		}
	}

	return bills

}

func createPartitions(lookFrom int, lookTill int, partitions int) [][]int {
	resRange := [][]int{}
	difference := lookTill - lookFrom
	size := difference / partitions
	for i := 0; i < partitions; i++ {
		start := lookFrom + (size * i)
		end := start + size
		if i == partitions-1 {
			end = lookTill + 1
		}
		resRange = append(resRange, []int{start, end - 1})
	}

	return resRange
}

func mapToBillEntity(s string) postgres.BillEntity {
	fields := strings.Split(s, ",")
	id, _ := strconv.ParseInt(fields[0], 10, 64)
	amt, _ := strconv.ParseFloat(fields[5], 64)
	bill := postgres.BillEntity{
		Id:       id,
		Name:     fields[1],
		Address:  fields[2],
		PlanName: fields[3],
		Date:     fields[4],
		Amount:   amt,
	}

	return bill
}
