package sa

type serviceAccount struct {
	id        int
	tenant_id int
	username  string
	usertype  string
}

type endpointGroup struct {
	id   int
	name string
}

func AddServiceAccount(id int, tenant_id int, username string, usertype string) {
	/* get id, tenant_id, username, type
	then add the object with these attributes */
	s := serviceAccount{}
	s.id = id
	s.tenant_id = tenant_id
	s.username = username
	s.usertype = usertype
}

func AddEnpointGroup(id int, name string) {
	/*get id and name of and then add it */
	eg := endpointGroup{}
	eg.id = id
	eg.name = name
}

func GetServiceAccount() {
	/*from tenantid and identityid get the service account
	corresponding to it */
}

func GetEndpointGroup() {
	/*from groupid and identityid get the endpoint group
	corresponding to it */
}

func GetServiceAccounts() {
	//return all the service accounts stored
}

func GetEndpointGroups() {
	//return all endpoint groups that are stored
}

func AssignSAtoWallet() {
	//send the service account to the wallet
}

func AssociateServiceAccountRules() {
	//check the if statements and then decide what permissions should be granted

}
