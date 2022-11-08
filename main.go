package main

import (
	"fmt"
	"github.com/go-ldap/ldap"
	"log"
)

const (
	BindUsername = "cn=admin,dc=example,dc=org"
	BindPassword = "adminpassword"
	FQDN         = "DC.example.org"
	BaseDN       = "dc=example,dc=org"
	Filter       = "(objectClass=*)"
)

func main() {
	// TLS Connection
	//l, err := ConnectTLS()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer l.Close()

	// Non-TLS Connection
	l, err := Connect()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Connection Successful\n")
	}

	defer l.Close()

	//// Anonymous Bind and Search
	//result, err := AnonymousBindAndSearch(l)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//result.Entries[0].Print()

	// Normal Bind and Search
	result, err := BindAndSearch(l)
	if err != nil {
		log.Fatal(err)
	}
	result.Entries[0].Print()

	//Create new Add request object to be added to LDAP server.
	//a := ldap.NewAddRequest("ou=groups,dc=example,dc=org", nil)
	//a.Attribute("cn", []string{"gotest"})
	//a.Attribute("objectClass", []string{"top"})
	//a.Attribute("description", []string{"this is a test to add an entry using golang"})
	//a.Attribute("sn", []string{"Google"})
	//
	//fmt.Println("Testing.")
	//add(a, l)
}

// Ldap Connection with TLS
func ConnectTLS() (*ldap.Conn, error) {
	// You can also use IP instead of FQDN
	l, err := ldap.DialURL(fmt.Sprintf("ldaps://%s:636", FQDN))
	if err != nil {
		return nil, err
	}

	return l, nil
}

// Ldap Connection without TLS
func Connect() (*ldap.Conn, error) {
	// You can also use IP instead of FQDN
	l, err := ldap.DialURL(fmt.Sprintf("ldap://%s:1389", "localhost"))
	if err != nil {
		return nil, err
	}

	return l, nil
}

// Anonymous Bind and Search
func AnonymousBindAndSearch(l *ldap.Conn) (*ldap.SearchResult, error) {
	l.UnauthenticatedBind("")

	anonReq := ldap.NewSearchRequest(
		"",
		ldap.ScopeBaseObject, // you can also use ldap.ScopeWholeSubtree
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		Filter,
		[]string{},
		nil,
	)
	result, err := l.Search(anonReq)
	if err != nil {
		return nil, fmt.Errorf("Anonymous Bind Search Error: %s", err)
	}

	if len(result.Entries) > 0 {
		result.Entries[0].Print()
		return result, nil
	} else {
		return nil, fmt.Errorf("Couldn't fetch anonymous bind search entries")
	}

}

// Normal Bind and Search
func BindAndSearch(l *ldap.Conn) (*ldap.SearchResult, error) {
	l.Bind(BindUsername, BindPassword)

	searchReq := ldap.NewSearchRequest(
		BaseDN,
		ldap.ScopeBaseObject, // you can also use ldap.ScopeWholeSubtree
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		Filter,
		[]string{},
		nil,
	)
	result, err := l.Search(searchReq)
	if err != nil {
		return nil, fmt.Errorf("Search Error: %s", err)
	}

	if len(result.Entries) > 0 {
		return result, nil
	} else {
		return nil, fmt.Errorf("Couldn't fetch search entries")
	}
}

func add(addRequest *ldap.AddRequest, l *ldap.Conn) {
	err := l.Add(addRequest)
	if err != nil {
		fmt.Println("Entry NOT done", err)
	} else {
		fmt.Println("Entry DONE", err)
	}
}
