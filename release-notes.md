### New Endpoints: None
-----------------------

### Deleted Endpoints: None
---------------------------

### Modified Endpoints: 2
-------------------------
POST /api/v1/azure/create
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - New property: existingSubnetName

GET /api/v1/azure/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Required changed
                  - New required property: existingSubnetName
                - Properties changed
                  - New property: existingSubnetName


