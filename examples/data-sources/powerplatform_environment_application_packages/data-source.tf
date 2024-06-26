terraform {
  required_providers {
    powerplatform = {
      source = "microsoft/power-platform"
    }
  }
}

provider "powerplatform" {
  use_cli = true
}

data "powerplatform_environments" "all_environments" {}

data "powerplatform_environment_application_packages" "all_applications" {
  environment_id = data.powerplatform_environments.all_environments.environments[0].id
}

data "powerplatform_environment_application_packages" "all_applications_from_publisher" {
  environment_id = data.powerplatform_environments.all_environments.environments[0].id
  publisher_name = "Power Platform Host Service"
}

data "powerplatform_environment_application_packages" "specific_application" {
  environment_id = data.powerplatform_environments.all_environments.environments[0].id
  publisher_name = "Microsoft Dynamics 365"
  name           = "Virtual connectors in Dataverse"
}
