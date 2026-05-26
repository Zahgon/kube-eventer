package elasticsearch

import (
	elastic6 "gopkg.in/olivere/elastic.v6"
)

type Elastic6Wrapper struct {
	client        *elastic6.Client
	pipeline      string
	bulkProcessor *elastic6.BulkProcessor
}

func NewEsClient6(config ElasticConfig, bulkWorkers int, pipeline string) (*Elastic6Wrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// commit if # requests >= 1000
// commit if size of requests >= 2 MB
// commit every 10s

func (es *Elastic6Wrapper) IndexExists(indices ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic6Wrapper) CreateIndex(name string, mapping string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic6Wrapper) getAliases(index string) (*elastic6.AliasesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (es *Elastic6Wrapper) AddAlias(index string, alias string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic6Wrapper) HasAlias(indexName string, aliasName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic6Wrapper) ErrorStats() int64 { _ = "STUB: not implemented"; return 0 }

func (es *Elastic6Wrapper) AddBulkReq(index, typeName string, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *Elastic6Wrapper) FlushBulk() error { _ = "STUB: not implemented"; return nil }

func bulkAfterCBV6(_ int64, _ []elastic6.BulkableRequest, response *elastic6.BulkResponse, err error) {
	_ = "STUB: not implemented"
	return
}
