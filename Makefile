all:

golang:
	rm -f ns_*.go
	rm -f service.go
	../../cinp/utils/cinp-codegen -l go -d . -s contractor http://127.0.0.1:8888/api/v1/
	sed s/NewContractor/NewContractorInt/ -i service.go
	echo -n "\n\
// OverrideCINPClient is for use during testing to replace the cinp clinet with a mocked out client\n\
func (s *Contractor) OverrideCINPClient(cinp cinp.CInPClient) {\n\
	s.cinp = cinp\n\
}" >> service.go
	go fmt ./...
	go vet ./...


PHONY: all golang
