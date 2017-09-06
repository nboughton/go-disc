# go-disc

go-disc is a basic MUD client built specifically for the [Discworld MUD](http://discworld.starturtle.net/lpc/). As such it does not and will not support some features that are common to other MUD clients such as scripting or triggers as both of these violate the rules of the Discworld.

Still in very early development. It can, theoretically, support other MUDs by implementing the sites.Site interface in a new go file under mud/sites (see Discworld.go for refernce) but I'm not personally interested in other MUDs so you might want to fork the project or submit a PR for other MUD support.

# To do
 - [x] Connect and receive data into the main window
 - [x] Editable command line in separate view
 - [ ] Password shadowing during login
 - [x] Scrollable command history
 - [ ] Tab completion - partially implemented
 - [ ] HP, GP, Burden meters in sidebar
 - [ ] Auto-mapping