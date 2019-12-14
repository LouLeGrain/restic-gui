# Restic GUI

Restic GUI is a web app written in Go that helps with backing up with restic cli program written in go. 
The interface design is somewhat borrowed from ARQ Backup which uses the same methodology for backing up. 
A simple go server with bootstrap and jquery is used to build the frontend interface. 
Later I like to replace that will a Vue or Svelte SPA. 

Initial implementation should allow to
* create repositories stored in a local sqlite db (implemented)
* add repository path (first phase only local will be supported after initial proofe of concept SFTP, AWS S3, Back Blaze B2)
* initialise repository with password (implemented)
* create backup backup form source loaction (implemented) 
* choose source and destination locations (implemented)
* create backup snapshots (implemented)
* list files stored in snapshots (implemented) 
* restore files and directories from snapshots ti specifig directory (implemented) 
* use separate scheduler settings for every backup
* add os specific wrapper app to allow for on click start of the go web server
