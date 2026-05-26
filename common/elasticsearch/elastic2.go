package elasticsearch

import (
	elastic2 "gopkg.in/olivere/elastic.v3"
)

type Elastic2Wrapper struct {
	client        *elastic2.Client
	bulkProcessor *elastic2.BulkProcessor
}

func NewEsClient2(config ElasticConfig, bulkWorkers int) (*Elastic2Wrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// commit if # requests >= 1000
// commit if size of requests >= 2 MB
// commit every 10s

func (es *Elastic2Wrapper) IndexExists(indices ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic2Wrapper) CreateIndex(name string, mapping string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic2Wrapper) getAliases(index string) (*elastic2.AliasesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (es *Elastic2Wrapper) AddAlias(index string, alias string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic2Wrapper) HasAlias(indexName string, aliasName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic2Wrapper) ErrorStats() int64 { _ = "STUB: not implemented"; return 0 }

func (es *Elastic2Wrapper) AddBulkReq(index, typeName string, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *Elastic2Wrapper) FlushBulk() error { _ = "STUB: not implemented"; return nil }

func bulkAfterCBV2(_ int64, _ []elastic2.BulkableRequest, response *elastic2.BulkResponse, err error) {
	_ = "STUB: not implemented"
	return
}
