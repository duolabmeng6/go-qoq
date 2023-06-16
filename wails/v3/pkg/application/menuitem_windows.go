//go:build windows

package application

import (
	"github.com/wailsapp/wails/v3/pkg/w32"
	"unsafe"
)

type windowsMenuItem struct {
	parent   *Menu
	menuItem *MenuItem

	hMenu     w32.HMENU
	id        int
	label     string
	disabled  bool
	checked   bool
	itemType  menuItemType
	hidden    bool
	submenu   w32.HMENU
	itemAfter *MenuItem
}

func (m *windowsMenuItem) setHidden(hidden bool) {
	m.hidden = hidden
	if m.hidden {
		// iterate the parent items and find the menu item before us
		for i, item := range m.parent.items {
			if item == m.menuItem {
				if i < len(m.parent.items)-1 {
					m.itemAfter = m.parent.items[i+1]
				} else {
					m.itemAfter = nil
				}
				break
			}
		}
		// Get the position of this menu item in the parent menu
		// m.pos = w32.GetMenuItemPosition(m.hMenu, uint32(m.id))
		// Remove from parent menu
		w32.RemoveMenu(m.hMenu, m.id, w32.MF_BYCOMMAND)
	} else {
		// Add to parent menu
		// Get the position of the item before us
		var pos int
		if m.itemAfter != nil {
			for i, item := range m.parent.items {
				if item == m.itemAfter {
					pos = i - 1
					break
				}
			}
			m.itemAfter = nil
		}
		w32.InsertMenuItem(m.hMenu, uint32(pos), true, m.getMenuInfo())
	}
}

func (m *windowsMenuItem) Checked() bool {
	return m.checked
}

func (m *windowsMenuItem) IsSeparator() bool {
	return m.itemType == separator
}

func (m *windowsMenuItem) IsCheckbox() bool {
	return m.itemType == checkbox
}

func (m *windowsMenuItem) Enabled() bool {
	return !m.disabled
}

func (m *windowsMenuItem) update() {
	w32.SetMenuItemInfo(m.hMenu, uint32(m.id), false, m.getMenuInfo())
}

func (m *windowsMenuItem) setLabel(label string) {
	m.label = label
	m.update()
}

func (m *windowsMenuItem) setDisabled(disabled bool) {
	m.disabled = disabled
	m.update()
}

func (m *windowsMenuItem) setChecked(checked bool) {
	m.checked = checked
	m.update()
}

func (m *windowsMenuItem) setAccelerator(accelerator *accelerator) {
	//// Set the keyboard shortcut of the menu item
	//var modifier C.int
	//var key *C.char
	//if accelerator != nil {
	//	modifier = C.int(toMacModifier(accelerator.Modifiers))
	//	key = C.CString(accelerator.Key)
	//}
	//
	//// Convert the key to a string
	//C.setMenuItemKeyEquivalent(m.nsMenuItem, key, modifier)
}

func newMenuItemImpl(item *MenuItem, parentMenu w32.HMENU, ID int) *windowsMenuItem {
	result := &windowsMenuItem{
		menuItem: item,
		hMenu:    parentMenu,
		id:       ID,
		disabled: item.disabled,
		checked:  item.checked,
		itemType: item.itemType,
		label:    item.label,
		hidden:   item.hidden,
	}

	return result
}

func newSpeechMenu() *MenuItem {
	panic("implement me")
}

func newHideMenuItem() *MenuItem {
	panic("implement me")

}

func newHideOthersMenuItem() *MenuItem {
	panic("implement me")

}

func newUnhideMenuItem() *MenuItem {
	panic("implement me")

}

func newUndoMenuItem() *MenuItem {
	panic("implement me")

}

// newRedoMenuItem creates a new menu item for redoing the last action
func newRedoMenuItem() *MenuItem {
	panic("implement me")

}

func newCutMenuItem() *MenuItem {
	panic("implement me")

}

func newCopyMenuItem() *MenuItem {
	panic("implement me")

}

func newPasteMenuItem() *MenuItem {
	panic("implement me")

}

func newPasteAndMatchStyleMenuItem() *MenuItem {
	panic("implement me")

}

func newDeleteMenuItem() *MenuItem {
	panic("implement me")

}

func newQuitMenuItem() *MenuItem {
	panic("implement me")

}

func newSelectAllMenuItem() *MenuItem {
	panic("implement me")

}

func newAboutMenuItem() *MenuItem {
	return newMenuItem("About " + globalApplication.options.Name).
		OnClick(func(ctx *Context) {
			globalApplication.ShowAboutDialog()
		})
}

func newCloseMenuItem() *MenuItem {
	panic("implement me")

}

func newReloadMenuItem() *MenuItem {
	panic("implement me")

}

func newForceReloadMenuItem() *MenuItem {
	panic("implement me")

}

func newToggleFullscreenMenuItem() *MenuItem {
	panic("implement me")

}

func newToggleDevToolsMenuItem() *MenuItem {
	panic("implement me")
}

func newZoomResetMenuItem() *MenuItem {
	panic("implement me")

}

func newZoomInMenuItem() *MenuItem {
	panic("implement me")

}

func newZoomOutMenuItem() *MenuItem {
	panic("implement me")

}

func newMinimizeMenuItem() *MenuItem {
	panic("implement me")
}

func newZoomMenuItem() *MenuItem {
	panic("implement me")
}

func newFullScreenMenuItem() *MenuItem {
	panic("implement me")
}

// ---------- unsupported on windows ----------

func (m *windowsMenuItem) setTooltip(_ string) {
	// Unsupported
}

func (m *windowsMenuItem) getMenuInfo() *w32.MENUITEMINFO {
	var mii w32.MENUITEMINFO
	mii.CbSize = uint32(unsafe.Sizeof(mii))
	mii.FMask = w32.MIIM_FTYPE | w32.MIIM_ID | w32.MIIM_STATE | w32.MIIM_STRING
	if m.IsSeparator() {
		mii.FType = w32.MFT_SEPARATOR
	} else {
		mii.FType = w32.MFT_STRING
		//var text string
		//if s := a.shortcut; s.Key != 0 {
		//	text = fmt.Sprintf("%s\t%s", a.text, s.String())
		//	shortcut2Action[a.shortcut] = a
		//} else {
		//	text = a.text
		//}
		mii.DwTypeData = w32.MustStringToUTF16Ptr(m.label)
		mii.Cch = uint32(len([]rune(m.label)))
	}
	mii.WID = uint32(m.id)
	if m.Enabled() {
		mii.FState &^= w32.MFS_DISABLED
	} else {
		mii.FState |= w32.MFS_DISABLED
	}

	if m.IsCheckbox() {
		mii.FMask |= w32.MIIM_CHECKMARKS
	}
	if m.Checked() {
		mii.FState |= w32.MFS_CHECKED
	}

	if m.menuItem.submenu != nil {
		mii.FMask |= w32.MIIM_SUBMENU
		mii.HSubMenu = m.submenu
	}
	return &mii
}
