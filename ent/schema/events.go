package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
			Comment("Bounding box coordinates of the plate"),
		field.Int("plate_channel").
			Optional().
			Comment("Channel of the plate"),
		field.Int("plate_confidence").
			Optional().
			Comment("Confidence score of the plate detection"),
		field.Bool("plate_is_exist").
			Optional().
			Comment("Indicates whether the plate exists"),
		field.String("plate_color").
			Optional().
			Comment("Color of the plate"),
		field.String("plate_number").
			Optional().
			Comment("Number on the plate"),
		field.String("plate_type").
			Optional().
			Comment("Type of the plate"),
		field.String("plate_region").
			Optional().
			Comment("Region of the plate"),
		field.Int("plate_upload_num").
			Optional().
			Comment("Upload number of the plate"),

		// SnapInfo fields
		field.String("snap_accurate_time").
			Optional().
			Comment("Accurate timestamp of the snap"),
		field.Bool("snap_allow_user").
			Optional().
			Comment("Indicates if user interaction is allowed"),
		field.String("snap_allow_user_end_time").
			Optional().
			Comment("End time for user interaction allowance"),
		field.Int("snap_dst_tune").
			Optional().
			Comment("DST adjustment"),
		field.String("snap_defence_code").
			Optional().
			Comment("Defence code"),
		field.String("snap_device_id").
			Optional().
			Comment("Device ID"),
		field.String("snap_direction").
			Optional().
			Comment("Direction of the vehicle"),
		field.Int("snap_in_car_people_num").
			Optional().
			Comment("Number of people in the car"),
		field.Int("snap_lan_no").
			Optional().
			Comment("LAN number"),
		field.Bool("snap_open_strobe").
			Optional().
			Comment("Indicates if strobe is open"),
		field.String("snap_snap_time").
			Optional().
			Comment("Snap timestamp"),
		field.Int("snap_time_zone").
			Optional().
			Comment("Time zone of the snap"),

		// Vehicle fields
		field.Int("vehicle_speed").
			Optional().
			Comment("Speed of the vehicle"),
		field.JSON("vehicle_bounding_box", []int{}).
			Optional().
			Comment("Bounding box coordinates of the vehicle"),
		field.String("vehicle_color").
			Optional().
			Comment("Color of the vehicle"),
		field.String("vehicle_series").
			Optional().
			Comment("Vehicle series"),
		field.String("vehicle_type").
			Optional().
			Comment("Type of the vehicle"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return nil
}

func (Event) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("plate_number"),
		index.Fields("plate_color"),
		index.Fields("plate_type"),
		index.Fields("plate_region"),
		index.Fields("plate_confidence"),

		index.Fields("snap_accurate_time"),
		index.Fields("snap_device_id"),
		index.Fields("snap_direction"),
		index.Fields("snap_time_zone"),

		index.Fields("vehicle_type"),
		index.Fields("vehicle_color"),
	}
}
