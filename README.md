# Restic GUI

Restic GUI is a web app written in Go that helps with backing up with restic cli program written in go .

Initial implementation should allow to
* create and modify repositories stored in a local sqlite db 
* add repository path (first phase only local will be supported after initial proofe of concept SFTP, AWS S3, Back Blaze B2)
* initialise repository with password
* create and modify backup deatils
* choose source and destination locations
* create backup snapshots (implemented)
* list files stored in snapshots (implemented) 
* restore files and directories from snapshots 