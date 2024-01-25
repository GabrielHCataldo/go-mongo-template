package mongo

import "github.com/GabrielHCataldo/go-errors/errors"

var MsgErrRefDocument = "mongo: ref document needs to be structure or slice of the struct"
var MsgErrDatabaseNotConfigured = "mongo: database not correct configured"
var MsgErrCollectionNotConfigured = "mongo: collection not correct configured"
var MsgErrDocumentIsNotPointer = "mongo: document param is not a pointer"
var MsgErrDocumentIsNotStruct = "mongo: document param is not a struct"
var MsgErrDocumentIsEmpty = "mongo: document param is empty"
var MsgErrDocumentsIsEmpty = "mongo: documents param is empty"
var MsgErrDestIsNotPointer = "mongo: dest param is not a pointer"
var MsgErrDestIsNotStruct = "mongo: dest param is not a struct"
var MsgErrNoDocuments = "mongo: no documents in result"
var MsgErrNoOpenSession = "mongo: no open session"

var ErrRefDocument = errors.New(MsgErrRefDocument)
var ErrDatabaseNotConfigured = errors.New(MsgErrDatabaseNotConfigured)
var ErrCollectionNotConfigured = errors.New(MsgErrCollectionNotConfigured)
var ErrDocumentIsNotPointer = errors.New(MsgErrDocumentIsNotPointer)
var ErrDocumentIsNotStruct = errors.New(MsgErrDocumentIsNotStruct)
var ErrDocumentIsEmpty = errors.New(MsgErrDocumentIsEmpty)
var ErrDocumentsIsEmpty = errors.New(MsgErrDocumentsIsEmpty)
var ErrDestIsNotPointer = errors.New(MsgErrDestIsNotPointer)
var ErrDestIsNotStruct = errors.New(MsgErrDestIsNotStruct)
var ErrNoDocuments = errors.New(MsgErrNoDocuments)
var ErrNoOpenSession = errors.New(MsgErrNoOpenSession)

func errRefDocument(skip int) error {
	ErrRefDocument = errors.NewSkipCaller(skip+1, MsgErrRefDocument)
	return ErrRefDocument
}

func errDatabaseNotConfigured(skip int) error {
	ErrDatabaseNotConfigured = errors.NewSkipCaller(skip+1, MsgErrDatabaseNotConfigured)
	return ErrDatabaseNotConfigured
}

func errCollectionNotConfigured(skip int) error {
	ErrCollectionNotConfigured = errors.NewSkipCaller(skip+1, MsgErrCollectionNotConfigured)
	return ErrCollectionNotConfigured
}

func errDocumentIsNotPointer(skip int) error {
	ErrDocumentIsNotPointer = errors.NewSkipCaller(skip+1, MsgErrDocumentIsNotPointer)
	return ErrDocumentIsNotPointer
}

func errDocumentIsNotStruct(skip int) error {
	ErrDocumentIsNotStruct = errors.NewSkipCaller(skip+1, MsgErrDocumentIsNotStruct)
	return ErrDocumentIsNotStruct
}

func errDocumentIsEmpty(skip int) error {
	ErrDocumentIsEmpty = errors.NewSkipCaller(skip+1, MsgErrDocumentIsEmpty)
	return ErrDocumentIsEmpty
}

func errDocumentsIsEmpty(skip int) error {
	ErrDocumentsIsEmpty = errors.NewSkipCaller(skip+1, MsgErrDocumentsIsEmpty)
	return ErrDocumentsIsEmpty
}

func errDestIsNotPointer(skip int) error {
	ErrDestIsNotPointer = errors.NewSkipCaller(skip+1, MsgErrDestIsNotPointer)
	return ErrDestIsNotPointer
}

func errDestIsNotStruct(skip int) error {
	ErrDestIsNotStruct = errors.NewSkipCaller(skip+1, MsgErrDestIsNotStruct)
	return ErrDestIsNotStruct
}

func errNoDocuments(skip int) error {
	ErrNoDocuments = errors.NewSkipCaller(skip+1, MsgErrNoDocuments)
	return ErrNoDocuments
}

func errNoOpenSession(skip int) error {
	ErrNoOpenSession = errors.NewSkipCaller(skip+1, MsgErrNoOpenSession)
	return ErrNoOpenSession
}
