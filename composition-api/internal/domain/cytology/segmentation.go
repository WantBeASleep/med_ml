package cytology

import (
	"time"

	"github.com/google/uuid"
)

type SegType string

const (
	SegTypeNIL SegType = "NIL"
	SegTypeNIR SegType = "NIR"
	SegTypeNIM SegType = "NIM"
	SegTypeCNO SegType = "CNO"
	SegTypeCGE SegType = "CGE"
	SegTypeC2N SegType = "C2N"
	SegTypeCPS SegType = "CPS"
	SegTypeCFC SegType = "CFC"
	SegTypeCLY SegType = "CLY"
	SegTypeSOS SegType = "SOS"
	SegTypeSDS SegType = "SDS"
	SegTypeSMS SegType = "SMS"
	SegTypeSTS SegType = "STS"
	SegTypeSPS SegType = "SPS"
	SegTypeSNM SegType = "SNM"
	SegTypeSTM SegType = "STM"
)

type GroupType string

const (
	GroupTypeCE GroupType = "CE"
	GroupTypeCL GroupType = "CL"
	GroupTypeME GroupType = "ME"
)

type SegmentationPoint struct {
	Id             int
	SegmentationID int
	X              int
	Y              int
	UID            int64
	CreateAt       time.Time
}

type Segmentation struct {
	Id                  int
	SegmentationGroupID int
	Points              []SegmentationPoint
	CreateAt            time.Time
}

type SegmentationGroup struct {
	Id         int
	CytologyID uuid.UUID
	SegType    SegType
	GroupType  GroupType
	IsAI       bool
	Details    *string
	CreateAt   time.Time
}
