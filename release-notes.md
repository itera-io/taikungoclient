### New Endpoints: None
-----------------------

### Deleted Endpoints: 2
------------------------
GET /api/v1/cloudcredentials/flavors/gpu-filter-metadata/{cloudId}  
POST /api/v1/cronjob/sync-dns-cert-configs  

### Modified Endpoints: 21
--------------------------
GET /api/v1/accounts/{accountId}/organizations/{id}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: projects
              - Items changed
                - Properties changed
                  - Modified property: status
                    - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/accounts/{accountId}/projects/offset-based/dropdown
- Modified query param: Status
  - Schema changed
    - Deleted enum values: [SyncDnsCert DisableDnsCert]
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: status
                    - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/accounts/{accountId}/projects/{id}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: status
              - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/catalog
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: boundProjects
                    - Items changed
                      - Properties changed
                        - Modified property: status
                          - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/catalog/{id}/details
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: boundProjects
              - Items changed
                - Properties changed
                  - Modified property: status
                    - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/dns-cert/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - New property: certManagerSynced
            - New property: clusterIssuerSynced
            - New property: externalDnsSynced
            - Deleted property: appliedConfigVersion
            - Deleted property: desiredConfigVersion
            - Deleted property: synced

GET /api/v1/generic-kubernetes/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: mainProject
                    - Properties changed
                      - Modified property: status
                        - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/imported-cluster/as-cloud-credential/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Properties changed
                - Modified property: status
                  - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/imported-cluster/as-fully-managed/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Properties changed
                - Modified property: status
                  - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/imported-cluster/as-read-only/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Properties changed
                - Modified property: status
                  - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/imported-cluster/{id}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: status
                    - Deleted enum values: [SyncDnsCert DisableDnsCert]
            - Modified property: project
              - Properties changed
                - Modified property: status
                  - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/notifications/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: category
                    - Deleted enum values: [StartedSyncDnsCert SyncedDnsCert StartedDisableDnsCert DisabledDnsCert]

GET /api/v1/projects
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: status
                    - Deleted enum values: [SyncDnsCert DisableDnsCert]

PUT /api/v1/projects/edit/status
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Modified property: status
            - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/projects/imported/details/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: status
              - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/projects/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Items changed
            - Properties changed
              - Modified property: status
                - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/projects/visibility/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Required changed
            - Deleted required property: disableDnsCert
            - Deleted required property: enableDnsCert
          - Properties changed
            - Deleted property: disableDnsCert
            - Deleted property: enableDnsCert

GET /api/v1/servers/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: project
              - Properties changed
                - Modified property: status
                  - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/sse/recent-events/{organizationId}
- Modified query param: Category
  - Schema changed
    - Deleted enum values: [StartedSyncDnsCert SyncedDnsCert StartedDisableDnsCert DisabledDnsCert]

GET /api/v1/standalone/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: project
              - Properties changed
                - Modified property: status
                  - Deleted enum values: [SyncDnsCert DisableDnsCert]

GET /api/v1/virtual-cluster/{parentProjectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Modified property: status
                    - Deleted enum values: [SyncDnsCert DisableDnsCert]
            - Modified property: project
              - Properties changed
                - Modified property: status
                  - Deleted enum values: [SyncDnsCert DisableDnsCert]


