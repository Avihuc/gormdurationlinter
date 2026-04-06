package example

import "time"

// Good: time.Duration with bigint and serializer:nanoduration
type ModelCorrect struct {
	ID       string
	Duration time.Duration `gorm:"type:bigint;serializer:nanoduration"`
}

// Good: *time.Duration with bigint and serializer:nanoduration
type ModelCorrectPtr struct {
	ID       string
	Duration *time.Duration `gorm:"type:bigint;serializer:nanoduration"`
}

// Good: time.Duration with no gorm tag
type ModelNoTag struct {
	ID       string
	Duration time.Duration
}

// Good: time.Duration with non-bigint gorm type
type ModelInterval struct {
	ID       string
	Duration time.Duration `gorm:"type:interval"`
}

// Good: int64 with bigint (not a Duration)
type ModelInt64 struct {
	ID    string
	Value int64 `gorm:"type:bigint"`
}

// Bad: time.Duration with bigint but missing serializer
type ModelMissing struct {
	ID       string
	Duration time.Duration `gorm:"type:bigint"` // want `time\.Duration field with gorm type:bigint must include serializer:nanoduration`
}

// Bad: *time.Duration with bigint but missing serializer
type ModelMissingPtr struct {
	ID       string
	Duration *time.Duration `gorm:"type:bigint"` // want `time\.Duration field with gorm type:bigint must include serializer:nanoduration`
}

// Bad: *time.Duration with bigint and other tags but missing serializer
type ModelMissingWithOtherTags struct {
	ID       string
	Duration *time.Duration `gorm:"type:bigint;not null"` // want `time\.Duration field with gorm type:bigint must include serializer:nanoduration`
}
