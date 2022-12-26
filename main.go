package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/blugelabs/bluge"
)

func main() {
	config := bluge.DefaultConfig("data/index_dir")
	index, err := bluge.OpenWriter(config)
	if err != nil {
		log.Panic(err)
	}
	defer index.Close()

	// index documents
	if err := indexDocuments(index); err != nil {
		log.Panic(err)
	}

	// search documents
	r, err := index.Reader()
	if err != nil {
		log.Panic(err)
	}
	defer r.Close()
	ctx := context.Background()
	query := bluge.NewBooleanQuery()
	query.AddMust(bluge.NewMatchAllQuery())
	searchReq := bluge.NewTopNSearch(10, query).WithStandardAggregations()
	dmi, err := r.Search(ctx, searchReq)
	if err != nil {
		log.Panic(err)
	}

	hits := make([]map[string]interface{}, 0)
	for {
		doc, err := dmi.Next()
		if err != nil {
			log.Panic(err)
		}
		if doc == nil {
			break
		}
		hit := make(map[string]interface{})
		err = doc.VisitStoredFields(func(field string, value []byte) bool {
			switch field {
			case "_id":
				hit["_id"] = string(value)
			case "@timestamp":
				hit["@timestamp"], _ = bluge.DecodeDateTime(value)
			default:
				hit[field] = string(value)
			}
			return true
		})
		hits = append(hits, hit)
	}

	took := int(dmi.Aggregations().Duration().Milliseconds())
	total := int(dmi.Aggregations().Count())
	log.Printf("took %d ms, found %d documents\n", took, total)
	for _, hit := range hits {
		log.Printf("%+v", hit)
	}
}

func indexDocuments(w *bluge.Writer) error {
	for i := 0; i < 10000; i++ {
		doc := bluge.NewDocument(strconv.Itoa(i))
		doc.AddField(bluge.NewKeywordField("_id", strconv.Itoa(i)).StoreValue().Sortable().Aggregatable())
		doc.AddField(bluge.NewKeywordField("gender", "male").StoreValue().Sortable().Aggregatable())
		doc.AddField(bluge.NewTextField("name", "John Doe").SearchTermPositions().HighlightMatches())
		doc.AddField(bluge.NewTextField("email", "jone@mail.com"))
		doc.AddField(bluge.NewTextField("phone", "1234567890"))
		doc.AddField(bluge.NewDateTimeField("@timestamp", time.Now()))
		err := w.Update(doc.ID(), doc)
		if err != nil {
			return err
		}
		if i%1000 == 0 {
			log.Printf("indexed %d documents", i)
			time.Sleep(time.Millisecond * 100)
		}
	}

	return nil
}
