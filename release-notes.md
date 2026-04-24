### New Endpoints: None
-----------------------

### Deleted Endpoints: None
---------------------------

### Modified Endpoints: 4
-------------------------
POST /api/v1/aws/subnet-list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Items changed
            - Properties changed
              - Deleted property: availableIpCount
              - Deleted property: isDefault
              - Deleted property: name
              - Deleted property: ownerId

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
                      - Deleted property: availableIpCount
                      - Deleted property: cidr
                      - Deleted property: isDefault
                      - Deleted property: name
                      - Deleted property: nodeCount
                      - Deleted property: ownerId
                      - Deleted property: zone

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
                      - Deleted property: availableIpCount
                      - Deleted property: cidr
                      - Deleted property: isDefault
                      - Deleted property: name
                      - Deleted property: nodeCount
                      - Deleted property: ownerId
                      - Deleted property: zone

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
                      - Deleted property: availableIpCount
                      - Deleted property: cidr
                      - Deleted property: isDefault
                      - Deleted property: name
                      - Deleted property: nodeCount
                      - Deleted property: ownerId
                      - Deleted property: zone


