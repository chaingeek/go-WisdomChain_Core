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

package config 

import (
	"fmt"
)

const (
	VersionMajor = 0          // Major version
	VersionMinor = 0          // Minor version
	VersionPatch = 1          // Patch version
	VersionStatus  = "unstable"    // status version "unstable" or "stable"
 )

 var Version = func() string {
	 return fmt.Sprintf("%d.%d.%d",VersionMajor,VersionMinor,VersionPatch)
 }()

 var VersionWithStatus=func() string {
	 vn:=Version
	 if VersionStatus!="" {
		 vn += "-" + VersionStatus
	 }
	 return vn
 } ()