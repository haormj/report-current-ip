package provider

import (
	"fmt"

	"github.com/haormj/report-current-ip/option"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type AliyunProvider struct {
	opt    option.ProviderOpt
	client *alidns.Client
}

func NewAliyunProvider(opt option.ProviderOpt) (*AliyunProvider, error) {
	client, err := alidns.NewClientWithAccessKey(
		opt.AliyunProviderOpt.RegionID,
		opt.AliyunProviderOpt.AccessKeyID,
		opt.AliyunProviderOpt.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	ap := AliyunProvider{
		opt:    opt,
		client: client,
	}
	return &ap, nil
}

func (a *AliyunProvider) ReportCurrentIP(ip string) error {
	recordIDs, err := a.GetRecordIDs()
	if err != nil {
		return err
	}
	for _, recordID := range recordIDs {
		if err := a.DeleteDomainRecord(recordID); err != nil {
			return err
		}
	}
	if err := a.AddDomainRecord(ip); err != nil {
		return err
	}
	return nil
}

func (a *AliyunProvider) GetRecordIDs() ([]string, error) {
	req := alidns.CreateDescribeDomainRecordsRequest()
	req.DomainName = a.opt.DomainRecordOpt.DomainName
	req.Type = a.opt.DomainRecordOpt.Type
	req.SearchMode = "EXACT"
	req.KeyWord = a.opt.DomainRecordOpt.RR
	rsp, err := a.client.DescribeDomainRecords(req)
	if err != nil {
		return nil, err
	}
	recordIDs := make([]string, 0)
	for _, domainRecord := range rsp.DomainRecords.Record {
		recordIDs = append(recordIDs, domainRecord.RecordId)
	}
	return recordIDs, nil
}

func (a *AliyunProvider) AddDomainRecord(value string) error {
	req := alidns.CreateAddDomainRecordRequest()
	req.DomainName = a.opt.DomainRecordOpt.DomainName
	req.RR = a.opt.DomainRecordOpt.RR
	req.Type = a.opt.DomainRecordOpt.Type
	req.Value = value
	_, err := a.client.AddDomainRecord(req)
	if err != nil {
		return fmt.Errorf("AliyunProvider.AddDomainRecord %w", err)
	}
	return nil
}

func (a AliyunProvider) DeleteDomainRecord(recordID string) error {
	req := alidns.CreateDeleteDomainRecordRequest()
	req.RecordId = recordID
	_, err := a.client.DeleteDomainRecord(req)
	if err != nil {
		return err
	}
	return nil
}
