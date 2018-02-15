# A Simple GUI for Rogue SSD's Data Cloaking Demo

This is a simple graphic UI for the RS data cloaking demo. For the demo to
run properly, it is important to set up the initial state of the machine 
as follows.

- The RS SSD device is NOT used for OS boot
- Visible namespace is /dev/nvme0n1 and formatted in ext4
- Visible namespace's mount point is /mnt/nvme, mounted.
- Hidden namespace is /dev/nvme0n2, and formatted in ext4
- hidden namespace's mount point /mnt/nvme2 exists, but NOT mounted

# Install dependencies
On Ubuntu (tested version 16.04), do 
```
sudo apt-get install build-essential libgtk-3-dev
go get github.com/andlabs/ui
```

On 64-bit Windows 10, install [Git](https://git-scm.com/download/win), and
[go1.9.3](https://dl.google.com/go/go1.9.3.windows-amd64.msi), then

```
go get github.com/andlabs/ui
```
Note that to make it work smoothly it is important to stick to go1.9.3.
The current stable Windows versions go1.9.4 and go1.8.7 have a bug that
prevents the go get from being succeed.

# Build GUI application
Use `make build` or simple `make` to build the GUI.

# Run GUI application

Run `./openclose` as root.
