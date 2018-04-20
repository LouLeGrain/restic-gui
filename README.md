# Restic GUI

Restic GUI is a web app written in Go that helps with backing up with restic cli program written in go .

Initial implementation should allow to
* create and modify repositories
** add repository path (first phase only local will be supported)
** initialise repository with password
* create and modify backup deatils
** choose source and destination locations
* create backup snapshots
* list files stored in snapshots
* restore files and directories from snapshots