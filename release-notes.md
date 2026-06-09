### New Endpoints: None
-----------------------

### Deleted Endpoints: 3
------------------------
POST /api/v1/dnscredentials/attach-project  
POST /api/v1/dnscredentials/detach-project  
GET /api/v1/internal/projects-for-billing  

### Modified Endpoints: 7
-------------------------
GET /api/v1/cloudcredentials
- Modified query param: IsAdmin
  - Required changed from false to true

POST /api/v1/dns-cert/enable
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Modified property: dnsCredentialId
            - Nullable changed from true to false

POST /api/v1/dns-cert/validate
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: azureClientId
          - Deleted property: azureClientSecret
          - Deleted property: azureResourceGroup
          - Deleted property: azureSubscriptionId
          - Deleted property: azureTenantId
          - Deleted property: gcpProject
          - Deleted property: gcpServiceAccountJson

GET /api/v1/dns-cert/{projectId}
- Responses changed
  - Modified response: 200
    - Content changed
      - Modified media type: application/json
        - Schema changed
          - Properties changed
            - Deleted property: dnsCredentialId
            - Deleted property: dnsCredentialName

POST /api/v1/dnscredentials
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: azureClientId
          - Deleted property: azureClientSecret
          - Deleted property: azureResourceGroup
          - Deleted property: azureSubscriptionId
          - Deleted property: azureTenantId
          - Deleted property: gcpProject
          - Deleted property: gcpServiceAccountJson

POST /api/v1/dnscredentials/validate
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: azureClientId
          - Deleted property: azureClientSecret
          - Deleted property: azureResourceGroup
          - Deleted property: azureSubscriptionId
          - Deleted property: azureTenantId
          - Deleted property: gcpProject
          - Deleted property: gcpServiceAccountJson

POST /api/v1/projects
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: dnsCredentialId


