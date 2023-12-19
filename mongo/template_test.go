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

func TestTemplateInsertMany(t *testing.T) {
	initMongoTemplate()
	for _, tt := range initListTestInsertMany() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.InsertMany(ctx, tt.value, tt.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertMany() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				t.Log("err expected:", err)
			}
			mongoTemplate.CloseSession(ctx, err)
		})
	}
	disconnectMongoTemplate()
}

func TestTemplateDeleteOne(t *testing.T) {
	initDocument()
	for _, tt := range initListTestDelete() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.DeleteOne(ctx, tt.filter, tt.ref, tt.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteOne() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				t.Log("err expected:", err)
			}
			mongoTemplate.CloseSession(ctx, err)
		})
	}
	disconnectMongoTemplate()
}

func TestTemplateDeleteMany(t *testing.T) {
	initDocument()
	for _, tt := range initListTestDelete() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.DeleteMany(ctx, tt.filter, tt.ref, tt.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteMany() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				t.Log("err expected:", err)
			}
		})
	}
	disconnectMongoTemplate()
}

func TestTemplateUpdateOneById(t *testing.T) {
	initDocument()
	for _, tt := range initListTestUpdateOneById() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.UpdateOneById(ctx, tt.id, tt.update, tt.ref, tt.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateOneById() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				t.Log("err expected:", err)
			}
		})
	}
	disconnectMongoTemplate()
}

func TestTemplateUpdateOne(t *testing.T) {
	initDocument()
	for _, tt := range initListTestUpdate() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.UpdateOne(ctx, tt.filter, tt.update, tt.ref, tt.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateOneById() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				t.Log("err expected:", err)
			}
			mongoTemplate.CloseSession(ctx, err)
		})
	}
	disconnectMongoTemplate()
}

func TestTemplateUpdateMany(t *testing.T) {
	initDocument()
	for _, tt := range initListTestUpdate() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.UpdateMany(ctx, tt.filter, tt.update, tt.ref, tt.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateMany() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				t.Log("err expected:", err)
			}
			mongoTemplate.CloseSession(ctx, err)
		})
	}
	disconnectMongoTemplate()
}

func TestTemplateReplaceOne(t *testing.T) {
	initDocument()
	for _, tt := range initListTestReplace() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.ReplaceOne(ctx, tt.filter, tt.replacement, tt.ref, tt.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReplaceOne() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil {
				t.Log("err expected:", err)
			}
			mongoTemplate.CloseSession(ctx, err)
		})
	}
	disconnectMongoTemplate()
}
