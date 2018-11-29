package domain

type Timerecord struct {
	Timestamp int64 `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Email string `json:"email,omitempty" bson:"email"`
	StartTime int64 `json:"start_time,omitempty" bson:"start_time"`
	EndTime int64 `json:"end_time,omitempty" bson:"end_time"`
	TimeDifference int64 `json:"time_difference,omitempty" bson:"time_difference"`
	BoatDriver string `json:"boat_driver,omitempty" bson:"boat_driver"`
}
