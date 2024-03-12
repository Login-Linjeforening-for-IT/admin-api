// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/guregu/null/zero"
)

type JobType string

const (
	JobTypeFull   JobType = "full"
	JobTypePart   JobType = "part"
	JobTypeSummer JobType = "summer"
	JobTypeVerv   JobType = "verv"
)

func (e *JobType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = JobType(s)
	case string:
		*e = JobType(s)
	default:
		return fmt.Errorf("unsupported scan type for JobType: %T", src)
	}
	return nil
}

type NullJobType struct {
	JobType JobType `json:"job_type"`
	Valid   bool    `json:"valid"` // Valid is true if JobType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullJobType) Scan(value interface{}) error {
	if value == nil {
		ns.JobType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.JobType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullJobType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.JobType), nil
}

func (e JobType) Valid() bool {
	switch e {
	case JobTypeFull,
		JobTypePart,
		JobTypeSummer,
		JobTypeVerv:
		return true
	}
	return false
}

type LocationType string

const (
	LocationTypeMazemap LocationType = "mazemap"
	LocationTypeCoords  LocationType = "coords"
	LocationTypeAddress LocationType = "address"
	LocationTypeNone    LocationType = "none"
)

func (e *LocationType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = LocationType(s)
	case string:
		*e = LocationType(s)
	default:
		return fmt.Errorf("unsupported scan type for LocationType: %T", src)
	}
	return nil
}

type NullLocationType struct {
	LocationType LocationType `json:"location_type"`
	Valid        bool         `json:"valid"` // Valid is true if LocationType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullLocationType) Scan(value interface{}) error {
	if value == nil {
		ns.LocationType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.LocationType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullLocationType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.LocationType), nil
}

func (e LocationType) Valid() bool {
	switch e {
	case LocationTypeMazemap,
		LocationTypeCoords,
		LocationTypeAddress,
		LocationTypeNone:
		return true
	}
	return false
}

type TimeTypeEnum string

const (
	TimeTypeEnumDefault  TimeTypeEnum = "default"
	TimeTypeEnumNoEnd    TimeTypeEnum = "no_end"
	TimeTypeEnumWholeDay TimeTypeEnum = "whole_day"
	TimeTypeEnumTbd      TimeTypeEnum = "tbd"
)

func (e *TimeTypeEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TimeTypeEnum(s)
	case string:
		*e = TimeTypeEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for TimeTypeEnum: %T", src)
	}
	return nil
}

type NullTimeTypeEnum struct {
	TimeTypeEnum TimeTypeEnum `json:"time_type_enum"`
	Valid        bool         `json:"valid"` // Valid is true if TimeTypeEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTimeTypeEnum) Scan(value interface{}) error {
	if value == nil {
		ns.TimeTypeEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TimeTypeEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTimeTypeEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TimeTypeEnum), nil
}

func (e TimeTypeEnum) Valid() bool {
	switch e {
	case TimeTypeEnumDefault,
		TimeTypeEnumNoEnd,
		TimeTypeEnumWholeDay,
		TimeTypeEnumTbd:
		return true
	}
	return false
}

type AdCityRelation struct {
	Ad   int32  `json:"ad"`
	City string `json:"city"`
}

type Audience struct {
	ID            int32       `json:"id"`
	NameNo        string      `json:"name_no"`
	NameEn        zero.String `json:"name_en"`
	DescriptionNo string      `json:"description_no"`
	DescriptionEn zero.String `json:"description_en"`
	UpdatedAt     time.Time   `json:"updated_at"`
	CreatedAt     time.Time   `json:"created_at"`
	DeletedAt     zero.Time   `json:"deleted_at"`
}

type Category struct {
	ID            int32       `json:"id"`
	Color         string      `json:"color"`
	NameNo        string      `json:"name_no"`
	NameEn        zero.String `json:"name_en"`
	DescriptionNo string      `json:"description_no"`
	DescriptionEn zero.String `json:"description_en"`
	UpdatedAt     time.Time   `json:"updated_at"`
	CreatedAt     time.Time   `json:"created_at"`
}

