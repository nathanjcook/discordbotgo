package helpcommand

func Help() (string, string, string, string, string, string, string, string) {
	var add_title = "Add Command"
	var add = "The Add Command Is Designed To Allow Admins To Submit The Details Required To Connect To A Particular Microservice\nTo Sucessfully Add A Microservice To The Database Then The Admin Must Use:\n\n!gobot add <microservice_name> <microservive_url> <microservice_timeout>\n\nThe command itself will be formatted as:\n\n\n!gobot calculator localhost/4356 66"
	var delete_title = "Delete Command"
	var delete = "The Delete Command Is Designed To Allow Admins To Delete instances of a microservice from the database by microservice name\n\nTo Sucessfully Delete A Microservice From The Database Then The Admin Must Use:\n\n!gobot delete <microservice_name>\nThe command itself will be formatted as:\n\n\n!gobot delete calculator"
	var info_title = "Info Command"
	var info = "The Info Command Is Design To List All Available Microservices\n\nTo Sucessfully Run The Command Then The User Must Use:\n\n\n!gobot info"
	var microservice_title = "Microservices"
	var microservice = "To Run An Instance Of The Microserver Then the User Must Use\n\n!gobot <microservice_name> <microservice_endpoint> <microservice_body>\n\n\n It Is Recommended To Run The Enforced Microservice Help Endpoint First To Understand What Endpoints Are Available And What Format The Body/Variables Have To Be\n\n You can do this by using: \n\n\n!gobot <microservice_name> help"

	return add_title, add, delete_title, delete, info_title, info, microservice_title, microservice
}
