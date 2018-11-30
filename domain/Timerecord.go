package domain

type Timerecord struct {
	Timestamp      int64  `json:"timestamp,omitempty" bson:"timestamp,omitempty" firestore:"timestamp,omitempty"`
	Email          string `json:"email,omitempty" bson:"email" firestore:"email,omitempty"`
	StartTime      int64  `json:"start_time,omitempty" bson:"start_time" firestore:"start_time,omitempty"`
	EndTime        int64  `json:"end_time,omitempty" bson:"end_time" firestore:"end_time,omitempty"`
	TimeDifference int64  `json:"time_difference,omitempty" bson:"time_difference" firestore:"time_difference,omitempty"`
	BoatDriver     string `json:"boat_driver,omitempty" bson:"boat_driver" firestore:"boat_driver,omitempty"`
}
