// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"
)

type Querier interface {
	AddAudienceToEvent(ctx context.Context, arg AddAudienceToEventParams) error
	AddCityToJob(ctx context.Context, arg AddCityToJobParams) error
	AddOrganizationToEvent(ctx context.Context, arg AddOrganizationToEventParams) error
	AddSkillToJob(ctx context.Context, arg AddSkillToJobParams) error
	// sqlc.embed(event)
	CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error)
	CreateJob(ctx context.Context, arg CreateJobParams) (JobAdvertisement, error)
	CreateLocation(ctx context.Context, arg CreateLocationParams) (Location, error)
	CreateOrganization(ctx context.Context, arg CreateOrganizationParams) (Organization, error)
	CreateRule(ctx context.Context, arg CreateRuleParams) (Rule, error)
	GetAddressLocations(ctx context.Context, arg GetAddressLocationsParams) ([]GetAddressLocationsRow, error)
	GetAllCities(ctx context.Context) ([]string, error)
	GetAudience(ctx context.Context, id int32) (Audience, error)
	GetAudiences(ctx context.Context) ([]GetAudiencesRow, error)
	GetAudiencesOfEvent(ctx context.Context, eventID int32) ([]Audience, error)
	GetCategories(ctx context.Context) ([]GetCategoriesRow, error)
	GetCategory(ctx context.Context, id int32) (Category, error)
	GetCoordsLocations(ctx context.Context, arg GetCoordsLocationsParams) ([]GetCoordsLocationsRow, error)
	GetEvent(ctx context.Context, id int32) (Event, error)
	GetEvents(ctx context.Context, arg GetEventsParams) ([]GetEventsRow, error)
	GetJob(ctx context.Context, id int32) (GetJobRow, error)
	GetJobs(ctx context.Context, arg GetJobsParams) ([]GetJobsRow, error)
	GetLocation(ctx context.Context, id int32) (Location, error)
	GetLocations(ctx context.Context, arg GetLocationsParams) ([]GetLocationsRow, error)
	GetMazemapLocations(ctx context.Context, arg GetMazemapLocationsParams) ([]GetMazemapLocationsRow, error)
	GetOrganization(ctx context.Context, shortname string) (Organization, error)
	GetOrganizations(ctx context.Context, arg GetOrganizationsParams) ([]GetOrganizationsRow, error)
	GetOrganizationsOfEvent(ctx context.Context, eventID int32) ([]GetOrganizationsOfEventRow, error)
	GetRule(ctx context.Context, id int32) (Rule, error)
	GetRules(ctx context.Context, arg GetRulesParams) ([]GetRulesRow, error)
	RemoveAudienceFromEvent(ctx context.Context, arg RemoveAudienceFromEventParams) error
	RemoveCityFromJob(ctx context.Context, arg RemoveCityFromJobParams) error
	RemoveOrganizationFromEvent(ctx context.Context, arg RemoveOrganizationFromEventParams) error
	RemoveSkillFromJob(ctx context.Context, arg RemoveSkillFromJobParams) error
	SoftDeleteEvent(ctx context.Context, id int32) (Event, error)
	SoftDeleteJob(ctx context.Context, id int32) (JobAdvertisement, error)
	SoftDeleteLocation(ctx context.Context, id int32) (Location, error)
	SoftDeleteOrganization(ctx context.Context, shortname string) (Organization, error)
	SoftDeleteRule(ctx context.Context, id int32) (Rule, error)
	UpdateEvent(ctx context.Context, arg UpdateEventParams) (Event, error)
	UpdateJob(ctx context.Context, arg UpdateJobParams) (JobAdvertisement, error)
	UpdateLocation(ctx context.Context, arg UpdateLocationParams) (Location, error)
	UpdateOrganization(ctx context.Context, arg UpdateOrganizationParams) (Organization, error)
	UpdateRule(ctx context.Context, arg UpdateRuleParams) (Rule, error)
}

var _ Querier = (*Queries)(nil)
