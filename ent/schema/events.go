package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"go-ent-project/utils/base"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

func (Event) Mixin() []ent.Mixin {
	return []ent.Mixin{
		base.BMixin{},
	}
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		// Plate fields
		field.JSON("plate_bounding_box", []int{}).
			Optional().
			Comment("Bounding box coordinates of the plate").
			Annotations(),
		field.Int("plate_channel").
			Optional().
			Comment("Channel of the plate").
			Annotations(),
		field.Bool("plate_is_exist").
			Optional().
			Comment("Indicates whether the plate exists").
			Annotations(),
		field.String("plate_color").
			Optional().
			Comment("Color of the plate").
			Annotations(),
		field.String("plate_number").
			Optional().
			Comment("Number on the plate").
			Annotations(),
		field.String("plate_type").
			Optional().
			Comment("Type of the plate").
			Annotations(),
		field.String("plate_region").
			Optional().
			Comment("Region of the plate").
			Annotations(),
		field.Int("plate_upload_num").
			Optional().
			Comment("Upload number of the plate").
			Annotations(),

		// SnapInfo fields
		field.Bool("snap_allow_user").
			Optional().
			Comment("Indicates if user interaction is allowed").
			Annotations(),
		field.String("snap_allow_user_end_time").
			Optional().
			Comment("End time for user interaction allowance").
			Annotations(),
		field.String("snap_defence_code").
			Optional().
			Comment("Defence code").
			Annotations(),
		field.String("snap_device_id").
			Optional().
			Comment("Device ID").
			Annotations(),
		field.Int("snap_in_car_people_num").
			Optional().
			Comment("Number of people in the car").
			Annotations(),
		field.Int("snap_lan_no").
			Optional().
			Comment("LAN number").
			Annotations(),
		field.Bool("snap_open_strobe").
			Optional().
			Comment("Indicates if strobe is open").
			Annotations(),

		// Vehicle fields
		field.JSON("vehicle_bounding_box", []int{}).
			Optional().
			Comment("Bounding box coordinates of the vehicle").
			Annotations(),
		field.String("vehicle_series").
			Optional().
			Comment("Vehicle series").
			Annotations(),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		// Define any relationships (edges) here if needed in the future.
	}
}
