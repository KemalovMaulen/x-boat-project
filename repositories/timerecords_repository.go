package repositories

import (
	"github.com/salambayev/x-boat-project/domain"
	"github.com/salambayev/x-boat-project/db"
	"context"
	"strconv"
	"google.golang.org/api/iterator"
)

type TimerecordsRepository interface {
	CreateTimerecord(timerecord *domain.Timerecord) error

	UpdateTimerecord(timestamp int64, timerecord *domain.Timerecord) error

	DeleteTimerecord(timestamp int64) error

	GetTimerecord(timestamp int64) (*domain.Timerecord, error)

	GetAllTimerecords() ([]*domain.Timerecord, error)
}

type timerecordsRepository struct {

}

func NewTimerecordsRepository() TimerecordsRepository {
	return &timerecordsRepository{}
}

func (tr *timerecordsRepository) CreateTimerecord(timerecord *domain.Timerecord) error {
	_, err := db.TimerecordCollection.Doc(strconv.FormatInt(timerecord.Timestamp, 10)).Create(context.Background(), timerecord)
	return err // TODO timeUid is better to use instead of timestamp
}

func (tr *timerecordsRepository) UpdateTimerecord(timestamp int64, timerecord *domain.Timerecord) error {
	_, err := db.TimerecordCollection.Doc(strconv.FormatInt(timerecord.Timestamp, 10)).Set(context.Background(),timerecord )
	return err //  db.TimerecordCollection.Update(bson.M{"timestamp": timestamp}, &timerecord)
}

func (tr *timerecordsRepository) DeleteTimerecord(timestamp int64) error {
	_, err := db.TimerecordCollection.Doc(strconv.FormatInt(timestamp, 10)).Delete(context.Background())
	return err //db.TimerecordCollection.Remove(bson.M{"timestamp": timestamp})
}

func (tr *timerecordsRepository) GetTimerecord(timestamp int64) (*domain.Timerecord, error) {
	dsnap, err := db.TimerecordCollection.Doc(strconv.FormatInt(timestamp, 10)).Get(context.Background())
	if err != nil {
		return nil, err
	}
	result := &domain.Timerecord{}
	err = dsnap.DataTo(result)
	return result, err
}

func (tr *timerecordsRepository) GetAllTimerecords() ([]*domain.Timerecord, error) {
	var result []*domain.Timerecord

	iter := db.TimerecordCollection.Documents(context.Background())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		timerecord := &domain.Timerecord{}
		doc.DataTo(timerecord)
		result = append(result, timerecord)
	}

	//err := db.TimerecordCollection.Find(nil).Sort("-timestamp").All(&result)
	return result, nil
}