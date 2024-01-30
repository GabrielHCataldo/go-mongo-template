package mongo

import (
	"context"
	"github.com/GabrielHCataldo/go-helper/helper"
	"github.com/GabrielHCataldo/go-logger/logger"
	"testing"
	"time"
)

func TestNewTemplate(t *testing.T) {
	for _, tt := range initListTestNewTemplate() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			temp, err := NewTemplate(ctx, tt.options)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("InsertOne() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(temp) {
				_ = temp.Disconnect(ctx)
			}
		})
	}

}

func TestTemplateInsertOne(t *testing.T) {
	for _, tt := range initListTestInsertOne() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			initMongoTemplate()
			_ = mongoTemplate.StartSession(ctx)
			if tt.beforeCloseMongoClient {
				_ = mongoTemplate.Disconnect(ctx)
			}
			err := mongoTemplate.InsertOne(ctx, tt.value, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("InsertOne() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
			if tt.forceErrCloseMongoClient {
				mongoTemplate.SimpleDisconnect(ctx)
				mongoTemplate.SimpleDisconnect(ctx)
			}
			_ = mongoTemplate.CloseSession(ctx, helper.IsNotNil(err))
			_ = mongoTemplate.CloseSession(ctx, helper.IsNotNil(err))
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
			err := mongoTemplate.InsertMany(ctx, tt.value, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("InsertMany() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateDeleteOne(t *testing.T) {
	initDocument()
	for _, tt := range initListTestDelete() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.DeleteOne(ctx, tt.filter, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("DeleteOne() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
			_ = mongoTemplate.CloseSession(ctx, helper.IsNotNil(err))
		})
	}
}

func TestTemplateDeleteOneById(t *testing.T) {
	initDocument()
	for _, tt := range initListTestDeleteById() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.DeleteOneById(ctx, tt.id, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("DeleteOneById() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
			_ = mongoTemplate.CloseSession(ctx, helper.IsNotNil(err))
		})
	}
}

func TestTemplateDeleteMany(t *testing.T) {
	initDocument()
	for _, tt := range initListTestDelete() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.DeleteMany(ctx, tt.filter, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("DeleteMany() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateUpdateOneById(t *testing.T) {
	initDocument()
	for _, tt := range initListTestUpdateOneById() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.UpdateOneById(ctx, tt.id, tt.update, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("UpdateOneById() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateUpdateOne(t *testing.T) {
	initDocument()
	for _, tt := range initListTestUpdate() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.UpdateOne(ctx, tt.filter, tt.update, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("UpdateOneById() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
			_ = mongoTemplate.CloseSession(ctx, helper.IsNotNil(err))
		})
	}
}

func TestTemplateUpdateMany(t *testing.T) {
	initDocument()
	for _, tt := range initListTestUpdate() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.UpdateMany(ctx, tt.filter, tt.update, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("UpdateMany() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
			_ = mongoTemplate.CloseSession(ctx, helper.IsNotNil(err))
		})
	}
}

func TestTemplateReplaceOne(t *testing.T) {
	initDocument()
	for _, tt := range initListTestReplaceOne() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.ReplaceOne(ctx, tt.filter, tt.replacement, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("ReplaceOne() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
			_ = mongoTemplate.CloseSession(ctx, helper.IsNotNil(err))
		})
	}
}

func TestTemplateReplaceOneById(t *testing.T) {
	initDocument()
	for _, tt := range initListTestReplaceOneById() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.ReplaceOneById(ctx, tt.id, tt.replacement, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("ReplaceOneById() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
			_ = mongoTemplate.CloseSession(ctx, helper.IsNotNil(err))
		})
	}
}

func TestTemplateFindOneById(t *testing.T) {
	initDocument()
	for _, tt := range initListTestFindOneById() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.FindOneById(ctx, tt.id, tt.dest, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("FindOneById() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateFindOne(t *testing.T) {
	initDocument()
	for _, tt := range initListTestFindOne() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.FindOne(ctx, tt.filter, tt.dest, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("FindOne() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateFindOneAndDelete(t *testing.T) {
	initDocument()
	for _, tt := range initListTestFindOneAndDelete() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.FindOneAndDelete(ctx, tt.filter, tt.dest, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("FindOneAndDelete() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
			_ = mongoTemplate.CloseSession(ctx, helper.IsNotNil(err))
		})
	}
}

func TestTemplateFindOneAndReplace(t *testing.T) {
	initDocument()
	for _, tt := range initListTestFindOneAndReplace() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.FindOneAndReplace(ctx, tt.filter, tt.replacement, tt.dest, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("FindOneAndReplace() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
			_ = mongoTemplate.CloseSession(ctx, helper.IsNotNil(err))
		})
	}
}

func TestTemplateFindOneAndUpdate(t *testing.T) {
	initDocument()
	for _, tt := range initListTestFindOneAndUpdate() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.FindOneAndUpdate(ctx, tt.filter, tt.update, tt.dest, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("FindOneAndUpdate() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
			_ = mongoTemplate.CloseSession(ctx, helper.IsNotNil(err))
		})
	}
}

func TestTemplateFind(t *testing.T) {
	initDocument()
	for _, tt := range initListTestFind() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.Find(ctx, tt.filter, tt.dest, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateFindAll(t *testing.T) {
	initDocument()
	for _, tt := range initListTestFind() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.FindAll(ctx, tt.dest, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateFindPageable(t *testing.T) {
	initDocument()
	for _, tt := range initListTestFindPageable() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			v, err := mongoTemplate.FindPageable(ctx, tt.filter, tt.pageInput, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("FindPageable() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			} else {
				var destContent []testStruct
				_ = v.Content.Decode(destContent)
				_ = v.Content.Decode(&destContent)
				logger.Info("result pageable:", v)
				if helper.IsGreaterThan(len(v.Content), 0) {
					var destItemContent testStruct
					_ = v.Content[0].Decode(&destItemContent)
					logger.Info("result item content:", destItemContent)
				}
			}
		})
	}
}

func TestTemplateExists(t *testing.T) {
	initDocument()
	for _, tt := range initListTestExists() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			v, err := mongoTemplate.Exists(ctx, tt.filter, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("Exists() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			} else {
				logger.Info("result:", v)
			}
		})
	}
}

func TestTemplateExistsById(t *testing.T) {
	initDocument()
	for _, tt := range initListTestExistsById() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			v, err := mongoTemplate.ExistsById(ctx, tt.id, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("ExistsById() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			} else {
				logger.Info("result:", v)
			}
		})
	}
}

func TestTemplateAggregate(t *testing.T) {
	initDocument()
	for _, tt := range initListTestAggregate() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.Aggregate(ctx, tt.pipeline, tt.dest, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("Aggregate() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateCountDocuments(t *testing.T) {
	initDocument()
	for _, tt := range initListTestCountDocuments() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.CountDocuments(ctx, tt.filter, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("CountDocuments() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateEstimatedDocumentCount(t *testing.T) {
	initDocument()
	for _, tt := range initListTestEstimatedDocumentCount() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.EstimatedDocumentCount(ctx, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("EstimatedDocumentCount() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateDistinct(t *testing.T) {
	initDocument()
	for _, tt := range initListTestDistinct() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.Distinct(ctx, tt.fieldName, tt.filter, tt.dest, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("Distinct() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateWatch(t *testing.T) {
	initDocument()
	for _, tt := range initListTestWatch() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.Watch(ctx, tt.pipeline, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("Watch() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateWatchWithHandler(t *testing.T) {
	initMongoTemplate()
	for _, tt := range initListTestWatchHandler() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			go func() {
				time.Sleep(2 * time.Second)
				initDocument()
			}()
			err := mongoTemplate.WatchWithHandler(ctx, tt.pipeline, tt.handler, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("WatchWithHandler() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateDropCollection(t *testing.T) {
	initDocument()
	time.Sleep(5 * time.Second)
	for _, tt := range initListTestDrop() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.DropCollection(ctx, tt.ref)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("DropCollection() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateDropDatabase(t *testing.T) {
	initDocument()
	time.Sleep(5 * time.Second)
	for _, tt := range initListTestDrop() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.DropDatabase(ctx, tt.ref)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("DropDatabase() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateCreateOneIndex(t *testing.T) {
	initDocument()
	clearIndexes()
	time.Sleep(5 * time.Second)
	for _, tt := range initListTestCreateOneIndex() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.CreateOneIndex(ctx, tt.input)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("CreateOneIndex() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateCreateManyIndex(t *testing.T) {
	initDocument()
	clearIndexes()
	time.Sleep(5 * time.Second)
	for _, tt := range initListTestCreateManyIndex() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			_, err := mongoTemplate.CreateManyIndex(ctx, tt.inputs)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("CreateManyIndex() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateDropOneIndex(t *testing.T) {
	initIndex()
	for _, tt := range initListTestDropIndex() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.DropOneIndex(ctx, tt.nameIndex, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("DropOneIndex() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateDropAllIndexes(t *testing.T) {
	initIndex()
	for _, tt := range initListTestDropIndex() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			err := mongoTemplate.DropAllIndexes(ctx, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("DropAllIndexes() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			}
		})
	}
}

func TestTemplateListIndexes(t *testing.T) {
	initIndex()
	for _, tt := range initListTestListIndexes() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			result, err := mongoTemplate.ListIndexes(ctx, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("ListIndexes() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			} else {
				logger.Info("result list indexes:", result)
			}
		})
	}
}

func TestTemplateListIndexSpecifications(t *testing.T) {
	initIndex()
	for _, tt := range initListTestListIndexes() {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), tt.durationTimeout)
			defer cancel()
			result, err := mongoTemplate.ListIndexSpecifications(ctx, tt.ref, tt.option, nil)
			if helper.IsNotEqualTo(helper.IsNotNil(err), tt.wantErr) {
				t.Errorf("ListIndexSpecifications() error = %v, wantErr %v", err, tt.wantErr)
			} else if helper.IsNotNil(err) {
				t.Log("err expected:", err)
			} else {
				logger.Info("result list indexes:", result)
			}
		})
	}
}

func TestTemplateGetClient(t *testing.T) {
	initMongoTemplate()
	logger.Info("result:", mongoTemplate.GetClient())
}

func TestTemplateCommitTransaction(t *testing.T) {
	initMongoTemplate()
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	err := mongoTemplate.CommitTransaction(ctx)
	logger.Info("result err:", err)
	_ = mongoTemplate.StartSession(ctx)
	err = mongoTemplate.CommitTransaction(ctx)
	logger.Info("result err:", err)
}

func TestTemplateAbortTransaction(t *testing.T) {
	initMongoTemplate()
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	err := mongoTemplate.AbortTransaction(ctx)
	logger.Info("result err:", err)
	_ = mongoTemplate.StartSession(ctx)
	err = mongoTemplate.AbortTransaction(ctx)
	logger.Info("result err:", err)
}
