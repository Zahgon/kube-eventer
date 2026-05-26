package elasticsearch

import (
	elastic7 "github.com/olivere/elastic/v7"
)

type Elastic7Wrapper struct {
	client        *elastic7.Client
	pipeline      string
	bulkProcessor *elastic7.BulkProcessor
}

func NewEsClient7(config ElasticConfig, bulkWorkers int, pipeline string) (*Elastic7Wrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// commit if # requests >= 1000
// commit if size of requests >= 2 MB
// commit every 10s

func (es *Elastic7Wrapper) IndexExists(indices ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic7Wrapper) CreateIndex(name string, mapping string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic7Wrapper) getAliases(index string) (*elastic7.AliasesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (es *Elastic7Wrapper) AddAlias(index string, alias string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic7Wrapper) HasAlias(indexName string, aliasName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic7Wrapper) ErrorStats() int64 { _ = "STUB: not implemented"; return 0 }

func (es *Elastic7Wrapper) AddBulkReq(index, typeName string, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *Elastic7Wrapper) FlushBulk() error { _ = "STUB: not implemented"; return nil }

func bulkAfterCBV7(_ int64, _ []elastic7.BulkableRequest, response *elastic7.BulkResponse, err error) {
	_ = "STUB: not implemented"
	return
}