type City struct {
	Name string `json:"name"`
}

type Event struct {
	ID                 int32        `json:"id"`
	Visible            bool         `json:"visible"`
	NameNo             string       `json:"name_no"`
	NameEn             zero.String  `json:"name_en"`
	DescriptionNo      string       `json:"description_no"`
	DescriptionEn      zero.String  `json:"description_en"`
	InformationalNo    zero.String  `json:"informational_no"`
	InformationalEn    zero.String  `json:"informational_en"`
	TimeType           TimeTypeEnum `json:"time_type"`
	TimeStart          time.Time    `json:"time_start"`
	TimeEnd            time.Time    `json:"time_end"`
	TimePublish        zero.Time    `json:"time_publish"`
	TimeSignupRelease  zero.Time    `json:"time_signup_release"`
	TimeSignupDeadline zero.Time    `json:"time_signup_deadline"`
	Canceled           bool         `json:"canceled"`
	Digital            bool         `json:"digital"`
	Highlight          bool         `json:"highlight"`
	ImageSmall         zero.String  `json:"image_small"`
	ImageBanner        string       `json:"image_banner"`
	LinkFacebook       zero.String  `json:"link_facebook"`
	LinkDiscord        zero.String  `json:"link_discord"`
	LinkSignup         zero.String  `json:"link_signup"`
	LinkStream         zero.String  `json:"link_stream"`
	Capacity           zero.Int     `json:"capacity"`
	Full               bool         `json:"full"`
	Category           int32        `json:"category"`
	Location           zero.Int     `json:"location"`
	Parent             zero.Int     `json:"parent"`
	Rule               zero.Int     `json:"rule"`
	UpdatedAt          time.Time    `json:"updated_at"`
	CreatedAt          time.Time    `json:"created_at"`
	DeletedAt          zero.Time    `json:"deleted_at"`
}

type EventAudienceRelation struct {
	Event    int32 `json:"event"`
	Audience int32 `json:"audience"`
}

type EventOrganizationRelation struct {
	Event        int32  `json:"event"`
	Organization string `json:"organization"`
}

type JobAdvertisement struct {
	ID                  int32       `json:"id"`
	Visible             bool        `json:"visible"`
	Highlight           bool        `json:"highlight"`
	TitleNo             string      `json:"title_no"`
	TitleEn             zero.String `json:"title_en"`
	PositionTitleNo     string      `json:"position_title_no"`
	PositionTitleEn     zero.String `json:"position_title_en"`
	DescriptionShortNo  string      `json:"description_short_no"`
	DescriptionShortEn  zero.String `json:"description_short_en"`
	DescriptionLongNo   string      `json:"description_long_no"`
	DescriptionLongEn   zero.String `json:"description_long_en"`
	JobType             JobType     `json:"job_type"`
	TimePublish         time.Time   `json:"time_publish"`
	TimeExpire          time.Time   `json:"time_expire"`
	ApplicationDeadline time.Time   `json:"application_deadline"`
	BannerImage         zero.String `json:"banner_image"`
	Organization        string      `json:"organization"`
	ApplicationUrl      zero.String `json:"application_url"`
	UpdatedAt           time.Time   `json:"updated_at"`
	CreatedAt           time.Time   `json:"created_at"`
	DeletedAt           zero.Time   `json:"deleted_at"`
}

type Location struct {
	ID              int32        `json:"id"`
	NameNo          string       `json:"name_no"`
	NameEn          zero.String  `json:"name_en"`
	Type            LocationType `json:"type"`
	MazemapCampusID zero.Int     `json:"mazemap_campus_id"`
	MazemapPoiID    zero.Int     `json:"mazemap_poi_id"`
	AddressStreet   zero.String  `json:"address_street"`
	AddressPostcode zero.Int     `json:"address_postcode"`
	CityName        zero.String  `json:"city_name"`
	CoordinateLat   zero.Float   `json:"coordinate_lat"`
	CoordinateLong  zero.Float   `json:"coordinate_long"`
	Url             zero.String  `json:"url"`
	UpdatedAt       time.Time    `json:"updated_at"`
	CreatedAt       time.Time    `json:"created_at"`
	DeletedAt       zero.Time    `json:"deleted_at"`
}

