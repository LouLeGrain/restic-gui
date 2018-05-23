# Restic GUI

Restic GUI is a web app written in Go that helps with backing up with restic cli program written in go. 
The interface design is somewhat borrowed fro ARQ Backup which uses the same methodology for backing up. 
A simple go server with bootstrap and jquery is used to build the frontend interface. 
Later I like to replace that will a Vue SAP. 

Initial implementation should allow to
* create and modify repositories stored in a local sqlite db 
* add repository path (first phase only local will be supported after initial proofe of concept SFTP, AWS S3, Back Blaze B2)
* initialise repository with password (currently manually created and entered into db)
* create and modify backup deatils 
* choose source and destination locations
* create backup snapshots (implemented)
* list files stored in snapshots (implemented) 
* restore files and directories from snapshots 
* use separate scheduler settings for every backup
* add os specific wrapper app to allow for on click start of the go web server
 
