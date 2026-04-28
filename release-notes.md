### New Endpoints: None
-----------------------

### Deleted Endpoints: None
---------------------------

### Modified Endpoints: 3
-------------------------
POST /api/v1/global-configurations/airgap/create
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: password
          - Deleted property: username

GET /api/v1/global-configurations/airgap/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Deleted property: password
                  - Deleted property: username

PUT /api/v1/global-configurations/airgap/update
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: password
          - Deleted property: username


