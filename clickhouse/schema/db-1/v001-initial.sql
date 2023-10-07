CREATE TABLE IF NOT EXISTS ba_metrics.metrics
(
    created_at DateTime,
    ip String,
    user_agent String,
    referer String,
)
ENGINE = MergeTree
PRIMARY KEY (created_at)