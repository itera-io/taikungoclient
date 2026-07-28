### New Endpoints: None
-----------------------

### Deleted Endpoints: None
---------------------------

### Modified Endpoints: 3
-------------------------
POST /api/v1/dns-cert/enable
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: enableTrustManager
          - Deleted property: stageIntentOnly

POST /api/v1/projects
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: certificateProfileId
          - Deleted property: isCertTrustEnabled
          - Deleted property: isDnsEnabled

GET /api/v1/servers/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: project
              - Properties changed
                - Deleted property: ctProfileId
                - Deleted property: ctProfileName