type Organization struct {
	Shortname     string      `json:"shortname"`
	NameNo        string      `json:"name_no"`
	NameEn        zero.String `json:"name_en"`
	DescriptionNo string      `json:"description_no"`
	DescriptionEn zero.String `json:"description_en"`
	Type          int32       `json:"type"`
	LinkHomepage  zero.String `json:"link_homepage"`
	LinkLinkedin  zero.String `json:"link_linkedin"`
	LinkFacebook  zero.String `json:"link_facebook"`
	LinkInstagram zero.String `json:"link_instagram"`
	Logo          zero.String `json:"logo"`
	UpdatedAt     time.Time   `json:"updated_at"`
	CreatedAt     time.Time   `json:"created_at"`
	DeletedAt     zero.Time   `json:"deleted_at"`
}

type PublicEvent struct {
	ID                 int32        `json:"id"`
	Visible            bool         `json:"visible"`
	NameNo             string       `json:"name_no"`
	NameEn             zero.String  `json:"name_en"`
	DescriptionNo      string       `json:"description_no"`
	DescriptionEn      zero.String  `json:"description_en"`
	InformationalNo    zero.String  `json:"informational_no"`
	InformationalEn    zero.String  `json:"informational_en"`
	TimeType           TimeTypeEnum `json:"time_type"`
	TimeStart          time.Time    `json:"time_start"`
	TimeEnd            time.Time    `json:"time_end"`
	TimePublish        zero.Time    `json:"time_publish"`
	TimeSignupRelease  zero.Time    `json:"time_signup_release"`
	TimeSignupDeadline zero.Time    `json:"time_signup_deadline"`
	Canceled           bool         `json:"canceled"`
	Digital            bool         `json:"digital"`
	Highlight          bool         `json:"highlight"`
	ImageSmall         zero.String  `json:"image_small"`
	ImageBanner        string       `json:"image_banner"`
	LinkFacebook       zero.String  `json:"link_facebook"`
	LinkDiscord        zero.String  `json:"link_discord"`
	LinkSignup         zero.String  `json:"link_signup"`
	LinkStream         zero.String  `json:"link_stream"`
	Capacity           zero.Int     `json:"capacity"`
	Full               bool         `json:"full"`
	Category           int32        `json:"category"`
	Location           zero.Int     `json:"location"`
	Parent             zero.Int     `json:"parent"`
	Rule               zero.Int     `json:"rule"`
	UpdatedAt          time.Time    `json:"updated_at"`
	CreatedAt          time.Time    `json:"created_at"`
	DeletedAt          zero.Time    `json:"deleted_at"`
}

type PublicJob struct {
	ID                  int32       `json:"id"`
	Visible             bool        `json:"visible"`
	Highlight           bool        `json:"highlight"`
	TitleNo             string      `json:"title_no"`
	TitleEn             zero.String `json:"title_en"`
	PositionTitleNo     string      `json:"position_title_no"`
	PositionTitleEn     zero.String `json:"position_title_en"`
	DescriptionShortNo  string      `json:"description_short_no"`
	DescriptionShortEn  zero.String `json:"description_short_en"`
	DescriptionLongNo   string      `json:"description_long_no"`
	DescriptionLongEn   zero.String `json:"description_long_en"`
	JobType             JobType     `json:"job_type"`
	TimePublish         time.Time   `json:"time_publish"`
	TimeExpire          time.Time   `json:"time_expire"`
	ApplicationDeadline time.Time   `json:"application_deadline"`
	BannerImage         zero.String `json:"banner_image"`
	Organization        string      `json:"organization"`
	ApplicationUrl      zero.String `json:"application_url"`
	UpdatedAt           time.Time   `json:"updated_at"`
	CreatedAt           time.Time   `json:"created_at"`
	DeletedAt           zero.Time   `json:"deleted_at"`
}

