package commands

var AddTitle = "\n\n***Add Command***\n"
var AddMsg = "Add Microservice To Database!\n\n!gobot add <microservice_name> <microservive_url> <microservive_timeout>"

var DeleteTitle = "\n\n***Delete Command***\n"
var DeleteMsg = "Delete Microservice From Database!\n\n!gobot delete <microservice_name>"

var InfoTitle = "\n\n***Info Command***\n"
var InfoMsg = "View All Microservices\n\n!gobot info"

var MicroserviceTitle = "\n\n***Microservices***\n"
var MicroserviceMsg = "Run Bot From Microservice\n\n!gobot <microservice_name> <microservice_endpoint> <microservice_body>\n\nIt Is Recommended To Run The Enforced Microservice Help Endpoint First To Understand What Endpoints Are Available And What Format The Body/Variables Have To Be\n\nYou can do this by using: \n\n!gobot <microservice_name> help"
