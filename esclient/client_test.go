package esclient

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	cfg := &elasticsearch.Config{
		Addresses: []string{"http://192.168.56.20:9200"},
		Username:  "elastic",
		Password:  "3qzW2Bgr4lIvJXYl2Vd0",
	}
	es := &EsClient{
		Config: cfg,
	}
	err := es.IntClient()
	if err != nil {
		t.Fatalf("init es client error")
	}

	t.Logf("init es client succeed")
}

var smsLogFmt = `
	{
		"createData": "%s",
		"sendData": "%s",
		"longCode": %d,
		"mobile": %d,
		"corpName": "魏巍教育有限公司2",
		"smsContent": "你收到好友 AAA 发送的消息",
		"state": 0,
		"operatorId": 1,
		"province": "guangdong",
		"ipAddr": "192.168.56.20",
		"replyTotal": 60,
		"fee": 10
	}
`

func TestIndexCreateV1(t *testing.T) {
	cfg := &elasticsearch.Config{
		Addresses: []string{"http://192.168.56.20:9200"},
		Username:  "elastic",
		Password:  "3qzW2Bgr4lIvJXYl2Vd0",
	}
	es := &EsClient{
		Config: cfg,
	}
	err := es.IntClient()
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	var wg sync.WaitGroup
	mobiles := []int64{18710565589, 18710565588, 18710565587, 18710565586,
		18710565585, 18710565584, 18710565583, 18710565582, 18710565581}
	for i, mobile := range mobiles {
		wg.Add(1)
		go func(i int, mobile int64) {
			defer wg.Done()

			now := time.Now()
			createData := now.Format("2020-01-01 00:00:00")
			sendData := now.Add(60 * time.Second)
			var longCode = 1000000000
			smsLogs := fmt.Sprintf(smsLogFmt, createData, sendData, longCode, mobile)

			cfg := &elasticsearch.Config{
				Addresses: []string{"http://192.168.56.20:9200"},
				Username:  "elastic",
				Password:  "3qzW2Bgr4lIvJXYl2Vd0",
			}
			es := &EsClient{
				Config: cfg,
			}
			err := es.IntClient()
			if err != nil {
				t.Fatalf("some error: %s", err)
			}

			req := esapi.IndexRequest{
				Index:      "sms-logs-index",
				DocumentID: fmt.Sprintf("%d", i),
				Body:       strings.NewReader(smsLogs),
				Refresh:    "true",
			}

			// Set up the request object.

			res, err := req.Do(context.Background(), es.Client)
			if err != nil {
				t.Fatalf("some error: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				t.Fatalf("some error: %s", res.Status())
			} else {
				t.Logf("succeed")
			}

		}(i, mobile)
	}
	wg.Wait()
}

func TestCreteIndexV2(t *testing.T) {
	now := time.Now()
	createData := now.Format("2020-01-01 00:00:00")
	sendData := now.Add(60 * time.Second)
	var longCode = 1000000000
	var mobile = 18710565589
	smsLogs := fmt.Sprintf(smsLogFmt, createData, sendData, longCode, mobile)

	cfg := &elasticsearch.Config{
		Addresses: []string{"http://192.168.56.20:9200"},
		Username:  "elastic",
		Password:  "3qzW2Bgr4lIvJXYl2Vd0",
	}
	es := &EsClient{
		Config: cfg,
	}
	err := es.IntClient()
	if err != nil {
		t.Fatalf("some error: %s", err)
	}

	req := esapi.IndexRequest{
		Index:      "sms-logs-index",
		DocumentID: "1",
		Body:       strings.NewReader(smsLogs),
		Refresh:    "true",
	}

	// Set up the request object.

	res, err := req.Do(context.Background(), es.Client)
	if err != nil {
		t.Fatalf("some error: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		t.Fatalf("some error: %s", res.Status())
	} else {
		t.Logf("succeed")
	}
}
