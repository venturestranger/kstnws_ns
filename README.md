## Posts Pool Server Documentation

###### REST Application Conventions and TLDR:

```
Abbreviations:
1. NS - Notification Service

HTTP Request Consists Of:
1. Destination (host:port/resource)
2. Method (GET, POST, PUT, DELETE)
3. Headers 
4. Request Body (JSON)

HTTP Response Consists of:
1. Status Code
2. Response Body (JSON)

Methods:
1. GET - read (returns status code and response body)
2. POST - create (requires request body, returns status code)
3. PUT - update (returns status code)
4. DELETE - delete (returns status code)

Status Codes:
1. 200 OK (successful request)
2. 400 Bad Request (wrong query format or non-existent fields)
3. 401 Unathorized (no access due to invalid token)
4. 405 Method Not Allowed (no method specified for a resource)
5. 500 Internal Server Error (server troubles to connect to its databases)
```

###### NS Routing:

```
- /notification/auth [GET] | returns a token if valid key 
* /notification [GET] | returns a list of notifications
* /notification [POST] | creates a notification
* /notification [PUT] | updates a notification
* /notification [DELETE] | removes a notification 
```

*- endpoints requiring authorization

###### Getting Access:

```
1. Access "/notification/auth" with "key" for fetching a token
2. Set a header "Authorization" to "Bearer @token"
```

###### Resources:

```
/notification [GET]

Description:
1. It returns a list of notifications in the chosen order (specified by order_way_)

* fields: 
1. id
2. id_user
3. order_way_
4. offset_ - specifies how many posts to skip
5. limit_ - specifies how many posts to return 
```

```
/notification [POST]

Description:
1. It creates a notification 
2. It requires request body (JSON) with the data specified

* fields (JSON in request body):
{
    "id_author": 1,
    "title": "Sport title",
    "lead": "",
    "picture_url": "",
    "content": "Hello here is some text",
    "date_publication": "1971-01-26",
    "date_edit": "1971-01-26",
    "category": "спорт",
    "hashtags": "kostanay sportinkazakhstan"
    "comment": ""
}
```

```
/validate [PUT]

Description:
1. It changes the status of a post. If verification is successful, it pushes a post to API and removes the local copy from PPS

* fields: 
1. id 
2. comment - this is the text that is displayed if a post has not been verified. It implements a notification systems and tells what is currently wrong with the post
3. pass - if it is passed with 'true', then the post is successfully verified, pushed to API, and removed from PPS
```

```
/validate [DELETE]

Description:
1. It removes a post copy from PPS by id

* fields: 
1. id 
```
