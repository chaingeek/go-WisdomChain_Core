// Copyright 2018 The go-WisdomChain_Core Authors
// This file is part of go-WisdomChain_Core.
//
// go-WisdomChain_Core is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-WisdomChain_Core is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-WisdomChain_Core. If not, see <http://www.gnu.org/licenses/>.

// wisdom is the official command-line client for Wisdom Chain.

package main 

import (

	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)


const (
	//client identifier over the main network
	clientIdentifier = "wisdom"
)

