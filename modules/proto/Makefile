# OS specific part
# -----------------
ifeq ($(OS),Windows_NT)
    CLEAR = cls
    LS = dir
    TOUCH =>> 
    RM = del /F /Q
    CPF = copy /y
    RMDIR = -RMDIR /S /Q
    MKDIR = -mkdir
    ERRIGNORE = 2>NUL || (exit 0)
    GO_PATH = $(subst \,/,${GOPATH})
    SEP=\\
else
    CLEAR = clear
    GO_PATH = ${GOPATH}
    LS = ls
    TOUCH = touch
    CPF = cp -f
    RM = rm -rf 
    RMDIR = rm -rf 
    MKDIR = mkdir -p
    ERRIGNORE = 2>/dev/null
    SEP=/
endif

DIRS = $(notdir $(patsubst %/,%,$(dir $(wildcard */.))))
PSEP = $(strip $(SEP))
.PHONY: clean build proto

clean:	
ifneq ("$(wildcard ./*.go)","")
	- $(RM) .$(PSEP)*.go  
endif
proto: clean
	- protoc \
	-I . \
	-I $(GOPATH)/src/ \
	-I $(GOPATH)/src/github.com/gogo/protobuf/protobuf \
	-I $(GOPATH)/src/github.com/bifrostcloud/protoc-gen-httpclient/proto \
	-I $(abspath $(lastword $(MAKEFILE_LIST))/../../) \
	-I $(abspath $(lastword $(MAKEFILE_LIST))/../../..) \
	-I $(abspath $(lastword $(MAKEFILE_LIST))/../../../..) \
	--gogo_out=\
	Mgithub.com/gogo/protobuf/gogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto,\
	Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
	:. \
	.$(PSEP)*.proto 

REPO_NAME = $(firstword $(subst /, ,$(subst ${GO_PATH}/src/,,$(abspath $(lastword $(MAKEFILE_LIST))))))
NAMESPACE_NAME = $(firstword $(subst /, ,$(subst ${GO_PATH}/src/${REPO_NAME},,$(abspath $(lastword $(MAKEFILE_LIST))/..))))
PROJECT_NAME = $(firstword $(subst /, ,$(subst ${GO_PATH}/src/${REPO_NAME}/${NAMESPACE_NAME},,$(abspath $(lastword $(MAKEFILE_LIST))/..))))
PARENT_FOLDER_NAME = $(firstword $(subst /, ,$(subst ${GO_PATH}/src/${REPO_NAME}/${NAMESPACE_NAME}/${PROJECT_NAME},,$(abspath $(lastword $(MAKEFILE_LIST))/..))))
FEATURE_NAME = $(firstword $(subst /, ,$(subst ${GO_PATH}/src/${REPO_NAME}/${NAMESPACE_NAME}/${PROJECT_NAME}/${PARENT_FOLDER_NAME},,$(abspath $(lastword $(MAKEFILE_LIST))/..))))


# Definitions for nullstring and space
# -------------------------------------
nullstring :=
space := $(nullstring) #End

