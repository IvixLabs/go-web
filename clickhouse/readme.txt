CREATE DATABASE IF NOT EXISTS goweb

CREATE TABLE IF NOT EXISTS entity
(
    EntityId              UUID,
    PropertyName          String,
    PropertyType          String,
    PropertyStringValue   String,
    PropertyIntValue      UInt64,
    PropertyFloatValue    Float64,
    PropertyBoolValue     Boolean,
    PropertyDateTimeValue DateTime,
    Date                  DateTime,
    IsNull                Boolean
)
ENGINE = ReplacingMergeTree(Date)
ORDER BY (EntityId, PropertyName);
