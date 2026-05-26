package elasticsearch

import (
	elastic5 "gopkg.in/olivere/elastic.v5"
)

type Elastic5Wrapper struct {
	client        *elastic5.Client
	pipeline      string
	bulkProcessor *elastic5.BulkProcessor
}

func NewEsClient5(config ElasticConfig, bulkWorkers int, pipeline string) (*Elastic5Wrapper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// commit if # requests >= 1000
// commit if size of requests >= 2 MB
// commit every 10s

func (es *Elastic5Wrapper) IndexExists(indices ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic5Wrapper) CreateIndex(name string, mapping string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic5Wrapper) getAliases(index string) (*elastic5.AliasesResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (es *Elastic5Wrapper) AddAlias(index string, alias string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic5Wrapper) HasAlias(indexName string, aliasName string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (es *Elastic5Wrapper) ErrorStats() int64 { _ = "STUB: not implemented"; return 0 }

func (es *Elastic5Wrapper) AddBulkReq(index, typeName string, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (es *Elastic5Wrapper) FlushBulk() error { _ = "STUB: not implemented"; return nil }

func bulkAfterCBV5(_ int64, _ []elastic5.BulkableRequest, response *elastic5.BulkResponse, err error) {
	_ = "STUB: not implemented"
	return
}