type Rule struct {
	ID            int32       `json:"id"`
	NameNo        string      `json:"name_no"`
	NameEn        zero.String `json:"name_en"`
	DescriptionNo string      `json:"description_no"`
	DescriptionEn zero.String `json:"description_en"`
	UpdatedAt     time.Time   `json:"updated_at"`
	CreatedAt     time.Time   `json:"created_at"`
	DeletedAt     zero.Time   `json:"deleted_at"`
}

type Skill struct {
	Ad    int32  `json:"ad"`
	Skill string `json:"skill"`
}

type VisibleEvent struct {
	ID                 int32        `json:"id"`
	Visible            bool         `json:"visible"`
	NameNo             string       `json:"name_no"`
	NameEn             zero.String  `json:"name_en"`
	DescriptionNo      string       `json:"description_no"`
	DescriptionEn      zero.String  `json:"description_en"`
	InformationalNo    zero.String  `json:"informational_no"`
	InformationalEn    zero.String  `json:"informational_en"`
	TimeType           TimeTypeEnum `json:"time_type"`
	TimeStart          time.Time    `json:"time_start"`
	TimeEnd            time.Time    `json:"time_end"`
	TimePublish        zero.Time    `json:"time_publish"`
	TimeSignupRelease  zero.Time    `json:"time_signup_release"`
	TimeSignupDeadline zero.Time    `json:"time_signup_deadline"`
	Canceled           bool         `json:"canceled"`
	Digital            bool         `json:"digital"`
	Highlight          bool         `json:"highlight"`
	ImageSmall         zero.String  `json:"image_small"`
	ImageBanner        string       `json:"image_banner"`
	LinkFacebook       zero.String  `json:"link_facebook"`
	LinkDiscord        zero.String  `json:"link_discord"`
	LinkSignup         zero.String  `json:"link_signup"`
	LinkStream         zero.String  `json:"link_stream"`
	Capacity           zero.Int     `json:"capacity"`
	Full               bool         `json:"full"`
	Category           int32        `json:"category"`
	Location           zero.Int     `json:"location"`
	Parent             zero.Int     `json:"parent"`
	Rule               zero.Int     `json:"rule"`
	UpdatedAt          time.Time    `json:"updated_at"`
	CreatedAt          time.Time    `json:"created_at"`
	DeletedAt          zero.Time    `json:"deleted_at"`
}

type VisibleJob struct {
	ID                  int32       `json:"id"`
	Visible             bool        `json:"visible"`
	Highlight           bool        `json:"highlight"`
	TitleNo             string      `json:"title_no"`
	TitleEn             zero.String `json:"title_en"`
	PositionTitleNo     string      `json:"position_title_no"`
	PositionTitleEn     zero.String `json:"position_title_en"`
	DescriptionShortNo  string      `json:"description_short_no"`
	DescriptionShortEn  zero.String `json:"description_short_en"`
	DescriptionLongNo   string      `json:"description_long_no"`
	DescriptionLongEn   zero.String `json:"description_long_en"`
	JobType             JobType     `json:"job_type"`
	TimePublish         time.Time   `json:"time_publish"`
	TimeExpire          time.Time   `json:"time_expire"`
	ApplicationDeadline time.Time   `json:"application_deadline"`
	BannerImage         zero.String `json:"banner_image"`
	Organization        string      `json:"organization"`
	ApplicationUrl      zero.String `json:"application_url"`
	UpdatedAt           time.Time   `json:"updated_at"`
	CreatedAt           time.Time   `json:"created_at"`
	DeletedAt           zero.Time   `json:"deleted_at"`
}
