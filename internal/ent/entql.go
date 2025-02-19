// Code generated by ent, DO NOT EDIT.

package ent

import (
	"go-ent-project/internal/ent/camera"
	"go-ent-project/internal/ent/car"
	"go-ent-project/internal/ent/event"
	"go-ent-project/internal/ent/permission"
	"go-ent-project/internal/ent/policestation"
	"go-ent-project/internal/ent/predicate"
	"go-ent-project/internal/ent/role"
	"go-ent-project/internal/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 7)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   camera.Table,
			Columns: camera.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: camera.FieldID,
			},
		},
		Type: "Camera",
		Fields: map[string]*sqlgraph.FieldSpec{
			camera.FieldCreatedAt:       {Type: field.TypeTime, Column: camera.FieldCreatedAt},
			camera.FieldUpdatedAt:       {Type: field.TypeTime, Column: camera.FieldUpdatedAt},
			camera.FieldName:            {Type: field.TypeString, Column: camera.FieldName},
			camera.FieldModel:           {Type: field.TypeString, Column: camera.FieldModel},
			camera.FieldImei:            {Type: field.TypeString, Column: camera.FieldImei},
			camera.FieldLocation:        {Type: field.TypeString, Column: camera.FieldLocation},
			camera.FieldActive:          {Type: field.TypeBool, Column: camera.FieldActive},
			camera.FieldPoliceStationID: {Type: field.TypeUUID, Column: camera.FieldPoliceStationID},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   car.Table,
			Columns: car.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: car.FieldID,
			},
		},
		Type: "Car",
		Fields: map[string]*sqlgraph.FieldSpec{
			car.FieldCreatedAt:       {Type: field.TypeTime, Column: car.FieldCreatedAt},
			car.FieldUpdatedAt:       {Type: field.TypeTime, Column: car.FieldUpdatedAt},
			car.FieldMake:            {Type: field.TypeString, Column: car.FieldMake},
			car.FieldModel:           {Type: field.TypeString, Column: car.FieldModel},
			car.FieldYear:            {Type: field.TypeInt, Column: car.FieldYear},
			car.FieldRegistration:    {Type: field.TypeString, Column: car.FieldRegistration},
			car.FieldColor:           {Type: field.TypeString, Column: car.FieldColor},
			car.FieldPoliceStationID: {Type: field.TypeUUID, Column: car.FieldPoliceStationID},
			car.FieldStolenDate:      {Type: field.TypeTime, Column: car.FieldStolenDate},
			car.FieldFirNumber:       {Type: field.TypeString, Column: car.FieldFirNumber},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: event.FieldID,
			},
		},
		Type: "Event",
		Fields: map[string]*sqlgraph.FieldSpec{
			event.FieldCreatedAt:            {Type: field.TypeTime, Column: event.FieldCreatedAt},
			event.FieldUpdatedAt:            {Type: field.TypeTime, Column: event.FieldUpdatedAt},
			event.FieldPlateBoundingBox:     {Type: field.TypeJSON, Column: event.FieldPlateBoundingBox},
			event.FieldPlateChannel:         {Type: field.TypeInt, Column: event.FieldPlateChannel},
			event.FieldPlateIsExist:         {Type: field.TypeBool, Column: event.FieldPlateIsExist},
			event.FieldPlateColor:           {Type: field.TypeString, Column: event.FieldPlateColor},
			event.FieldPlateNumber:          {Type: field.TypeString, Column: event.FieldPlateNumber},
			event.FieldPlateType:            {Type: field.TypeString, Column: event.FieldPlateType},
			event.FieldPlateRegion:          {Type: field.TypeString, Column: event.FieldPlateRegion},
			event.FieldPlateUploadNum:       {Type: field.TypeInt, Column: event.FieldPlateUploadNum},
			event.FieldSnapAllowUser:        {Type: field.TypeBool, Column: event.FieldSnapAllowUser},
			event.FieldSnapAllowUserEndTime: {Type: field.TypeString, Column: event.FieldSnapAllowUserEndTime},
			event.FieldSnapDefenceCode:      {Type: field.TypeString, Column: event.FieldSnapDefenceCode},
			event.FieldSnapDeviceID:         {Type: field.TypeString, Column: event.FieldSnapDeviceID},
			event.FieldSnapInCarPeopleNum:   {Type: field.TypeInt, Column: event.FieldSnapInCarPeopleNum},
			event.FieldSnapLanNo:            {Type: field.TypeInt, Column: event.FieldSnapLanNo},
			event.FieldSnapOpenStrobe:       {Type: field.TypeBool, Column: event.FieldSnapOpenStrobe},
			event.FieldVehicleBoundingBox:   {Type: field.TypeJSON, Column: event.FieldVehicleBoundingBox},
			event.FieldVehicleSeries:        {Type: field.TypeString, Column: event.FieldVehicleSeries},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   permission.Table,
			Columns: permission.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: permission.FieldID,
			},
		},
		Type: "Permission",
		Fields: map[string]*sqlgraph.FieldSpec{
			permission.FieldCreatedAt: {Type: field.TypeTime, Column: permission.FieldCreatedAt},
			permission.FieldUpdatedAt: {Type: field.TypeTime, Column: permission.FieldUpdatedAt},
			permission.FieldName:      {Type: field.TypeString, Column: permission.FieldName},
			permission.FieldCanRead:   {Type: field.TypeBool, Column: permission.FieldCanRead},
			permission.FieldCanCreate: {Type: field.TypeBool, Column: permission.FieldCanCreate},
			permission.FieldCanUpdate: {Type: field.TypeBool, Column: permission.FieldCanUpdate},
			permission.FieldCanDelete: {Type: field.TypeBool, Column: permission.FieldCanDelete},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   policestation.Table,
			Columns: policestation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: policestation.FieldID,
			},
		},
		Type: "PoliceStation",
		Fields: map[string]*sqlgraph.FieldSpec{
			policestation.FieldCreatedAt:       {Type: field.TypeTime, Column: policestation.FieldCreatedAt},
			policestation.FieldUpdatedAt:       {Type: field.TypeTime, Column: policestation.FieldUpdatedAt},
			policestation.FieldName:            {Type: field.TypeString, Column: policestation.FieldName},
			policestation.FieldLocation:        {Type: field.TypeString, Column: policestation.FieldLocation},
			policestation.FieldCode:            {Type: field.TypeString, Column: policestation.FieldCode},
			policestation.FieldIdentifier:      {Type: field.TypeString, Column: policestation.FieldIdentifier},
			policestation.FieldParentStationID: {Type: field.TypeUUID, Column: policestation.FieldParentStationID},
		},
	}
	graph.Nodes[5] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   role.Table,
			Columns: role.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: role.FieldID,
			},
		},
		Type: "Role",
		Fields: map[string]*sqlgraph.FieldSpec{
			role.FieldCreatedAt: {Type: field.TypeTime, Column: role.FieldCreatedAt},
			role.FieldUpdatedAt: {Type: field.TypeTime, Column: role.FieldUpdatedAt},
			role.FieldName:      {Type: field.TypeString, Column: role.FieldName},
		},
	}
	graph.Nodes[6] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldCreatedAt:       {Type: field.TypeTime, Column: user.FieldCreatedAt},
			user.FieldUpdatedAt:       {Type: field.TypeTime, Column: user.FieldUpdatedAt},
			user.FieldName:            {Type: field.TypeString, Column: user.FieldName},
			user.FieldEmail:           {Type: field.TypeString, Column: user.FieldEmail},
			user.FieldPassword:        {Type: field.TypeString, Column: user.FieldPassword},
			user.FieldPhone:           {Type: field.TypeString, Column: user.FieldPhone},
			user.FieldActive:          {Type: field.TypeBool, Column: user.FieldActive},
			user.FieldRoleID:          {Type: field.TypeUUID, Column: user.FieldRoleID},
			user.FieldPoliceStationID: {Type: field.TypeUUID, Column: user.FieldPoliceStationID},
		},
	}
	graph.MustAddE(
		"police_station",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   camera.PoliceStationTable,
			Columns: []string{camera.PoliceStationColumn},
			Bidi:    false,
		},
		"Camera",
		"PoliceStation",
	)
	graph.MustAddE(
		"police_station",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.PoliceStationTable,
			Columns: []string{car.PoliceStationColumn},
			Bidi:    false,
		},
		"Car",
		"PoliceStation",
	)
	graph.MustAddE(
		"users",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   policestation.UsersTable,
			Columns: []string{policestation.UsersColumn},
			Bidi:    false,
		},
		"PoliceStation",
		"User",
	)
	graph.MustAddE(
		"camera",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   policestation.CameraTable,
			Columns: []string{policestation.CameraColumn},
			Bidi:    false,
		},
		"PoliceStation",
		"Camera",
	)
	graph.MustAddE(
		"car",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   policestation.CarTable,
			Columns: []string{policestation.CarColumn},
			Bidi:    false,
		},
		"PoliceStation",
		"Car",
	)
	graph.MustAddE(
		"parent",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   policestation.ParentTable,
			Columns: []string{policestation.ParentColumn},
			Bidi:    false,
		},
		"PoliceStation",
		"PoliceStation",
	)
	graph.MustAddE(
		"child_stations",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   policestation.ChildStationsTable,
			Columns: []string{policestation.ChildStationsColumn},
			Bidi:    false,
		},
		"PoliceStation",
		"PoliceStation",
	)
	graph.MustAddE(
		"permissions",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.PermissionsTable,
			Columns: []string{role.PermissionsColumn},
			Bidi:    false,
		},
		"Role",
		"Permission",
	)
	graph.MustAddE(
		"users",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   role.UsersTable,
			Columns: []string{role.UsersColumn},
			Bidi:    false,
		},
		"Role",
		"User",
	)
	graph.MustAddE(
		"role",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.RoleTable,
			Columns: []string{user.RoleColumn},
			Bidi:    false,
		},
		"User",
		"Role",
	)
	graph.MustAddE(
		"police_station",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.PoliceStationTable,
			Columns: []string{user.PoliceStationColumn},
			Bidi:    false,
		},
		"User",
		"PoliceStation",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (cq *CameraQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CameraQuery builder.
func (cq *CameraQuery) Filter() *CameraFilter {
	return &CameraFilter{config: cq.config, predicateAdder: cq}
}

// addPredicate implements the predicateAdder interface.
func (m *CameraMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CameraMutation builder.
func (m *CameraMutation) Filter() *CameraFilter {
	return &CameraFilter{config: m.config, predicateAdder: m}
}

// CameraFilter provides a generic filtering capability at runtime for CameraQuery.
type CameraFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CameraFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CameraFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(camera.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *CameraFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(camera.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *CameraFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(camera.FieldUpdatedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *CameraFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(camera.FieldName))
}

// WhereModel applies the entql string predicate on the model field.
func (f *CameraFilter) WhereModel(p entql.StringP) {
	f.Where(p.Field(camera.FieldModel))
}

// WhereImei applies the entql string predicate on the imei field.
func (f *CameraFilter) WhereImei(p entql.StringP) {
	f.Where(p.Field(camera.FieldImei))
}

// WhereLocation applies the entql string predicate on the location field.
func (f *CameraFilter) WhereLocation(p entql.StringP) {
	f.Where(p.Field(camera.FieldLocation))
}

// WhereActive applies the entql bool predicate on the active field.
func (f *CameraFilter) WhereActive(p entql.BoolP) {
	f.Where(p.Field(camera.FieldActive))
}

// WherePoliceStationID applies the entql [16]byte predicate on the police_station_id field.
func (f *CameraFilter) WherePoliceStationID(p entql.ValueP) {
	f.Where(p.Field(camera.FieldPoliceStationID))
}

// WhereHasPoliceStation applies a predicate to check if query has an edge police_station.
func (f *CameraFilter) WhereHasPoliceStation() {
	f.Where(entql.HasEdge("police_station"))
}

// WhereHasPoliceStationWith applies a predicate to check if query has an edge police_station with a given conditions (other predicates).
func (f *CameraFilter) WhereHasPoliceStationWith(preds ...predicate.PoliceStation) {
	f.Where(entql.HasEdgeWith("police_station", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (cq *CarQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CarQuery builder.
func (cq *CarQuery) Filter() *CarFilter {
	return &CarFilter{config: cq.config, predicateAdder: cq}
}

// addPredicate implements the predicateAdder interface.
func (m *CarMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CarMutation builder.
func (m *CarMutation) Filter() *CarFilter {
	return &CarFilter{config: m.config, predicateAdder: m}
}

// CarFilter provides a generic filtering capability at runtime for CarQuery.
type CarFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CarFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CarFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(car.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *CarFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(car.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *CarFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(car.FieldUpdatedAt))
}

// WhereMake applies the entql string predicate on the make field.
func (f *CarFilter) WhereMake(p entql.StringP) {
	f.Where(p.Field(car.FieldMake))
}

// WhereModel applies the entql string predicate on the model field.
func (f *CarFilter) WhereModel(p entql.StringP) {
	f.Where(p.Field(car.FieldModel))
}

// WhereYear applies the entql int predicate on the year field.
func (f *CarFilter) WhereYear(p entql.IntP) {
	f.Where(p.Field(car.FieldYear))
}

// WhereRegistration applies the entql string predicate on the registration field.
func (f *CarFilter) WhereRegistration(p entql.StringP) {
	f.Where(p.Field(car.FieldRegistration))
}

// WhereColor applies the entql string predicate on the color field.
func (f *CarFilter) WhereColor(p entql.StringP) {
	f.Where(p.Field(car.FieldColor))
}

// WherePoliceStationID applies the entql [16]byte predicate on the police_station_id field.
func (f *CarFilter) WherePoliceStationID(p entql.ValueP) {
	f.Where(p.Field(car.FieldPoliceStationID))
}

// WhereStolenDate applies the entql time.Time predicate on the stolen_date field.
func (f *CarFilter) WhereStolenDate(p entql.TimeP) {
	f.Where(p.Field(car.FieldStolenDate))
}

// WhereFirNumber applies the entql string predicate on the fir_number field.
func (f *CarFilter) WhereFirNumber(p entql.StringP) {
	f.Where(p.Field(car.FieldFirNumber))
}

// WhereHasPoliceStation applies a predicate to check if query has an edge police_station.
func (f *CarFilter) WhereHasPoliceStation() {
	f.Where(entql.HasEdge("police_station"))
}

// WhereHasPoliceStationWith applies a predicate to check if query has an edge police_station with a given conditions (other predicates).
func (f *CarFilter) WhereHasPoliceStationWith(preds ...predicate.PoliceStation) {
	f.Where(entql.HasEdgeWith("police_station", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (eq *EventQuery) addPredicate(pred func(s *sql.Selector)) {
	eq.predicates = append(eq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the EventQuery builder.
func (eq *EventQuery) Filter() *EventFilter {
	return &EventFilter{config: eq.config, predicateAdder: eq}
}

// addPredicate implements the predicateAdder interface.
func (m *EventMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the EventMutation builder.
func (m *EventMutation) Filter() *EventFilter {
	return &EventFilter{config: m.config, predicateAdder: m}
}

// EventFilter provides a generic filtering capability at runtime for EventQuery.
type EventFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *EventFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *EventFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(event.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *EventFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(event.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *EventFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(event.FieldUpdatedAt))
}

// WherePlateBoundingBox applies the entql json.RawMessage predicate on the plate_bounding_box field.
func (f *EventFilter) WherePlateBoundingBox(p entql.BytesP) {
	f.Where(p.Field(event.FieldPlateBoundingBox))
}

// WherePlateChannel applies the entql int predicate on the plate_channel field.
func (f *EventFilter) WherePlateChannel(p entql.IntP) {
	f.Where(p.Field(event.FieldPlateChannel))
}

// WherePlateIsExist applies the entql bool predicate on the plate_is_exist field.
func (f *EventFilter) WherePlateIsExist(p entql.BoolP) {
	f.Where(p.Field(event.FieldPlateIsExist))
}

// WherePlateColor applies the entql string predicate on the plate_color field.
func (f *EventFilter) WherePlateColor(p entql.StringP) {
	f.Where(p.Field(event.FieldPlateColor))
}

// WherePlateNumber applies the entql string predicate on the plate_number field.
func (f *EventFilter) WherePlateNumber(p entql.StringP) {
	f.Where(p.Field(event.FieldPlateNumber))
}

// WherePlateType applies the entql string predicate on the plate_type field.
func (f *EventFilter) WherePlateType(p entql.StringP) {
	f.Where(p.Field(event.FieldPlateType))
}

// WherePlateRegion applies the entql string predicate on the plate_region field.
func (f *EventFilter) WherePlateRegion(p entql.StringP) {
	f.Where(p.Field(event.FieldPlateRegion))
}

// WherePlateUploadNum applies the entql int predicate on the plate_upload_num field.
func (f *EventFilter) WherePlateUploadNum(p entql.IntP) {
	f.Where(p.Field(event.FieldPlateUploadNum))
}

// WhereSnapAllowUser applies the entql bool predicate on the snap_allow_user field.
func (f *EventFilter) WhereSnapAllowUser(p entql.BoolP) {
	f.Where(p.Field(event.FieldSnapAllowUser))
}

// WhereSnapAllowUserEndTime applies the entql string predicate on the snap_allow_user_end_time field.
func (f *EventFilter) WhereSnapAllowUserEndTime(p entql.StringP) {
	f.Where(p.Field(event.FieldSnapAllowUserEndTime))
}

// WhereSnapDefenceCode applies the entql string predicate on the snap_defence_code field.
func (f *EventFilter) WhereSnapDefenceCode(p entql.StringP) {
	f.Where(p.Field(event.FieldSnapDefenceCode))
}

// WhereSnapDeviceID applies the entql string predicate on the snap_device_id field.
func (f *EventFilter) WhereSnapDeviceID(p entql.StringP) {
	f.Where(p.Field(event.FieldSnapDeviceID))
}

// WhereSnapInCarPeopleNum applies the entql int predicate on the snap_in_car_people_num field.
func (f *EventFilter) WhereSnapInCarPeopleNum(p entql.IntP) {
	f.Where(p.Field(event.FieldSnapInCarPeopleNum))
}

// WhereSnapLanNo applies the entql int predicate on the snap_lan_no field.
func (f *EventFilter) WhereSnapLanNo(p entql.IntP) {
	f.Where(p.Field(event.FieldSnapLanNo))
}

// WhereSnapOpenStrobe applies the entql bool predicate on the snap_open_strobe field.
func (f *EventFilter) WhereSnapOpenStrobe(p entql.BoolP) {
	f.Where(p.Field(event.FieldSnapOpenStrobe))
}

// WhereVehicleBoundingBox applies the entql json.RawMessage predicate on the vehicle_bounding_box field.
func (f *EventFilter) WhereVehicleBoundingBox(p entql.BytesP) {
	f.Where(p.Field(event.FieldVehicleBoundingBox))
}

// WhereVehicleSeries applies the entql string predicate on the vehicle_series field.
func (f *EventFilter) WhereVehicleSeries(p entql.StringP) {
	f.Where(p.Field(event.FieldVehicleSeries))
}

// addPredicate implements the predicateAdder interface.
func (pq *PermissionQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PermissionQuery builder.
func (pq *PermissionQuery) Filter() *PermissionFilter {
	return &PermissionFilter{config: pq.config, predicateAdder: pq}
}

// addPredicate implements the predicateAdder interface.
func (m *PermissionMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PermissionMutation builder.
func (m *PermissionMutation) Filter() *PermissionFilter {
	return &PermissionFilter{config: m.config, predicateAdder: m}
}

// PermissionFilter provides a generic filtering capability at runtime for PermissionQuery.
type PermissionFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PermissionFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *PermissionFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(permission.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *PermissionFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(permission.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *PermissionFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(permission.FieldUpdatedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *PermissionFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(permission.FieldName))
}

// WhereCanRead applies the entql bool predicate on the can_read field.
func (f *PermissionFilter) WhereCanRead(p entql.BoolP) {
	f.Where(p.Field(permission.FieldCanRead))
}

// WhereCanCreate applies the entql bool predicate on the can_create field.
func (f *PermissionFilter) WhereCanCreate(p entql.BoolP) {
	f.Where(p.Field(permission.FieldCanCreate))
}

// WhereCanUpdate applies the entql bool predicate on the can_update field.
func (f *PermissionFilter) WhereCanUpdate(p entql.BoolP) {
	f.Where(p.Field(permission.FieldCanUpdate))
}

// WhereCanDelete applies the entql bool predicate on the can_delete field.
func (f *PermissionFilter) WhereCanDelete(p entql.BoolP) {
	f.Where(p.Field(permission.FieldCanDelete))
}

// addPredicate implements the predicateAdder interface.
func (psq *PoliceStationQuery) addPredicate(pred func(s *sql.Selector)) {
	psq.predicates = append(psq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PoliceStationQuery builder.
func (psq *PoliceStationQuery) Filter() *PoliceStationFilter {
	return &PoliceStationFilter{config: psq.config, predicateAdder: psq}
}

// addPredicate implements the predicateAdder interface.
func (m *PoliceStationMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PoliceStationMutation builder.
func (m *PoliceStationMutation) Filter() *PoliceStationFilter {
	return &PoliceStationFilter{config: m.config, predicateAdder: m}
}

// PoliceStationFilter provides a generic filtering capability at runtime for PoliceStationQuery.
type PoliceStationFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PoliceStationFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *PoliceStationFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(policestation.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *PoliceStationFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(policestation.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *PoliceStationFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(policestation.FieldUpdatedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *PoliceStationFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(policestation.FieldName))
}

// WhereLocation applies the entql string predicate on the location field.
func (f *PoliceStationFilter) WhereLocation(p entql.StringP) {
	f.Where(p.Field(policestation.FieldLocation))
}

// WhereCode applies the entql string predicate on the code field.
func (f *PoliceStationFilter) WhereCode(p entql.StringP) {
	f.Where(p.Field(policestation.FieldCode))
}

// WhereIdentifier applies the entql string predicate on the identifier field.
func (f *PoliceStationFilter) WhereIdentifier(p entql.StringP) {
	f.Where(p.Field(policestation.FieldIdentifier))
}

// WhereParentStationID applies the entql [16]byte predicate on the parent_station_id field.
func (f *PoliceStationFilter) WhereParentStationID(p entql.ValueP) {
	f.Where(p.Field(policestation.FieldParentStationID))
}

// WhereHasUsers applies a predicate to check if query has an edge users.
func (f *PoliceStationFilter) WhereHasUsers() {
	f.Where(entql.HasEdge("users"))
}

// WhereHasUsersWith applies a predicate to check if query has an edge users with a given conditions (other predicates).
func (f *PoliceStationFilter) WhereHasUsersWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("users", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasCamera applies a predicate to check if query has an edge camera.
func (f *PoliceStationFilter) WhereHasCamera() {
	f.Where(entql.HasEdge("camera"))
}

// WhereHasCameraWith applies a predicate to check if query has an edge camera with a given conditions (other predicates).
func (f *PoliceStationFilter) WhereHasCameraWith(preds ...predicate.Camera) {
	f.Where(entql.HasEdgeWith("camera", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasCar applies a predicate to check if query has an edge car.
func (f *PoliceStationFilter) WhereHasCar() {
	f.Where(entql.HasEdge("car"))
}

// WhereHasCarWith applies a predicate to check if query has an edge car with a given conditions (other predicates).
func (f *PoliceStationFilter) WhereHasCarWith(preds ...predicate.Car) {
	f.Where(entql.HasEdgeWith("car", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasParent applies a predicate to check if query has an edge parent.
func (f *PoliceStationFilter) WhereHasParent() {
	f.Where(entql.HasEdge("parent"))
}

// WhereHasParentWith applies a predicate to check if query has an edge parent with a given conditions (other predicates).
func (f *PoliceStationFilter) WhereHasParentWith(preds ...predicate.PoliceStation) {
	f.Where(entql.HasEdgeWith("parent", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasChildStations applies a predicate to check if query has an edge child_stations.
func (f *PoliceStationFilter) WhereHasChildStations() {
	f.Where(entql.HasEdge("child_stations"))
}

// WhereHasChildStationsWith applies a predicate to check if query has an edge child_stations with a given conditions (other predicates).
func (f *PoliceStationFilter) WhereHasChildStationsWith(preds ...predicate.PoliceStation) {
	f.Where(entql.HasEdgeWith("child_stations", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (rq *RoleQuery) addPredicate(pred func(s *sql.Selector)) {
	rq.predicates = append(rq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the RoleQuery builder.
func (rq *RoleQuery) Filter() *RoleFilter {
	return &RoleFilter{config: rq.config, predicateAdder: rq}
}

// addPredicate implements the predicateAdder interface.
func (m *RoleMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the RoleMutation builder.
func (m *RoleMutation) Filter() *RoleFilter {
	return &RoleFilter{config: m.config, predicateAdder: m}
}

// RoleFilter provides a generic filtering capability at runtime for RoleQuery.
type RoleFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *RoleFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[5].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *RoleFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(role.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *RoleFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(role.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *RoleFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(role.FieldUpdatedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *RoleFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(role.FieldName))
}

// WhereHasPermissions applies a predicate to check if query has an edge permissions.
func (f *RoleFilter) WhereHasPermissions() {
	f.Where(entql.HasEdge("permissions"))
}

// WhereHasPermissionsWith applies a predicate to check if query has an edge permissions with a given conditions (other predicates).
func (f *RoleFilter) WhereHasPermissionsWith(preds ...predicate.Permission) {
	f.Where(entql.HasEdgeWith("permissions", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasUsers applies a predicate to check if query has an edge users.
func (f *RoleFilter) WhereHasUsers() {
	f.Where(entql.HasEdge("users"))
}

// WhereHasUsersWith applies a predicate to check if query has an edge users with a given conditions (other predicates).
func (f *RoleFilter) WhereHasUsersWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("users", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{config: uq.config, predicateAdder: uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{config: m.config, predicateAdder: m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[6].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *UserFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(user.FieldID))
}

// WhereCreatedAt applies the entql time.Time predicate on the created_at field.
func (f *UserFilter) WhereCreatedAt(p entql.TimeP) {
	f.Where(p.Field(user.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql time.Time predicate on the updated_at field.
func (f *UserFilter) WhereUpdatedAt(p entql.TimeP) {
	f.Where(p.Field(user.FieldUpdatedAt))
}

// WhereName applies the entql string predicate on the name field.
func (f *UserFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(user.FieldName))
}

// WhereEmail applies the entql string predicate on the email field.
func (f *UserFilter) WhereEmail(p entql.StringP) {
	f.Where(p.Field(user.FieldEmail))
}

// WherePassword applies the entql string predicate on the password field.
func (f *UserFilter) WherePassword(p entql.StringP) {
	f.Where(p.Field(user.FieldPassword))
}

// WherePhone applies the entql string predicate on the phone field.
func (f *UserFilter) WherePhone(p entql.StringP) {
	f.Where(p.Field(user.FieldPhone))
}

// WhereActive applies the entql bool predicate on the active field.
func (f *UserFilter) WhereActive(p entql.BoolP) {
	f.Where(p.Field(user.FieldActive))
}

// WhereRoleID applies the entql [16]byte predicate on the role_id field.
func (f *UserFilter) WhereRoleID(p entql.ValueP) {
	f.Where(p.Field(user.FieldRoleID))
}

// WherePoliceStationID applies the entql [16]byte predicate on the police_station_id field.
func (f *UserFilter) WherePoliceStationID(p entql.ValueP) {
	f.Where(p.Field(user.FieldPoliceStationID))
}

// WhereHasRole applies a predicate to check if query has an edge role.
func (f *UserFilter) WhereHasRole() {
	f.Where(entql.HasEdge("role"))
}

// WhereHasRoleWith applies a predicate to check if query has an edge role with a given conditions (other predicates).
func (f *UserFilter) WhereHasRoleWith(preds ...predicate.Role) {
	f.Where(entql.HasEdgeWith("role", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasPoliceStation applies a predicate to check if query has an edge police_station.
func (f *UserFilter) WhereHasPoliceStation() {
	f.Where(entql.HasEdge("police_station"))
}

// WhereHasPoliceStationWith applies a predicate to check if query has an edge police_station with a given conditions (other predicates).
func (f *UserFilter) WhereHasPoliceStationWith(preds ...predicate.PoliceStation) {
	f.Where(entql.HasEdgeWith("police_station", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
