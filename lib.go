package main

import (
	"fmt"
	"github.com/denverdino/aliyungo/dns"
	"github.com/gutengo/shell"
	"net/http"
	"regexp"
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

func getIp() string {
	res, err := http.Get("http://pv.sohu.com/cityjson?ie=utf-8")
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	ip := regexp.MustCompile(`\d\d\d\.\d\d\d.\d\d\d.\d\d\d`).FindString(string(body))
	return ip
}
