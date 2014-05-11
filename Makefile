#
# Copyright 2014 Hong Miao. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


GO=$(shell which go)
#CLOC=$(shell which cloc)
GOFMT=$(GO) fmt

BIN=./bin

DIRS=$(shell find ./ -type d)
SRCS=$(foreach dir_var,$(DIRS),$(wildcard $(dir_var)/*.go))
SRC_DIRS=$(sort $(dir $(SRCS)))

# OBJS=$(patsubst %.cpp,$(BUILD_DIR)/%.o,$(notdir $(SRCS)))
# DEPS=$(patsubst %.o,%.d,$(OBJS))
# INCLUDES=$(foreach dir_var,$(DIRS), -I $(dir_var))

TARGETS=$(patsubst %.go,$(BIN)/%,$(wildcard *.go))

build: $(TARGETS)
	@echo $(SRCS)
	@echo $(TARGETS)

$(TARGETS): $(BIN)/%: %.go
	@echo "building $<..."
	@$(GO) build -o $@ $<

clean:
	@$(RM) -f $(TARGETS)
	
fmt:
	$(foreach var,$(SRCS),$(GOFMT) $(var)) 

# cloc:
	# @$(CLOC) . --exclude-dir=webclient/assets