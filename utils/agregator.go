package utils

func (s *BulkInsertStruct) Test() (i *BulkInsertStruct) {
	return
}

func (t *BulkInsertStruct) SetColumn(columns string) (s *BulkInsertStruct) {
	t.Columns = &columns
	return t
}

func (t *BulkInsertStruct) SetValues(values string) (s *BulkInsertStruct) {
	t.Values = &values
	return t
}

func (t *BulkInsertStruct) SetQuery(query string) (s *BulkInsertStruct) {
	t.Query = &query
	return t
}

func (t *BulkInsertStruct) GetColumn() string {
	return *t.Columns
}

func (t *BulkInsertStruct) GetValues() string {
	return *t.Values
}

func (t *BulkInsertStruct) GetQuery() string {
	return *t.Query
}
