package main

import (
	"fmt"
	"ivixlabs.com/goweb/internal/clickhouse"
	"ivixlabs.com/goweb/internal/model"
)

func main() {
	conn := clickhouse.NewConn([]string{"192.168.1.12:9000"})

	entityPropertyRepo := model.NewClickHouseEntityPropertyRepository(conn)
	entityRepo := model.NewClickHouseEntityRepository(conn, entityPropertyRepo)

	for j := 0; j < 100; j++ {
		batch := entityRepo.CreateBatch()
		for i := 0; i < 10000; i++ {
			en := model.NewEntity()
			strVal := "StringValue"
			en.SetProperty("StringProp", "string", &strVal)

			intVal := fmt.Sprintf("%v", i)
			en.SetProperty("IntProp", "int", &intVal)

			entityRepo.Add(batch, en)
		}
		entityRepo.SendBatch(batch)
	}

}
