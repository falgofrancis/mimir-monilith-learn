# Goal
Try out mimir [guide](https://grafana.com/docs/mimir/latest/get-started/), not for production use.

# Changes from Guide
1. Deployed Prometheus in the same network and used Remote write endpoint for Prometheus - `http://mimir:9009/api/v1/push`. Kept seeing `Failed to send batch, retrying` + `lame referral` error when endpoint is not reachable.
1. Overridden Grafana password to be `kidding`.
1. Grafana [datasource](https://github.com/grafana/grafana/tree/main/conf/provisioning/datasources) in yaml.
1. Sample Application is in the folder app and it keeps generating one metrics and you can generate a metric with 10 labels
1. No LB

# Learnings
1. Grafana Mimir is an open source, horizontally scalable, highly available, multi-tenant, long-term storage for Prometheus.
1. components: distributor, compactor, ingester, querier, query-frontend and store-gateway
1. Component split helps to horizontally scale and allows maintenance with zero downtime
1. Support multiple object store implementations. You may want to keep data in pvc
1. check possibilities of using external Memcached
1. [Mimir doesn’t do deduplication by default, like when Prometheus with 2 replicas](https://tech.loveholidays.com/grafana-mimir-our-journey-towards-infinite-wisdom-with-5m-active-time-series-7a262ba53a3f). 
1. [Cortex seems to miscalculate quorum when one ingester is unhealthy](https://github.com/cortexproject/cortex/issues/4654).