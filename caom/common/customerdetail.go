package common
import "time"

type CustomerDetail struct{
	ID int
	NAME string
	CREATE_DATE time.Time
	CLOSE_DATE time.Time
	CLOSED int
}