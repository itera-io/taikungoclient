### New Endpoints: None
-----------------------

### Deleted Endpoints: None
---------------------------

### Modified Endpoints: 6
-------------------------
GET /api/v1/accounts/{accountId}/user/offset-based/dropdown
- Modified query param: FilterRole
  - Schema changed
    - Type changed from 'string' to 'integer'
    - Format changed from '' to 'int32'
    - Deleted enum values: [None Admin AccountAdmin AccountOwner Member Manager]

POST /api/v1/aws/subnet-list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Items changed
            - Properties changed
              - Deleted property: vpcId

GET /api/v1/common/enumvalues
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Required changed
            - Deleted required property: filteredUserRoles
          - Properties changed
            - Deleted property: filteredUserRoles

GET /api/v1/servers/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: project
              - Properties changed
                - Modified property: cloudSubnets
                  - Items changed
                    - Properties changed
                      - Deleted property: vpcId

GET /api/v1/standalone/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: project
              - Properties changed
                - Modified property: cloudSubnets
                  - Items changed
                    - Properties changed
                      - Deleted property: vpcId

GET /api/v1/virtual-cluster/{parentProjectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: project
              - Properties changed
                - Modified property: cloudSubnets
                  - Items changed
                    - Properties changed
                      - Deleted property: vpcId


