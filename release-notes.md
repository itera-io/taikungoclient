### New Endpoints: None
-----------------------

### Deleted Endpoints: 28
-------------------------
POST /api/v1/azure/vnet-list  
POST /api/v1/azure/vnet-resource-group-list  
POST /api/v1/azure/vnet-subnet-list  
GET /api/v1/custom-cas  
POST /api/v1/custom-cas  
PUT /api/v1/custom-cas  
GET /api/v1/custom-cas/list  
POST /api/v1/custom-cas/lockmanager  
POST /api/v1/custom-cas/makedefault  
POST /api/v1/custom-cas/validate  
DELETE /api/v1/custom-cas/{id}  
POST /api/v1/dns-cert/disable  
POST /api/v1/dns-cert/enable  
POST /api/v1/dns-cert/sync  
POST /api/v1/dns-cert/validate  
GET /api/v1/dns-cert/{projectId}  
GET /api/v1/dnscredentials  
POST /api/v1/dnscredentials  
PUT /api/v1/dnscredentials  
GET /api/v1/dnscredentials/list  
POST /api/v1/dnscredentials/lockmanager  
POST /api/v1/dnscredentials/makedefault  
POST /api/v1/dnscredentials/validate  
DELETE /api/v1/dnscredentials/{id}  
POST /api/v1/global-configurations/dns-cert/create  
GET /api/v1/global-configurations/dns-cert/list  
POST /api/v1/global-configurations/dns-cert/status  
DELETE /api/v1/global-configurations/dns-cert/{id}  

### Modified Endpoints: 9
-------------------------
POST /api/v1/aws/vpc-list
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: ipMode

POST /api/v1/azure/create
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: existingSubnetName
          - Deleted property: existingVirtualNetworkName
          - Deleted property: existingVirtualNetworkResourceGroupName
          - Deleted property: vnetMode

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
                  - Deleted required property: existingSubnetName
                  - Deleted required property: existingVirtualNetworkName
                  - Deleted required property: existingVirtualNetworkResourceGroupName
                  - Deleted required property: vnetDeploymentMode
                - Properties changed
                  - Deleted property: existingSubnetName
                  - Deleted property: existingVirtualNetworkName
                  - Deleted property: existingVirtualNetworkResourceGroupName
                  - Deleted property: vnetDeploymentMode

POST /api/v1/azure/subscriptions
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - New property: clientId
          - New property: clientSecret
          - New property: tenantId
          - Deleted property: azureClientId
          - Deleted property: azureClientSecret
          - Deleted property: azureTenantId

GET /api/v1/cloudcredentials/flavors/{cloudId}
- Deleted query param: GpuCount
- Deleted query param: GpuModel
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Deleted property: gpuDetails

GET /api/v1/flavors/aws/{cloudId}
- Deleted query param: GpuCount
- Deleted query param: GpuModel
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Deleted property: gpuDetails

GET /api/v1/flavors/azure/{cloudId}
- Deleted query param: GpuCount
- Deleted query param: GpuModel
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Deleted property: gpuDetails

GET /api/v1/flavors/google/{cloudId}
- Deleted query param: GpuCount
- Deleted query param: GpuModel
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Deleted property: gpuDetails

GET /api/v1/flavors/zadara/{cloudId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Modified property: data
              - Items changed
                - Properties changed
                  - Deleted property: gpuDetails


