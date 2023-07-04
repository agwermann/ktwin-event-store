package repository

import (
	"agwermann/dt-service/internal/app/context/twinevent/domain"
	"agwermann/dt-service/internal/app/infra/db"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

var TWIN_EVENT_METADATA = table.Metadata{
	Name:    "twin_event",
	Columns: []string{"instance_id", "interface_id", "event_data", "created_at"},
	PartKey: []string{"interface_id", "instance_id"},
	SortKey: []string{},
}

var TWIN_EVENT_TABLE = table.New(TWIN_EVENT_METADATA)

type TwinEvent struct {
	InstanceId  string    `db:"instance_id"`
	InterfaceId string    `db:"interface_id"`
	EventData   string    `db:"event_data"`
	CreatedAt   time.Time `db:"created_at"`
}

type TwinEventRepository interface {
	GetAllTwinEvents() ([]domain.TwinEvent, error)
	GetTwinEvents(interfaceId string, instanceId string) ([]domain.TwinEvent, error)
	CreateTwinEvent(twinInterface domain.TwinEvent) error
	DeleteTwinEvent(interfaceId string, id string) error
}

func NewTwinEventRepository(
	mapper TwinEventMapper,
	dbConnection db.DBConnection,
) TwinEventRepository {
	return &twinEventRepository{
		dbConnection: dbConnection,
		mapper:       mapper,
	}
}

type twinEventRepository struct {
	dbConnection db.DBConnection
	mapper       TwinEventMapper
}

func (t *twinEventRepository) GetAllTwinEvents() ([]domain.TwinEvent, error) {

	var twinEvents []TwinEvent

	err := t.dbConnection.GetManyWithoutParameters(TWIN_EVENT_TABLE, qb.M{}, &twinEvents)

	if err != nil {
		return []domain.TwinEvent{}, err
	}

	return t.mapper.ToDomainList(twinEvents), err
}

func (t *twinEventRepository) GetTwinEvents(interfaceId string, instanceId string) ([]domain.TwinEvent, error) {
	var twinEvents []TwinEvent

	err := t.dbConnection.GetManyWithParameters(TWIN_EVENT_TABLE, t.getConditions(interfaceId, instanceId), &twinEvents)

	if err != nil {
		return []domain.TwinEvent{}, err
	}

	return t.mapper.ToDomainList(twinEvents), err
}

func (t *twinEventRepository) CreateTwinEvent(twinEvent domain.TwinEvent) error {
	twinInstanceDB := TwinEvent(twinEvent)
	twinInstanceDB.CreatedAt = time.Now()

	return t.dbConnection.InsertQueryDB(TWIN_EVENT_TABLE, twinInstanceDB)
}

func (t *twinEventRepository) DeleteTwinEvent(interfaceId string, instanceId string) error {
	return t.dbConnection.DeleteQuery(TWIN_EVENT_TABLE, t.getConditions(interfaceId, instanceId))
}

func (t *twinEventRepository) getConditions(interfaceId string, instanceId string) qb.M {
	return qb.M{"interface_id": interfaceId, "instance_id": instanceId}
}