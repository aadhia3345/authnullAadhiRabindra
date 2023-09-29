package main

import (
	sa "github.com/aadhia3345/authnullAadhiRabindra.git/ServiceAccount"
	wi "github.com/aadhia3345/authnullAadhiRabindra.git/WorkloadIdentity"
)

func main() {
	sa.AddEnpointGroup()
	sa.AddServiceAccount()
	wi.AddIdentity()
}
