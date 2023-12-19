package mongo

import (
	"context"
	"testing"
	"time"
)

func TestTemplateInsertOne(t *testing.T) {
	for _, tt := range initListTestInsertOne() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cancel()
			err := mongoTemplate.InsertOne(ctx, tt.value, tt.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertOne() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
