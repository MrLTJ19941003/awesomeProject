package persistclient

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler_distributed/config"
	"awesomeProject/crawler_distributed/rpcsupport"
	"log"
)

//RPC client
func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host) //创建一个连接hostRPC的client
	if err != nil {
		panic(err)
	}
	out := make(chan engine.Item) //创建一个chan engine.Item
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : Got item #%d : %v ", itemCount, item)
			itemCount++

			//call RPC to save item
			result := ""
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver : error saving item #%v : %v ", item, err)
			}
		}
	}()
	return out, nil
}
