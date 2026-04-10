### New Endpoints: None
-----------------------

### Deleted Endpoints: 1
------------------------
POST /api/v1/users/manage-role  

### Modified Endpoints: 4
-------------------------
DELETE /api/v1/organizations/{id}
- Summary changed from 'Delete the specified organization. Only available for global roles.' to 'Delete the specified organization. Only available for admins.'
- Description changed from 'This endpoint allows deleting an organization by its ID. Requires global role permissions.' to ''

POST /api/v1/projectapp/install
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - New property: labels

GET /api/v1/robot/scope-list
- Deleted query param: GlobalRobot

POST /api/v1/servers/create
- Request body changed
  - Content changed
    - Modified media type: application/json
      - Schema changed
        - Properties changed
          - Deleted property: subnetId


