### New Endpoints: None
-----------------------

### Deleted Endpoints: 3
------------------------
POST /api/v1/googlecloud/filestore-list  
POST /api/v1/googlecloud/subnetwork-list  
POST /api/v1/googlecloud/vpc-list  

### Modified Endpoints: 2
-------------------------
POST /api/v1/googlecloud/create
- Request body changed
  - Content changed
    - Modified media type: multipart/form-data
      - Schema changed
        - Properties changed
          - Deleted property: importedFilestoreId
          - Deleted property: importedGcpProjectId
          - Deleted property: importedSubnetworkName
          - Deleted property: importedVpcName
          - Deleted property: vpcMode

GET /api/v1/googlecloud/list
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Deleted property: importedFilestoreId
                  - Deleted property: importedFilestoreName
                  - Deleted property: importedGcpProjectId
                  - Deleted property: importedSubnetworkName
                  - Deleted property: importedVpcName
                  - Deleted property: ipMode
                  - Deleted property: sharedFileSystemEnabled
                  - Deleted property: vpcMode


