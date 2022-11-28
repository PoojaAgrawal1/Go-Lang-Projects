--------------- GoMock libraries required --------

go get github.com/golang/mock/gomock 
go get github.com/golang/mock/mockgen


------- to create mockRunner.go file use below command ------------

mockgen -destination=mocks/mockRunner.go -package=mocks sample/test/IUser IUserInterface

where package name = sample/test
Interface = IUser/IUserInterface