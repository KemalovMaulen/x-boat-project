package repositories
//
//import (
//	"github.com/salambayev/x-boat-project/domain"
//	"github.com/salambayev/x-boat-project/db"
//	"gopkg.in/mgo.v2/bson"
//)
//
//type TimerecordsRepository interface {
//	CreateTimerecord(timerecord *domain.Timerecord) error
//
//	UpdateTimerecord(timestamp int64, timerecord *domain.Timerecord) error
//
//	DeleteTimerecord(timestamp int64) error
//
//	GetTimerecord(timestamp int64) (*domain.Timerecord, error)
//
//	GetAllTimerecords() ([]*domain.Timerecord, error)
//}
//
//type timerecordsRepository struct {
//
//}
//
//func NewTimerecordsRepository() TimerecordsRepository {
//	return &timerecordsRepository{}
//}
//
//func (tr *timerecordsRepository) CreateTimerecord(timerecord *domain.Timerecord) error {
//	return db.TimerecordCollection.Insert(&timerecord)
//}
//
//func (tr *timerecordsRepository) UpdateTimerecord(timestamp int64, timerecord *domain.Timerecord) error {
//	return db.TimerecordCollection.Update(bson.M{"timestamp": timestamp}, &timerecord)
//}
//
//func (tr *timerecordsRepository) DeleteTimerecord(timestamp int64) error {
//	return db.TimerecordCollection.Remove(bson.M{"timestamp": timestamp})
//}
//
//func (tr *timerecordsRepository) GetTimerecord(timestamp int64) (*domain.Timerecord, error) {
//	result := domain.Timerecord{}
//	err := db.TimerecordCollection.Find(bson.M{"timestamp": timestamp}).One(&result)
//	return &result, err
//}
//
//func (tr *timerecordsRepository) GetAllTimerecords() ([]*domain.Timerecord, error) {
//	var result []*domain.Timerecord
//	err := db.TimerecordCollection.Find(nil).Sort("-timestamp").All(&result)
//	return result, err
//}