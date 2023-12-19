package mongo

import (
	"context"
	"testing"
)

func TestNewTemplate(t *testing.T) {
	for _, tt := range initListTestNewTemplate() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			temp, err := NewTemplate(ctx, tt.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertOne() error = %v, wantErr %v", err, tt.wantErr)
			} else if temp != nil {
				temp.Disconnect()
			}
			disconnectMongoTemplate()
		})
	}

}

func TestTemplateInsertOne(t *testing.T) {
	for _, tt := range initListTestInsertOne() {
		t.Run(tt.name, func(t *testing.T) {
			initMongoTemplate()
			if tt.beforeStartSession {
				mongoTemplate.StartSession(true)
			}
			if tt.beforeCloseMongoClient {
				mongoTemplate.Disconnect()
			}
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.InsertOne(ctx, tt.value, tt.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertOne() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				t.Log("err expected:", err)
			}
			if tt.forceErrCloseMongoClient {
				mongoTemplate.Disconnect()
			}
			mongoTemplate.CloseSession(ctx, err)
			disconnectMongoTemplate()
		})
	}
}
