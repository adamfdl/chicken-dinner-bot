# /messages

Retrieve messages

## GET
+ Response 200 (text/plain)

        Hello World!
        Hello Other World


# /messages/{id}
Represents a particular message by it's *id*.

## Retrieve message [GET]
+ Response 200 (text/plain)

        Hello World!

## DELETE MESSAGE DELETE
+ Response 204

# /messages/{id}.json

## Retrieve message in JSON [GET]
+ Response 200 (application/json)

        {"message": "Hello World!"}