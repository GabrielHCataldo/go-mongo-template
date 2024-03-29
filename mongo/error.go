package mongo

import "errors"

var ErrRefDocument = errors.New("mongo: ref document needs to be structure or slice of the struct")
var ErrDatabaseNotConfigured = errors.New("mongo: database not correct configured")
var ErrCollectionNotConfigured = errors.New("mongo: collection not correct configured")
var ErrDocumentIsNotPointer = errors.New("mongo: document param is not a pointer")
var ErrDocumentIsNotStruct = errors.New("mongo: document param is not a struct")
var ErrDocumentIsEmpty = errors.New("mongo: document param is empty")
var ErrDocumentsIsEmpty = errors.New("mongo: documents param is empty")
var ErrDestIsNotPointer = errors.New("mongo: dest param is not a pointer")
var ErrDestIsNotStruct = errors.New("mongo: dest param is not a struct")
var ErrNoDocuments = errors.New("mongo: no documents in result")
var ErrNoOpenSession = errors.New("mongo: no open session")
