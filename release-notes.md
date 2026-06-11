### New Endpoints: None
-----------------------

### Deleted Endpoints: None
---------------------------

### Modified Endpoints: 2
-------------------------
POST /api/v1/aws/create
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Required changed
          - New required property: awsAccessKeyId
          - New required property: awsSecretAccessKey
        - Properties changed
          - Deleted property: externalId
          - Deleted property: roleArn

POST /api/v1/aws/update
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: externalId
          - Deleted property: roleArn


