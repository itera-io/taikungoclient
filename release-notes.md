### New Endpoints: None
-----------------------

### Deleted Endpoints: 1
------------------------
POST /api/v1/azure/file-share-list  

### Modified Endpoints: 6
-------------------------
POST /api/v1/azure/create
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: importedFileShareId
          - Deleted property: ipMode
          - Deleted property: sharedFileSystemEnabled
          - Deleted property: subnets
          - Modified property: vnetMode
            - Deleted enum values: [Import]

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
                  - Deleted required property: importedFileShareId
                  - Deleted required property: importedFileShareName
                  - Deleted required property: sharedFileSystemEnabled
                - Properties changed
                  - Deleted property: importedFileShareId
                  - Deleted property: importedFileShareName
                  - Deleted property: sharedFileSystemEnabled

POST /api/v1/azure/vnet-list
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: ipMode
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Items changed
            - Properties changed
              - Deleted property: ipMode

POST /api/v1/azure/vnet-subnet-list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Items changed
            - Properties changed
              - Deleted property: ipMode
              - Deleted property: isPublic

GET /api/v1/dnscredentials
- Deleted query param: IncludeLocked
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Items changed
            - Required changed
              - Deleted required property: isLocked
            - Properties changed
              - Deleted property: isLocked

GET /api/v1/servers/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: project
              - Required changed
                - Deleted required property: dnsCredentialId
                - Deleted required property: dnsCredentialName
              - Properties changed
                - Deleted property: dnsCredentialId
                - Deleted property: dnsCredentialName


