// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2016-2018 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package builtin

const cifsMountControlSummary = `allows to mount and unmount CIFS shares`

const cifsMountControlBaseDeclarationSlots = `
  cifs-mount-control:
    allow-installation:
      slot-snap-type:
        - core
    deny-auto-connection: true
`

const cifsMountControlConnectedPlugSecComp = `
# Description: Allow mount and umount access.

mount
umount
`

const cifsMountControlConnectedPlugAppArmor = `
# Description: Allow to mount and unmount CIFS filesystems.

# Required for mounts
capability sys_admin,

# Allow mounts to our snap-specific writable directories
mount fstype=cifs ** -> /home/*/snap/@{SNAP_NAME}/@{SNAP_REVISION}/{,**/},
mount fstype=cifs ** -> /var/snap/@{SNAP_NAME}/@{SNAP_REVISION}/{,**/},
mount fstype=cifs ** -> /home/*/snap/@{SNAP_NAME}/common/{,**/},
mount fstype=cifs ** -> /var/snap/@{SNAP_NAME}/common/{,**/},

umount fstype=cifs /home/*/snap/@{SNAP_NAME}/@{SNAP_REVISION}/{,**/},
umount fstype=cifs /var/snap/@{SNAP_NAME}/@{SNAP_REVISION}/{,**/},
umount fstype=cifs /home/*/snap/@{SNAP_NAME}/common/{,**/},
umount fstype=cifs /var/snap/@{SNAP_NAME}/common/{,**/},
`

func init() {
	registerIface(&commonInterface{
		name:                  "cifs-mount-control",
		summary:               cifsMountControlSummary,
		implicitOnCore:        true,
		implicitOnClassic:     true,
		baseDeclarationSlots:  cifsMountControlBaseDeclarationSlots,
		connectedPlugAppArmor: cifsMountControlConnectedPlugAppArmor,
		connectedPlugSecComp:  cifsMountControlConnectedPlugSecComp,
		reservedForOS:         true,
	})
}