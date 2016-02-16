package main

import (
	"fmt"
	"github.com/denverdino/aliyungo/dns"
	"github.com/gutengo/shell"
)

func List(domain string) {
	client := dns.NewClient(rc.ACCESS_KEY_ID, rc.ACCESS_KEY_SECRET)
	res, err := client.DescribeDomainRecords(&dns.DescribeDomainRecordsArgs{
		DomainName: domain,
	})
	if err != nil {
		shell.ErrorExit(err)
	}

	for _, v := range res.DomainRecords.Record {
		fmt.Println(v.RecordId, v.RR, v.Value, v.DomainName, v.Type)
	}
}

func Update(recordId, rr, value string) error {
	client := dns.NewClient(rc.ACCESS_KEY_ID, rc.ACCESS_KEY_SECRET)
	_, err := client.UpdateDomainRecord(&dns.UpdateDomainRecordArgs{
		RecordId: recordId,
		RR:       rr,
		Value:    value,
		Type:     dns.ARecord,
	})
	return err
}
